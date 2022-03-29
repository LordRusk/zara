package modules

import (
	"bytes"
	"fmt"
	"os/exec"
)

/* cpu module that gets average cpu temp between all cores */
type Cpu struct {
	ii, sig, pos int
}

/* returns cpu module that gets average cpu temp between all cores */
func NewCpu(interv, sig int) *Cpu {
	return &Cpu{
		ii:  interv,
		sig: sig,
	}
}

func (_ Cpu) Run() (string, error) {
	out, err := exec.Command("sh", "-c", "sensors | tail").Output()
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
	average /= len(temps)

	return fmt.Sprintf("🌡%d°C\n", average), nil
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
