package model

import "time"

// User represent user of social media
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// AuthUser represent data of user authenticate
type AuthUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Token represent data token to return
type Token struct {
	Token string    `json:"token"`
	User  UserToken `json:"user"`
}

// UserToken represent data token of user
type UserToken struct {
	Username string `json:"username"`
}
