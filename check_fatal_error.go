package check_error

import (
	"log"
)

func Check_fatal_error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
