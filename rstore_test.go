package main

import (
	"github.com/morya/utils/log"
	"testing"
)

func TestOK(t *testing.T) {
    var symbols = []string{
        "eos_usdt",
        "ltc_usdt",
    }
    for _, s := range symbols {
        v := GetLastStatus(s)
        log.Infof(v)
    }
}
