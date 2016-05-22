package models

import (
	"database/sql"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//User -
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//AllUsers -
func AllUsers(ctx *gin.Context) ([]*User, error) {
	db := ctx.MustGet("db").(*sql.DB)
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		user := new(User)
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

//FindUser -
func FindUser(ctx *gin.Context) (*User, error) {
	db := ctx.MustGet("db").(*sql.DB)
	email := strings.ToLower(ctx.PostForm("email"))
	user := new(User)
	err := db.QueryRow("SELECT * FROM users WHERE email=$1;", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//AddUser -
func AddUser(ctx *gin.Context) (*User, error) {
	db := ctx.MustGet("db").(*sql.DB)
	user := new(User)
	user.Name = ctx.PostForm("name")
	user.Email = strings.ToLower(ctx.PostForm("email"))
	b, err := bcrypt.GenerateFromPassword([]byte(ctx.PostForm("password")+"my secret pepper"), bcrypt.DefaultCost) // TODO MUST replace pepper string

	if err != nil {
		return nil, err
	}
	user.Password = string(b)
	err = db.QueryRow("INSERT INTO users(name,email,password) VALUES($1,$2,$3) returning id, created_at, updated_at;", &user.Name, &user.Email, &user.Password).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	user.Password = ""
	return user, nil
}

//RemoveUser -
func RemoveUser(ctx *gin.Context) (bool, error) {
	db := ctx.MustGet("db").(*sql.DB)
	email := strings.ToLower(ctx.Param("email"))
	_, err := db.Exec("DELETE FROM users WHERE email=$1", email)
	if err != nil {
		return false, err
	}
	return true, nil
}

/*
CREATE TABLE "users" (
	id bigserial primary key,
	name varchar(50) NOT NULL,
	email varchar(50) NOT NULL,
	password varchar(200) NOT NULL,
	created_at timestamp DEFAULT current_timestamp,
	updated_at timestamp DEFAULT current_timestamp,
	unique(email)
);

INSERT INTO "users" (name,email, password) VALUES ('rick','plumbus@fleeb.com','hashedpassword'), ('morty', 'dumbus@fleeb.com', 'hashedpassword');

*/
