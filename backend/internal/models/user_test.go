package models_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"user-management-api/internal/models"
)

func TestUser_JSONMarshalling(t *testing.T) {
	// given
	user := models.User{
		ID:        1,
		Username:  "johndoe",
		Email:     "johndoe@example.com",
		CreatedAt: "2022-01-01 00:00:00",
		UpdatedAt: "2022-01-01 00:00:00",
	}
	expectedJSON := `{"id":1,"username":"johndoe","email":"johndoe@example.com","created_at":"2022-01-01 00:00:00","updated_at":"2022-01-01 00:00:00"}`

	// when
	jsonData, err := json.Marshal(user)

	// then
	assert.Nil(t, err)
	assert.Equal(t, expectedJSON, string(jsonData))
}
