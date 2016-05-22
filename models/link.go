package models

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//Link -
type Link struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Subject   string    `json:"string"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//AllLinks -
func AllLinks(ctx *gin.Context) ([]*Link, error) {
	db := ctx.MustGet("db").(*sql.DB)
	email := ctx.Param("email")
	rows, err := db.Query("SELECT * FROM links WHERE email=$1", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var links []*Link

	for rows.Next() {
		link := new(Link)
		err = rows.Scan(&link.ID,
			&link.Title,
			&link.Subject,
			&link.URL,
			&link.CreatedAt,
			&link.UpdatedAt)
		if err != nil {
			return nil, err
		}
		links = append(links, link)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return links, nil
}

//AddLink -
func AddLink(ctx *gin.Context) (*Link, error) {
	db := ctx.MustGet("db").(*sql.DB)
	link := new(Link)
	link.Title = ctx.PostForm("title")
	link.Subject = ctx.PostForm("subject")
	link.URL = ctx.PostForm("url")

	err := db.QueryRow(`
		INSERT INTO
		links(title, subject, url)
		VALUES($1,$2,$3) returning id;`,
		&link.Title,
		&link.Subject,
		&link.URL).Scan(&link.ID)

	if err != nil {
		return nil, err
	}
	return link, nil
}

// //RemoveLink -
// func RemoveLink(ctx *gin.Context) (bool, error) {
// 	db := ctx.MustGet("db").(*sql.DB)
// 	id := ctx.Param("id")
// 	_, err := db.Exec("DELETE FROM links WHERE id=$1", id)
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }

//UpdateLink -
func UpdateLink(ctx *gin.Context) (*Link, error) {
	db := ctx.MustGet("db").(*sql.DB)
	link := new(Link)
	link.ID, _ = strconv.Atoi(ctx.Param("id"))
	link.Title = ctx.PostForm("title")
	link.Subject = ctx.PostForm("subject")
	link.UpdatedAt = time.Now()

	err := db.QueryRow("UPDATE links SET (title, subject, updated_at) = ($1,$2,$3) WHERE id=$4 returning id;",
		&link.Title,
		&link.Subject,
		&link.UpdatedAt,
		&link.ID).Scan(&link.ID)

	if err != nil {
		return nil, err
	}
	return link, nil
}

/*
CREATE TABLE "links" (
	id bigserial primary key,
	title varchar(100) NOT NULL,
	subject varchar(100) NOT NULL,
	url varchar(100),
	created_at timestamp DEFAULT current_timestamp,
	updated_at timestamp DEFAULT current_timestamp
);


CREATE TABLE "notes" (
	id bigserial primary key,
	link_id bigserial references links,
	user_id bigserial references users,
	entry text NOT NULL,
	created_at timestamp DEFAULT current_timestamp,
	updated_at timestamp DEFAULT current_timestamp,
);

/// user links
Priority          int       `json:"priority"` // 1- > low, 2 -> medium, 3 -> high
Checkoff          int       `json:"checkoff"` //1 -> daily, 2 -> weekly, 3 -> monthly
CheckoffTimeStamp time.Time `json:"checkoff_time"`
Note              string    `json:"note"`

priority integer NOT NULL,
checkoff integer NOT NULL,
checkoff_time timestamp DEFAULT current_timestamp - interval '1 day',
note text,

INSERT INTO "links" (title, subject ,url, priority, checkoff, note) VALUES ('Mylink','Math','www.linkofmine.com',1,2,'This is my favorite link'), ('Mylink2','Science man','www.anotherlink.com',2,2,'This is ok'),('A Link for smart people','Science again','www.quantumawesome.com',1,1,'I will do it!');
*/
