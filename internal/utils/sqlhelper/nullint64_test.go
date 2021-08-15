package sqlhelper

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNullInt64(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		t.Run("Null", func(t *testing.T) {
			r := NewNullInt64(0)

			assert.Equal(t, sql.NullInt64{}, r)
		})

		t.Run("NotNull", func(t *testing.T) {
			r := NewNullInt64(10)

			assert.Equal(t, sql.NullInt64{Valid: true, Int64: 10}, r)
		})
	})

	t.Run("Get", func(t *testing.T) {
		t.Run("Null", func(t *testing.T) {
			r := GetNullInt64(sql.NullInt64{})

			assert.Equal(t, int64(0), r)
		})

		t.Run("NotNull", func(t *testing.T) {
			r := GetNullInt64(sql.NullInt64{Valid: true, Int64: 10})

			assert.Equal(t, int64(10), r)
		})
	})
}
