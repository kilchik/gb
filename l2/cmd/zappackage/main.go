package main

import (
	"go.uber.org/zap"
	"strconv"
)

func main()  {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger = logger.With(zap.String("host", "srv42"))

	logger.With(zap.Uint64("uid", 100500)).Info("file successfully uploaded")
	logger.With(zap.Uint64("uid", 200512)).Warn("libjpeg: invalid format")
	logger.With(zap.Uint64("uid", 101345)).Error("file corrupted")
	logger.Info("storage space left: " + strconv.Itoa(1024))
}
