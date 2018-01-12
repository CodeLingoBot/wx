package main

import (
	"github.com/morya/utils/log"
	"testing"
)

func TestOK(t *testing.T) {
	v := GetLastStatus("eos_usdt")
	log.Infof(v)
}
