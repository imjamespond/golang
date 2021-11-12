package main

import (
	"auto-curl/utils"
	"strings"

	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
)

type MainWindow = declarative.MainWindow
type Size = declarative.Size
type PushButton = declarative.PushButton
type LineEdit = declarative.LineEdit
type Label = declarative.Label
type StatusBarItem = declarative.StatusBarItem

func main() {
	var inTE, outTE *walk.TextEdit
	var bash *walk.StatusBarItem
	var status *walk.StatusBarItem
	var proxy *walk.LineEdit

	cfg := utils.GetCfg()

	mw := new(MyMainWindow)

	MainWindow{
		AssignTo: &mw.MainWindow,

		Title:   "Auto Curl",
		MinSize: Size{Width: 800, Height: 600},
		Layout:  declarative.VBox{},
		Children: []declarative.Widget{
			declarative.HSplitter{
				Children: []declarative.Widget{
					declarative.TextEdit{AssignTo: &inTE},
					declarative.TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},

			declarative.Composite{
				Layout: declarative.Grid{Columns: 4},
				Children: []declarative.Widget{
					Label{
						Text: "Proxy:",
					},
					LineEdit{
						AssignTo:  &proxy,
						Text:      cfg.Proxy,
						Alignment: declarative.AlignHCenterVCenter,
						// OnEditingFinished:,
						OnKeyPress: func(key walk.Key) {
							if key == walk.KeyReturn {
								cfg.Proxy = proxy.Text()
								utils.WriteCfg()
								status.SetText("保存成功!")
							}
						},
					},
					PushButton{
						// MaxSize: Size{Width: 100},
						Text: "Config Bash",
						OnClicked: func() {
							mw.FileOpen(func(path string) {
								bash.SetText(getBash(path))
							})
						},
					},
				},
			},

			PushButton{
				Text: "请求",
				OnClicked: func() {
					request := inTE.Text()
					if len(cfg.Bash) == 0 {
						// dialog.NewError(fmt.Errorf("please set bash location"), win).Show()
					} else if len(request) > 0 && strings.Index(request, "curl") == 0 {
						status.SetText("请求中...")
						response := utils.Curl(cfg.Bash, request)
						outTE.SetText(response)
						status.SetText("请求完成!")
					}
				},
			},
		},

		StatusBarItems: []StatusBarItem{
			{
				AssignTo:    &status,
				Text:        "",
				ToolTipText: "no tooltip for me",
			},
			{
				AssignTo: &bash,
				Text:     getBash(cfg.Bash),
				Width:    400,
			},
		},
	}.Run()
}

func getBash(path string) string {
	return "bash: " + path
}
