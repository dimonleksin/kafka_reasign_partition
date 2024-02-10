package main

import (
	"log"
	"os"

	"github.com/dimonleksin/kafka_reasign_partition/actions"
	"github.com/dimonleksin/kafka_reasign_partition/pkg"
)

func main() {
	settings := pkg.Settings{}
	err := settings.GetSettings()
	if err != nil {
		log.Fatal(err)
	}
	// log.Println(*settings.H, *settings.Help)
	if *settings.H || *settings.Help {
		pkg.PrintHelp()
		os.Exit(0)
	}

	switch *settings.Action {
	case "rebalance":
		actions.Reasign(settings)
	case "move":
		if len(*settings.Topic) != 0 {
			err := actions.MoveTopic(settings)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
