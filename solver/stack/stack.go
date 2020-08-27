package stack

import (
	"github.com/patrick246/ballsortinggame-solver/game"
	"sync"
)

type LevelStateStack struct {
	content []game.LevelState
	mutex   sync.Mutex
}

func New() *LevelStateStack {
	return &LevelStateStack{
		content: nil,
	}
}

func (stack *LevelStateStack) Push(elem game.LevelState) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	stack.content = append(stack.content, elem)
}

func (stack *LevelStateStack) Pop() (game.LevelState, bool) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	if len(stack.content) == 0 {
		return game.LevelState{}, false
	}

	elem := stack.content[len(stack.content)-1]
	stack.content = stack.content[0 : len(stack.content)-1]
	return elem, true
}
