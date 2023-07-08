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

func createCenteredButton(res *UIResources, text string) *widget.Container {
	btnContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{Stretch: true})),
		widget.ContainerOpts.BackgroundImage(res.background),
	)

	btn := widget.NewButton(
		widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: widget.AnchorLayoutPositionCenter,
		})),
		widget.ButtonOpts.Image(res.button.image),
		widget.ButtonOpts.Text(text, res.button.face, res.button.text),
		widget.ButtonOpts.TextPadding(res.button.padding),
	)

	btnContainer.AddChild(btn)

	return btnContainer
}

func titleScreenContainer(res *UIResources, ui func() *ebitenui.UI) widget.PreferredSizeLocateableWidget {
	container := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.Insets{Top: screenHeight * 0.6}),
		),
		),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{StretchHorizontal: true})),
	)

	container.AddChild(createCenteredButton(res, "Story"))
	container.AddChild(createCenteredButton(res, "Arcade"))
	container.AddChild(createCenteredButton(res, "Options"))

	return container
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

	var ui *ebitenui.UI

	ui = &ebitenui.UI{
		Container: rootContainer,
	}

	rootContainer.AddChild(titleScreenContainer(res, func() *ebitenui.UI {
		return ui
	}))

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
