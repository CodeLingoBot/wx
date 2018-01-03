package main

import (
	"github.com/morya/utils/log"
	"testing"
)

func TestOK(t *testing.T) {
	v := GetLastStatus()
	log.Infof(v)
}
