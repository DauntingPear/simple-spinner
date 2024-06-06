package animation

import (
	"time"
)

type Animation int

const (
	Dots Animation = iota
)

type Property struct {
	frameDuration time.Duration
	chars         []string
}

var lookup = map[Animation]Property{
	Dots: {80, []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}},
}

func GetAnimation(a Animation) (time.Duration, []string) {
	return lookup[a].GetDuration(), lookup[a].GetCharacters()
}

func (p Property) GetDuration() time.Duration {
	return p.frameDuration * time.Millisecond
}

func (p Property) GetCharacters() []string {
	return p.chars
}
