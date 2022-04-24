package request

import "errors"

type Registration struct {
	Id  string `json:"id"`
	Name     string `json:"name"`
	Password  string `json:"password"`
}

func (registration *Registration) Validate() error {
	if registration.Id == "" {
		return errors.New("empty id")
	}
	if registration.Name == "" {
		return errors.New("empty name")
	}
	if registration.Password == "" {
		return errors.New("empty password")
	}
	return nil
}
