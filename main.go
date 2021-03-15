package main

import (
	"fmt"

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
	processes, errs := getGolangProcesses()
	if len(errs) > 0 {
		// you want to handle errors here.
		// the most common error you'll see are :
		// - "not a Go executable" & "unrecognized executable format"
		// i guess those can be filtered out
	}

	for _, process := range processes {
		fmt.Println(process)
	}
}
