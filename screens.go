package main

import (
	"github.com/ebitenui/ebitenui/widget"
)

func createCenteredButton(res *UIResources, text string, disabled bool) *widget.Container {
	btnContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{Stretch: true})),
	)

	btn := widget.NewButton(
		widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: widget.AnchorLayoutPositionCenter,
		})),
		widget.ButtonOpts.Image(res.button.image),
		widget.ButtonOpts.Text(text, res.button.face, res.button.text),
		widget.ButtonOpts.TextPadding(res.button.padding),
	)
	btn.GetWidget().Disabled = disabled

	btnContainer.AddChild(btn)

	return btnContainer
}

func titleScreenContainer(res *UIResources) widget.PreferredSizeLocateableWidget {
	titleScreenContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
		),
		),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{StretchHorizontal: true})),
	)

	titleContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(
			widget.AnchorLayoutOpts.Padding(widget.Insets{Top: screenHeight * 0.2}),
		)),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{Stretch: true})),
	)
	title := widget.NewText(
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: widget.AnchorLayoutPositionCenter,
		})),
		widget.TextOpts.Text("Demo Game", res.text.titleFace, res.colour.teal))
	titleContainer.AddChild(title)
	titleScreenContainer.AddChild(titleContainer)

	btnContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.Insets{Top: screenHeight * 0.25}),
		),
		),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{Stretch: true})),
	)

	btnContainer.AddChild(createCenteredButton(res, "Story", true))
	btnContainer.AddChild(createCenteredButton(res, "Arcade", false))
	btnContainer.AddChild(createCenteredButton(res, "Options", false))

	titleScreenContainer.AddChild(btnContainer)
	return titleScreenContainer
}

func arcadeScreen(res *UIResources) widget.PreferredSizeLocateableWidget {
	return widget.NewContainer()
}

func optionsScreen(res *UIResources) widget.PreferredSizeLocateableWidget {
	return widget.NewContainer()
}