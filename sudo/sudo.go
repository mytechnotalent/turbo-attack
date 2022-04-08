// Terms Of Use
// ------------
// Do NOT use this on any computer you do not own or are not allowed to run this on.
// You may NEVER attempt to sell this, it is free and open source.
// The authors and publishers assume no responsibility.
// For educational purposes only.

// Package sudo ensures application has proper permissions.
package sudo

import (
	"errors"
	"os/exec"
	"strconv"
)

// Check ensures application is run as root.
// It will return an error if one occurred.
func Check(i int) error {
	cmd := exec.Command("id", "-u")
	output, _ := cmd.Output()
	j, _ := strconv.Atoi(string(output[:len(output)-1])) // remove the \n
	// 0 = root, 501 = non-root user
	if i != j {
		return errors.New("application will only run as root (sudo)")
	}
	return nil
}
