package config

import (
	"fmt"
)

func (C Config) SetUser(user string) {
	C.CurrentUserName = user
	err := write(C)
	if err != nil {
		fmt.Println(err)
	}
}
