// Terms Of Use
// ------------
// Do NOT use this on any computer you do not own or are not allowed to run this on.
// You may NEVER attempt to sell this, it is free and open source.
// The authors and publishers assume no responsibility.
// For educational purposes only.

// Package sudo ensures application has proper permissions.
package sudo

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

// Check ensures application has proper permissions before proceeding.
func Check() {
	cmd := exec.Command("id", "-u")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	i, err := strconv.Atoi(string(output[:len(output)-1])) // remove the \n
	if err != nil {
		log.Fatal(err)
	}
	// 0 = root, 501 = non-root user
	if i != 0 {
		fmt.Println("application will only run as root (sudo)")
		os.Exit(0)
	}
}
