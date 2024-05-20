package utils

import "github.com/jackc/pgx/v5/pgtype"

func NullStrings(s string) pgtype.Text {
	return pgtype.Text{
		String: s,
		Valid:  true,
	}
}

func NullInt64(i int64) pgtype.Int8 {
	return pgtype.Int8{
		Int64: i,
		Valid: true,
	}
}
