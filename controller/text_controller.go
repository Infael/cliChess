package controller

import (
	"bufio"
	"cli-chess/model"
	"fmt"
	"log"
	"os"
	"regexp"
)

type TextController struct {
	IsRunning    bool
	engine       Engine
	view         View
	reader       *bufio.Reader
	currentFen   *model.FenState
	recentFens   []*model.FenState
	historyIndex int
}

func NewTextController(view View, engine Engine) *TextController {
	startingFen := model.NewGame()

	return &TextController{
		IsRunning:    false,
		view:         view,
		engine:       engine,
		reader:       bufio.NewReader(os.Stdin),
		currentFen:   startingFen,
		recentFens:   []*model.FenState{startingFen},
		historyIndex: 0,
	}
}

func (tc *TextController) Start() {
	tc.IsRunning = true
	tc.view.SetBoard(*tc.currentFen)
}

func (tc *TextController) HandleInput() {
	// read input from user
	fmt.Print("Player: ")
	nextMove, _, err := tc.reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}

	switch string(nextMove) {
	case "q":
		tc.stop()
	case "hint":
		tc.hint()
	case "b":
		tc.backMove()
	case "n":
		tc.nextMove()
	case "new":
		tc.newGame()
	default:
		if tc.historyIndex == 0 && string(nextMove) == "" {
			// let the engine make the first move a.k.a play as white
			tc.move(string(nextMove))
			return
		}
		regexp := regexp.MustCompile(`^[a-h][1-8][a-h][1-8]$`)
		if !regexp.MatchString(string(nextMove)) {
			fmt.Println("Invalid move")
			return
		}
		tc.move(string(nextMove))
	}

}

func (tc *TextController) move(nextMove string) {
	newFen, _, err := tc.engine.Move(*tc.currentFen, nextMove)
	if err != nil {
		log.Fatal(err)
	}
	tc.currentFen = model.NewFenState(newFen)

	// update history
	if tc.historyIndex < len(tc.recentFens)-1 {
		tc.recentFens = tc.recentFens[:tc.historyIndex+1]
	}
	tc.recentFens = append(tc.recentFens, tc.currentFen)
	tc.historyIndex++

	tc.view.SetBoard(*tc.currentFen)

	evaluation, err := tc.engine.Evaluate()
	if err != nil {
		log.Fatal(err)
	}
	tc.view.ShowEvaluation(evaluation)
}

func (tc *TextController) stop() {
	tc.IsRunning = false
}

func (tc *TextController) hint() {
	fmt.Println("Hint: " + tc.engine.SuggestNextMove(*tc.currentFen))
}

func (tc *TextController) backMove() {
	if tc.historyIndex == 0 {
		return
	}

	tc.historyIndex--
	tc.currentFen = tc.recentFens[tc.historyIndex]
	tc.view.SetBoard(*tc.currentFen)
}

func (tc *TextController) nextMove() {
	if tc.historyIndex == len(tc.recentFens)-1 {
		return
	}

	tc.historyIndex++
	tc.currentFen = tc.recentFens[tc.historyIndex]
	tc.view.SetBoard(*tc.currentFen)
}

func (tc *TextController) newGame() {
	tc.engine.NewGame()
	tc.currentFen = model.NewGame()
	tc.view.SetBoard(*tc.currentFen)
	tc.recentFens = nil
}
