package isvg

import (
	"os"
	"testing"

	"github.com/toukii/bytes"
	"github.com/toukii/goutils"
	"github.com/toukii/icat"
)

func TestIsvg(t *testing.T) {
	bs := goutils.ReadFile("github.svg")
	dbs, err := Decode(bs)
	if err != nil {
		t.Errorf("%s", err)
	}
	if err := icat.ICatRead(bytes.NewReader(dbs), os.Stdout); err != nil {
		t.Errorf("%s", err)
	}
}
