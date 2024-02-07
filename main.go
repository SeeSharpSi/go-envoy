package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Hello world")
	if _, err := os.Stat("./releases"); os.IsNotExist(err) {
		fmt.Println("creating ./releases dir")
		err := os.Mkdir("./releases", 0755)
		check(err)
	}
	file_name := time.Now().Format("20060102150405")
	err := os.Mkdir("./releases/"+file_name, 0755)
	check(err)
	remote := get_remote("origin")
	clone_repo(remote, file_name)
}

func get_remote(remote string) string {
	cmd := exec.Command("git", "remote", "get-url", remote)
	output, err := cmd.Output()
	check(err)
	return string(output)
}

// git clone --depth=1 --branch=master git://someserver/somerepo dirformynewrepo
// rm -rf ./dirformynewrepo/.git

func clone_repo(url string, file_name string) {
	cmd := exec.Command("git", "clone", "--depth=1", "--branch=master", url, file_name)
	output, err := cmd.Output()
	check(err)
	fmt.Println("clone_repo() | " + string(output))
	cmd = exec.Command("rm", "-rf", file_name+".git")
	output, err = cmd.Output()
	check(err)
	fmt.Println("clone_repo() | " + string(output))
}
