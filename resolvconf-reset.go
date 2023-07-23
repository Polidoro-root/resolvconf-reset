package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	path, err := filepath.Abs("/etc/resolvconf.conf")

	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile(path, os.O_TRUNC|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	flag.Parse()

	nameservers := fmt.Sprintf(
		"name_servers=\"%s\"\n",
		strings.Join(flag.Args(), " "),
	)

	file.WriteString(nameservers)

	file.Close()

	cmd := exec.Command("resolvconf", "-u")

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
