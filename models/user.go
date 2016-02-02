package models

import (
	"bytes"
	"database/sql"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func AllUsers(ctx *gin.Context) ([]*User, error) {
	db := ctx.MustGet("db").(*sql.DB)
	rows, err := db.Query("SELECT * FROM classmates")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*User, 0)
	for rows.Next() {
		user := new(User)
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
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

func FindUser(ctx *gin.Context) (*User, error) {
	db := ctx.MustGet("db").(*sql.DB)
	email := ctx.PostForm("email")
	user := new(User)
	err := db.QueryRow("SELECT * FROM classmates WHERE email=$1;", email).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func AddUser(ctx *gin.Context) (*User, error) {
	db := ctx.MustGet("db").(*sql.DB)
	user := new(User)
	user.Name = ctx.PostForm("name")
	user.Email = ctx.PostForm("email")
	b, err := bcrypt.GenerateFromPassword([]byte(ctx.PostForm("password")+"my secret pepper"), bcrypt.DefaultCost) // TODO MUST replace pepper string
	n := bytes.IndexByte(b, 0)
	user.Password = string(b[:n])
	err = db.QueryRow("INSERT INTO classmates(name,email,password) VALUES($1,$2,$3) returning id;", &user.Name, &user.Email, &user.Password).Scan(&user.Id)
	if err != nil {
		return nil, err
	}
	user.Password = ""
	return user, nil
}

func RemoveUser(ctx *gin.Context) (bool, error) {
	db := ctx.MustGet("db").(*sql.DB)
	email := ctx.Param("email")
	_, err := db.Exec("DELETE FROM classmates WHERE email=$1", email)
	if err != nil {
		return false, err
	}
	return true, nil
}

/*
CREATE TABLE "classmates" (
	id bigserial primary key,
	name varchar(50) NOT NULL,
	email varchar(50) NOT NULL,
	password varchar(200) NOT NULL,
	created_at timestamp DEFAULT current_timestamp,
	updated_at timestamp DEFAULT current_timestamp,
	unique(email)
);

INSERT INTO "classmates" (name,email, password) VALUES ('rick','plumbus@fleeb.com','hashedpassword'), ('morty', 'dumbus@fleeb.com', 'hashedpassword');

*/
