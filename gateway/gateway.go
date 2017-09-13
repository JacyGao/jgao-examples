package gateway

import (
	"encoding/json"
)

type Config struct {
	GetUserURL  string
	PostUserURL string
	Token       string
}

type gateway struct {
	c  Config
	st Store
	f  Fetch
}

type Store interface {
	Insert(doc *User) error
	Get(key string, doc interface{}) error
}

type Fetch interface {
	Get(key string, url string, res interface{}) error
	Post(body []byte, url string, res interface{}) error
}

func NewGateway(c Config, st Store, http Fetch) *gateway {
	return &gateway{
		c,
		st,
		http,
	}
}

func (gw *gateway) GetUser(pbu *UserReq) (*User, error) {
	// Get user info from rest endpoint
	user := &User{}
	if err := gw.f.Get(pbu.Id, gw.c.GetUserURL, &user); err != nil {
		return nil, err
	}
	// Store user
	if err := gw.st.Insert(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (gw *gateway) AddUser(key string) (*UserRes, error) {
	// Get user info from store
	user := &User{}
	if err := gw.st.Get(key, &user); err != nil {
		return nil, err
	}
	// Making POST request to add user to remote service
	body, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	res := &UserRes{}
	if err := gw.f.Post(body, gw.c.PostUserURL, &res); err != nil {
		return nil, err
	}
	return res, nil
}
