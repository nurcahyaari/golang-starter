package sqlhelper

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNullString(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		t.Run("Null", func(t *testing.T) {
			r := NewNullString("")

			assert.Equal(t, sql.NullString{}, r)
		})

		t.Run("NotNull", func(t *testing.T) {
			r := NewNullString("test")

			assert.Equal(t, sql.NullString{Valid: true, String: "test"}, r)
		})
	})

	t.Run("Get", func(t *testing.T) {
		t.Run("Null", func(t *testing.T) {
			r := GetNullString(sql.NullString{})

			assert.Equal(t, "", r)
		})

		t.Run("NotNull", func(t *testing.T) {
			r := GetNullString(sql.NullString{Valid: true, String: "test"})

			assert.Equal(t, "test", r)
		})
	})
}
