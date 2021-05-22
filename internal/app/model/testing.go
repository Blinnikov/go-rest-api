package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@domain.com",
		Password: "Tryt0|=|c1cKm3",
	}
}
