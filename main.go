package main

import (
	"bytes"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
	"github.com/lordrusk/zara/modules"
)

var update = make(chan struct{})
var bar = make([]string, len(mods))

/*
wrapper to handle output and error
of module.Moudle.Run()
*/
func run(m modules.Module) {
	out, err := m.Run()
	if err != nil {
		fmt.Printf("Failed to run '%s': %s\n", m.Name(), err)
		return
	}

	bar[m.Pos()] = strings.Split(out, "\n")[0]
	update <- struct{}{}
}

func main() {
	x, err := xgb.NewConn() /* connect to X */
	if err != nil {
		fmt.Printf("Cannot connect to X: %s\n", err)
		return
	}
	defer x.Close()
	root := xproto.Setup(x).DefaultScreen(x).Root

	go func() {
		var buf bytes.Buffer
		for range update {
			for i := 0; i < len(mods); i++ {
				if bar[i] != "" {
					buf.WriteString(delim)
					buf.WriteString(bar[i])
				}
			}

			if buf.Len() > 0 {
				xproto.ChangeProperty(x, xproto.PropModeReplace, root, xproto.AtomWmName, xproto.AtomString, 8, uint32(buf.Len()-len(delim)), buf.Bytes()[len(delim):]) /* set the root window name */
				buf.Reset()
			}
		}
	}()

	sigs := make(chan os.Signal, 1024)
	sigMap := make(map[os.Signal][]modules.Module)
	for i := 0; i < len(mods); i++ {
		go func(i int) {
			mods[i].SetPos(i)

			if mods[i].Sig() != 0 {
				sig := syscall.Signal(34 + mods[i].Sig())
				signal.Notify(sigs, sig)
				sigMap[sig] = append(sigMap[sig], mods[i])
			}

			run(mods[i]) /* build bar */
			if mods[i].Int() != 0 {
				interval := time.Duration(mods[i].Int()) * time.Second
				for {
					time.Sleep(interval)
					run(mods[i])
				}
			}
		}(i)
	}

	for sig := range sigs {
		go func(sig *os.Signal) {
			ms := sigMap[*sig]
			for _, m := range ms {
				go run(m)
			}
		}(&sig)
	}
}
