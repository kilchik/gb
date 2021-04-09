package log

import (
	"io"
	"log"
)

var (
	I *log.Logger
	W *log.Logger
	E *log.Logger
	D *log.Logger
)

func Init(wInfo, wWarn, wErr, wDbg io.Writer)  {
	I = log.New(wInfo, "[I] ", log.Ldate | log.Ltime)
	W = log.New(wWarn, "[W] ", log.Ldate | log.Ltime)
	E = log.New(wErr, "[E] ", log.Ldate | log.Ltime)
	D = log.New(wDbg, "[D] ", log.Ldate | log.Ltime | log.Lshortfile)
}
