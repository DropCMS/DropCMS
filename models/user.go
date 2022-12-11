package models

import (
	"context"
	"database/sql"
	"dropCms/db"
	"errors"
	"net/mail"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       uint
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Passowrd string `json:"password"`
}

func (u User) Create() error {
	if len(u.Name) == 0 && len(u.Email) == 0 && len(u.Username) == 0 && len(u.Passowrd) < 8 {
		if len(u.Name) == 0 {
			return errors.New("Input valid name.")
		}
		if len(u.Email) == 0 {
			return errors.New("Input valid email address.")
		}
		if len(u.Username) == 0 {
			return errors.New("Input valid username.")
		}
		if len(u.Passowrd) < 8 {
			return errors.New("Password must be more than or equal to 8 digit.")
		}
	}
	_, ok := mail.ParseAddress(u.Email)
	if ok != nil {
		return ok
	}
	hlod, errh := bcrypt.GenerateFromPassword([]byte(u.Passowrd), bcrypt.DefaultCost)
	if errh != nil {
		return errh
	}
	query := "INSERT INTO users(name, email, username, password) VALUES (?, ?, ?, ?)"
	ctx, cancelCtx := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelCtx()
	stmt, err := db.Db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, rrr := stmt.ExecContext(ctx, u.Name, u.Email, u.Username, hlod)
	if rrr != nil {
		return rrr
	}
	return nil
}

func (u User) Get() (interface{}, error) {
	if len(u.Username) == 0 {
		return nil, errors.New("Invalid input")
	}
	query := "SELECT id, username, name, email, password FROM users WHERE username = ? LIMIT 1"
	ctx, cancelCtx := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelCtx()
	stmt, err := db.Db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, u.Username)
	user := new(User)
	errr := row.Scan(&user.Id, &user.Username, &user.Name, &user.Email, &user.Passowrd)
	switch {
	case errr == sql.ErrNoRows:
		return nil, errors.New("No User Found.")
	case errr != nil:
		return nil, errr
	default:
		return user, nil
	}
}
