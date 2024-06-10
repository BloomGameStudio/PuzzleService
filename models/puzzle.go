package models

type Puzzle struct {
	ID        [32]byte
	Solution  []byte
	Title     string
	Committed bool
	Revealed  bool
}
