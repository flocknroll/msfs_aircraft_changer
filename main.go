package main

import (
	"log"

	"github.com/flocknroll/azga/go_installer/msfstools"
	"gopkg.in/ini.v1"
)

func main() {
	path, ok := msfstools.GetPackageFolderPath()
	log.Print(path, ok)

	cfg, _ := ini.Load("test_data/asobo-bushtrip-chile/Missions/Asobo/BushTrips/Chile/Chile.FLT")

	// log.Print(cfg, err)
	log.Print(cfg.Section("Sim.0").Key("Sim"))
}
