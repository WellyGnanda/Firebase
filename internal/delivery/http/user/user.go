package user

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	userEntity "go-tutorial-2020/internal/entity/user"
	"go-tutorial-2020/pkg/response"
)

// IUserSvc is an interface to User Service
type IUserSvc interface {
	GetAllUsers(ctx context.Context) ([]userEntity.User, error)
	InsertUser(ctx context.Context, user userEntity.User) error
	UpdateUser(ctx context.Context, user userEntity.User) error
	DeleteByNIP(ctx context.Context, nip string) error
	SelectUser(ctx context.Context) ([]userEntity.User, error)
}

type (
	// Handler ...
	Handler struct {
		userSvc IUserSvc
	}
)

// New for user domain handler initialization
func New(is IUserSvc) *Handler {
	return &Handler{
		userSvc: is,
	}
}

// UserHandler will return user data
func (h *Handler) UserHandler(w http.ResponseWriter, r *http.Request) {
	var (
		resp     *response.Response
		metadata interface{}
		result   interface{}
		err      error
		errRes   response.Error
		user     userEntity.User
		_type    string
		//sua      []userEntity.User
	)
	// Make new response object
	resp = &response.Response{}
	body, _ := ioutil.ReadAll(r.Body)
	// Defer will be run at the end after method finishes
	defer resp.RenderJSON(w, r)

	switch r.Method {
	// Check if request method is GET
	case http.MethodGet:
		if _, x := r.URL.Query()["type"]; x {
			_type = r.FormValue("type")
			if _type == "getUserFirebase" {
				result, err = h.userSvc.SelectUser(context.Background())
			} else if _type == "getUserMySQL" {
				// Ambil semua data user
				result, err = h.userSvc.GetAllUsers(context.Background())

			}
		}

	case http.MethodPost:
		json.Unmarshal(body, &user)
		err = h.userSvc.InsertUser(context.Background(), user)

	case http.MethodPut:
		json.Unmarshal(body, &user)
		err = h.userSvc.UpdateUser(context.Background(), user)
	case http.MethodDelete:
		if _, nipOK := r.URL.Query()["NIP"]; nipOK {
			err = h.userSvc.DeleteByNIP(context.Background(), r.FormValue("NIP"))
		} else {
			err = errors.New("400")
		}
	default:
		err = errors.New("400")
	}

	// If anything from service or data return an error
	if err != nil {
		// Error response handling
		errRes = response.Error{
			Code:   101,
			Msg:    "Data Not Found",
			Status: true,
		}
		// If service returns an error
		if strings.Contains(err.Error(), "service") {
			// Replace error with server error
			errRes = response.Error{
				Code:   201,
				Msg:    "Failed to process request due to server error",
				Status: true,
			}
		}

		// Logging
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		resp.Error = errRes
		return
	}

	// Inserting data to response
	resp.Data = result
	resp.Metadata = metadata
	// Logging
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
}
