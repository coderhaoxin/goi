package main

import "testing"

func TestUtils(t *testing.T) {
	gopath := mkTmpDir("test")
	setEnv(gopath)
}
