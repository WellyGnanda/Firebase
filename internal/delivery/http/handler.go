package http

import "github.com/gorilla/mux"

// Handler will initialize mux router and register handler
func (s *Server) Handler() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", s.User.UserHandler).Methods("GET")
	r.HandleFunc("/insertUser", s.User.UserHandler).Methods("POST")
	r.HandleFunc("/updateUser", s.User.UserHandler).Methods("PUT")
	r.HandleFunc("/deleteUser", s.User.UserHandler).Methods("DELETE")
	return r
}
