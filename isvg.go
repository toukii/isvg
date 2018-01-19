package isvg

import (
	"os/exec"

	"github.com/toukii/bytes"
)

func Decode(bs []byte) ([]byte, error) {
	cmd := exec.Command("rsvg-convert")

	w := bytes.NewWriter(make([]byte, 0, 1024))
	cmd.Stdout = w
	cmd.Stdin = bytes.NewReader(bs)
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return w.Bytes(), nil
}
