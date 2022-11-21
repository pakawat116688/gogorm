package repository

import "gorm.io/gorm"

type UserDB struct {
	userdb *gorm.DB
}

func NewUserDB(userdb *gorm.DB) UserRepository  {
	return UserDB{userdb: userdb}
}

func (u User) TableName() string  {
	return "User"
}

func (r UserDB) CreateUserTable() error {

	query := `CREATE TABLE "User" (
		"uid"	INTEGER,
		"username"	text UNIQUE,
		"password"	text,
		PRIMARY KEY("uid" AUTOINCREMENT)
	);`
	tx := r.userdb.Raw(query).Scan(&User{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r UserDB) InsertUserData(user User) (*User, error) {
	tx := r.userdb.Create(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	userResponse := User{}
	tx = r.userdb.Last(&userResponse)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &userResponse, nil
}

func (r UserDB) GetAllUser() ([]User, error)  {
	userResponse := []User{}
	tx := r.userdb.Order("uid").Find(&userResponse)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return userResponse, nil
}

func (r UserDB) GetOneUser(uname string) (*User, error)  {
	userResponse := User{}
	tx := r.userdb.Where("username=?", uname).Find(&userResponse)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &userResponse, nil
}