package main

import (
	"testing"
)

func TestMsgs(*testing.T) {
	i := int32(11)
	btn := &ButtonMessage{
		Building: &i,
	}
	btnMsgs = append(btnMsgs, btn)
}
