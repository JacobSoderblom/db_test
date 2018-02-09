package user

import "encoding/json"

type Service interface {
	Create(name, email, password string, addr Address) (*User, error)
	Save(id UserID, name, email string, addr Address) (*User, error)
	Find(id UserID) (*User, error)
	All() ([]*userListItem, error)
}

func NewService(users Repository) Service {
	return &service{users}
}

type service struct {
	users Repository
}

func (s *service) Create(name, email, password string, addr Address) (*User, error) {
	u := NewUser(name, email, password, addr)

	if err := s.users.Store(u); err != nil {
		return nil, err
	}

	return u, nil
}

func (s *service) Save(id UserID, name, email string, addr Address) (*User, error) {
	u, err := s.users.Find(id)
	if err != nil {
		return nil, err
	}

	u.SetName(name)
	u.SetEmail(email)
	u.SetAddress(addr)

	if err := s.users.Store(u); err != nil {
		return nil, err
	}

	return u, nil
}

func (s *service) Find(id UserID) (*User, error) {
	return s.users.Find(id)
}

func (s *service) All() ([]*userListItem, error) {
	list := []*userListItem{}
	var listData struct {
		Data []userListItem `json:"data"`
	}

	b, err := s.users.FindList()
	if err != nil {
		return list, err
	}

	if err := json.Unmarshal(*b, &listData); err != nil {
		return list, err
	}

	for _, i := range listData.Data {
		list = append(list, &i)
	}

	return list, nil
}

type userListItem struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}