package main

import (
	"log"
	"strconv"
	"strings"
	"testing"
)

// TestSendStats sends the stats for one light bulb
func TestSplitID(t *testing.T) {
	id := "1-1-back"
	s := strings.Split(id, "-")
	r, _ := strconv.Atoi(s[0])
	b, _ := strconv.Atoi(s[1])
	if r != 1 || b != 1 || s[2] != "back" {
		log.Fatal("test failed")
	}
}
