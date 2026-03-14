package game

import (
	"bytes"

	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text/v2"
)

var gameFont ebitentext.Face

func init() {
	src, err := ebitentext.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))
	if err != nil {
		panic(err)
	}
	gameFont = &ebitentext.GoTextFace{Source: src, Size: FontSize}
}
