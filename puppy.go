package puppy

import (
	"fmt"

	"github.com/thejaswa/dog"
)

func Bark() string {
	return "Bow!"
}

func Braks() string {
	return "Bow! Bow! Bow!"
}

func BigBark() string {
	//return "BowBOW"
	return dog.BigDogBark(Bark())
}
func BigBarks() string {
	//return "BowBOW"
	return dog.BigDogBark(Braks())
}
func Tagging() {
	fmt.Println("im from v1.0.0")
}



