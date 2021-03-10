package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

const bufferSize = 1024

func FileExist(path string) bool  {
	stat, err := os.Lstat(path)
	if stat.IsDir() {
		return false
	}
	return !os.IsNotExist(err)
}

func UserInput() (logFile string){
	flag.StringVar(&logFile,"f","","Please enter the log file to view")

	flag.Usage = func() {
		fmt.Printf("Usage: %v -f  logFile \n",os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	return
}

func LogFileView(logFile string){
    file,err := os.Open(logFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	content := make([]byte, bufferSize)
	for {
		n, err := file.Read(content)
		if err != nil {
			if err == io.EOF {
				//等待1s
				time.Sleep(time.Second * 1)
				continue
			} else {
				fmt.Println(err)
				break
			}
		} else {
			fmt.Print(string(content[:n]))
		}
	}
}

func main()  {
	logFile := UserInput()
	if logFile == "" {
		flag.Usage()
		return
	}
	if FileExist(logFile) {
		LogFileView(logFile)
	}
}