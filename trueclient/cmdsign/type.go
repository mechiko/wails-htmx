package cmdsign

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const modError = "cmdexec"

var paths = []string{
	`C:\Program Files (x86)\Crypto Pro\CSP`,
	`C:\Program Files\Crypto Pro\CSP`,
}

var FileIn, FileOut string

type cmdexec struct {
	// *exec.Cmd
	FileIn  string
	FileOut string
	Hash    string
	Command []string
}

func init() {
	path := os.Getenv("PATH")
	path = strings.Join(paths, ";") + ";" + path
	os.Setenv("PATH", path)
	fileIn, err := os.CreateTemp("", "lite_")
	if err != nil {
		panic(fmt.Sprintf("error create IN temp file %s", err.Error()))
	}
	FileIn = fileIn.Name()
	defer fileIn.Close()
	fileOut, err := os.CreateTemp("", "lite_")
	if err != nil {
		panic(fmt.Sprintf("error create IN temp file %s", err.Error()))
	}
	defer fileOut.Close()
	FileOut = fileOut.Name()
}

func New(hash string) *cmdexec {
	// fmt.Println("IN file: ", fileIn.Name())
	command := strings.Split(fmt.Sprintf("-sfsign -sign -in %s -out %s -my %s -base64 -add -addsigtime -cades_strict -verify -alg GOST12_256", FileIn, FileOut, hash), " ")
	return &cmdexec{
		FileIn:  FileIn,
		FileOut: FileOut,
		Command: command,
		Hash:    hash,
	}
}

func (c *cmdexec) Sign(in string) (string, error) {
	err := os.WriteFile(c.FileIn, []byte(in), os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("error write IN temp file %s", err.Error()))
	}
	cmd := exec.Command("csptest.exe", c.Command...)
	if out, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("%s %s Sign() %w", modError, out, err)
	}
	// if err := cmd.Run(); err != nil {
	// 	return "", fmt.Errorf("%s Sign() %w", modError, err)
	// }
	sign, err := os.ReadFile(c.FileOut)
	if err != nil {
		return "", fmt.Errorf("%s Sifn() %w", modError, err)
	}
	clearSign := bytes.ReplaceAll(sign, []byte("\r\n"), []byte{})
	// FilesRemove()
	return string(clearSign), nil
}

func FilesRemove() {
	os.Remove(FileIn)
	os.Remove(FileOut)
}
