package main

import (
	"log"
	"os"

	ps "github.com/keybase/go-ps"

	goversion "rsc.io/goversion/version"
)

func isGo(pr ps.Process) (ok bool, err error) {
	if pr.Pid() == 0 {
		return
	}
	path, err := pr.Path()
	if err != nil {
		return false, err
	}
	_, err = goversion.ReadExe(path)
	return err == nil, err
}

func getGolangProcesses() ([]ps.Process, []error) {
	var (
		errGroup  []error
		processes []ps.Process
	)

	processList, err := ps.Processes()
	if err != nil {
		return processes, errGroup
	}

	for _, process := range processList {
		ok, err := isGo(process)
		if err != nil {
			errGroup = append(errGroup, err)
			continue
		} else if ok {
			processes = append(processes, process)
		}
	}

	return processes, errGroup
}

func main() {

	whitelist := map[string]bool{
		"gopls": true,
		"go":    true,
	}

	processes, errs := getGolangProcesses()

	if len(errs) > 0 {
		// you want to handle errors here.
		// the most common errors you'll see are :
		// - "not a Go executable" & "unrecognized executable format"
		// i guess those can be filtered out cos they're not really errors
	}

	for _, process := range processes {
		// a whitelist of programs i don't want to kill
		// feel free to modify it (at your own risk)
		_, ok := whitelist[process.Executable()]
		if ok {
			continue
		}

		// if this isn't obvious, terminating the process running this program is a bit conter-intuitive.
		// if you really want to stay true to the cause, you can terminate yourself
		// when you're done terminating everybody else. (after the for-loop) 😏
		if os.Getpid() == process.Pid() {
			continue
		}

		// there were just soo many ways to kill a process, so i decided to go with this approach.
		// it looks clumsy to me though, but it works fine.
		os_process, err := os.FindProcess(process.Pid())
		if err != nil {
			log.Println(err)
			continue
		}

		if err = os_process.Signal(os.Kill); err != nil {
			log.Println(err)
			continue
		}
	}

	log.Println("extermination complete.")
}
