package modules

import (
	"fmt"

	volume "github.com/itchyny/volume-go"
)

/* audio system independent volume module */
type Audio struct {
	ii, sig, pos int
}

/* returns audio system independent volume module */
func NewAudio(inter, sig int) *Audio {
	return &Audio{
		ii:  inter,
		sig: sig,
	}
}

func (_ Audio) Run() (string, error) {
	vol, err := volume.GetVolume()
	if err != nil {
		return "", fmt.Errorf("Failed to get volume: %s", err)
	}

	var emoji string
	if vol >= 70 {
		emoji = "🔊"
	} else if vol <= 30 {
		emoji = "🔈"
	} else {
		emoji = "🔉"
	}

	return fmt.Sprintf("%s%d%%", emoji, vol), nil
}

func (a *Audio) Int() int {
	return a.ii
}

func (a *Audio) Sig() int {
	return a.sig
}

func (a *Audio) SetPos(pos int) {
	a.pos = pos
}

func (a *Audio) Pos() int {
	return a.pos
}

func (_ Audio) Name() string {
	return "Audio Module"
}
