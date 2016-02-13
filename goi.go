package main

import "github.com/pkg4go/convert"
import "github.com/pkg4go/execx"
import "strings"
import "time"
import "path"
import "fmt"
import "os"

func main() {
	pkg, args := getArgs(os.Args)

	gopath := mkTmpDir(path.Base(pkg))

	setEnv(gopath)
	goget(args)

	err := os.RemoveAll(gopath)
	if err != nil {
		panic(err)
	}
}

func mkTmpDir(name string) string {
	tmp := os.TempDir()
	s := convert.String(time.Now().UnixNano())
	tmpdir := path.Join(tmp, "goi-"+name+"-"+s)
	err := os.Mkdir(tmpdir, 0700)
	if err != nil {
		// TODO: if os.IsExist(err), retry
		panic(err)
	}
	return tmpdir
}

func setEnv(gopath string) {
	gobin := os.Getenv("GOBIN")

	err := os.Setenv("GOPATH", gopath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("GOBIN : %s \nGOPATH: %s \n", gobin, gopath)
}

func getArgs(cmdArgs []string) (pkg string, args []string) {
	flags := make([]string, 0)
	args = []string{"get"}

	for _, v := range cmdArgs[1:] {
		if strings.HasPrefix(v, "-") {
			flags = append(flags, v)
		} else {
			pkg = resolvePath(v)
		}
	}

	if len(flags) != 0 {
		args = append(args, flags[:]...)
	}
	args = append(args, pkg)

	return
}

func goget(args []string) {
	out, err := execx.Run("go", args[:]...)
	if err != nil {
		panic(err)
	}

	fmt.Println(out)
}

func resolvePath(pkgParh string) string {
	if strings.Count(pkgParh, "/") == 1 {
		return path.Join("github.com", pkgParh)
	}
	return pkgParh
}
