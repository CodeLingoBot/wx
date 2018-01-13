package main

import (
	"github.com/morya/utils/log"
	"testing"
)

func TestValidUser(t *testing.T) {
    var users = []string{
    "o-A1l0zVAgf51kqtZY-oyFGUBi6Y",
    "o-A1l03mCLRjTP09Z6UZdOVLUBLs",
    }
    for _, u := range users {
        if !IsValidUser(u) {
            t.Fatal(u)
        }
    }
}
