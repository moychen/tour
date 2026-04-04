package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("应用管理器")

	// 顶部搜索栏
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("搜索应用...")

	// 应用展示区
	appGrid := container.NewGridWrap(fyne.NewSize(100, 100))
	appNames := []string{"应用1", "应用2", "应用3", "应用4", "应用5"}

	for _, appName := range appNames {
		appLabel := widget.NewLabel(appName)
		appLabel.Alignment = fyne.TextAlignCenter

		// 创建一个垂直容器，将标签居中放置
		buttonContent := container.NewCenter(appLabel)

		// 创建一个自定义按钮
		customButton := &customButton{
			button: widget.NewButton("", func() {
				openAppDetail(appName)
			}),
			content: buttonContent,
		}
		customButton.button.Importance = widget.LowImportance
		customButton.button.Resize(fyne.NewSize(100, 100))

		// 扩展按钮的渲染器
		customButton.ExtendBaseWidget(customButton)

		appGrid.Add(customButton)
	}

	// 布局
	content := container.NewVBox(
		searchEntry,
		appGrid,
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}

func openAppDetail(appName string) {
	fmt.Printf("打开应用详情: %s\n", appName)
	// 这里可以添加打开应用详情页的逻辑
}

type customButton struct {
	widget.Button
	content fyne.CanvasObject
}

func (b *customButton) CreateRenderer() fyne.WidgetRenderer {
	return &customButtonRenderer{
		widget.NewButtonRenderer(b.Button),
		b.content,
	}
}

type customButtonRenderer struct {
	widget.ButtonRenderer
	content fyne.CanvasObject
}

func (r *customButtonRenderer) Layout(size fyne.Size) {
	r.ButtonRenderer.Layout(size)
	r.content.Resize(size)
	r.content.Move(fyne.NewPos(0, 0))
}

func (r *customButtonRenderer) MinSize() fyne.Size {
	return r.content.MinSize()
}

func (r *customButtonRenderer) Refresh() {
	r.ButtonRenderer.Refresh()
}
