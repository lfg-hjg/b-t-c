package main

import (
	"bytes"
	"fmt"
	"github.com/harrisongerst/b-t-c/assets"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	SCREEN_HEIGHT = 241
	SCREEN_WIDTH  = 241

	TILE_SIZE = 16
)

var tilesImage *ebiten.Image

func init() {
	var err error
	img, _, err := image.Decode(bytes.NewReader(assets.Shore_png))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = ebiten.NewImageFromImage(img)
}

type Game struct {
	layers [][]int
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	w := tilesImage.Bounds().Dx()
	tileXCount := w / TILE_SIZE

	const xCount = SCREEN_WIDTH / TILE_SIZE

	for _, l := range g.layers {
		for i, t := range l {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate((float64(i%xCount) * TILE_SIZE), float64((i/xCount)*TILE_SIZE))
			sx := (t % tileXCount) * TILE_SIZE
			sy := (t / tileXCount) * TILE_SIZE
			screen.DrawImage(tilesImage.SubImage(image.Rect(sx, sy, sx+TILE_SIZE, sy+TILE_SIZE)).(*ebiten.Image), op)
		}
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %1.2f", ebiten.ActualTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func main() {

	g := &Game{
		layers: [][]int{
			{
				3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
				3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
				3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
				3, 218, 3, 3, 3, 3, 3, 3, 3, 3, 3, 218, 3, 244, 3,
				3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,

				3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
				3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
				3, 3, 244, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
				3, 3, 3, 3, 3, 3, 3, 3, 3, 219, 3, 3, 3, 219, 3,
				3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,

				3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
				3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
				3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
				3, 218, 3, 3, 3, 3, 3, 3, 3, 3, 3, 244, 3, 3, 3,
				3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
			},
			{
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
				1, 1, 1, 1, 1, 1, 1, 1, 54, 55, 56, 1, 1, 1, 1,
				1, 1, 1, 1, 1, 76, 77, 78, 79, 81, 81, 1, 1, 1, 1,
				1, 1, 1, 1, 1, 111, 112, 113, 114, 115, 116, 1, 1, 1, 1,

				1, 1, 1, 1, 1, 126, 127, 128, 129, 131, 131, 1, 1, 1, 1,
				1, 1, 1, 1, 1, 313, 313, 245, 242, 313, 313, 1, 1, 1, 1,
				1, 1, 1, 1, 1, 1, 1, 245, 242, 1, 1, 1, 1, 1, 1,
				1, 1, 1, 1, 1, 1, 1, 245, 242, 1, 1, 1, 1, 1, 1,
				1, 1, 1, 1, 1, 1, 1, 245, 242, 1, 1, 1, 1, 1, 1,

				1, 1, 1, 1, 1, 1, 1, 245, 242, 1, 1, 1, 1, 1, 1,
				1, 1, 1, 1, 1, 1, 1, 245, 242, 1, 1, 1, 1, 1, 1,
				1, 1, 1, 1, 1, 1, 1, 245, 242, 1, 1, 1, 1, 1, 1,
				1, 1, 1, 1, 1, 1, 1, 245, 242, 1, 1, 1, 1, 1, 1,
				1, 1, 1, 1, 1, 1, 1, 245, 242, 1, 1, 1, 1, 1, 1,
			},
		},
	}
	ebiten.SetWindowSize(961, 961)
	ebiten.SetWindowTitle("beneath the castle")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
