package modules

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

/*
module that shows number of pacman packages
that could be upgraded, assuming repos are synced */
type Pacpackages struct {
	ii, sig, pos int
}

/*
returns a module that shows number of pacman
packages that could be upgraded, assuming repos are synced */
func NewPacPackages(interv, sig int) *Pacpackages {
	return &Pacpackages{
		ii:  interv,
		sig: sig,
	}
}

func (_ Pacpackages) Run() (string, error) {
	out, err := exec.Command("sh", "-c", "pacman -Qu | grep -Fcv \"[ignored]\"").CombinedOutput()
	if err != nil && string(out) != "0\n" {
		return "", fmt.Errorf("%s: %s", bytes.TrimSpace(out), err)
	}

	i, err := strconv.Atoi(strings.TrimSpace(string(out)))
	if err != nil {
		return "", err
	}
	if i == 0 {
		return "", nil
	} else {
		return fmt.Sprintf("📦%d", i), nil
	}
}

func (p *Pacpackages) Int() int {
	return p.ii
}

func (p *Pacpackages) Sig() int {
	return p.sig
}

func (p *Pacpackages) SetPos(pos int) {
	p.pos = pos
}

func (p *Pacpackages) Pos() int {
	return p.pos
}

func (_ Pacpackages) Name() string {
	return "Pacman Packages"
}
