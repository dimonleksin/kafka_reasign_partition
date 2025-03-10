package actions

import (
	"fmt"
	"log"

	"github.com/dimonleksin/kafka_reasign_partition/pkg"
)

func Reasign(settings pkg.Settings) {
	var (
		err            error
		RebalancePlane pkg.Cluster
		numberOfTopics int
		responce       string
	)

	admin, err := settings.Conf()
	if err != nil {
		log.Fatal(err)
	}
	defer admin.Close()

	r := pkg.Cluster{}
	err = r.GetNumberOfBrokers(admin)
	if err != nil {
		log.Fatal(err)
	}
	err = r.GetCurrentBalance(admin, *settings.MoveSetting.From)
	if err != nil {
		log.Fatal(err)
	}
	r.CreateBackup()
	if len(settings.MoveSetting.To) == 0 {
		RebalancePlane, numberOfTopics, err = r.CreateRebalancePlane(nil)
	} else {
		RebalancePlane, numberOfTopics, err = r.CreateRebalancePlane(settings.MoveSetting.To)
	}
	if err != nil {
		log.Fatal(err)
	}
	r = RebalancePlane

	fmt.Println()

	fmt.Println("# # # # # # # Assign after rebalance # # # # # # #")
	fmt.Println(pkg.MakeTable(r.Brokers))
	fmt.Println("# # # # # # # # # # # # # # # # # #  # # # # # # #")
	fmt.Print("\n\nPlane to reassign. Are you sure?[y/n]: ")
	_, err = fmt.Scan(&responce)
	if err != nil {
		log.Fatal("Error read you responce", err)
	}
	if responce == "y" {
		err = r.Rebalance(admin, numberOfTopics, *settings.MoveSetting.Treads, settings)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Print("So, goodby")
	}

}
