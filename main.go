package main

import (
	"log"
	"strings"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami"
	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
	"github.com/pnguyen215/gobase-voip-core/pkg/ami/utils"
	"github.com/pnguyen215/gobase-voip-core/pkg/bot/telegram"
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

	collection := make(map[string]string)

	collection["key1"] = "1"
	collection["key2"] = "2"

	log.Printf("key found = %v", utils.TakeKeyFromValue(collection, "1"))
	log.Printf("equal = %v", strings.EqualFold("value_1", "Value_1"))

	instance := &telegram.TelegramBot{
		Token:     "6097117689:AAEInfj6wwy1rOryFUQFslan36q4E0ic56o",
		ChatId:    1127479716,
		DebugMode: true,
	}

	_, err := instance.SendMessage(collection)

	if err != nil {
		log.Fatal(err)
	}
}
