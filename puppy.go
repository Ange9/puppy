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
	return dog.WhenPuppyGrows(Bark())
}

func BigBarks() string {
	return dog.WhenPuppyGrows(Barks())
}
