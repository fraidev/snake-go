package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Apple struct {
	Body rl.Rectangle
}

func NewApple(x, y float32) *Apple {
	return &Apple{
		Body: rl.NewRectangle(x, y, 40, 40),
	}
}

func (a *Apple) regen() {
	a.Body.X = float32(rl.GetRandomValue(0, 19)) * 40
	a.Body.Y = float32(rl.GetRandomValue(0, 19)) * 40
}

func (s *Apple) Draw() {
	rl.DrawRectangleRec(s.Body, rl.Red)
}

func (s *Apple) Update() {

}

