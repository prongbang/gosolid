package dip

import (
	"fmt"
)

type User struct {
	ID int
}

type UserRepository interface {
	GetUser(id int) (User, error)
}

type DatabaseUserRepository struct {
	// TODO implement database
}

func (r *DatabaseUserRepository) GetUser(id int) (User, error) {
	if id == 1 {
		return User{ID: 1}, nil
	}
	return User{}, nil
}

type UserService struct {
	repo UserRepository
}

func (s *UserService) RegisterUser(user User) error {
	u, err := s.repo.GetUser(user.ID)
	if u.ID == 0 {
		fmt.Println(" > Registered user id:", user.ID)
		return nil
	}
	return err
}

func Example() {
	fmt.Println("Dependency Inversion Principle")

	usrRepo := &DatabaseUserRepository{}
	usrSvc := &UserService{
		repo: usrRepo,
	}
	usr := User{}
	usrSvc.RegisterUser(usr)
}
