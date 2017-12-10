package db

type User struct {
	ID       int    `gorm:"primary_key"`
	Name     string `gorm:"unique;size:32"`
	IsGroup  bool
	Password string
}

type Repository struct {
	ID     int `gorm:"primary_key"`
	User   User
	UserID int
	Name   string
}
