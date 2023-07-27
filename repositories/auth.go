package repositories

import (
	"database/sql"
)

type User struct {
	ID             int
	Email          string
	HashedPassword string
}

func CheckUserExist(db *sql.DB, email string) (bool, error) {
	const sqlStr = `select exists (select 1 from users where email = ?);`
	var exists bool
	err := db.QueryRow(sqlStr, email).Scan(&exists)
	return exists, err
}

func CreateUser(db *sql.DB, email, hashedPassword string) (sql.Result, error) {
	const sqlStr = `insert into users (email, password, created_at) values (?, ?, now()); `
	return db.Exec(sqlStr, email, hashedPassword)
}

func GetUser(db *sql.DB, email string) (User, error) {
	const sqlStr = `select id, email, password from users where email = ?;`
	var user User
	err := db.QueryRow(sqlStr, email).Scan(&user.ID, &user.Email, &user.HashedPassword)
	return user, err
}
