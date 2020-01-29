package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func readDir(dir string) (map[string]string, error) {
	filesInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	envMap := make(map[string]string, len(filesInfo))
	for _, fileInfo := range filesInfo {
		if fileInfo.IsDir() {
			continue
		}

		file, err := os.Open(filepath.Join(dir, fileInfo.Name()))
		if err != nil {
			return nil, err
		}

		reader := bufio.NewReader(file)
		value, _, err := reader.ReadLine()
		if err != nil {
			return nil, err
		}

		envMap[fileInfo.Name()] = strings.TrimSpace(string(value))

		err = file.Close()
		if err != nil {
			return nil, err
		}
	}

	return envMap, nil
}

func runCmd(cmdArgs []string, env map[string]string) int {
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)

	for key, value := range env {
		arr := []string{key, value}
		v := strings.Join(arr, "=")
		cmd.Env = append(cmd.Env, v)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
	return cmd.ProcessState.ExitCode()
}

func main() {
	if len(os.Args) < 3 {
		println("Args count is incorrect, must be: path_to_dir command arg1 arg2 arg3")
		os.Exit(-1)
	}

	env, err := readDir(os.Args[1])
	if err != nil {
		os.Exit(111)
	}

	exitCode := runCmd(os.Args[2:], env)
	os.Exit(exitCode)
}
