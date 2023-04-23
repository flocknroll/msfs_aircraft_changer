package main

import (
	"log"

	"github.com/flocknroll/azga/go_installer/msfstools"
)

func main() {
	path, ok := msfstools.GetPackageFolderPath()

	log.Print(path, ok)
}
