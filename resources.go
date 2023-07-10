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
	smallFace font.Face
	titleFace font.Face
}

type ColourResourses struct {
	teal color.Color
}

type PanelResources struct {
	image   *image.NineSlice
	padding widget.Insets
}

type CheckboxResources struct {
	image   *widget.ButtonImage
	graphic *widget.CheckboxGraphicImage
	spacing int
}

type SliderResources struct {
	trackImage *widget.SliderTrackImage
	handle     *widget.ButtonImage
	handleSize int
}

type UIResources struct {
	background *image.NineSlice
	fonts      *fonts
	button     *ButtonResources
	text       *TextResources
	colour     *ColourResourses
	panel      *PanelResources
	checkbox   *CheckboxResources
	slider     *SliderResources
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

func loadCheckboxResources() (*CheckboxResources, error) {
	idle, err := loadImageNineSlice("assets/graphics/checkbox-idle.png", 20, 0)
	if err != nil {
		return nil, err
	}

	hover, err := loadImageNineSlice("assets/graphics/checkbox-hover.png", 20, 0)
	if err != nil {
		return nil, err
	}

	disabled, err := loadImageNineSlice("assets/graphics/checkbox-disabled.png", 20, 0)
	if err != nil {
		return nil, err
	}

	checked, err := loadGraphicImages("assets/graphics/checkbox-checked-idle.png", "assets/graphics/checkbox-checked-disabled.png")
	if err != nil {
		return nil, err
	}

	unchecked, err := loadGraphicImages("assets/graphics/checkbox-unchecked-idle.png", "assets/graphics/checkbox-unchecked-disabled.png")
	if err != nil {
		return nil, err
	}

	greyed, err := loadGraphicImages("assets/graphics/checkbox-greyed-idle.png", "assets/graphics/checkbox-greyed-disabled.png")
	if err != nil {
		return nil, err
	}

	return &CheckboxResources{
		image: &widget.ButtonImage{
			Idle:     idle,
			Hover:    hover,
			Pressed:  hover,
			Disabled: disabled,
		},

		graphic: &widget.CheckboxGraphicImage{
			Checked:   checked,
			Unchecked: unchecked,
			Greyed:    greyed,
		},

		spacing: 10,
	}, nil
}

func loadSliderResources() (*SliderResources, error) {
	idle, err := newImageFromFile("assets/graphics/slider-track-idle.png")
	if err != nil {
		return nil, err
	}

	disabled, err := newImageFromFile("assets/graphics/slider-track-disabled.png")
	if err != nil {
		return nil, err
	}

	handleIdle, err := newImageFromFile("assets/graphics/slider-handle-idle.png")
	if err != nil {
		return nil, err
	}

	handleHover, err := newImageFromFile("assets/graphics/slider-handle-hover.png")
	if err != nil {
		return nil, err
	}

	handleDisabled, err := newImageFromFile("assets/graphics/slider-handle-disabled.png")
	if err != nil {
		return nil, err
	}

	return &SliderResources{
		trackImage: &widget.SliderTrackImage{
			Idle:     image.NewNineSlice(idle, [3]int{0, 30, 0}, [3]int{8, 0, 0}),
			Hover:    image.NewNineSlice(idle, [3]int{0, 30, 0}, [3]int{8, 0, 0}),
			Disabled: image.NewNineSlice(disabled, [3]int{0, 30, 0}, [3]int{8, 0, 0}),
		},

		handle: &widget.ButtonImage{
			Idle:     image.NewNineSliceSimple(handleIdle, 5, 8),
			Hover:    image.NewNineSliceSimple(handleHover, 5, 8),
			Pressed:  image.NewNineSliceSimple(handleHover, 5, 8),
			Disabled: image.NewNineSliceSimple(handleDisabled, 5, 8),
		},

		handleSize: 6,
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
		smallFace: fonts.smallFace,
		titleFace: fonts.titleFace,
	}

	panel, err := loadPanelResources()
	if err != nil {
		return nil, err
	}

	checkbox, err := loadCheckboxResources()
	if err != nil {
		return nil, err
	}

	slider, err := loadSliderResources()
	if err != nil {
		return nil, err
	}

	return &UIResources{
		background: background,
		fonts:      fonts,
		button:     button,
		text:       &text,
		colour:     &colours,
		panel:      panel,
		checkbox:   checkbox,
		slider:     slider,
	}, nil
}
