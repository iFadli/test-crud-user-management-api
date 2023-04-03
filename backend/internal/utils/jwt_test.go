package utils_test

import (
	"errors"
	"testing"
	"user-management-api/internal/models"
	"user-management-api/internal/utils"

	"github.com/dgrijalva/jwt-go"
)

func TestGenerateToken(t *testing.T) {
	// prepare test data
	user := &models.User{
		ID:       100,
		Username: "unittest",
		Email:    "unit@test.ok",
	}

	// execute the function
	token, err := utils.GenerateToken(user)

	// check the result
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if token == "" {
		t.Error("token should not be empty")
	}

	// validate the token
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(utils.JWT_SECRET), nil
	})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !parsedToken.Valid {
		t.Error("token should be valid")
	}
}

func TestVerifyToken(t *testing.T) {
	// prepare test data
	user := &models.User{
		ID:       122,
		Username: "verifyunit",
		Email:    "verify@unit.xyz",
	}
	token, err := utils.GenerateToken(user)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// execute the function
	parsedToken, err := utils.VerifyToken(token)

	// check the result
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if parsedToken == nil {
		t.Error("parsedToken should not be nil")
	}

	// validate the token
	if !parsedToken.Valid {
		t.Error("token should be valid")
	}

	// check the token claims
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		t.Error("token claims should be of type MapClaims")
	}

	if claims["id"] != float64(user.ID) {
		t.Errorf("unexpected id: token-%v | user-%v", claims["id"], user.ID)
	}
	if claims["username"] != user.Username {
		t.Errorf("unexpected username: %v", claims["username"])
	}
	if claims["email"] != user.Email {
		t.Errorf("unexpected email: %v", claims["email"])
	}
}

//
//func checkInterface(i interface{}) {
//	// memeriksa tipe data pada interface{}
//	switch i.(type) {
//	case int:
//		fmt.Println("nilai bertipe int")
//	case float64:
//		fmt.Println("nilai bertipe float64")
//	case string:
//		fmt.Println("nilai bertipe string")
//	default:
//		fmt.Println("tipe data tidak dikenali")
//	}
//}
