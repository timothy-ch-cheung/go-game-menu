package main

import (
	"image/color"
	"strconv"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"golang.org/x/image/font"
)

const (
	backgroundColour = "131a22"

	textIdleColour     = "dff4ff"
	textDisabledColour = "5a7a91"
)

type ButtonResources struct {
	image   *widget.ButtonImage
	text    *widget.ButtonTextColor
	face    font.Face
	padding widget.Insets
}

type UIResources struct {
	background *image.NineSlice
	fonts *fonts
	button      *ButtonResources
}

func (res *UIResources) close() {
	res.fonts.close()
}

func hexToColor(hex string) color.Color {
	u, err := strconv.ParseUint(hex, 16, 0)
	if err != nil {
		panic(err)
	}

	return color.NRGBA{
		R: uint8(u & 0xff0000 >> 16),
		G: uint8(u & 0xff00 >> 8),
		B: uint8(u & 0xff),
		A: 255,
	}
}

func loadButtonResources(fonts *fonts) (*ButtonResources, error) {
	idle, err := loadImageNineSlice("assets/graphics/button-idle.png", 12, 0)
	if err != nil {
		return nil, err
	}

	hover, err := loadImageNineSlice("assets/graphics/button-hover.png", 12, 0)
	if err != nil {
		return nil, err
	}
	pressed, err := loadImageNineSlice("assets/graphics/button-pressed.png", 12, 0)
	if err != nil {
		return nil, err
	}

	disabled, err := loadImageNineSlice("assets/graphics/button-disabled.png", 12, 0)
	if err != nil {
		return nil, err
	}

	img := &widget.ButtonImage{
		Idle:         idle,
		Hover:        hover,
		Pressed:      pressed,
		PressedHover: hover,
		Disabled:     disabled,
	}

	return &ButtonResources{
		image: img,

		text: &widget.ButtonTextColor{
			Idle:     hexToColor(textIdleColour),
			Disabled: hexToColor(textDisabledColour),
		},

		face: fonts.face,

		padding: widget.Insets{
			Left:  30,
			Right: 30,
			Top: 10,
			Bottom: 10,
		},
	}, nil
}

func loadUIResources() (*UIResources, error) {
	background := image.NewNineSliceColor(hexToColor(backgroundColour))

	fonts, err := loadFonts()
	if err != nil {
		return nil, err
	}

	button, err := loadButtonResources(fonts)
	if err != nil {
		return nil, err
	}

	return &UIResources{
		background: background,
		fonts: fonts,
		button: button,
	}, nil
}