package pkg

import (
	"github.com/in-memory-rest/internal/domain"
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

func TestMinStrLen(t *testing.T) {
	t.Run("with wrong min length field", func(t *testing.T) {
		u := domain.User{
			FirstName: "John",
		}

		changeset := setupChangeset()
		changeset.MinStrLen("FirstName", u.FirstName, 8)

		isValid := changeset.IsValid
		errors := changeset.Errors

		assert.Falsef(t, isValid, "changeset should be invalid")
		assert.Equal(t, errors["FirstName"], "minimum length of 'FirstName' is 8. Got: '4'")
	})

	t.Run("with valid min length field", func(t *testing.T) {
		u := domain.User{
			FirstName: "John",
		}

		changeset := setupChangeset()
		changeset.MinStrLen("FirstName", u.FirstName, 4)

		isValid := changeset.IsValid
		errors := changeset.Errors

		assert.True(t, isValid)
		assert.Empty(t, errors["FirstName"])
	})
}

func TestMaxStrLen(t *testing.T) {
	t.Run("with wrong max length field", func(t *testing.T) {
		u := domain.User{
			FirstName: "John Doe",
		}

		changeset := setupChangeset()
		changeset.MaxStrLen("FirstName", u.FirstName, 4)

		isValid := changeset.IsValid
		errors := changeset.Errors

		assert.Falsef(t, isValid, "changeset should be invalid")
		assert.Equal(t, errors["FirstName"], "maximum length of 'FirstName' is 4. Got: '8'")
	})

	t.Run("with valid max length field", func(t *testing.T) {
		u := domain.User{
			FirstName: "John Doe",
		}

		changeset := setupChangeset()
		changeset.MaxStrLen("FirstName", u.FirstName, 12)

		isValid := changeset.IsValid
		errors := changeset.Errors

		assert.True(t, isValid)
		assert.Empty(t, errors["FirstName"])
	})
}
