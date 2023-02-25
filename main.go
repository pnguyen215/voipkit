package main

import (
	"log"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami"
	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
	"github.com/pnguyen215/gobase-voip-core/pkg/ami/utils"
)

func main() {
	next := ami.NewDictionary()

	log.Printf("len begin = %v", next.Length())
	next.AddKeyTranslator("test001", "test001").AddKeyTranslator("test002", "test002")
	log.Printf("dictionaries begin = %v", next.Json())
	data := *next.FindDictionariesByKey(config.AmiListenerEventCommon)
	log.Printf("dictionaries after = %v", utils.ToJson(data))
	next.Reset()
	data = *next.FindDictionariesByKey(config.AmiListenerEventCommon)
	log.Printf("dictionaries reset = %v", utils.ToJson(data))
}
