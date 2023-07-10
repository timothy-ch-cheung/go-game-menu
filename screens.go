package main

import (
	"image/color"

	"github.com/ebitenui/ebitenui/widget"

	"fmt"
)

func createCenteredButton(res *UIResources, text string, disabled bool, handler widget.ButtonClickedHandlerFunc) *widget.Container {
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
		widget.ButtonOpts.ClickedHandler(handler),
	)
	btn.GetWidget().Disabled = disabled

	btnContainer.AddChild(btn)

	return btnContainer
}

func getLabelColour(colour color.Color) *widget.LabelColor {
	return &widget.LabelColor{
		Idle:     colour,
		Disabled: colour,
	}
}

func createCheckbox(res *UIResources, label string) *widget.LabeledCheckbox {
	return widget.NewLabeledCheckbox(
		widget.LabeledCheckboxOpts.Spacing(res.checkbox.spacing),
		widget.LabeledCheckboxOpts.CheckboxOpts(
			widget.CheckboxOpts.ButtonOpts(widget.ButtonOpts.Image(res.checkbox.image)),
			widget.CheckboxOpts.Image(res.checkbox.graphic)),
		widget.LabeledCheckboxOpts.LabelOpts(widget.LabelOpts.Text(label, res.text.smallFace, getLabelColour(res.colour.teal))))
}

func titleScreenContainer(res *UIResources, switchScreen SwitchScreenFunc) widget.PreferredSizeLocateableWidget {
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

	btnContainer.AddChild(createCenteredButton(res, "Story", true, func(args *widget.ButtonClickedEventArgs) {}))
	btnContainer.AddChild(createCenteredButton(res, "Arcade", false, func(args *widget.ButtonClickedEventArgs) { switchScreen(Arcade) }))
	btnContainer.AddChild(createCenteredButton(res, "Options", false, func(args *widget.ButtonClickedEventArgs) { switchScreen(Options) }))

	titleScreenContainer.AddChild(btnContainer)
	return titleScreenContainer
}

func arcadeScreenContainer(res *UIResources, switchScreen SwitchScreenFunc) widget.PreferredSizeLocateableWidget {
	arcadeContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(res.panel.padding),
		)),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{StretchHorizontal: true})),
	)
	title := widget.NewText(
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: widget.AnchorLayoutPositionCenter,
		})),
		widget.TextOpts.Text("Arcade Page", res.text.titleFace, res.colour.teal))
	arcadeContainer.AddChild(title)

	backBtn := widget.NewButton(
		widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: widget.AnchorLayoutPositionCenter,
		})),
		widget.ButtonOpts.Image(res.button.image),
		widget.ButtonOpts.Text("Back", res.button.face, res.button.text),
		widget.ButtonOpts.TextPadding(res.button.padding),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			switchScreen(Title)
		}),
	)
	arcadeContainer.AddChild(backBtn)

	return arcadeContainer
}

func optionsScreenContainer(res *UIResources, switchScreen SwitchScreenFunc) widget.PreferredSizeLocateableWidget {
	optionsContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(res.panel.padding),
		)),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{StretchHorizontal: true, StretchVertical: true})),
	)

	backBtn := widget.NewButton(
		widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: widget.AnchorLayoutPositionCenter,
		})),
		widget.ButtonOpts.Image(res.button.image),
		widget.ButtonOpts.Text("Back", res.button.face, res.button.text),
		widget.ButtonOpts.TextPadding(res.button.padding),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			switchScreen(Title)
		}),
	)
	optionsContainer.AddChild(backBtn)

	optionsPanel := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Padding(res.panel.padding),
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(5),
		)),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{Stretch: true})),
		widget.ContainerOpts.BackgroundImage(res.panel.image),
	)
	optionsContainer.AddChild(optionsPanel)

	optionsPanel.AddChild(createCheckbox(res, "Setting Alpha"))
	optionsPanel.AddChild(createCheckbox(res, "Setting Beta"))
	optionsPanel.AddChild(createCheckbox(res, "Setting Gamma"))
	optionsPanel.AddChild(createCheckbox(res, "Setting Delta"))

	var volumeText *widget.Label
	volumeContainer := widget.NewContainer(widget.ContainerOpts.Layout(widget.NewRowLayout(
		widget.RowLayoutOpts.Spacing(10),
	)))
	volumeSlider := widget.NewSlider(
		widget.SliderOpts.WidgetOpts(widget.WidgetOpts.MinSize(250, 10)),
		widget.SliderOpts.MinMax(1, 100),
		widget.SliderOpts.Images(res.slider.trackImage, res.slider.handle),
		widget.SliderOpts.FixedHandleSize(res.slider.handleSize),
		widget.SliderOpts.TrackOffset(5),
		widget.SliderOpts.ChangedHandler(func(args *widget.SliderChangedEventArgs) {
			volumeText.Label = fmt.Sprintf("Volume: %d", args.Current)
		}),
	)
	volumeContainer.AddChild(volumeSlider)
	volumeText = widget.NewLabel(
		widget.LabelOpts.Text(fmt.Sprintf("Volume: %d", volumeSlider.Current), res.text.smallFace, getLabelColour(res.colour.teal)),
	)
	volumeContainer.AddChild(volumeText)

	optionsPanel.AddChild(volumeContainer)

	return optionsContainer
}
