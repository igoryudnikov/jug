package algorithm

import "fmt"

// Represents current state of the system
type State struct {
	X    int `json:"x"`
	CapX int `json:"-"`
	Y    int `json:"y"`
	CapY int `json:"-"`
	Z    int `json:"-"`
}

func (s State) String() string {
	return fmt.Sprintf("x:%v \t y:%v \t", s.X, s.Y)
}

// Factory method
func NewState(x, y, z int) State {
	return State{
		X:    0,
		CapX: x,
		Y:    0,
		CapY: y,
		Z:    z,
	}
}

// Is invariant for any cases
func (s State) IsValid() bool {
	return (0 <= s.X && s.X <= s.CapX) || (0 <= s.Y && s.Y <= s.CapY)
}
