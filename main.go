package main

import (
	"log"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami"
)

func main() {
	log.Printf("translate key found = %v", ami.NewDictionary().TranslateField("Linkedid"))
}
