package structs

import "database/sql"

type Blog struct {
	Id      int
	Title   string
	Content string
}

type User struct {
	Id       int
	Username string
	Password string
	CreateAt string
	UpdateAt sql.NullString
}

type NotificationStateSignup struct {
	Announcement string
	Message      string
}
