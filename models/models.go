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
type Theme struct {
	ID   int    `json:"theme_id"`
	Name string `json:"name"`
}

type Quiz struct {
	ID        int       `json:"quiz_id"`
	Quest     string    `json:"quest"`
	Answer    string    `json:"answer"`
	Descript  string    `json:"descript"`
	ThemeID   int       `json:"theme_id"`
	CreatedAt time.Time `json:"created_at"`
}
