package main

import "testing"

func TestGoodFunc(t *testing.T) {

	if goodFunc() != "looks good" {
		t.Error("ERROR")
	}
}

func TestBadFunc(t *testing.T) {

	if badFunc() != "looks good" {
		t.Error("ERROR")
	}
}
