// an example time module
package modules

import (
	"strings"
	"time"
)

/* simply formatted time module */
type Time struct {
	ii, sig, pos int
}

/* returns a simply formatted time module */
func NewTime(inter, sig int) *Time {
	return &Time{
		ii:  inter,
		sig: sig,
	}
}

func (_ Time) Run() (string, error) {
	return strings.ReplaceAll(time.Now().Format(time.UnixDate), "  ", " "), nil
}

func (t *Time) Int() int {
	return t.ii
}

func (t *Time) Sig() int {
	return t.sig
}

func (t *Time) SetPos(pos int) {
	t.pos = pos
}

func (t *Time) Pos() int {
	return t.pos
}

func (_ Time) Name() string {
	return "time"
}
