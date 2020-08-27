package tube

import (
	"errors"
	"fmt"
)

const tubeMaxContent = 4

type TubeInfo struct {
	Content []string `json:"content"`
}

type Tube struct {
	index    int
	contents [tubeMaxContent]string
}

func New(info TubeInfo) (Tube, error) {
	tubeSize := len(info.Content)
	if tubeSize > tubeMaxContent {
		return Tube{}, errors.New(fmt.Sprintf("a tube can not be fuller than %d balls", tubeMaxContent))
	}

	var tubeContent [tubeMaxContent]string

	for i, c := range info.Content {
		tubeContent[i] = c
	}

	return Tube{
		index:    tubeSize,
		contents: tubeContent,
	}, nil
}

func (t *Tube) Push(element string) error {
	if t.index == tubeMaxContent {
		return errors.New("tube is already full")
	}

	if t.index != 0 {
		contentBelow := t.contents[t.index-1]
		if contentBelow != element {
			return errors.New("it is only possible to push onto the same content")
		}
	}

	t.contents[t.index] = element
	t.index++
	return nil
}

func (t *Tube) Peek() (string, bool) {
	if t.index == 0 {
		return "", false
	}

	return t.contents[t.index-1], true
}

func (t *Tube) Pop() (string, error) {
	if t.index == 0 {
		return "", errors.New("can't pop empty tube")
	}
	elem := t.contents[t.index-1]
	t.contents[t.index-1] = ""
	t.index--
	return elem, nil
}

func (t *Tube) Len() int {
	return t.index
}

func (t *Tube) IsSolved() bool {
	if t.index != tubeMaxContent && t.index != 0 {
		return false
	}

	var last string
	for i, content := range t.contents {
		if i != 0 && content != last {
			return false
		}

		last = content
	}
	return true
}

func (t *Tube) IsFull() bool {
	return t.index == tubeMaxContent
}

func (t *Tube) Equals(other *Tube) bool {
	if t.index != other.index {
		return false
	}

	for i := 0; i < t.index; i++ {
		if t.contents[i] != other.contents[i] {
			return false
		}
	}
	return true
}

func (t *Tube) Print() {
	fmt.Print("| ")
	for _, c := range t.contents {
		fmt.Print(c + " ")
	}
	fmt.Print("\n")
}
