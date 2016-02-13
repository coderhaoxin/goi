package main

import . "github.com/pkg4go/assert"
import "testing"
import "strings"
import "os"

func TestDirAndEnv(t *testing.T) {
	a := A{t}

	gopath := mkTmpDir("test")
	setEnv(gopath)
	// TODO: check tmp dir exists
	a.Equal(os.Getenv("GOPATH"), gopath)
}

func TestResolvePath(t *testing.T) {
	a := A{t}
	a.Equal(resolvePath("coderhaoxin/goi"), "github.com/coderhaoxin/goi")
}

func TestGetArgs(t *testing.T) {
	a := A{t}

	pkg, args := getArgs([]string{"goi", "coderhaoxin/hp"})
	a.Equal(len(args), 2)
	a.Equal(pkg, "github.com/coderhaoxin/hp")
	a.Equal(strings.Join(args, ""), "getgithub.com/coderhaoxin/hp")

	pkg, args = getArgs([]string{"goi", "-u", "coderhaoxin/hp"})
	a.Equal(len(args), 3)
	a.Equal(pkg, "github.com/coderhaoxin/hp")
	a.Equal(strings.Join(args, ""), "get-ugithub.com/coderhaoxin/hp")

	pkg, args = getArgs([]string{"goi", "coderhaoxin/hp", "-u"})
	a.Equal(len(args), 3)
	a.Equal(pkg, "github.com/coderhaoxin/hp")
	a.Equal(strings.Join(args, ""), "get-ugithub.com/coderhaoxin/hp")
}
