package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	pid := os.Getpid()
	argError := flag.Bool("error", false, "Should the process exit with error")
	argRepeat := flag.Int("repeat", 0, "How many times should it run (0 means infinite)")
	argCode := flag.Int("code", 0, "Exit code of the process")
	argDebug := flag.Bool("debug", true, "Print extra debug info")
	argSleep := flag.Int("sleep", 1, "How many seconds the process should sleep after each loop")
	flag.Parse()
	argMessage := strings.Join(flag.Args(), " ")
	if len(argMessage) < 1 {
		argMessage = "test message goes here"
	}

	if *argDebug {
		fmt.Println("pid:", pid)
		fmt.Println("error:", *argError)
		fmt.Println("repeat:", *argRepeat)
		fmt.Println("code:", *argCode)
		fmt.Println("message:", argMessage)
	}

	for i := 0; i < *argRepeat || *argRepeat < 1; i++ {
		fmt.Println(pid, "#", i, ":", argMessage)
		time.Sleep(time.Second * time.Duration(*argSleep))
	}

	fmt.Println("done, exiting with code", *argCode)

	exitCode := *argCode
	if *argError == true && exitCode == 0 {
		exitCode = 255
	}
	os.Exit(exitCode)
}
