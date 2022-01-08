// Output the args, separated by spaces, terminated with a newline.
// The return status is 0 unless a write error occurs.
// If -n is specified, the trailing newline is suppressed
// If the -e option is given, interpretation of the following backslash-escaped characters is enabled.
// The -E option disables the interpretation of these escape characters, even on systems where they are interpreted by default.
// The xpg_echo shell option may be used to dynamically determine whether or not echo expands these escape characters by default.

package main

import (
	"flag"
	"fmt"
)

func usage() {
	fmt.Printf("" +
		"NAME:\n echo -- write arguments to the standard output")
	flag.PrintDefaults()
}
