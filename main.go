package main

import (
	"context"
	"log"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami"
)

func main() {

	_, err := ami.NewAMISocketWith(context.Background(), "27.71.226.238:5038")

	if err != nil {
		log.Fatalf(err.Error())
	}

}
