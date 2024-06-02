package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	flagVersion bool
	flagHelp    bool
)

const (
	versionYes = `yes (GNU coreutils) 0.1
Copyright (C) 2024 rendick, Inc.
License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Written by rendick.
`
	helpYes = `Usage: yes [STRING]...
  or:  yes OPTION
Repeatedly output a line with all specified STRING(s), or 'y'.

      --help        display this help and exit
      --version     output version information and exit

GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
Full documentation <https://www.gnu.org/software/coreutils/yes>
or available locally via: info '(coreutils) yes invocation'`
)

func main() {
	flag.BoolVar(&flagVersion, "version", false, "display version information")
	flag.BoolVar(&flagHelp, "help", false, "display help information")
	flag.Parse()

	if flagVersion {
		fmt.Printf("%s", versionYes)
		os.Exit(0)
	}

	if flagHelp {
		fmt.Printf("%s", helpYes)
		os.Exit(0)
	}

	nonFlagsArgs := flag.Args()
	if len(nonFlagsArgs) == 0 {
		for {
			fmt.Println("y")
		}
	}
}
