package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func main() {

	service := UserService{}

	find, err := service.find("abc")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println(&User{"v", 10})
			return
		}
	}

	fmt.Println(find)

	all := service.findAll()
	for s := range all {
		log.Println(s)
	}

}