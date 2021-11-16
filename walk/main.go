package main

import (
	"auto-curl/utils"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
)

type MainWindow = declarative.MainWindow
type Size = declarative.Size
type PushButton = declarative.PushButton
type LineEdit = declarative.LineEdit
type Label = declarative.Label
type StatusBarItem = declarative.StatusBarItem

var (
	Trace   *log.Logger // 记录所有日志
	Info    *log.Logger // 重要的信息
	Warning *log.Logger // 需要注意的信息
	Error   *log.Logger // 非常严重的问题
)

func init() {
	utils.ReadCfg()

	file, err := os.OpenFile("trace.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	Trace = log.New(io.MultiWriter(file, os.Stdout),
		"Trace: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	var inTE, outTE *walk.TextEdit
	var bash *walk.StatusBarItem
	var status *walk.StatusBarItem
	var running *walk.StatusBarItem
	var proxy *walk.LineEdit
	var interval *walk.LineEdit

	var runningTicker *time.Ticker

	var count int

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
					declarative.TextEdit{AssignTo: &inTE, VScroll: true},
					declarative.TextEdit{AssignTo: &outTE, VScroll: true, ReadOnly: true},
				},
			},

			declarative.Composite{
				Layout: declarative.Grid{Columns: 6},
				Children: []declarative.Widget{
					Label{
						Text: "Proxy:",
					},
					LineEdit{
						AssignTo:  &proxy,
						MaxSize:   Size{Width: 300},
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
					Label{
						Text: "Interval:",
					},
					LineEdit{
						AssignTo:  &interval,
						Text:      strconv.Itoa(cfg.Interval),
						MaxSize:   Size{Width: 100},
						Alignment: declarative.AlignHCenterVCenter,
						// OnEditingFinished:,
						OnKeyPress: func(key walk.Key) {
							if key == walk.KeyReturn {
								str, err := strconv.Atoi(interval.Text())
								if err == nil {
									cfg.Interval = str
									utils.WriteCfg()
									status.SetText("保存成功!")
								}
							}
						},
					},
					PushButton{
						Text:    "Paste",
						MaxSize: Size{Width: 100},
						OnClicked: func() {
							if text, err := walk.Clipboard().Text(); err != nil {
								log.Print("Paste: ", err)
							} else {
								inTE.SetText(text)
							}
						},
					},
					PushButton{
						Text:    "Config Bash",
						MaxSize: Size{Width: 100},
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
					// stop previous ticker
					if runningTicker != nil {
						runningTicker.Stop()
						runningTicker = nil
						running.SetText("已停止")
					} else {
						request := inTE.Text()
						reqFunc := func() {
							count++
							if len(cfg.Bash) == 0 {
								rs := walk.MsgBox(mw, "Error", "please set bash location", walk.MsgBoxOKCancel)
								fmt.Println(rs)
							} else if len(request) > 0 && strings.Index(request, "curl") == 0 {
								now := time.Now()
								status.SetText("请求中...")
								running.SetText(fmt.Sprintf("运行中...%d", count))
								response := utils.Curl(cfg.Bash, request)
								outTE.SetText(response)
								Trace.Println(response)
								status.SetText(fmt.Sprintf("请求完成! %vms\n", time.Since(now).Milliseconds()))
							}
						}

						count = 0

						if cfg.Interval > 0 {
							runningTicker = utils.StartTicker(int64(cfg.Interval), reqFunc)
						} else {
							reqFunc()
						}
					}
				},
			},
		},

		StatusBarItems: []StatusBarItem{
			{
				AssignTo: &status,
				Text:     "",
				Width:    200,
			},
			{
				AssignTo: &running,
				Text:     "",
				Width:    200,
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
