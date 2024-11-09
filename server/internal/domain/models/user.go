package models

import (
	"errors"
	"regexp"
)

type User struct {
	ID        int
	Username  string
	Email     string
	Password  string
	CreatedAt string
}

func (self *User) Validate() (err []error) {
  if len(self.Username) > 50 {
    err = append(err, errors.New("username must contain less than 50 characters"))
  }

  if len(self.Username) < 4 {
    err = append(err, errors.New("username must contain at least 4 characters"))
  }

  if self.ValidateEmail() == false {
    err = append(err, errors.New("email is invalid"))
  }

  if len(self.Password) > 24 {
    err = append(err, errors.New("password must contain less than 24 characters"))
  }

  if len(self.Password) < 4 {
    err = append(err, errors.New("password must contain at least 4 characters"))
  }

  return
}

func (self *User) ValidateEmail() bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(self.Email)
}
