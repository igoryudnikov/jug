package test

import (
	"github.com/stretchr/testify/require"
	"jug/algorithm"
	"testing"
)

func TestSolve(t *testing.T) {

	t.Run("Solve should find 2 solutions with 4 and 6 length for x=10 y=2 z=4", func(t *testing.T) {
		solutions := algorithm.Solve(algorithm.SolveInput{
			X: 10,
			Y: 2,
			Z: 4,
		})
		require.Len(t, solutions, 2)
		require.Equal(t, solutions[0].Length, 4)
		require.Equal(t, solutions[1].Length, 6)
	})

	t.Run("it should find 2 solutions with n and m length for x=4 y=3 z=2", func(t *testing.T) {
		solutions := algorithm.Solve(algorithm.SolveInput{
			X: 4,
			Y: 3,
			Z: 2,
		})
		require.Len(t, solutions, 2)
		require.Equal(t, solutions[0].Length, 4)
		require.Equal(t, solutions[1].Length, 6)
	})

	t.Run("it should not find solutions for x=6 y=3 z=1", func(t *testing.T) {
		solutions := algorithm.Solve(algorithm.SolveInput{
			X: 6,
			Y: 3,
			Z: 1,
		})
		require.Len(t, solutions, 0)
	})

	t.Run("it should find solutions for x=100000 y=99998 z=2", func(t *testing.T) {
		solutions := algorithm.Solve(algorithm.SolveInput{
			X: 1000000,
			Y: 999998,
			Z: 2,
		})
		require.Len(t, solutions, 2)
	})

	t.Run("ShortestSolutionIfExists should return solution with shortest length", func(t *testing.T) {
		s1 := algorithm.SolutionCandidate{
			Length: 1,
		}
		s3 := algorithm.SolutionCandidate{
			Length: 3,
		}
		s5 := algorithm.SolutionCandidate{
			Length: 5,
		}
		result := algorithm.ShortestSolutionIfExists([]algorithm.SolutionCandidate{s1, s3, s5})
		require.Equal(t, result.Length, 1)
	})

	t.Run("it should return nil on empty array", func(t *testing.T) {
		result := algorithm.ShortestSolutionIfExists([]algorithm.SolutionCandidate{})
		require.Nil(t, result)
	})

}
