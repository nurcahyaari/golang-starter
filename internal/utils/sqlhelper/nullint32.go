package sqlhelper

import "database/sql"

func NewNullInt32(i int32) sql.NullInt32 {
	if i == 0 {
		return sql.NullInt32{}
	}

	return sql.NullInt32{
		Valid: true,
		Int32: i,
	}
}

func GetNullInt32(i sql.NullInt32) int32 {
	if i.Valid {
		return i.Int32
	}

	return 0
}
