package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	// Write Unit file
	tmpl := `[Unit]
Description=%s
After=network.target

[Service]
ExecStart=%s
Restart=always

[Install]
WantedBy=multi-user.target
`
	cmd := strings.Join(os.Args[1:], " ")
	prog := strings.Split(cmd, " ")[0]
	name := filepath.Base(prog)
	f, err := os.Create(fmt.Sprintf("/etc/systemd/system/%s.service", name))
	if err != nil {
		panic(err)
	}
	_, err = fmt.Fprintf(f, tmpl, name, cmd)
	if err != nil {
		panic(err)
	}

	// Start service
	err = reloadDaemon()
	if err != nil {
		panic(err)
	}
	err = startAndEnable(name)
	if err != nil {
		panic(err)
	}
}

func reloadDaemon() error {
	cmd := exec.Command("systemctl", "daemon-reload")
	return cmd.Run()
}

func startAndEnable(name string) error {
	cmd := exec.Command("systemctl", "enable", "--now", name)
	return cmd.Run()
}
