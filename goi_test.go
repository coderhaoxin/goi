package main

import "testing"

func TestUtils(t *testing.T) {
	gopath := mkTmpDir("test")
	setEnv(gopath)
}

func TestResolvePath(t *testing.T) {
	if resolvePath("pkg4go/goi") != "github.com/pkg4go/goi" {
		t.Fail()
	}
}
