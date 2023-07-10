package main

import (
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func newImageFromFile(path string) (*ebiten.Image, error) {
	file, err := embeddedAssets.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := ebitenutil.NewImageFromReader(file)
	return img, err
}

func loadImageNineSlice(path string, centerWidth int, centerHeight int) (*image.NineSlice, error) {
	img, err := newImageFromFile(path)
	if err != nil {
		return nil, err
	}
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	return image.NewNineSlice(img,
			[3]int{(width - centerWidth) / 2, centerWidth, width - (width-centerWidth)/2 - centerWidth},
			[3]int{(height - centerHeight) / 2, centerHeight, height - (height-centerHeight)/2 - centerHeight}),
		nil
}

func loadGraphicImages(idle string, disabled string) (*widget.ButtonImageImage, error) {
	idleImage, err := newImageFromFile(idle)
	if err != nil {
		return nil, err
	}

	disabledImage, err := newImageFromFile(disabled)
	if err != nil {
		return nil, err
	}

	return &widget.ButtonImageImage{
		Idle:     idleImage,
		Disabled: disabledImage,
	}, nil
}
