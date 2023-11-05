package users

import (
	"bytes"
	"database/sql"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

//  ==== USER SCHEMAS ====

type UserSerializer struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

//  ==== USER STRUCT ====

type User struct {
	Username     string // Primary key
	passwordHash []byte
	Created_at   time.Time // Managed by db with default value
	Updated_at   time.Time // Managed by db with trigger
}

func (user *User) setPassword(pwd string) error {
	// bcrypt generate the salt from the pwd, no need to handle it ourself.
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.passwordHash = hash
	return nil
}

func (user *User) checkPassword(pwd string) (bool, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return false, err
	}
	if bytes.Equal(hash, user.passwordHash) {
		return true, nil
	}
	return false, nil
}

//  ==== USER DB INTERACTIONS ====

func SelectUsers(db *sql.DB) ([]User, error) {

	rows, err := db.Query(`SELECT * FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var user User
		if err := rows.Scan(
			&user.Username,
			&user.passwordHash,
			&user.Created_at,
			&user.Updated_at); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func SelectUser(db *sql.DB, username string) (User, error) {
	user := User{}

	err := db.QueryRow(`
	SELECT * FROM users WHERE username = ?
	`, username,
	).Scan(&user.Username, &user.passwordHash, &user.Created_at, &user.Updated_at)

	switch {
	case err == sql.ErrNoRows:
		return user, fmt.Errorf("no user with username %s", username)
	default:
		return user, err
	}
}

func InsertUser(db *sql.DB, user UserSerializer) (User, error) {
	var newUser User

	if user.Password == "" {
		return newUser, fmt.Errorf("cannot insert user: password is not set")
	}

	newUser.Username = user.Username
	newUser.setPassword(user.Password)

	tx, err := db.Begin()
	if err != nil {
		return newUser, err
	}

	err = tx.QueryRow(
		`INSERT INTO users (username, passwordHash) VALUES (?, ?) RETURNING username`,
		user.Username, newUser.passwordHash).Scan(&newUser.Username)
	if err != nil {
		return newUser, err
	}

	if tx.Commit() != nil {
		return newUser, err
	}

	return newUser, nil
}

func UpdateUser(db *sql.DB, user User) (User, error) {
	if user.passwordHash == nil {
		return user, fmt.Errorf("cannot update user: password is not set")
	}

	tx, err := db.Begin()
	if err != nil {
		return user, err
	}

	_, err = tx.Exec(`
		UPDATE users 
		SET Username = ?
		SET passwordHash = ?
		WHERE Username = ?`,
		user.Username,
		user.passwordHash,
		user.Username,
	)

	if err != nil {
		return user, err
	}

	err = tx.Commit()
	if err != nil {
		return user, err
	}

	return user, nil
}

func DeleteUser(db *sql.DB, user User) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		DELETE FROM users 
		WHERE username = $1
	`, user.Username)

	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
