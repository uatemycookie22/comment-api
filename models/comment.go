package models

import (
	"database/sql"
	"strconv"
)

type Comment struct {
	Id      int    `json:"id""`
	Message string `json:"message"`
}

var DB *sql.DB

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./identifier.sqlite")

	if err != nil {
		return err
	}

	DB = db

	return nil
}

// Returns the id of the created row
func CreateComment(newComment Comment) (int, error) {

	res, err := DB.Exec("INSERT INTO comments VALUES(NULL, ?)", newComment.Message)

	if err != nil {
		return 0, err
	}

	var id int64

	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}

	return int(id), nil
}

// Returns the id of the created row
func GetComments(count int) ([]Comment, error) {
	limitCount := strconv.Itoa(count)
	rows, err := DB.Query("SELECT id, message from comments LIMIT ?", limitCount)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	comments := make([]Comment, 0)

	for rows.Next() {
		singleComment := Comment{}
		err = rows.Scan(&singleComment.Id, &singleComment.Message)

		if err != nil {
			return nil, err
		}

		comments = append(comments, singleComment)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return comments, err
}
