package main

/**
 * @author Oliver Kelton, oakelton@gmail.com
 * @date May 28, 2019
 * Benchmarks programs from the command line.
 */

import (
	"os"
	"fmt"
	"log"
	"time"
	"strings"
)

func GetExecutableName (name string) (string, error) {
	path, exists := os.LookupEnv("PATH")
	if !exists { log.Fatal("PATH environment var not found") }
	paths := strings.Split(path, ":")
	paths = append([]string{ "." }, paths...)
	var (
		filename string
		err error
	)
	for i := range paths {
		filename = paths[i] + "/" + name
		_, err = os.Stat(filename)
		if err == nil { return filename, nil }
	}//-- end for range paths
	return "", fmt.Errorf("'%s' not found", name)
}//-- end func GetExecutableName

func main () {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr,
			"\x1b[31mbenchmark requires at least one arg\x1b[0m\n")
		os.Exit(1)
	}
	procName, err := GetExecutableName(os.Args[1])
	if err != nil { log.Fatal(err) }
	fmt.Fprintf(os.Stderr, "\x1b[32mBenchmarking %s...\x1b[0m\n", procName)
	attribs := os.ProcAttr{
		Files: []*os.File{ os.Stdin, os.Stdout, os.Stderr }}
	start := time.Now()
	proc, err := os.StartProcess(procName, os.Args[1:], &attribs)
	if err != nil { log.Fatal(err) }
	_, err = proc.Wait()
	if err != nil { log.Fatal(err) }
	end := time.Now()
	dur := end.Sub(start)
	fmt.Fprintf(os.Stderr, "\x1b[32mTime taken: %v\x1b[0m\n", dur)
	os.Exit(0)
}//-- end func main

