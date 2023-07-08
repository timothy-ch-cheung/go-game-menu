package main

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

const (
	fontFace = "assets/fonts/PrintChar21.ttf"
)

type fonts struct {
	face         font.Face
	smallFace    font.Face
}

func (f *fonts) close() {
	if f.face != nil {
		_ = f.face.Close()
	}
	if f.smallFace != nil {
		_ = f.smallFace.Close()
	}
}

func loadFont(path string, size float64) (font.Face, error) {
	fontData, err := embeddedAssets.ReadFile(path)
	if err != nil {
		return nil, err
	}

	ttfFont, err := truetype.Parse(fontData)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(ttfFont, &truetype.Options{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	}), nil
}

func loadFonts() (*fonts, error) {
	regularfontFace, err := loadFont(fontFace, 16)
	if err != nil {
		return nil, err
	}

	smallFontFace, err := loadFont(fontFace, 12)

	return &fonts{
		face:         regularfontFace,
		smallFace:    smallFontFace,
	}, nil
}