package fileserver

import (
	"log"
	"os/exec"

	"github.com/basi-a/useless-applet/config"
)

func FileServerStart() *exec.Cmd {
	cmd := exec.Command( config.Shell, config.ShellOption, "python -m http.server "+config.FileServerPort+" -d "+config.FileServerSourceDir)
	if err := cmd.Start(); err != nil {
		log.Println(err)
	}
	return cmd
}

func FileServerStop(cmd *exec.Cmd)  {
	if err := cmd.Process.Kill(); err != nil {
		log.Println(err)
	}
	cmd.Wait()
}

