package modules

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const rxPath = "/sys/class/net/*/statistics/rx_bytes"
const txPath = "/sys/class/net/*/statistics/tx_bytes"

/* network traffic module */
type Nettraf struct {
	rx, rxold, tx, txold int

	ii, sig, pos int
}

/*
returns a network traffic module

needs to refresh every second for
accurate reports so only takes sig
*/
func NewNettraf(sig int) *Nettraf {
	return &Nettraf{
		ii:  1,
		sig: sig,
	}
}

func (n *Nettraf) Run() (string, error) {
	/* rx || down */
	dirs, err := os.ReadDir(strings.Split(rxPath, "*")[0])
	if err != nil {
		return "", err
	}

	var sum int
	for _, dir := range dirs {
		bites, err := os.ReadFile(strings.ReplaceAll(rxPath, "*", dir.Name()))
		if err != nil {
			return "", err
		}

		ibuf, err := strconv.Atoi(strings.TrimSpace(string(bites)))
		if err != nil {
			return "", fmt.Errorf("Failed to convert %s to int: %s", bites, err)
		}
		sum += ibuf
	}
	n.rx = (sum - n.rxold) / 1024
	n.rxold = sum

	/* tx || up */
	dirs, err = os.ReadDir(strings.Split(txPath, "*")[0])
	if err != nil {
		return "", err
	}

	sum = 0
	for _, dir := range dirs {
		bites, err := os.ReadFile(strings.ReplaceAll(txPath, "*", dir.Name()))
		if err != nil {
			return "", err
		}

		ibuf, err := strconv.Atoi(strings.TrimSpace(string(bites)))
		if err != nil {
			return "", fmt.Errorf("Failed to convert %s to int: %s", bites, err)
		}
		sum += ibuf
	}
	n.tx = (sum - n.txold) / 1024
	n.txold = sum

	return fmt.Sprintf("🔻%dKiB 🔺%dKiB\n", n.rx, n.tx), nil
}

func (n *Nettraf) Int() int {
	return n.ii
}

func (n *Nettraf) Sig() int {
	return n.sig
}

func (n *Nettraf) SetPos(pos int) {
	n.pos = pos
}

func (n *Nettraf) Pos() int {
	return n.pos
}

func (_ Nettraf) Name() string {
	return "Network Traffic"
}
