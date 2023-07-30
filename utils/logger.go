package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func InitLogger() {
	// formatter := &log.JSONFormatter{
	// 	FieldMap: log.FieldMap{
	// 		log.FieldKeyTime:  "timestamp",
	// 		log.FieldKeyLevel: "level",
	// 		log.FieldKeyMsg:   "message",
	// 		log.FieldKeyFunc:  "caller",
	// 	},
	// }

	formatter := &log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
		FieldMap: log.FieldMap{
			log.FieldKeyTime:  "timestamp",
			log.FieldKeyLevel: "level",
			log.FieldKeyMsg:   "message",
			log.FieldKeyFunc:  "caller",
		},
	}

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(formatter)

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Log everything
	log.SetLevel(log.TraceLevel)
}
