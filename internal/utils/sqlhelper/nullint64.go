package sqlhelper

import "database/sql"

func NewNullInt64(i int64) sql.NullInt64 {
	if i == 0 {
		return sql.NullInt64{}
	}

	return sql.NullInt64{
		Valid: true,
		Int64: i,
	}
}

func GetNullInt64(i sql.NullInt64) int64 {
	if i.Valid {
		return i.Int64
	}

	return 0
}
