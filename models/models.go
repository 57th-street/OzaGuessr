package models

import "time"

type User struct {
	ID       int    `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Profile
	CreatedAt time.Time `json:"created_at"`
}

type Profile struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
	Intro    string `json:"intro"`
}
