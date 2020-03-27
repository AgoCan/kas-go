package logging

import (
	"das-go/config"
	"fmt"
	"testing"
)

func TestAppendLog(t *testing.T) {
	Init(config.LogInfoFile)
	fmt.Println(config.LogInfoFile)
	Append(InfoLevel, "this is test logging")
}