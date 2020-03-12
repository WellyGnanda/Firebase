package skeleton

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"

	userEntity "go-tutorial-2020/internal/entity/user"
	"go-tutorial-2020/pkg/errors"
	"go-tutorial-2020/pkg/kafka"
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
	InsertFirebase(ctx context.Context, user userEntity.User) error
	UpdateByNipFirebase(ctx context.Context, nip string, user userEntity.User) error
	GetUserClient(ctx context.Context, header http.Header) ([]userEntity.User, error)
	InsertUserClient(ctx context.Context, headers http.Header, user userEntity.User) error
	DeleteUserByNipFirebase(ctx context.Context, nip string) error
	GetUserPage(ctx context.Context, page int, length int) ([]userEntity.User, error)
}

// Service ...
type Service struct {
	userData UserData
	kafka    *kafka.Kafka
}

// New ...
func New(userData UserData, kafka *kafka.Kafka) Service {
	return Service{
		userData: userData,
		kafka:    kafka,
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

//SelectUser ...
func (s Service) SelectUser(ctx context.Context) ([]userEntity.User, error) {
	result, err := s.userData.SelectUser(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return result, err
}

//InsertFirebase ...
func (s Service) InsertFirebase(ctx context.Context, user userEntity.User) error {

	err := s.userData.InsertFirebase(ctx, user)
	return err
}

//PublishFirebase ...
func (s Service) PublishFirebase(user userEntity.User) error {
	err := s.kafka.SendMessageJSON("New_User", user)
	if err != nil {
		return errors.Wrap(errors.New("PUBLISH GAGAL"), "[SERVICE][PublishUser]")
	}
	return err
}

//UpdateByNipFirebase ...
func (s Service) UpdateByNipFirebase(ctx context.Context, nip string, user userEntity.User) error {
	err := s.userData.UpdateByNipFirebase(ctx, nip, user)
	return err
}

// DeleteUserByNipFirebase ...
func (s Service) DeleteUserByNipFirebase(ctx context.Context, nip string) error {
	err := s.userData.DeleteUserByNipFirebase(ctx, nip)
	return err
}

// GetUserClient ...
func (s Service) GetUserClient(ctx context.Context, headers http.Header) ([]userEntity.User, error) {
	UserClient, err := s.userData.GetUserClient(ctx, headers)
	return UserClient, err
}

// InsertUserClient ...
func (s Service) InsertUserClient(ctx context.Context, headers http.Header, user userEntity.User) error {
	err := s.userData.InsertUserClient(ctx, headers, user)
	return err
}

// GetUserPage ...
func (s Service) GetUserPage(ctx context.Context, page int, length int) ([]userEntity.User, error) {
	userList, err := s.userData.GetUserPage(ctx, page, length)
	return userList, err
}
