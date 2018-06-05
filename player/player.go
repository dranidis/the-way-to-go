package main

import (
	"fmt"
)

// INTER OMIT
type Player interface {
	Play()
}
// END OMIT

// STRUCT OMIT
type MP3Player struct {
	Song string
}
// END OMIT

// METHOD OMIT
func (m MP3Player) Play() {
	fmt.Println("Now playing song: " + m.Song)
}
// END OMIT

// CLIENT OMIT
func start(p Player) {
	p.Play()
}

func main() {
	m := MP3Player{"We will rock you!"}
	start(m)
}
// END OMIT