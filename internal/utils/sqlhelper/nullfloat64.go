package sqlhelper

import "database/sql"

func NewNullFloat64(f float64) sql.NullFloat64 {
	if f == 0 {
		return sql.NullFloat64{}
	}

	return sql.NullFloat64{
		Valid:   true,
		Float64: f,
	}
}

func GetNullFloat64(i sql.NullFloat64) float64 {
	if i.Valid {
		return i.Float64
	}

	return 0
}
