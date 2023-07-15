package modules

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/distatus/battery"
)

type Bat struct {
	ii, sig, pos int
}

/* Displays percentage of all connected batteries */
func NewBat(inter, sig int) *Bat {
	return &Bat{
		ii: inter,
		sig: sig,
	}
}

func (_ *Bat) Run() (string, error) {
	str := ""
	batteries, err := battery.GetAll()
	if err != nil {
		return "", errors.Wrap(err, "could not get battery info")
	}

	for i, battery := range batteries {
		if i != 0 {
			str += " "
		}
		perc := int(battery.Current / battery.Full * 100)
		if perc > 90 {
			str += fmt.Sprint("")
		} else if perc > 80 {
			str += fmt.Sprint("")
		} else if perc > 50 {
			str += fmt.Sprint("")
		} else if perc > 30 {
			str += fmt.Sprint("")
		} else if perc < 10 {
			str += fmt.Sprint("")
		}
		str += fmt.Sprintf("%d%%", int(battery.Current/battery.Full*100))
		if battery.State.String() == "Charging" {
			str += fmt.Sprint("🔌")
		} else {
			str += fmt.Sprint("")
		}
	}
	return str, nil
}

func (b *Bat) Int() int {
	return b.ii
}

func (b *Bat) Sig() int {
	return b.sig
}

func (b *Bat) SetPos(pos int) {
	b.pos = pos
}

func (b *Bat) Pos() int {
	return b.pos
}

func (b *Bat) Name() string {
	return "Batteries"
}
