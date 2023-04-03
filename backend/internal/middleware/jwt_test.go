package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"user-management-api/internal/models"
	"user-management-api/internal/utils"
)

func TestValidateToken(t *testing.T) {
	// Create a dummy handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	// Create a new request and recorder
	req, err := http.NewRequest("GET", "https://api.fadli.dev/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	// Create a new ValidateToken middleware and wrap the dummy handler
	middleware := ValidateToken(handler)

	// Test unauthorized request
	middleware.ServeHTTP(rr, req)
	if rr.Code != http.StatusUnauthorized {
		t.Errorf("Status code = %v; want %v", rr.Code, http.StatusUnauthorized)
	}

	// Test invalid token request
	req.Header.Set("Authorization", "Bearer invalidation-token")
	rr = httptest.NewRecorder()
	middleware.ServeHTTP(rr, req)
	if rr.Code != http.StatusUnauthorized {
		t.Errorf("Status code = %v; want %v", rr.Code, http.StatusUnauthorized)
	}

	var dummyUser models.User
	dummyUser.ID = 69
	dummyUser.Username = "UnitDummy"
	dummyUser.Password = "unitDPassword"
	dummyUser.Email = "dummy@user.test"

	// Test valid token request
	tokenString, _ := utils.GenerateToken(&dummyUser)
	req.Header.Set("Authorization", "Bearer "+tokenString)
	rr = httptest.NewRecorder()
	middleware.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Status code = %v; want %v", rr.Code, http.StatusOK)
	}

}
