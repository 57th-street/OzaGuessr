package repositories

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/57th-street/oza-gueser/models"
	"github.com/57th-street/oza-gueser/utils"
)

func CheckUserExist(db *sql.DB, email string) (bool, error) {
	const sqlStr = `select exists (select 1 from users where email = ?);`
	var exists bool
	err := db.QueryRow(sqlStr, email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func Register(db *sql.DB, user models.User) (models.User, error) {
	const sqlStr = `insert into users (email, password, created_at) values (?, ?, now()); `
	encryptPw, err := utils.EncryptPassword(user.Password)
	result, err := db.Exec(sqlStr, user.Email, encryptPw)
	if err != nil {
		return models.User{}, err
	}
	id, _ := result.LastInsertId()
	var newUser models.User
	newUser.ID, newUser.Email, newUser.Password = int(id), user.Email, user.Password
	return newUser, nil
}

func Login(db *sql.DB, user models.User) (models.User, error) {
	const sqlStr = `select * from users where email ?;`
	row := db.QueryRow(sqlStr, user.Email)
	if err := row.Err(); err != nil {
		return models.User{}, err
	}
	var loginUser models.User
	err := row.Scan(&loginUser.ID, &loginUser.Email, &loginUser.Password, &loginUser.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, fmt.Errorf("invalid email or password")
		}
		return models.User{}, err
	}
	err = utils.CompareHashAndPassword(loginUser.Password, user.Password)
	if err != nil {
		return models.User{}, fmt.Errorf("invalid email or password")
	}
	return loginUser, nil
}
