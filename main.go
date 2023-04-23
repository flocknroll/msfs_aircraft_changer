package main

import (
	"log"
	"os"

	"github.com/flocknroll/azga/go_installer/msfstools"
	"github.com/rivo/tview"
	"gopkg.in/ini.v1"
)

func main() {
	path, ok := msfstools.GetPackageFolderPath()
	log.Print(path, ok)

	cfg, _ := ini.Load("test_data/test.cfg")

	cfg.Section("Sim.0").NewKey("# Sim", "test")
	cfg.WriteTo(os.Stdout)

	// log.Print(cfg, err)
	log.Print(cfg.Section("Sim.0").Key("# Sim"))

	box := tview.NewBox().SetBorder(true).SetTitle("MSFS bush trips aircrafts changer")
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}
