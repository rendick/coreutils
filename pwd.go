package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	flagVersionPwd bool
	flagHelpPwd    bool
)

const (
	helpPwd = `Usage: yes [STRING]...
  or:  yes OPTION
Repeatedly output a line with all specified STRING(s), or 'y'.

      --help        display this help and exit
      --version     output version information and exit

GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
Full documentation <https://www.gnu.org/software/coreutils/yes>
or available locally via: info '(coreutils) yes invocation'`
	versionPwd = `pwd (GNU coreutils) 0.1
Copyright (C) 2024 rendick.
License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Written by rendick.`
)

func main() {
	flag.BoolVar(&flagHelpPwd, "help", false, "display help information")
	flag.BoolVar(&flagVersionPwd, "versoin", false, "display version information")
	flag.Parse()

	if flagHelpPwd {
		fmt.Printf("%s\n", helpPwd)
		os.Exit(0)
	}

	if flagVersionPwd {
		fmt.Printf("%s\n", versionPwd)
		os.Exit(0)
	}

	nonFlagsArgs := flag.Args()
	if len(nonFlagsArgs) == 0 {
		pwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", pwd)
	}
}
