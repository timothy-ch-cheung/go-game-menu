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
	ui *ebitenui.UI
}

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
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{StretchHorizontal: true})),
	))
	rootContainer.AddChild(flipBook)

	var ui *ebitenui.UI

	ui = &ebitenui.UI{
		Container: rootContainer,
	}

	titleScreen := titleScreenContainer(res)
	flipBook.SetPage(titleScreen)

	return ui, func() {
		res.close()
	}, nil
}

func (game *Game) Update() error {
	game.ui.Update()
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.ui.Draw(screen)
	msg := fmt.Sprintf("FPS: %0.2f", ebiten.ActualFPS())
	ebitenutil.DebugPrint(screen, msg)
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

	defer closeUI()

	if err := ebiten.RunGame(&Game{ui: ui}); err != nil {
		log.Fatal(err)
	}
}
