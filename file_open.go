package main

import (
	"log"

	"github.com/lxn/walk"

	"auto-curl/utils"
)

type MyMainWindow struct {
	*walk.MainWindow
	prevFilePath string
}

func (mw *MyMainWindow) FileOpen(cb func(path string)) error {
	cfg := utils.GetCfg()

	dlg := new(walk.FileDialog)

	dlg.FilePath = mw.prevFilePath
	dlg.Filter = "Bash (*.exe;)|*.exe;"
	dlg.Title = "Select Bash.exe"

	if ok, err := dlg.ShowOpen(mw); err != nil {
		return err
	} else if !ok {
		return nil
	}

	log.Println(dlg.FilePath)

	cfg.Bash = dlg.FilePath
	utils.WriteCfg()
	if cb != nil {
		cb(cfg.Bash)
	}

	return nil
}
