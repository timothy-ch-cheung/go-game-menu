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

type TextResources struct {
	titleFace font.Face
}

type ColourResourses struct {
	teal color.Color
}

type PanelResources struct {
	image   *image.NineSlice
	padding widget.Insets
}

type UIResources struct {
	background *image.NineSlice
	fonts      *fonts
	button     *ButtonResources
	text       *TextResources
	colour     ColourResourses
	panel      PanelResources
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

func loadColourResources() *ColourResourses {
	return &ColourResourses{
		teal: hexToColor("008080"),
	}
}

func loadButtonResources(fonts *fonts) (*ButtonResources, error) {
	idle, err := loadImageNineSlice("assets/graphics/button-idle.png", 34, 20)
	if err != nil {
		return nil, err
	}

	hover, err := loadImageNineSlice("assets/graphics/button-hover.png", 34, 20)
	if err != nil {
		return nil, err
	}
	pressed, err := loadImageNineSlice("assets/graphics/button-pressed.png", 34, 20)
	if err != nil {
		return nil, err
	}

	disabled, err := loadImageNineSlice("assets/graphics/button-disabled.png", 34, 20)
	if err != nil {
		return nil, err
	}

	img := &widget.ButtonImage{
		Idle:     idle,
		Hover:    hover,
		Pressed:  pressed,
		Disabled: disabled,
	}

	return &ButtonResources{
		image: img,

		text: &widget.ButtonTextColor{
			Idle:     hexToColor(textIdleColour),
			Disabled: hexToColor(textDisabledColour),
		},

		face: fonts.face,

		padding: widget.Insets{
			Left:   30,
			Right:  30,
			Top:    10,
			Bottom: 10,
		},
	}, nil
}

func loadPanelResources() (*PanelResources, error) {
	img, err := loadImageNineSlice("assets/graphics/panel-idle.png", 48, 48)
	if err != nil {
		return nil, err
	}

	return &PanelResources{
		image: img,
		padding: widget.Insets{
			Top:    20,
			Bottom: 20,
			Left:   20,
			Right:  20,
		},
	}, nil
}

func loadUIResources() (*UIResources, error) {
	background := image.NewNineSliceColor(hexToColor(backgroundColour))
	colours := *loadColourResources()

	fonts, err := loadFonts()
	if err != nil {
		return nil, err
	}

	button, err := loadButtonResources(fonts)
	if err != nil {
		return nil, err
	}

	text := TextResources{
		titleFace: fonts.titleFace,
	}

	panel, err := loadPanelResources()
	if err != nil {
		return nil, err
	}

	return &UIResources{
		background: background,
		fonts:      fonts,
		button:     button,
		text:       &text,
		colour:     colours,
		panel:      *panel,
	}, nil
}
