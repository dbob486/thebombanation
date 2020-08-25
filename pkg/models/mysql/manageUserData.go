package mysql

import (
	"database/sql"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) uploadData() {
	//upload video to DB
}

func (m *UserModel) getData() {
	//get video from DB
}

func (m *UserModel) deleteData() {
	//delete video from DB
}