package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"crud-test-go/models"
	"crud-test-go/services"

	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	// Arrange
	services.AddUser(models.User{Name: "Test User", Email: "test@example.com"})
	req := httptest.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()

	// Act
	GetUsers(w, req)

	// Assert
	res := w.Result()
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Contains(t, string(body), "Test User 1")
}

func TestCreateUser(t *testing.T) {
	user := models.User{Name: "New User", Email: "new@example.com"}
	userJSON, _ := json.Marshal(user)
	req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	CreateUser(w, req)
	res := w.Result()
	defer res.Body.Close()
	var createdUser models.User
	json.NewDecoder(res.Body).Decode(&createdUser)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "New User", createdUser.Name)
	assert.Equal(t, "new@example.com", createdUser.Email)
	assert.NotZero(t, createdUser.ID)
}


func TestCreateUser_DuplicateName(t *testing.T) {
	// Arrange: first create the user
	user := models.User{Name: "duplicate", Email: "dupe@example.com"}
	userJSON, _ := json.Marshal(user)
	req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	CreateUser(w, req)

	// Act: try creating the same user again
	req2 := httptest.NewRequest("POST", "/users", bytes.NewBuffer(userJSON))
	req2.Header.Set("Content-Type", "application/json")
	w2 := httptest.NewRecorder()
	CreateUser(w2, req2)

	// Assert
	res := w2.Result()
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Contains(t, string(body), "username already taken")
}
