package main

import (
	"flag"
	"sample_server/init/cmd"
)

var configPathFlag = flag.String("config", "./config.yaml", "path to config.yaml file")

func main() {
	flag.Parse()
	cmd.NewCmd(*configPathFlag)
}
