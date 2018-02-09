package inmem

import (
	"sync"
	"github.com/JacobSoderblom/db_test/pkg/user"
	"encoding/json"
	"fmt"
)

type userRepository struct {
	mtx    sync.RWMutex
	users  map[user.UserID]*user.User
}

func (r *userRepository) Store(u *user.User) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.users[u.ID] = u
	return nil
}

func (r *userRepository) Find(id user.UserID) (*user.User, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if val, ok := r.users[id]; ok {
		return val, nil
	}

	return nil, user.ErrUnknown
}

func (r *userRepository) FindList() (*[]byte, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	list := userList{}

	for _, u := range r.users {
		list.Data = append(list.Data, userListItem{
			ID: int(u.ID),
			Name: u.Name,
			Email: u.Email,
		})
	}

	b, err := json.Marshal(list)
	if err != nil {
		return nil, err
	}

	return &b, nil
}

func NewUserRepository() user.Repository {
	r := &userRepository{
		users: make(map[user.UserID]*user.User),
	}

	r.mtx.Lock()
	defer r.mtx.Unlock()

	for i := 0; i < 1000; i++ {
		u := user.NewUser(
			fmt.Sprintf("user-%v", i),
			fmt.Sprintf("email-%v", i),
			fmt.Sprintf("pass-%v", i),
			user.Address{
				Street: fmt.Sprintf("street-%v", i),
				City: fmt.Sprintf("city-%v", i),
			},
		)
		u.ID = user.UserID(i)

		r.users[u.ID] = u
	}

	return r
}

type userList struct {
	Data []userListItem `json:"data"`
}

type userListItem struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}