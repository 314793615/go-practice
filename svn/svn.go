package svn

import (
	"bytes"
	"fmt"
	"os/exec"
)

func shell(command string) (string , error) {
	c := exec.Command("svn", command)
	var stdout, stderr bytes.Buffer
	c.Stdout, c.Stderr = &stdout, &stderr
	e := c.Run()
	if e != nil{
		fmt.Printf("happend a error: %s", e.Error())
		return "", e
	}
	return stdout.String(), nil 

}