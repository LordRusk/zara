package modules

import (
	"bytes"
	"fmt"
	"os/exec"

	errs "github.com/pkg/errors"
)

var errDiv0 = errs.New("Divide by 0! Install LM_Sensors!")

/* cpu module that gets average cpu temp between all cores */
type Cpu struct {
	sensorStr, sensorLetter string

	ii, sig, pos int
}

/* returns cpu module that gets average cpu temp between all cores */
func NewCpu(interv, sig int, fahrenheit bool) *Cpu {
	cpu := &Cpu{}
	if fahrenheit {
		cpu.sensorStr = "sensors -f | head -8 | tail -6"
		cpu.sensorLetter = "F"

	} else {
		cpu.sensorStr = "sensors | head -8 | tail -6"
		cpu.sensorLetter = "C"
	}
	cpu.ii = interv
	cpu.sig = sig
	return cpu
}

func (c *Cpu) Run() (string, error) {
	out, err := exec.Command("sh", "-c", c.sensorStr).Output()
	if err != nil {
		return "", err
	}

	sout := bytes.Split(out, []byte("\n"))
	var temps []int
	for i := 0; i < len(sout); i++ {
		ints := extractInts(string(sout[i])) /* modules/weather.go */
		if len(ints) > 2 {
			temps = append(temps, ints[1])
		}
	}

	var average int
	for _, i := range temps {
		average += i
	}
	if len(temps) < 1 {
		return "", errDiv0 /* fixes divide by 0 error */
	}
	average /= len(temps)

	return fmt.Sprintf("🌡%d°%s\n", average, c.sensorLetter), nil
}

func (c *Cpu) Int() int {
	return c.ii
}

func (c *Cpu) Sig() int {
	return c.sig
}

func (c *Cpu) SetPos(pos int) {
	c.pos = pos
}

func (c *Cpu) Pos() int {
	return c.pos
}

func (_ Cpu) Name() string {
	return "Cpu"
}
