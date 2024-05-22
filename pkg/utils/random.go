package utils

import (
	"github.com/google/uuid"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstwxyz"

func init() {
	rand.NewSource(time.Now().UnixNano())
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomBytes(n int) []byte {
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, n)
	k := len(alphabet)
	for i := 0; i < n; i++ {
		result[i] = alphabet[rand.Intn(k)]
	}
	return result
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// RandomEmail generates a random email address
func RandomEmail() string {
	return RandomString(6) + "@" + RandomString(6) + ".com"
}

func RandomUsername() string {
	return RandomString(6)
}

func RandomFullName() string {
	return RandomString(6) + " " + RandomString(6)
}

func RandomStatus() string {
	status := []string{"available", "sold"}
	n := len(status)
	return status[rand.Intn(n)]
}

func RandomPrice() int {
	return RandomInt(1000, 100000)
}

func RandomRole() string {
	roles := []string{"admin", "user", "agent", "client"}
	n := len(roles)
	return roles[rand.Intn(n)]
}

func GenerateRandomUserID() uuid.UUID {
	return uuid.New()
}
