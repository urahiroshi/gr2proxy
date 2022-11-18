package main

import (
	"log"

	"github.com/urahiroshi/gr2proxy/config"
	"github.com/urahiroshi/gr2proxy/record"
	"go.uber.org/zap/zapcore"
)

func main() {
	logger, err := config.NewLogger(zapcore.DebugLevel)
	if err != nil {
		log.Fatal(err.Error())
	}
	c := &record.RecordConfig{
		Config: config.Config{
			LocalPort: 50051,
			AdminPort: 50151,
			Logger:    logger,
		},
		RemoteUrl: "http://localhost:50052",
	}
	record.Run(c)
}
