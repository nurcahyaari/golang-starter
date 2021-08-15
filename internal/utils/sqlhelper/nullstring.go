package sqlhelper

import "database/sql"

func NewNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}

	return sql.NullString{
		Valid:  true,
		String: s,
	}
}

func GetNullString(s sql.NullString) string {
	if s.Valid {
		return s.String
	}

	return ""
}
