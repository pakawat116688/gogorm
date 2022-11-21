package repository

type User struct {
	ID       int    `gorm:"column:uid;autoIncrement;primaryKey"`
	Username string `gorm:"column:username;unique"`
	Password string `gorm:"column:password"`
}

type CreateUser struct {
	Username string `gorm:"column:username;unique"`
	Password string `gorm:"column:password"`
}

type UserRepository interface {
	CreateUserTable() error
	InsertUserData(User) (*User, error)
	GetAllUser() ([]User, error)
	GetOneUser(string) (*User, error)
}