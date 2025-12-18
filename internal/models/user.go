package models

import "time"

type User struct {
	ID   int       `json:"id"`
	Name string    `json:"name"`
	Dob  time.Time `json:"dob"` 
	Age  int       `json:"age,omitempty"`
}
