package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

func OsUsage() {
	fm, err := os.Stat("log.log")
	if os.IsNotExist(err) {
		fmt.Println(err)
		return
	}
	fmt.Println(fm.Name(), fm.Mode(), fm.Size())
}

func OsUsageWith() {
	file, _ := os.OpenFile("os.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	fmt.Println(file.Name())
	file.WriteString("Hello")
	file.WriteString("Hello, World")
}

func OsPathUsage() {
	dir, _ := os.Getwd()
	fmt.Println(dir, path.Base(dir))
	fmt.Println(dir, path.Dir(dir))
	parentDir := path.Dir(dir)
	fmt.Println(dir, path.Join(parentDir, "Chapter3"))
}

func OsDirUsage() {
	path, _ := os.Getwd()
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		fmt.Println("file", info.Name(), "in dir: ", path)
		return nil
	})
}

func OsExecUsage() {
	dockerPath, err := exec.LookPath("docker")
	if err != nil {
		return
	}
	fmt.Println(dockerPath)

	pwdPath, err:= exec.LookPath("pwd")
	if err!= nil {
		return
	}
	fmt.Println(pwdPath)
	
	fmt.Println("**********")

	cmd:= exec.Command("docker", "ps")
	stdout, _:=cmd.StdoutPipe()
	cmd.Start()
	opBytes, _:= ioutil.ReadAll(stdout)
	fmt.Println(cmd.Dir, cmd.Path, string(opBytes))

	fmt.Println("**********")

	pwd, _:= os.Getwd()
	cmd2 := exec.Command("ls", pwd)
	var buf bytes.Buffer
	cmd2.Stdout = &buf
	cmd2.Start()
	cmd2.Wait()
	fmt.Println(buf.String())

	fmt.Println("**********")

	cmd3 :=exec.Command("cat", "os.log")
	out, _:=cmd3.Output()
	fmt.Println(string(out))

	fmt.Println("**********")

	cmd4:=exec.Command("sh", "os.sh")
	out4, _:= cmd4.CombinedOutput()
	fmt.Println(string(out4))
}

func main() {
	OsUsage()
	fmt.Println("----------")
	OsUsageWith()
	fmt.Println("----------")
	OsPathUsage()
	fmt.Println("----------")
	OsDirUsage()
	fmt.Println("----------")
	OsExecUsage()
}
