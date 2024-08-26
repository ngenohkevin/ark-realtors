package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/ngenohkevin/ark-realtors/db/sqlc"
	"github.com/ngenohkevin/ark-realtors/internal/token"
	"github.com/ngenohkevin/ark-realtors/pkg/utils"
	"net/http"
	"time"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"` //once created cannot be edited
	FullName string `json:"full_name" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

type userResponse struct {
	ID                uuid.UUID `json:"id"`
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	Role              string    `json:"role"`
	CreatedAt         time.Time `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		ID:                user.ID,
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		Role:              user.Role,
		CreatedAt:         user.CreatedAt,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	userID := uuid.New()

	arg := db.CreateUserParams{
		ID:             userID,
		Username:       req.Username,
		FullName:       req.FullName,
		Email:          req.Email,
		HashedPassword: hashedPassword,
	}

	user, err := server.Store.CreateUser(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := newUserResponse(user)
	ctx.JSON(http.StatusOK, resp)
}

type getUserRequest struct {
	Username string `uri:"username" binding:"required"`
}

type getUserResponse struct {
	Id                uuid.UUID `json:"id"`
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	Role              string    `json:"role"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.Store.GetUser(ctx, req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := getUserResponse{
		Id:                user.ID,
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		Role:              user.Role,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
	authPayload := ctx.MustGet(AuthorizationPayloadKey).(*token.Payload)
	if authPayload.Role != utils.AdminRole && authPayload.Username != user.Username {
		err := errors.New("restricted access, you don't have the required permissions")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, resp)

}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	SessionID             uuid.UUID    `json:"session_id"`
	AccessToken           string       `json:"access_token"`
	AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
	RefreshToken          string       `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
	User                  userResponse `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.Store.GetUser(ctx, req.Username)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	err = utils.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	//create an access token for the user
	accessToken, accessPayload, err := server.TokenMaker.CreateToken(
		user.Username,
		user.Role,
		server.Config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	//create a refresh token for the user
	refreshToken, refreshPayload, err := server.TokenMaker.CreateToken(
		user.Username,
		user.Role,
		server.Config.RefreshTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	session, err := server.Store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := loginUserResponse{
		SessionID:             session.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User:                  newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, rsp)
}

type updateUserRequest struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// trying to update users
func (server *Server) updateUser(ctx *gin.Context) {

	// get the user id from the uri
	var uriReq struct {
		ID string `uri:"id" binding:"required"`
	}

	if err := ctx.ShouldBindUri(&uriReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// parse Id(string) to uuid
	ID, err := uuid.Parse(uriReq.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var req updateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Get the authenticated user's detail
	authPayload := ctx.MustGet(AuthorizationPayloadKey).(*token.Payload)
	authUser, err := server.Store.GetUser(ctx, authPayload.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Get the user to be updated
	targetUser, err := server.Store.GetUserById(ctx, ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	// Only an admin can update another user's role to admin
	if req.Role == utils.AdminRole && authUser.Role != utils.AdminRole {
		err := errors.New("only admin can perform this action")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	// Only an admin can update another user's details
	if authUser.Role != utils.AdminRole && authUser.Username != targetUser.Username {
		err := errors.New("restricted access, you don't have the required permissions")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	//Admin can only update their own details
	if authUser.Role == utils.AdminRole && authUser.Username != targetUser.Username {
		err := errors.New("you can only update your own details")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	// hash the password and update the password changed at
	var HashedPassword string
	var PasswordChangedAt pgtype.Timestamptz
	if req.Password != "" {
		HashedPassword, err = utils.HashPassword(req.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		PasswordChangedAt = pgtype.Timestamptz{Time: time.Now(), Valid: true}
	} else {
		HashedPassword = ""
		PasswordChangedAt = pgtype.Timestamptz{Valid: false}
	}

	arg := db.UpdateUserParams{
		Username:          pgtype.Text{String: req.Username, Valid: req.Username != ""},
		FullName:          pgtype.Text{String: req.FullName, Valid: req.FullName != ""},
		Email:             pgtype.Text{String: req.Email, Valid: req.Email != ""},
		HashedPassword:    pgtype.Text{String: HashedPassword, Valid: HashedPassword != ""},
		PasswordChangedAt: PasswordChangedAt,
		Role:              pgtype.Text{String: req.Role, Valid: req.Role != ""},
		ID:                ID,
	}

	user, err := server.Store.UpdateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	//Only the authenticated user can update their details, and an admin can update any user details

	//authPayload := ctx.MustGet(AuthorizationPayloadKey).(*token.Payload)
	//if authPayload.Role != utils.UserRole && authPayload.Username != user.Username {
	//	err := errors.New("restricted access, you don't have the required permissions")
	//	ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	//	return
	//}

	//if authPayload.Role == utils.AdminRole {
	//	update.Role = utils.NullStrings(req.Username)
	//}

	ctx.JSON(http.StatusOK, user)
}
