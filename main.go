package main

import (
	"fmt"
	"log"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	ui    *ebitenui.UI
	arrow *character
}

type Screen int

type SwitchScreenFunc func(Screen)

const (
	Title   Screen = 0
	Arcade  Screen = 1
	Options Screen = 2
)

var currentScreen Screen = Title

func createUI() (*ebitenui.UI, func(), error) {
	res, err := loadUIResources()
	if err != nil {
		return nil, nil, err
	}

	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(widget.AnchorLayoutOpts.Padding(widget.Insets{
			Top:    20,
			Bottom: 20,
		}))),

		widget.ContainerOpts.BackgroundImage(res.background))

	flipBook := widget.NewFlipBook(widget.FlipBookOpts.ContainerOpts(
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{StretchHorizontal: true, StretchVertical: true})),
	))
	rootContainer.AddChild(flipBook)

	var ui *ebitenui.UI

	ui = &ebitenui.UI{
		Container: rootContainer,
	}

	var titleScreen, arcadeScreen, optionsScreen widget.PreferredSizeLocateableWidget

	switchScreen := func(screen Screen) {
		currentScreen = screen
		switch screen {
		case Title:
			flipBook.SetPage(titleScreen)
		case Arcade:
			flipBook.SetPage(arcadeScreen)
		case Options:
			flipBook.SetPage(optionsScreen)
		}
	}

	titleScreen = titleScreenContainer(res, switchScreen)
	arcadeScreen = arcadeScreenContainer(res, switchScreen)
	optionsScreen = optionsScreenContainer(res, switchScreen)

	flipBook.SetPage(titleScreen)

	return ui, func() {
		res.close()
	}, nil
}

func (game *Game) Update() error {
	game.ui.Update()

	switch currentScreen {
	case Arcade:
		game.arrow.update()
	}

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.ui.Draw(screen)
	msg := fmt.Sprintf("FPS: %0.2f", ebiten.ActualFPS())
	ebitenutil.DebugPrint(screen, msg)

	switch currentScreen {
	case Arcade:
		game.arrow.draw(screen)
	}
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Game Menu Demo")

	ui, closeUI, err := createUI()
	if err != nil {
		log.Fatal(err)
	}

	arrow, err := createArrow()

	defer closeUI()

	if err := ebiten.RunGame(&Game{ui: ui, arrow: arrow}); err != nil {
		log.Fatal(err)
	}
}
