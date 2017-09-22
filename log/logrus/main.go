package main

import (
	log "github.com/sirupsen/logrus"
	"time"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	writer,_ := rotatelogs.New("d:/var/log/main3.log",
		//rotatelogs.WithLinkName("d:/var/log/main3.log"),
		rotatelogs.WithMaxAge(24 * time.Hour),
		rotatelogs.WithRotationTime(time.Hour),)


	log.AddHook(lfshook.NewHook(lfshook.WriterMap{
		log.InfoLevel: writer,
		//log.ErrorLevel: writer,
	}))

	log.SetOutput(writer)
	// Only log the warning severity or above.
	//log.SetLevel(log.WarnLevel)
	log.Info("A walrus appears")
}
