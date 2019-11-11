package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

func execute() {
	//fmt.Println(testPortConnection(os.Args[1], "22"))
	maxPort := 100
	minPort := 1
	var wg sync.WaitGroup
	wg.Add((maxPort - minPort + 1))
	for run := minPort; run <= maxPort; run++ {
		go func(host, port string) string {
			if testPortConnection(host, port) == true {
				fmt.Println(("Port " + port + " OPEN"))
				wg.Done()
				return ("Port " + port + " OPEN")
			}
			fmt.Println(("Port " + port + " CLOSED"))
			wg.Done()
			return ("Port " + port + " CLOSED")
		}(os.Args[1], strconv.FormatInt(int64(run), 10))
		//testPortConnection(os.Args[1], strconv.FormatInt(int64(run), 10))
	}
	wg.Wait()
}

func testPortConnection(host, port string) bool {
	fmt.Println("Testing connection...." + port)
	//out, err := exec.Command("ls").Output()
	out, err := exec.Command("timeout", "1", "telnet", host, port).Output()
	if err != nil {
		//fmt.Println("Coming in error for PORT " + port)
		//fmt.Printf("%s", err)
	}
	fmt.Println(string(out))
	if strings.HasSuffix(string(err.Error()), "exit status 1") {
		return true
	}
	return false
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		execute()
	}
}
