package puppy

import (
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
