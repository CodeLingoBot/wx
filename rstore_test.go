package main

import (
	"testing"

	"github.com/morya/utils/log"
)

func TestOK(t *testing.T) {
	v := GetLastStatus()
	log.Infof(v)
}
