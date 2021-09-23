package creditcard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type UserHandler struct{}

func (u *UserHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw, err)
		return
	}
	user.CreatedAt = time.Now()
	data, _ := json.Marshal(user)
	rw.Header().Add("content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	fmt.Fprint(rw, string(data))
}
