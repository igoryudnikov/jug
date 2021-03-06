package algorithm

import (
	"fmt"
	"sort"
	"strings"
)

// Represent name of operation and it's result
type StateWithOperationName struct {
	Name  string `json:"name"`
	State State  `json:"state"`
}

// Keeps current state and all previous states with applied state change's human readable names
type SolutionCandidate struct {
	State        State                    `json:"-"`
	StateChanges []StateWithOperationName `json:"state_changes"`
	Length       int                      `json:"length"`
}

func (sc SolutionCandidate) String() string {
	var ss []string
	for _, stateChange := range sc.StateChanges {
		ss = append(ss, fmt.Sprintf("%v\t%v\n", stateChange.State.String(), stateChange.Name))
	}
	return strings.Join(ss, "")
}

// Applies new state with human readable state change name to solution candidate, increments its length
func (sc SolutionCandidate) Updated(state State, name string) SolutionCandidate {
	sc.State = state
	sc.StateChanges = append(sc.StateChanges, StateWithOperationName{
		Name:  name,
		State: state,
	})
	sc.Length = len(sc.StateChanges)
	return sc
}

type SolveInput struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

// Entry point with default solution predicate
func Solve(in SolveInput) []SolutionCandidate {
	initState := NewState(in.X, in.Y, in.Z)
	anyOfJugsEqualsZ := func(state State) bool {
		return state.X == state.Z || state.Y == state.Z
	}
	return SolveWithPredicate(initState, anyOfJugsEqualsZ)
}

// Does the job for initial state and injectable solution predicate
func SolveWithPredicate(initState State, solutionPredicate func(State) bool) []SolutionCandidate {

	possibleOperations := GetPossibleOperations()
	states := make(map[State]struct{})
	states[initState] = struct{}{}

	solutionCandidates := []SolutionCandidate{
		{
			State:        initState,
			StateChanges: []StateWithOperationName{},
		},
	}

	var solutions []SolutionCandidate

	for {
		var currentCandidates []SolutionCandidate
		for _, solutionCandidate := range solutionCandidates {
			for _, operation := range possibleOperations {
				newState := operation.StateChange(solutionCandidate.State)
				_, stateExists := states[newState]
				if solutionPredicate(newState) {
					solutions = append(solutions, solutionCandidate.Updated(newState, operation.Name))
				} else if !stateExists && newState.IsValid() {
					states[newState] = struct{}{}
					currentCandidates = append(currentCandidates, solutionCandidate.Updated(newState, operation.Name))
				}
			}
		}
		if len(currentCandidates) == 0 {
			return solutions
		}
		solutionCandidates = currentCandidates
	}

}

func ShortestSolutionIfExists(solutions []SolutionCandidate) *SolutionCandidate {
	if len(solutions) == 0 {
		return nil
	}
	sort.SliceStable(solutions[:], func(i, j int) bool {
		return solutions[i].Length < solutions[j].Length
	})
	return &solutions[0]
}
