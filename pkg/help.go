package pkg

import "fmt"

func PrintHelp() {
	textHelp := "--bootstrap-server [string] for set addres of brokers\n" +
		"\t format: host:port, like 127.0.0.1:9092\n\n" +
		"--topic [string/[]string] for set topics name for move\n" +
		"\tif u send some topics - separator ','\n\n" +
		"--action [string] Set action of u needed (move/return/rebalance)\n\n" +
		"--user [string] set username, if u dont set this arg, used PLAINTEXT\n" +
		"\tif set --user, u need set and --password\n\n" +
		"--password [string] set password for connect to kafka\n" +
		"\tif u set password without --user, this call panic\n\n" +
		"--mechanism [string] scram-sha-256 or scram-sha-512\n" +
		"\tdefining only if you use auth with loggin and password\n\n" +
		"--tls define if need use TLS\n\n" +
		"--ca [string] path to CA file\n\n" +
		"--cert [string] path to cert file\n" +
		"\tif you define --cert, dont forget defined --key\n\n" +
		"--key [string] path to key file\n\n" +
		"--to [int] set brokers ids for desctination brokers (sep ','). For example 1,2,3\n\n" +
		"--from [int] set source broker id\n\n" +
		"--treads [int] seted number of treads for reassign\n\n" +
		"--sync sync/async work with topic. If true - krpg wait when all replicas for partition moved in desired state\n" +
		"\tif --treads > 1 then anyone thed wait mine topic\n\n" +
		"--api-version set version of brokers in format 2.1.0 (default)\n\n" +
		"--version print version from this build\n\n" +
		"-h or --help for print this help"
	fmt.Println(textHelp)
}
