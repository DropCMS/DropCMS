package utils

import (
	"database/sql"
	"time"
)

func NullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
func NullInt(i uint) sql.NullInt64 {
	u := int64(i)
	if u == 0 {
		return sql.NullInt64{}
	}
	return sql.NullInt64{
		Int64: u,
		Valid: true,
	}
}
func NullTime(t time.Time) sql.NullTime {
	if t.IsZero() {
		return sql.NullTime{}
	}
	return sql.NullTime{
		Time:  t,
		Valid: true,
	}
}
