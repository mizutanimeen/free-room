package logs

import (
	"io"
	"log"
	"os"
)

func Settings(aFilename string) (*os.File, error) {
	tLogFile, tError := os.OpenFile(aFilename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if tError != nil {
		return nil, tError
	}
	tMultiLogFile := io.MultiWriter(os.Stdout, tLogFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(tMultiLogFile)

	return tLogFile, nil
}
