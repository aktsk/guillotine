package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
	"os"
	"os/exec"
	"time"
)

type Specification struct {
	WatchedFile string `required:"true" split_words:"true"`
	Interval    int    `default:"1"`
}

func main() {
	var s Specification

	if err := envconfig.Process("guillotine", &s); err != nil {
		log.Fatal(err.Error())
		return
	}

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: cmd args...")
		return
	}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	fmt.Println("Watch File:", s.WatchedFile)
	if err := cmd.Start(); err != nil {
		return
	}

	go func() {
		for {
			time.Sleep(time.Duration(s.Interval) * time.Second)
			if _, err := os.Stat(s.WatchedFile); err == nil {
				cmd.Process.Kill()
				fmt.Println(s.WatchedFile, "exists")
				break
			}
		}
	}()

	cmd.Wait()
}
