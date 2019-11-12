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

func executeBatch(minP, maxP int) {
	//fmt.Println(testPortConnection(os.Args[1], "22"))
	maxPort := maxP
	minPort := minP
	var wg sync.WaitGroup
	wg.Add((maxPort - minPort + 1))
	for run := minPort; run <= maxPort; run++ {
		go func(host, port string, wg *sync.WaitGroup) {
			if testPortConnection(host, port) == true {
				fmt.Println(("Port " + port + " OPEN"))
			} else {
				fmt.Println(("Port " + port + " CLOSED"))
			}
			wg.Done()
		}(os.Args[1], strconv.FormatInt(int64(run), 10), &wg)
		//testPortConnection(os.Args[1], strconv.FormatInt(int64(run), 10))
	}
	wg.Wait()
}

func execute() {
	batchSize := 200
	for run := 1; run < 65536; run++ {
		higherEnd := run + batchSize
		if (higherEnd) > 65535 {
			higherEnd = 65535
		}
		executeBatch(run, higherEnd)
		run = run + batchSize
	}
}

func testPortConnection(host, port string) bool {
	//fmt.Println("Testing connection...." + port)
	out, err := exec.Command("timeout", "5", "nc", "-z", "-v", host, port).Output()
	if err != nil {
		//fmt.Println("Coming in error for PORT " + port)
		//fmt.Printf("%s", err)
		if strings.HasSuffix(string(err.Error()), "exit status 1") {
			return true
		}
		if strings.HasSuffix(string(err.Error()), "exit status 124") {
			return false
		}
	}
	fmt.Println(string(out))
	return true
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		execute()
	}
}
