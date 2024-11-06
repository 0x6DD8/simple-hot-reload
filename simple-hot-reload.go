package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	var args = os.Args
	if len(args) < 3 {
		var exeName = filepath.Base(args[0])
		fmt.Println("Usage:", exeName, "<path-to-watch> <program-to-run> [args...]")
		return
	}

	var pathToWatch = args[1]
	var programToRun = args[2]
	var programToRunArgs = args[3:] //also pass the args to the program to run

	fmt.Println("Watching path:", pathToWatch)
	fmt.Println("Program to run:", programToRun, "", strings.Join(programToRunArgs, ""))

	var cmd = runCommand(programToRun, programToRunArgs)

	watchForChanges(pathToWatch, func() {
		var currentFormattedTime = time.Now().Format("2006-01-02 15:04:05")
		fmt.Println(" [", currentFormattedTime, "]", "🔥 File changed, initiating hot reload...")
		cmd.Process.Kill()
		cmd = runCommand(programToRun, programToRunArgs)
	})
}

func runCommand(command string, args []string) *exec.Cmd {
	var cmd = exec.Command(command, args...)
	//also pass the stdout and stderr to the parent process
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	return cmd
}

func watchForChanges(rootPath string, onChange func()) error {
	var initialStats, err = getAllFileStats(rootPath)
	if err != nil {
		return err
	}

	for {
		// sleep for a second before checking again
		// without it, it will like consume all the cpu lol
		time.Sleep(1 * time.Second)

		var currentStats, err = getAllFileStats(rootPath)
		if err != nil {
			return err
		}

		if checkIfSomethingHasChanged(initialStats, currentStats) {
			onChange()
			initialStats = currentStats
		}
	}
}

func getAllFileStats(rootPath string) (map[string]os.FileInfo, error) {
	var stats = make(map[string]os.FileInfo)

	// WalkDir should be somewhat more performant than Walk? i havent tested it tho
	var err = filepath.WalkDir(rootPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			stats[path], err = d.Info()
		}
		return nil
	})

	return stats, err
}

func checkIfSomethingHasChanged(initialStats, currentStats map[string]os.FileInfo) bool {
	if len(initialStats) != len(currentStats) {
		return true
	}

	for path, initialStat := range initialStats {
		var currentStat = currentStats[path]

		// wont mod time be enough? if the file size changes, the mod time will change too...
		// not sure if this is the best way to check for changes
		if currentStat.ModTime() != initialStat.ModTime() || currentStat.Size() != initialStat.Size() {
			return true
		}
	}
	return false
}
