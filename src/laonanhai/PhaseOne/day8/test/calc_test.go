package main

import (
	"testing"
)

func Testadd(t *testing.T)  {
	m :=add(2,4)
	if m !=6{
		t.Fatalf("add(2,4) errror ,expect%d,%s",6,m)
	}
	t.Logf("test success")
}