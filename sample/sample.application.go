package main

import (
	"github.com/markdicksonjr/nibbler"
	"github.com/markdicksonjr/nibbler-redis"
	"log"
	"time"
)

func main() {

	// allocate logger and configuration
	logger := nibbler.DefaultLogger{}
	config, err := nibbler.LoadConfiguration()
	nibbler.LogFatalNonNil(logger, err)

	// prepare extensions for initialization
	redisExtension := redis.Extension{}

	// initialize the application, provide config, logger, extensions
	app := nibbler.Application{}
	nibbler.LogFatalNonNil(logger, app.Init(config, logger, []nibbler.Extension{
		&redisExtension,
	}))

	cmd := redisExtension.Client.Set("test", "sd", time.Minute)
	if cmd.Err() != nil {
		log.Fatal(cmd.Err())
	}

	strCmd := redisExtension.Client.Get("test")
	nibbler.LogFatalNonNil(logger, strCmd.Err())

	log.Println(strCmd.Val() == "sd")

	nibbler.LogFatalNonNil(logger, app.Run())
}
