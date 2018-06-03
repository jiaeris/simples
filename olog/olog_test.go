package olog

import "testing"

func TestSimple(t *testing.T) {
	SetLogLevel(ALL)
	SetPrintTime(false)
	Simple()
}
