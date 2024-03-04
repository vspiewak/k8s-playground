package main

import (
	"fmt"
	"os/user"
	"time"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	for {
		fmt.Println("user: " + user.Username + " id: " + user.Uid)
		time.Sleep(1 * time.Second)
	}
}
