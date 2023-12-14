package puppy

import (
	"github.com/Ange9/dog"
)

func Bark() string {
	return "Boof!"
}
func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.BigBark(Bark())
}

func BigBarks() string {
	return dog.BigBarks(Barks())
}
