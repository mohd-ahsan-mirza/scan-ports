# Description
I built this project mostly to try out Go lang cocurrency feature. Due to go lang multithreading capabilites I am able to scan `65535` ports in about 30 minutes
# Prerequisites
Go lang installed
# Usage
* Clone Repo
* ``` go install scan-ports ```
* ``` go run main.go > output.txt ```
* ``` cat output.txt | grep OPEN ```
