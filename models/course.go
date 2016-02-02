package models

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type Course struct {
	Id                int       `json:"id"`
	Name              string    `json:"name"`
	Email             string    `json:"email"` // classmate email key
	Url               string    `json:"url"`
	Priority          int       `json:"priority"` // 1- > low, 2 -> medium, 3 -> high
	Checkoff          int       `json:"checkoff"` //1 -> daily, 2 -> weekly, 3 -> monthly
	CheckoffTimeStamp time.Time `json:"checkoff_time"`
	Note              string    `json:"note"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func AllCourses(ctx *gin.Context) ([]*Course, error) {
	db := ctx.MustGet("db").(*sql.DB)
	email := ctx.Param("email")
	rows, err := db.Query("SELECT * FROM courses WHERE email=$1", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	courses := make([]*Course, 0)

	for rows.Next() {
		course := new(Course)
		err := rows.Scan(&course.Id,
			&course.Name,
			&course.Email,
			&course.Url,
			&course.Priority,
			&course.Checkoff,
			&course.CheckoffTimeStamp,
			&course.Note,
			&course.CreatedAt,
			&course.UpdatedAt)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return courses, nil
}

func AddCourse(ctx *gin.Context) (*Course, error) {
	db := ctx.MustGet("db").(*sql.DB)
	course := new(Course)
	course.Name = ctx.PostForm("name")
	course.Email = ctx.PostForm("email")
	course.Url = ctx.PostForm("url")
	course.Priority, _ = strconv.Atoi(ctx.PostForm("priority"))
	course.Checkoff, _ = strconv.Atoi(ctx.PostForm("checkoff"))
	course.Note = ctx.PostForm("note")

	err := db.QueryRow(`
		INSERT INTO 
		courses(name, email, url, priority, checkoff, note) 
		VALUES($1,$2,$3,$4,$5,$6) returning id;`,
		&course.Name,
		&course.Email,
		&course.Url,
		&course.Priority,
		&course.Checkoff,
		&course.Note).Scan(&course.Id)

	if err != nil {
		return nil, err
	}
	return course, nil
}

func RemoveCourse(ctx *gin.Context) (bool, error) {
	db := ctx.MustGet("db").(*sql.DB)
	id := ctx.Param("id")
	_, err := db.Exec("DELETE FROM courses WHERE id=$1", id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func UpdateCourse(ctx *gin.Context) (*Course, error) {
	db := ctx.MustGet("db").(*sql.DB)
	course := new(Course)
	course.Id, _ = strconv.Atoi(ctx.Param("id"))
	course.Name = ctx.PostForm("name")
	course.Url = ctx.PostForm("url")
	course.Priority, _ = strconv.Atoi(ctx.PostForm("priority"))
	course.Checkoff, _ = strconv.Atoi(ctx.PostForm("checkoff"))
	checkoffTimeStamp := ctx.PostForm("checkoff_time") // , _ = time.Parse(time.RFC3339, ctx.PostForm("checkoff_time"))
	course.Note = ctx.PostForm("note")
	course.UpdatedAt = time.Now()

	err := db.QueryRow("UPDATE courses SET (name, url, priority, checkoff, checkoff_time, note, updated_at) = ($1,$2,$3,$4,$5,$6,$7) WHERE id=$8 returning id;",
		&course.Name,
		&course.Url,
		&course.Priority,
		&course.Checkoff,
		checkoffTimeStamp,
		&course.Note,
		&course.UpdatedAt,
		&course.Id).Scan(&course.Id)

	if err != nil {
		return nil, err
	}
	return course, nil
}

/*
CREATE TABLE "courses" (
	id bigserial primary key,
	name varchar(100) NOT NULL,
	email varchar(50) references classmates(email),
	url varchar(100),
	priority integer NOT NULL,
	checkoff integer NOT NULL,
	checkoff_time timestamp DEFAULT current_timestamp - interval '1 day',
	note text,
	created_at timestamp DEFAULT current_timestamp,
	updated_at timestamp DEFAULT current_timestamp
);


CREATE TABLE "notes" (
	id bigserial primary key,
	course_id bigserial references courses,
	entry text NOT NULL,
	created_at timestamp DEFAULT current_timestamp,
	updated_at timestamp DEFAULT current_timestamp,
);

CREATE TABLE "urls" (
	id bigserial primary key,
	course_id bigserial references courses,
	url varchar(100) NOT NULL,
	created_at timestamp DEFAULT current_timestamp,
	updated_at timestamp DEFAULT current_timestamp,
);

INSERT INTO "courses" (name, email, url, priority, checkoff, note) VALUES ('Mycourse','plumbus@fleeb.com','www.courseofmine.com',1,2,'This is my favorite course'), ('Mycourse2','plumbus@fleeb.com','www.anothercourse.com',2,2,'This is ok'),('A Course for smart people','dumbus@fleeb.com','www.quantumawesome.com',1,1,'I will do it!');
*/
