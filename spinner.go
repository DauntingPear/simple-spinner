package spinner

import (
	"fmt"
	"time"
)

type spinner struct {
	message       string
	chars         []string
	frame         int
	frameDuration time.Duration
	done          chan bool
	ticks         time.Ticker
}

type Spinner interface {
	Start()
	Stop()
}

func NewSpinner(message string, a Animation) *spinner {
	duration, chars := GetAnimation(a)
	return &spinner{
		message:       message,
		chars:         chars,
		frame:         0,
		frameDuration: duration,
		done:          make(chan bool),
	}
}

func (s *spinner) Start() {
	s.ticks = *time.NewTicker(s.frameDuration)
	go s.render()
}

func (s *spinner) Stop() {
	s.done <- true
	s.ticks.Stop()
	show()
}

func (s *spinner) render() {

	hide()
	fmt.Print(s.message)

outer:
	for {
		select {
		case <-s.done:
			fmt.Print("\r")
			break outer
		case <-s.ticks.C:
			s.renderFrame(true)
		}
	}
}

func (s *spinner) renderFrame(animate bool) {
	fmt.Printf("%s%c", s.chars[s.frame], 8)

	if animate {
		s.setNextFrame()
	}
}

func (s *spinner) setNextFrame() {
	s.frame++
	if s.frame >= len(s.chars) {
		s.frame = 0
	}
}

func show() {
	fmt.Print("\x1b[?25h")
}

func hide() {
	fmt.Print("\x1b[?25l")
}
