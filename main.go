package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/flocknroll/azga/go_installer/msfstools"
	"github.com/rivo/tview"
	"gopkg.in/ini.v1"
)

func main() {
	path, ok := msfstools.GetPackageFolderPath()
	log.Print(path, ok)

	bushtrips := make([]string, 20)
	aircrafts := make([]string, 20)

	filepath.WalkDir("test_data", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, ".FLT") {
			bushtrips = append(bushtrips, path)
		}
		if strings.HasSuffix(path, "aircraft.cfg") {
			aircrafts = append(aircrafts, path)
		}

		return nil
	})

	fmt.Println(bushtrips, aircrafts)

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
