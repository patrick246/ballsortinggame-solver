package tube

import (
	"github.com/patrick246/ballsortinggame-solver/game"
	"testing"
)

func TestTube_PopPushEmpty(t *testing.T) {
	tube, err := New(game.TubeInfo{Content: []string{}})
	if err != nil {
		t.Fatal(err)
	}

	err = tube.Push("a")
	if err != nil {
		t.Fatal(err)
	}

	elem, err := tube.Pop()
	if err != nil {
		t.Fatal(err)
	}

	if elem != "a" {
		t.Fatal("elem != a. elem: ", elem)
	}

	length := tube.Len()
	if length != 0 {
		t.Fatal("len != 0. len: ", length)
	}
}

func TestTube_PushPopAlreadyFull(t *testing.T) {
	tube, err := New(game.TubeInfo{Content: []string{"a", "a"}})
	if err != nil {
		t.Fatal(err)
	}

	err = tube.Push("a")
	if err != nil {
		t.Fatal(err)
	}

	length := tube.Len()
	if length != 3 {
		t.Fatal("length != 3. length", length)
	}

	elem, err := tube.Pop()
	if err != nil {
		t.Fatal(err)
	}

	if elem != "a" {
		t.Fatal("elem != a. elem: ", elem)
	}

	length = tube.Len()
	if length != 2 {
		t.Fatal("length != 0. length: ", length)
	}
}

func TestTube_PopEmpty(t *testing.T) {
	tube, err := New(game.TubeInfo{Content: []string{}})
	if err != nil {
		t.Fatal(err)
	}

	_, err = tube.Pop()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestTube_PushFull(t *testing.T) {
	tube, err := New(game.TubeInfo{Content: []string{"a", "a", "a", "a"}})
	if err != nil {
		t.Fatal(err)
	}

	err = tube.Push("a")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestTube_PushNotMatching(t *testing.T) {
	tube, err := New(game.TubeInfo{Content: []string{"a"}})
	if err != nil {
		t.Fatal(err)
	}

	err = tube.Push("b")
	if err == nil {
		t.Fatal("expected err, got nil")
	}
}
