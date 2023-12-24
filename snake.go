package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Snake struct {
	Body      []rl.Rectangle
	Direction Direction
}

func NewSnake(x, y float32) *Snake {
	return &Snake{
		Body: []rl.Rectangle{
			rl.NewRectangle(x, y, 40, 40),
			rl.NewRectangle(x, y-40, 40, 40),
			rl.NewRectangle(x, y-80, 40, 40),
		},
		Direction: Down,
	}
}

func (s *Snake) Draw() {
	for _, bodyPart := range s.Body {
		rl.DrawRectangleRec(bodyPart, rl.Green)
	}
}

func (s *Snake) Update(apple *Apple, isGameOver *bool, score *int) {
	lastMove := s.Body[0]
	newMove := getNewMove(lastMove, s.Direction)

	// check wall collision
	if newMove.X < 0 || newMove.X > 760 || newMove.Y < 0 || newMove.Y > 760 {
		*isGameOver = true
		return
	}

	// check snake collision
	for _, bodyPart := range s.Body {
		if newMove.X == bodyPart.X && newMove.Y == bodyPart.Y {
			*isGameOver = true
			return
		}
	}

	// add new snake part
	s.Body = append([]rl.Rectangle{newMove}, s.Body...)

	// check if snake ate apple
	// if so, regenerate apple and increase score
	// else, remove last snake part
	if s.Body[0].X == apple.Body.X && s.Body[0].Y == apple.Body.Y {
		apple.regen()
		*score++
	} else {
		s.Body = s.Body[:len(s.Body)-1]
	}
}

func (s *Snake) SetDirectionLeft() {
	s.Direction = Left
}

func (s *Snake) SetDirectionRight() {
	s.Direction = Right
}

func (s *Snake) SetDirectionUp() {
	s.Direction = Up
}

func (s *Snake) SetDirectionDown() {
	s.Direction = Down
}

func getNewMove(lastMove rl.Rectangle, direction Direction) rl.Rectangle {
	switch direction {
	case Up:
		return rl.NewRectangle(lastMove.X, lastMove.Y-40, 40, 40)
	case Down:
		return rl.NewRectangle(lastMove.X, lastMove.Y+40, 40, 40)
	case Left:
		return rl.NewRectangle(lastMove.X-40, lastMove.Y, 40, 40)
	case Right:
		return rl.NewRectangle(lastMove.X+40, lastMove.Y, 40, 40)
	}

	panic("Invalid direction")
}
