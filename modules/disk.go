package modules

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	awk "github.com/benhoyt/goawk/interp"
	awkP "github.com/benhoyt/goawk/parser"
)

/* simple disk usage module */
type Disk struct {
	mountPath, icon string
	prog            *awkP.Program

	ii, sig, pos int
}

var diskAwkBytes = []byte(` /[0-9]/ {print $3 "/" $2}`)

/*
returns a new simple disk usage module

mountPath: i.E "/home", "/mnt"
different icons for "/home" and "/mnt"

if mountPath is empty string, it is
treated as "/"
*/
func NewDisk(interv, sig int, mountPath string) *Disk {
	prog, err := awkP.ParseProgram(diskAwkBytes, nil)
	if err != nil {
		fmt.Printf("Failed to parse 'disk' awk bytes: %s\n", err)
		os.Exit(1)
	}

	var icon string
	if mountPath == "/home" {
		icon = "🏠"
	} else if mountPath == "/mnt" {
		icon = "💾"
	} else {
		icon = "🖥"
	}

	if mountPath == "" {
		mountPath = "/"
	}

	return &Disk{
		mountPath: mountPath,
		icon:      icon,
		prog:      prog,
		ii:        interv,
		sig:       sig,
	}
}

func (d *Disk) Run() (string, error) {
	out, err := exec.Command("df", "-h", d.mountPath).Output()
	if err != nil {
		return "", fmt.Errorf("Failed to run 'df -h %s': %s", d.mountPath, err)
	}

	var output, awkErr bytes.Buffer
	awkConfig := &awk.Config{
		Stdin:  bytes.NewReader(out),
		Output: &output,
		Error:  &awkErr,
	}

	if _, err := awk.ExecProgram(d.prog, awkConfig); err != nil {
		return "", fmt.Errorf("Failed to run awk: %s", err)
	}
	if awkErr.Bytes() != nil {
		return "", fmt.Errorf("Awk error: %s", err)
	}

	return fmt.Sprintf("%s: %s", d.icon, output.String()), nil
}

func (d *Disk) Int() int {
	return d.ii
}

func (d *Disk) Sig() int {
	return d.sig
}

func (d *Disk) SetPos(pos int) {
	d.pos = pos
}

func (d *Disk) Pos() int {
	return d.pos
}

func (_ Disk) Name() string {
	return "Disk"
}
