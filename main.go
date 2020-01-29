package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ReadDir(dir string) (map[string]string, error) {
	filesInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	envMap := make(map[string]string, len(filesInfo))
	for _, fileInfo := range filesInfo {
		file, err := os.Open(filepath.Join(dir, fileInfo.Name()))
		if err != nil {
			return nil, err
		}

		reader := bufio.NewReader(file)
		value, _, err := reader.ReadLine()
		if err != nil {
			return nil, err
		}

		envMap[file.Name()] = string(value)

		err = file.Close()
		if err != nil {
			return nil, err
		}
	}

	return envMap, nil
}

func RunCmd(cmd []string, env map[string]string) int {
	return 0
}

func main() {
	env, err := ReadDir("/home/max/envdir")
	if err != nil {
		os.Exit(-1)
	}

	cmd := exec.Command("ls", "arg1", "arg2")

	for s, s2 := range env {
		arr := []string{s, s2}
		cmd.Env = strings.Join(arr, "=")
	}

	err = cmd.Run()
	if err != nil {
		os.Exit(111)
	}
}
