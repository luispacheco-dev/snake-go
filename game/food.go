package game

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Food struct {
	x      float32
	y      float32
	width  float32
	height float32
}

func NewFood(g *Game) *Food {
	f := &Food{
		width:  10,
		height: 10,
	}
	f.randomXY(g.board)

	return f

}

func (f *Food) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, f.x, f.y, f.width, f.height, color.White, false)
}

// ---

func (f *Food) randomXY(b Rect) {
	f.x = randRange(b.x, b.x+b.width-f.width)
	f.y = randRange(b.y, b.y+b.height-f.height)
}

//---

func randRange(min, max float32) float32 {
	return float32(rand.Intn(int(max)-int(min))) + min
}
