package solver

import (
	"github.com/patrick246/ballsortinggame-solver/game"
	"github.com/patrick246/ballsortinggame-solver/solver/stack"
	"log"
	"time"
)

type Solver struct {
	initialState      game.LevelInfo
	levelStateChannel chan game.LevelState
	stateStack        *stack.LevelStateStack
	resultChannel     chan []game.Move
}

func New(initialState game.LevelInfo) *Solver {
	return &Solver{
		initialState:      initialState,
		stateStack:        stack.New(),
		levelStateChannel: make(chan game.LevelState, 1),
		resultChannel:     make(chan []game.Move),
	}
}

func (s *Solver) Run() []game.Move {
	state, err := game.NewGame(s.initialState)
	if err != nil {
		return nil
	}

	s.stateStack.Push(*state)

	done := make(chan struct{})
	for i := 0; i < 4; i++ {
		go solverThread(s.stateStack, s.resultChannel, done)
	}

	result := <-s.resultChannel
	close(done)
	return result
}

func solverThread(stack *stack.LevelStateStack, result chan []game.Move, done chan struct{}) {
	for {
		select {
		case _, ok := <-done:
			if !ok {
				return
			}
		default:

		}

		var state game.LevelState
		ok := false
		for !ok {
			state, ok = stack.Pop()
			if !ok {
				time.Sleep(100 * time.Millisecond)
			}
		}

		state.Print()

		moves := state.GenerateMoves()
		for _, move := range moves {
			possible := state.IsMovePossible(move.From, move.To)
			if !possible {
				continue
			}

			nextState, err := state.DoMove(move.From, move.To)
			if err != nil {
				log.Printf("error while doing move: %v", err)
				continue
			}

			if nextState.IsLoop() {
				continue
			}

			if nextState.IsSolved() {
				result <- nextState.Moves
				return
			}

			stack.Push(*nextState)
		}
	}
}
