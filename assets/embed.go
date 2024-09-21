package assets

import (
	_ "embed"
)

var (
	//go:embed Shore.png
	Shore_png []byte

	//go:embed wallandfloor.png
	Wallandfloor_png []byte

	//go:embed TexturedGrass.png
	Grass_png []byte

	//go:embed my_guy.png
	MyGuy_png []byte

	//go:embed Baddie.png
	Baddie_png []byte
)
