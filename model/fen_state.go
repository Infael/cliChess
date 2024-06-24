package model

type FenState struct {
	Fen string
}

func NewFenState(fen string) *FenState {
	return &FenState{
		Fen: fen,
	}
}

func NewGame() *FenState {
	return &FenState{
		Fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
	}
}

func (fs *FenState) GetFen() string {
	return fs.Fen
}

func (fs *FenState) SetFen(fen string) {
	fs.Fen = fen
}
