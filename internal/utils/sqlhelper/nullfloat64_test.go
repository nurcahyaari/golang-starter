package sqlhelper

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNullFloat64(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		t.Run("Null", func(t *testing.T) {
			r := NewNullFloat64(0)

			assert.Equal(t, sql.NullFloat64{}, r)
		})

		t.Run("NotNull", func(t *testing.T) {
			r := NewNullFloat64(10)

			assert.Equal(t, sql.NullFloat64{Valid: true, Float64: 10}, r)
		})
	})

	t.Run("Get", func(t *testing.T) {
		t.Run("Null", func(t *testing.T) {
			r := GetNullFloat64(sql.NullFloat64{})

			assert.Equal(t, float64(0), r)
		})

		t.Run("NotNull", func(t *testing.T) {
			r := GetNullFloat64(sql.NullFloat64{Valid: true, Float64: 10})

			assert.Equal(t, float64(10), r)
		})
	})
}
