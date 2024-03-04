package models

import (
	"database/sql"
)

type Comment struct {
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
