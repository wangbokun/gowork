package exec

import (
	"os"
	"os/exec"

	"github.com/wangbokun/gowork/color"
)

// const (
// 	CheckSymbol = "\u2714 "
// 	CrossSymbol = "\u2716 "
// )

func Execute(workDir, script string, args ...string) bool {
	cmd := exec.Command(script, args...)

	if workDir != "" {
		cmd.Dir = workDir
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		color.Red(err.Error())
		return false
	}

	return true
}
