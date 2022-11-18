package repository

import "gorm.io/gorm"

type UserDB struct {
	userdb *gorm.DB
}

func NewUserDB(userdb *gorm.DB) UserRepository  {
	return UserDB{userdb: userdb}
}

func (r UserDB) CreateUserTable() error {
	err := r.userdb.AutoMigrate(User{})
	if err != nil {
		return err
	}
	return nil
}

func (r UserDB) InsertUserData(user CreateUser) (*User, error) {
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
	tx := r.userdb.Find(&userResponse)
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