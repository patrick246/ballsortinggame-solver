package game

import (
	"errors"
	"fmt"
	"github.com/patrick246/ballsortinggame-solver/game/tube"
)

type LevelInfo struct {
	Tubes []tube.TubeInfo `json:"tubes"`
}

type Move struct {
	From int
	To   int
}

type LevelState struct {
	TubeHistory [][]tube.Tube
	TubeState   []tube.Tube
	Moves       []Move
}

func NewGame(initialState LevelInfo) (*LevelState, error) {
	var tubes []tube.Tube
	for _, info := range initialState.Tubes {
		t, err := tube.New(info)
		if err != nil {
			return nil, err
		}
		tubes = append(tubes, t)
	}

	return &LevelState{
		TubeHistory: [][]tube.Tube{},
		TubeState:   tubes,
		Moves:       []Move{},
	}, nil
}

func (ls *LevelState) GenerateMoves() []Move {
	var moves []Move

	for from := 0; from < len(ls.TubeState); from++ {
		for to := 0; to < len(ls.TubeState); to++ {
			if !ls.IsMovePossible(from, to) {
				continue
			}

			moves = append(moves, Move{
				From: from,
				To:   to,
			})
		}
	}
	return moves
}

func (ls *LevelState) IsMovePossible(from, to int) bool {
	if from == to {
		return false
	}

	if ls.TubeState[to].IsFull() {
		return false
	}

	topFrom, fromHasContent := ls.TubeState[from].Peek()
	if !fromHasContent {
		return false
	}

	topTo, toHasContent := ls.TubeState[to].Peek()
	if toHasContent && topTo != topFrom {
		return false
	}
	return true
}

func (ls *LevelState) DoMove(from, to int) (*LevelState, error) {
	if !ls.IsMovePossible(from, to) {
		return nil, errors.New("move not possible")
	}

	nextLevelHistory := make([][]tube.Tube, len(ls.TubeHistory))
	for i := range ls.TubeHistory {
		nextLevelHistory[i] = make([]tube.Tube, len(ls.TubeHistory[i]))
		copy(nextLevelHistory[i], ls.TubeHistory[i])
	}

	nextLevelState := LevelState{
		TubeHistory: nextLevelHistory,
		TubeState:   make([]tube.Tube, len(ls.TubeState)),
		Moves:       make([]Move, len(ls.Moves), len(ls.Moves)+1),
	}

	copy(nextLevelState.TubeState, ls.TubeState)
	copy(nextLevelState.Moves, ls.Moves)
	nextLevelState.TubeHistory = append(nextLevelState.TubeHistory, ls.TubeState)

	elem, err := nextLevelState.TubeState[from].Pop()
	if err != nil {
		return nil, err
	}

	err = nextLevelState.TubeState[to].Push(elem)
	if err != nil {
		return nil, err
	}

	nextLevelState.Moves = append(nextLevelState.Moves, Move{
		From: from,
		To:   to,
	})
	return &nextLevelState, nil
}

func (ls *LevelState) IsSolved() bool {
	for _, t := range ls.TubeState {
		if !t.IsSolved() {
			return false
		}
	}
	return true
}

func TubeStateEquals(ts1, ts2 []tube.Tube) bool {
	if len(ts1) != len(ts2) {
		return false
	}

	for i := 0; i < len(ts1); i++ {
		if !ts1[i].Equals(&ts2[i]) {
			return false
		}
	}
	return true
}

func (ls *LevelState) IsLoop() bool {
	for _, history := range ls.TubeHistory {
		if TubeStateEquals(history, ls.TubeState) {
			return true
		}
	}
	return false
}

func (ls *LevelState) Print() {
	fmt.Println("Level State:")
	for _, t := range ls.TubeState {
		t.Print()
	}
}
