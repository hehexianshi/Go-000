package main

type UserService struct {

}

func (user *UserService) find(id string) (*User, error) {
	dao := UserDao{}
	return dao.findOne(id)

}

func (user *UserService) findAll() []*User {
	dao := UserDao{}
	all, _ := dao.findAll()
	return all
}
