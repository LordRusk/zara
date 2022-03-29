/*
maybe when I make a fifo statuscmd
patch I'll re-add the horoscope
aspect of this script
*/

package modules

import (
	"fmt"
	"io"
	"net/http"
)

const moonLink = "https://wttr.in/"

/* simple moonphase module */
type Moonphase struct {
	location string

	ii, sig, pos int
}

/*
returns simple moonphase module

location inserted here: wttr.in/'location'?format=%m
see wttr.in for more info

if no location is given, it use's your ip,
this can be inaccurate like in my case.
*/
func NewMoonphase(interv, sig int, location string) *Moonphase {
	return &Moonphase{
		location: location,
		ii:       interv,
		sig:      sig,
	}
}

func (m *Moonphase) Run() (string, error) {
	resp, err := http.Get(moonLink + m.location + "?format=%m")
	if err != nil {
		return "", err
	} else if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Reponse from '%s' not ok: %d", moonLink+m.location+"?format=%m", resp.StatusCode)
	}
	moon, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Failed to read body of response: %s", err)
	}

	return string(moon), nil
}

func (m *Moonphase) Int() int {
	return m.ii
}

func (m *Moonphase) Sig() int {
	return m.sig
}

func (m *Moonphase) SetPos(pos int) {
	m.pos = pos
}

func (m *Moonphase) Pos() int {
	return m.pos
}

func (_ Moonphase) Name() string {
	return "Moonphase"
}
