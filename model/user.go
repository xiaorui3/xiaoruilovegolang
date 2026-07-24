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
