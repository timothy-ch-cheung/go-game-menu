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

type Game struct{
	ui *ebitenui.UI
}

func createButton(res *UIResources, text string) (*widget.Button) {
	return widget.NewButton(
		widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Stretch: false,
		})),
		widget.ButtonOpts.Image(res.button.image),
		widget.ButtonOpts.Text(text, res.button.face, res.button.text),
		widget.ButtonOpts.TextPadding(res.button.padding),
	)
}

func titleScreenContainer(res *UIResources, ui func() *ebitenui.UI) widget.PreferredSizeLocateableWidget {
	container := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(widget.RowLayoutOpts.Direction(widget.DirectionVertical))))

	container.AddChild(createButton(res, "Story"))
	container.AddChild(createButton(res, "Arcade"))
	container.AddChild(createButton(res, "Options"))

	return container
}

func createUI() (*ebitenui.UI, func(), error) {
	res, err := loadUIResources()
	if err != nil {
		return nil, nil, err
	}

	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(1),
			widget.GridLayoutOpts.Stretch([]bool{true}, []bool{false, true, false}),
			widget.GridLayoutOpts.Padding(widget.Insets{
				Top:    20,
				Bottom: 20,
			}),
			widget.GridLayoutOpts.Spacing(0, 20))),

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