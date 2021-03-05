package algorithm

// Represents state transformation
type StateChange func(State) State

// Represents state transformation with human readable name
type Operation struct {
	Name        string
	StateChange StateChange
}

// Get all possible state transformations
func GetPossibleOperations() []Operation {

	return []Operation{
		{
			Name: "Empty X",
			StateChange: func(state State) State {
				state.X = 0
				return state
			},
		},
		{
			Name: "Empty Y",
			StateChange: func(state State) State {
				state.Y = 0
				return state
			},
		},
		{
			Name: "Fill X",
			StateChange: func(state State) State {
				state.X = state.CapX
				return state
			},
		},
		{
			Name: "Fill Y",
			StateChange: func(state State) State {
				state.Y = state.CapY
				return state
			},
		},
		{
			Name: "Transfer Y to X",
			StateChange: func(state State) State {
				freeX := state.CapX - state.X
				if freeX >= state.Y {
					state.X = state.X + state.Y
					state.Y = 0
				} else {
					restY := state.Y - freeX
					state.X = state.CapX
					state.Y = restY
				}
				return state
			},
		},
		{
			Name: "Transfer X to Y",
			StateChange: func(state State) State {
				freeY := state.CapY - state.Y
				if freeY >= state.X {
					state.Y = state.Y + state.X
					state.X = 0
				} else {
					restX := state.X - freeY
					state.Y = state.CapY
					state.X = restX
				}
				return state
			},
		},
	}

}
