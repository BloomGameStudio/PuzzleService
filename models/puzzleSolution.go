package models

type PuzzleSolution struct {
    PuzzleID int    `json:"puzzle_id"`
    Solution string `json:"solution"`
    Solved   bool   `json:"solved"`
}