package main

import (
	log "github.com/sirupsen/logrus"
)

func main()  {
	log.SetFormatter(&log.JSONFormatter{})

	standardFields := log.Fields{
		"host": "srv42",
	}

	hlog := log.WithFields(standardFields)



	hlog.WithFields(log.Fields{"uid": 100500}).Info("file successfully uploaded")
	hlog.WithFields(log.Fields{"uid": 200512}).Warn("libjpeg: invalid format")
	hlog.WithFields(log.Fields{"uid": 101345}).Error("file corrupted")
	hlog.Infof("storage space left: %d", 1024)
}
