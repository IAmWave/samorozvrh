// An interface to the Python solver.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
)

const SOLVER_COMMAND = "../env/bin/samorozvrh_solver --time-limit %d"

func Solve(query []byte, timeLimit int) ([]byte, error) {
	// Create a temporary file with the query, feed it to the solver
	// and run it

	tempfile, err := createQueryFile(query)
	if err != nil {
		return nil, err
	}
	defer os.Remove(tempfile.Name()) // clean up

	commandParts := strings.Split(fmt.Sprintf(SOLVER_COMMAND, timeLimit)+" "+tempfile.Name(), " ")
	subProcess := exec.Command(commandParts[0], commandParts[1:]...)
	// Run relative to the directory given by -rootdir
	subProcess.Dir = path.Join(rootDir, "solver")

	res, err := subProcess.CombinedOutput()
	if err != nil {
		err2 := fmt.Errorf(strings.Join([]string{err.Error(), string(res)}, ";"))
		return nil, err2
	}
	return res, nil
}

func createQueryFile(query []byte) (*os.File, error) {
	tempfile, err := ioutil.TempFile("", "query")
	if err != nil {
		return nil, err
	}
	if _, err := tempfile.Write(query); err != nil {
		return nil, err
	}
	if err := tempfile.Close(); err != nil {
		return nil, err
	}
	return tempfile, nil
}
