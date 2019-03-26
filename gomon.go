package main

import (
	"os"
	"os/exec"
	"os/signal"
	"fmt"
	"log"
	"syscall"
	"github.com/fsnotify/fsnotify"
)

func main()  {
	if len(os.Args) < 2 {
		log.Fatal("Wtf les amis pas de fichier")
	}

	path := os.Args[1]

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}

	defer watcher.Close()
	defer fmt.Println("")

	done := make(chan bool)
	signalChan := make(chan os.Signal, 1)

	signal.Notify(signalChan, os.Interrupt)

	var pgid int
	pgid = runCmd(path, true)

	go func () {
		<-signalChan
		syscall.Kill(-pgid, 15)
		done <- true
	}()

	go func() {
		for {
			select {
				case event, ok := <-watcher.Events:
					if !ok {
						return
					}
					
					if event.Op&fsnotify.Write == fsnotify.Write {
						syscall.Kill(-pgid, 15)
						pgid = runCmd(path, false)
					}
				case err, ok := <-watcher.Errors:
					if !ok {
						return
					}
					fmt.Println("Error happened 😢", err)
			}
		}
	}()

	err = watcher.Add(path)
	if err != nil {
		panic(err)
	}

	<- done
}

func runCmd(path string, first bool) (pgid int) {
	cmd := exec.Command("go", "run", path)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	cmd.Start()

	if first {
		fmt.Println("Starting")
	} else {
		fmt.Println("Restarting due to change")
	}
	id, _ := syscall.Getpgid(cmd.Process.Pid)

	return id
}