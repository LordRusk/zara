package modules

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	awk "github.com/benhoyt/goawk/interp"
	awkP "github.com/benhoyt/goawk/parser"
)

const cwireless = "/sys/class/net/w*/operstate"
const cether = "/sys/class/net/e*/operstate"

/*
effectivly replaces * with correct
file for above consts
*/
func parsePath(path string) string {
	matches, err := filepath.Glob(path)
	if err != nil {
		return fmt.Sprintf("Failed to Glob path: %s: %s", path, err)
	}

	if len(matches) < 1 {
		return ""
	} else {
		return matches[0]
	}
}

var (
	wireless         = parsePath(cwireless)
	ether            = parsePath(cether)
	procWireless     = "/proc/net/wireless"
	internetAwkBytes = []byte(`/^\s*w/ { print "📶", int($3 * 100 / 70) "%" }`)
)

/*
internet module that shows wireless
and wired connections
*/
type Internet struct {
	prog *awkP.Program

	ii, sig, pos int
}

/*
returns internet module that shows wireless
and wired connections
*/
func NewInternet(interv, sig int) *Internet {
	prog, err := awkP.ParseProgram(internetAwkBytes, nil)
	if err != nil {
		fmt.Printf("Failed to parse 'internet' awk bytes: %s\n", err)
		os.Exit(1)
	}

	return &Internet{
		prog: prog,
		ii:   interv,
		sig:  sig,
	}
}

func (inter *Internet) Run() (string, error) {
	/* wireless */
	wisUp, err := os.ReadFile(wireless)
	if err != nil {
		return "", fmt.Errorf("Failed to read '%s', %s\n", wireless, err)
	}

	var out string
	if string(wisUp) != "up\n" {
		out += "📡"
	} else { /* get connection percentage */
		f, err := os.OpenFile(procWireless, os.O_RDONLY, 0666)
		if err != nil {
			return "", fmt.Errorf("Failed to open file '%s': %s", procWireless, err)
		}
		defer f.Close()

		var output, awkErr bytes.Buffer
		awkConfig := &awk.Config{
			Stdin:  f,
			Output: &output,
			Error:  &awkErr,
		}

		if _, err := awk.ExecProgram(inter.prog, awkConfig); err != nil {
			return "", fmt.Errorf("Failed to run awk: %s", err)
		}

		if awkErr.Bytes() != nil {
			return "", fmt.Errorf("Failed to run awk: %s", awkErr.Bytes())
		}
		out += strings.TrimSpace(output.String())
	}

	/* ether */
	eisUp, err := os.ReadFile(ether)
	if err != nil {
		return "", fmt.Errorf("Failed to read '%s': %s", wireless, err)
	}

	if string(eisUp) != "up\n" {
		out += " ❎"
	} else {
		out += " 🌐"

	}

	return out, nil
}

func (i *Internet) Int() int {
	return i.ii
}

func (i *Internet) Sig() int {
	return i.sig
}

func (i *Internet) SetPos(pos int) {
	i.pos = pos
}

func (i *Internet) Pos() int {
	return i.pos
}

func (_ Internet) Name() string {
	return "internet"
}
