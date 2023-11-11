package types

import (
	"time"
)

type Person struct {
	ID        int        `json:"id"`
	FirstName string     `json:"firstname"`
	LastName  string     `json:"lastname"`
	UserName  string     `json:"username"`
	Verified  bool       `json:"verified"`
	Joined    *time.Time `json:"joined"`
}
