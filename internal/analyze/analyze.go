package analyze

import (
	"io"
	"os/exec"
)

func RunTree(dir string, w io.Writer) error {
	cmd := exec.Command("tree", "-I", "__pycache__")
	cmd.Dir = dir
	cmd.Stdout = w
	cmd.Stderr = w
	return cmd.Run()
}

func RunCloc(dir string, w io.Writer) error {
	cmd := exec.Command("cloc", ".")
	cmd.Dir = dir
	cmd.Stdout = w
	cmd.Stderr = w
	return cmd.Run()
}
