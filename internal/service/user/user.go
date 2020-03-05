package skeleton

import (
	"context"
	"log"
	"strconv"
	"strings"

	userEntity "go-tutorial-2020/internal/entity/user"
	"go-tutorial-2020/pkg/errors"
)

// UserData ...
type UserData interface {
	GetAllUsers(ctx context.Context) ([]userEntity.User, error)
	InsertUser(ctx context.Context, user userEntity.User) error
	GetUserByName(ctx context.Context, userNama string) (userEntity.User, error)
	UpdateUser(ctx context.Context, user userEntity.User) error
	GetMaxNIP(ctx context.Context) (int, error)
	DeleteByNIP(ctx context.Context, nip string) error
	SelectUser(ctx context.Context) ([]userEntity.User, error)
}

// Service ...
type Service struct {
	userData UserData
}

// New ...
func New(userData UserData) Service {
	return Service{
		userData: userData,
	}
}

// GetAllUsers ...
func (s Service) GetAllUsers(ctx context.Context) ([]userEntity.User, error) {
	// Panggil method GetAllUsers di data layer user
	users, err := s.userData.GetAllUsers(ctx)
	// Error handling
	if err != nil {
		return users, errors.Wrap(err, "[SERVICE][GetAllUsers]")
	}
	// Return users array
	return users, err
}

// InsertUser ...
func (s Service) InsertUser(ctx context.Context, user userEntity.User) error {
	var (
		userValidasi userEntity.User
		err          error
		maxNip       int
	)
	userValidasi, err = s.userData.GetUserByName(ctx, user.Nama)
	if strings.Contains(userValidasi.Nama, user.Nama) {
		return errors.Wrap(errors.New("data already exists"), "[SERVICE][InsertUser]")
	}
	maxNip, err = s.userData.GetMaxNIP(ctx)
	user.Nip = "P" + strconv.Itoa(maxNip+1)
	log.Println(user.Nip)
	log.Println(maxNip)
	err = s.userData.InsertUser(ctx, user)
	return err
}

//UpdateUser ...
func (s Service) UpdateUser(ctx context.Context, user userEntity.User) error {
	var (
		//userValidasi userEntity.User
		err error
	)
	// userValidasi, err = s.userData.GetUserByName(ctx, user.Nama)
	// if userValidasi.Nama != user.Nama {
	// 	return errors.Wrap(errors.New("data already exists"), "[SERVICE][InsertUser]")
	// }
	err = s.userData.UpdateUser(ctx, user)
	return err
}

// GetUserByName ...
func (s Service) GetUserByName(ctx context.Context, userNama string) (userEntity.User, error) {
	result, err := s.userData.GetUserByName(ctx, userNama)
	return result, err
}

//DeleteByNIP ...
func (s Service) DeleteByNIP(ctx context.Context, nip string) error {
	err := s.userData.DeleteByNIP(ctx, nip)
	return err
}

func (s Service) SelectUser(ctx context.Context) ([]userEntity.User, error) {
	result, err := s.userData.SelectUser(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return result, err
}
