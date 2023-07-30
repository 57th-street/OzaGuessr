package repositories

import (
	"database/sql"

	"github.com/57th-street/oza-gueser/models"
)

func GetProfile(db *sql.DB, userId int) (models.Profile, error) {
	var profile models.Profile
	const sqlStr = `select name, image_url, intro form users where id = ?;`
	err := db.QueryRow(sqlStr, userId).Scan(&profile.Name, &profile.ImageUrl, &profile.Intro)
	return profile, err
}

func UpdateProfile(db *sql.DB, userId int, profile models.Profile) error {
	const sqlStr = `update users set name = ?, image_url = ?, intro = ? where id = ?;`
	_, err := db.Exec(sqlStr, profile.Name, profile.ImageUrl, profile.Intro, userId)
	return err
}
