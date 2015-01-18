package main

import "github.com/pkg4go/pkgs/command"
import "github.com/pkg4go/pkgs/convert"
import "time"
import "path"
import "fmt"
import "os"

func main() {
	pkg := os.Args[1]
	name := path.Base(pkg)
	gopath := mkTmpDir(name)
	setEnv(gopath)
	get(pkg)
	err := os.RemoveAll(gopath)
	if err != nil {
		panic(err)
	}
}

func mkTmpDir(name string) string {
	tmp := os.TempDir()
	s, _ := convert.String(time.Now().UnixNano())
	tmpdir := path.Join(tmp, "goi-"+name+"-"+s)
	err := os.Mkdir(tmpdir, 0700)
	if err != nil {
		// TODO: if os.IsExist(err), retry
		panic(err)
	}
	return tmpdir
}

func rmTmpDir() {

}

func setEnv(gopath string) {
	gobin := os.Getenv("GOBIN")

	err := os.Setenv("GOPATH", gopath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("GOBIN : %s \nGOPATH: %s \n", gobin, gopath)
}

func get(pkg string) {
	out, err := command.Run("go", "get", pkg)
	if err != nil {
		panic(err)
	}

	fmt.Println(out)
}
