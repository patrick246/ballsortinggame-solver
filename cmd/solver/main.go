package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/patrick246/ballsortinggame-solver/game"
	solver2 "github.com/patrick246/ballsortinggame-solver/solver"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var file = flag.String("input-file", "level.json", "Path to the file with the level definition for the solver")

func main() {
	flag.Parse()
	start := time.Now()

	jsonFile, err := os.Open(*file)
	if err != nil {
		log.Fatalf("error while opening file: %v", err)
	}

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("error while readin file: %v", err)
	}

	var levelInfo game.LevelInfo
	err = json.Unmarshal(bytes, &levelInfo)

	solver := solver2.New(levelInfo)
	moves := solver.Run()

	for _, m := range moves {
		fmt.Printf("%d -> %d\n", m.From+1, m.To+1)
	}
	fmt.Printf("Took %v", time.Now().Sub(start))
}
