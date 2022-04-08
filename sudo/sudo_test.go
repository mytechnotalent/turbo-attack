// Terms Of Use
// ------------
// Do NOT use this on any computer you do not own or are not allowed to run this on.
// You may NEVER attempt to sell this, it is free and open source.
// The authors and publishers assume no responsibility.
// For educational purposes only.

package sudo

import (
	"testing"
)

// Verify Check has root user.
func TestCheckHasRootUser(t *testing.T) {
	// Calls
	got := Check(0)
	// Asserts
	if got != nil {
		t.Fatalf(`Check() = %v`, got)
	}
}

// Verify Check does not have root user.
func TestCheckDoesNotHaveRootUser(t *testing.T) {
	// Calls
	got := Check(501)
	// Asserts
	if got == nil {
		t.Fatalf(`Check() = %v`, got)
	}
}
