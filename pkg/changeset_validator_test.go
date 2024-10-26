package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func setupChangeset() *ChangesetValidator {
	changeset := NewChangesetValidator()

	return changeset
}

func TestCastToString(t *testing.T) {
	t.Run("with wrong type field", func(t *testing.T) {
		m := map[string]any{
			"Name":     "John Doe",
			"NickName": "JD",
			"Age":      31,
		}

		changeset := setupChangeset()
		changeset.CastToString(m)

		isValid := changeset.IsValid
		errors := changeset.Errors

		assert.Falsef(t, isValid, "changeset should be invalid")
		assert.Equal(t, errors["Age"], "field 'Age' is not string type")
	})

	t.Run("with all valid types", func(t *testing.T) {
		m := map[string]any{
			"Name":     "John Doe",
			"NickName": "JD",
		}

		changeset := setupChangeset()
		changeset.CastToString(m)

		isValid := changeset.IsValid
		errors := changeset.Errors

		assert.True(t, isValid)
		assert.Empty(t, errors)
	})

}

func TestValidateRequired(t *testing.T) {
	t.Run("with missing field", func(t *testing.T) {
		m := map[string]any{
			"Name":     "John Doe",
			"NickName": "JD",
		}

		fields := []string{"Name", "NickName", "Age"}

		changeset := setupChangeset()
		changeset.CastToString(m)
		changeset.ValidateRequired(m, fields)

		isValid := changeset.IsValid
		errors := changeset.Errors

		assert.False(t, isValid)
		assert.Equal(t, errors["Age"], "Missing required field 'Age'")
	})

	t.Run("with empty field", func(t *testing.T) {
		m := map[string]any{
			"Name":     "",
			"NickName": "JD",
		}

		fields := []string{"Name", "NickName", "Age"}

		changeset := setupChangeset()
		changeset.CastToString(m)
		changeset.ValidateRequired(m, fields)

		isValid := changeset.IsValid
		errors := changeset.Errors

		assert.False(t, isValid)
		assert.Equal(t, errors["Age"], "Missing required field 'Age'")
		assert.Equal(t, errors["Name"], "Empty required field 'Name'")
	})
}
