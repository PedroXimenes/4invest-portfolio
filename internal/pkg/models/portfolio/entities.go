package portfolio

import "fmt"

type Portfolio struct {
	ID     int64
	UserID int64  `json:"user_id"`
	Name   string `json:"name"`
}

func (p *Portfolio) ValidateInput() (key string, err error) {
	if p.Name == "" {
		err = fmt.Errorf("Missing key: email")
		key = "email"
	} else if p.UserID == 0 {
		err = fmt.Errorf("Missing or invalid key: userID")
		key = "userID"
	}
	return
}
