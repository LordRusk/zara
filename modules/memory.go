package modules

import (
	"bytes"
	"fmt"
	"os/exec"
)

/* memory module that can read swap as well */
type Memory struct {
	swap bool

	ii, sig, pos int
}

/*
returns memory module that can read swap as well

if swap is false, system memory is read.
if swap is true, system swap is read.
*/
func NewMemory(interv, sig int, swap bool) *Memory {
	return &Memory{
		swap: swap,
		ii:   interv,
		sig:  sig,
	}
}

const ( /* important line numbers */
	mem  = 1
	swap = 2
)

func (m *Memory) Run() (string, error) {
	out, err := exec.Command("free", "--mebi").Output()
	if err != nil {
		return "", err
	}

	var outstr string
	sout := bytes.Split(out, []byte("\n"))
	if m.swap {
		if len(sout) >= 3 {
			ints := extractInts(string(sout[swap])) /* modules/weather.go */
			outstr = fmt.Sprintf("🔃%.2fGiB/%.2fGiB", float64(ints[1])/float64(1024), float64(ints[0])/float64(1024))
		}
	} else {
		ints := extractInts(string(sout[mem])) /* modules/weather.go */
		outstr = fmt.Sprintf("🧠%.2fGiB/%.2fGiB", float64(ints[1])/float64(1024), float64(ints[0])/float64(1024))
	}

	return outstr, nil
}

func (m *Memory) Int() int {
	return m.ii
}

func (m *Memory) Sig() int {
	return m.sig
}

func (m *Memory) SetPos(pos int) {
	m.pos = pos
}

func (m *Memory) Pos() int {
	return m.pos
}

func (_ Memory) Name() string {
	return "Memory"
}
