package repositories

import (
	"database/sql"
)

func CheckUserExist(db *sql.DB, email string) (bool, error) {
	const sqlStr = `select exists (select 1 from users where email = ?);`
	var exists bool
	err := db.QueryRow(sqlStr, email).Scan(&exists)
	return exists, err
}

func CreateUser(db *sql.DB, email, hashedPassword string) error {
	const sqlStr = `insert into users (email, password, created_at) values (?, ?, now()); `
	_, err := db.Exec(sqlStr, email, hashedPassword)
	return err
}

func GetUserPassword(db *sql.DB, email string) (string, error) {
	const sqlStr = `select password from users where email = ?;`
	var hashedPassword string
	err := db.QueryRow(sqlStr, email).Scan(&hashedPassword)
	return hashedPassword, err
}
