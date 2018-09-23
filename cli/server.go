package main

import (
	"flag"
	"os"

	"github.com/tcoupin/rok4go/objects"
	"github.com/tcoupin/rok4go/objects/backend"
	"github.com/tcoupin/rok4go/server"
	"github.com/tcoupin/rok4go/utils/cli"
	"github.com/tcoupin/rok4go/utils/log"
)

func main() {

	listen := cli.FlagListen()
	mongourl := cli.FlagMongoDB()
	flag.Parse()

	log.SetLevel(log.LEVEL_TRACE)

	log.DEBUG("Create backend storage")
	var be objects.Backend

	be = &backend.MongoDB{}
	err := be.Init(*mongourl)
	if err != nil {
		log.ERROR("Error init backend storage: %v", err)
		os.Exit(1)
	}

	config := objects.Config{}
	config.SetBackend(be)

	log.INFO("Starting Rok4go!")
	svr := server.NewServer(*listen, &config)
	log.INFO("Listen %s", *listen)
	log.ERROR("%v", svr.ListenAndServe())
	log.INFO("Rok4go is shuting down")

}
