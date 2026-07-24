package model

import "awesomeProject/dao"

type User struct {
	ID   int
	Name string
}

func (User) TypeName() string {
	return "user"
}
func (u User) GetUserTest(id int) (User, error) {
	var jjj User
	err := dao.Db.Table("user").Where("id = ?", id).First(&jjj).Error
	return jjj, err
}

func (u User) GetUserTestToDataBaseTable(id int) ([]User, error) {
	if id == -1 {
		var jjj []User
		err := dao.Db.Raw("select id, name from user").Scan(&jjj).Error
		return jjj, err
	}
	var jjj []User
	err := dao.Db.Raw("select id, name from user where id = ?", id).Scan(&jjj).Error
	return jjj, err
}
