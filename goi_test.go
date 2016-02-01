package main

import "testing"
import "strings"
import "os"

func TestDirAndEnv(t *testing.T) {
	gopath := mkTmpDir("test")
	setEnv(gopath)
	// TODO: check tmp dir exists
	if os.Getenv("GOPATH") != gopath {
		t.Errorf("GOPATH: %s, gopath: %s", os.Getenv("GOPATH"), gopath)
	}
}

func TestResolvePath(t *testing.T) {
	if resolvePath("coderhaoxin/goi") != "github.com/coderhaoxin/goi" {
		t.Errorf("resolved path: %s", resolvePath("coderhaoxin/goi"))
	}
}

func TestGetArgs(t *testing.T) {
	pkg, args := getArgs([]string{"goi", "-u", "coderhaoxin/hp"})
	if pkg != "github.com/coderhaoxin/hp" {
		t.Errorf("pkg:", pkg)
	}
	if strings.Join(args, "") != "get-ugithub.com/coderhaoxin/hp" {
		t.Errorf("args:", strings.Join(args, ""))
	}

	pkg, args = getArgs([]string{"goi", "coderhaoxin/hp", "-u"})
	if pkg != "github.com/coderhaoxin/hp" {
		t.Errorf("pkg:", pkg)
	}
	if strings.Join(args, "") != "get-ugithub.com/coderhaoxin/hp" {
		t.Errorf("args:", strings.Join(args, ""))
	}
}
