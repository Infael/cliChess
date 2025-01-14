package engineconnector

import (
	"bufio"
	"cli-chess/model"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

const searchDepthCommand = "go movetime 1000"

// const searchDepthCommand = "go depth 4"

type EngineConnector struct {
	cmd    *exec.Cmd
	reader *bufio.Reader
	writer *bufio.Writer
}

func NewEngineConnector(engineCommand string) *EngineConnector {
	cmd := exec.Command(engineCommand)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	return &EngineConnector{
		cmd:    cmd,
		reader: bufio.NewReader(stdout),
		writer: bufio.NewWriter(stdin),
	}
}

func (ec *EngineConnector) Start() error {
	if err := ec.cmd.Start(); err != nil {
		return err
	}
	fmt.Println("Engine started")

	// ec.sendCommand("isready")

	// line, _, err := ec.reader.ReadLine()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if string(line) != "readyok" {
	// 	log.Fatal("Engine did not respond with readyok")
	// }

	return nil
}

func (ec *EngineConnector) NewGame() {
	ec.sendCommand("ucinewgame")
}

func (ec *EngineConnector) Move(fen model.FenState, move string) (string, string, error) {
	ec.sendCommand("position fen " + fen.GetFen() + " moves " + move)
	currentFen := ec.getFen()

	nextMove := ec.SuggestNextMove(*model.NewFenState(currentFen))

	fmt.Println("Engine move: ", nextMove)
	ec.sendCommand("position fen " + currentFen + " moves " + nextMove)
	return ec.getFen(), nextMove, nil
}

func (ec *EngineConnector) SuggestNextMove(fen model.FenState) string {
	ec.sendCommand(searchDepthCommand)

	line, _, err := ec.reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	for strings.Split(string(line[:]), " ")[0] != "bestmove" {
		line, _, err = ec.reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
	}

	return strings.Split(string(line[:]), " ")[1]
}

func (ec *EngineConnector) sendCommand(command string) {
	_, err := ec.writer.WriteString(command + "\n")
	if err != nil {
		log.Fatal(err)
	}
	ec.writer.Flush()
}

func (ec *EngineConnector) getFen() string {
	ec.sendCommand("d")
	// example stockfish output:
	//
	//  +---+---+---+---+---+---+---+---+
	//  | r | n | b | q | k | b |   | r | 8
	//  +---+---+---+---+---+---+---+---+
	//  | p | p | p | p | p | p | p | p | 7
	//  +---+---+---+---+---+---+---+---+
	//  |   |   |   |   |   | n |   |   | 6
	//  +---+---+---+---+---+---+---+---+
	//  |   |   |   |   |   |   |   |   | 5
	//  +---+---+---+---+---+---+---+---+
	//  |   |   |   |   | P |   |   |   | 4
	//  +---+---+---+---+---+---+---+---+
	//  |   |   |   |   |   |   |   |   | 3
	//  +---+---+---+---+---+---+---+---+
	//  | P | P | P | P |   | P | P | P | 2
	//  +---+---+---+---+---+---+---+---+
	//  | R | N | B | Q | K | B | N | R | 1
	//  +---+---+---+---+---+---+---+---+
	//    a   b   c   d   e   f   g   h

	// Fen: rnbqkb1r/pppppppp/5n2/8/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 1 2
	// Key: 16EA26742B9F7B69
	// Checkers:

	line, _, err := ec.reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	for strings.Split(string(line[:]), " ")[0] != "Fen:" {
		line, _, err = ec.reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
	}
	return strings.Join(strings.Split(string(line[:]), " ")[1:], " ")
}

func (ec *EngineConnector) Evaluate() (string, error) {
	// stockfish evaluation example:
	// 	info string NNUE evaluation using nn-b1a57edbea57.nnue
	// info string NNUE evaluation using nn-baff1ede1f90.nnue

	//  NNUE derived piece values:
	// +-------+-------+-------+-------+-------+-------+-------+-------+
	// |   r   |   n   |   b   |   q   |   k   |   b   |   n   |   r   |
	// | -5.70 | -4.48 | -5.07 | -9.91 |       | -4.88 | -4.49 | -5.60 |
	// +-------+-------+-------+-------+-------+-------+-------+-------+
	// |   p   |   p   |   p   |   p   |   p   |   p   |   p   |   p   |
	// | -0.63 | -1.12 | -1.27 | -1.28 | -1.46 | -1.66 | -1.63 | -0.63 |
	// +-------+-------+-------+-------+-------+-------+-------+-------+
	// |       |       |       |       |       |       |       |       |
	// |       |       |       |       |       |       |       |       |
	// +-------+-------+-------+-------+-------+-------+-------+-------+
	// |       |       |       |       |       |       |       |       |
	// |       |       |       |       |       |       |       |       |
	// +-------+-------+-------+-------+-------+-------+-------+-------+
	// |       |       |       |       |       |       |       |       |
	// |       |       |       |       |       |       |       |       |
	// +-------+-------+-------+-------+-------+-------+-------+-------+
	// |       |       |       |       |       |       |       |       |
	// |       |       |       |       |       |       |       |       |
	// +-------+-------+-------+-------+-------+-------+-------+-------+
	// |   P   |   P   |   P   |   P   |   P   |   P   |   P   |   P   |
	// | +0.45 | +0.81 | +0.96 | +1.04 | +1.16 | +1.30 | +1.27 | +0.47 |
	// +-------+-------+-------+-------+-------+-------+-------+-------+
	// |   R   |   N   |   B   |   Q   |   K   |   B   |   N   |   R   |
	// | +4.81 | +3.73 | +4.49 | +9.58 |       | +4.23 | +3.74 | +4.89 |
	// +-------+-------+-------+-------+-------+-------+-------+-------+

	//  NNUE network contributions (White to move)
	// +------------+------------+------------+------------+
	// |   Bucket   |  Material  | Positional |   Total    |
	// |            |   (PSQT)   |  (Layers)  |            |
	// +------------+------------+------------+------------+
	// |  0         |     0.00   |  -  0.26   |  -  0.26   |
	// |  1         |     0.00   |  +  0.87   |  +  0.87   |
	// |  2         |     0.00   |  +  0.26   |  +  0.26   |
	// |  3         |     0.00   |  +  0.10   |  +  0.10   |
	// |  4         |     0.00   |  +  0.28   |  +  0.28   |
	// |  5         |     0.00   |  +  0.07   |  +  0.07   |
	// |  6         |     0.00   |  +  0.08   |  +  0.08   |
	// |  7         |     0.00   |  +  0.09   |  +  0.09   | <-- this bucket is used
	// +------------+------------+------------+------------+

	// NNUE evaluation        +0.09 (white side)
	// Final evaluation       +0.11 (white side) [with scaled NNUE, ...]

	// get last line and return final evaluation value
	ec.sendCommand("eval")
	line, _, err := ec.reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	for strings.Split(string(line[:]), " ")[0] != "Final" {
		line, _, err = ec.reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
	}
	return strings.Fields(string(line[:]))[2], nil
}
