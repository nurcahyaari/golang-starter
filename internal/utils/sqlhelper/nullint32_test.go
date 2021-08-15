package sqlhelper

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNullInt32(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		t.Run("Null", func(t *testing.T) {
			r := NewNullInt32(0)

			assert.Equal(t, sql.NullInt32{}, r)
		})

		t.Run("NotNull", func(t *testing.T) {
			r := NewNullInt32(10)

			assert.Equal(t, sql.NullInt32{Valid: true, Int32: 10}, r)
		})
	})

	t.Run("Get", func(t *testing.T) {
		t.Run("Null", func(t *testing.T) {
			r := GetNullInt32(sql.NullInt32{})

			assert.Equal(t, int32(0), r)
		})

		t.Run("NotNull", func(t *testing.T) {
			r := GetNullInt32(sql.NullInt32{Valid: true, Int32: 10})

			assert.Equal(t, int32(10), r)
		})
	})
}
