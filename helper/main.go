package helper

import "log"

func HaltOn(err error) {
	if err != nil {
		log.Fatal("Error here: ", err)
	}
}
