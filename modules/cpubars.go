package modules

import (
	"fmt"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

type scpu struct {
	total, idle uint64
}

/*
module that prints a Heavy Vertical Bar for each core
with height correspoding to the current load on that core
*/
type Cpubar struct {
	cache []scpu

	ii, sig, pos int
}

/*
returns a module that prints a Heavy Vertical Bar for each core
with height correspoding to the current load on that core
*/
func NewCpubar(interv, sig int) *Cpubar {
	return &Cpubar{
		ii:  interv,
		sig: sig,
	}
}

var bars = map[int]string{
	0: "▁",
	1: "▂",
	2: "▃",
	3: "▄",
	4: "▅",
	5: "▆",
	6: "▇",
	7: "█",
	8: "█",
}

func (c *Cpubar) Run() (string, error) {
	stat, err := linuxproc.ReadStat("/proc/stat")
	if err != nil {
		return "", fmt.Errorf("stat read fail: %s", err)
	}

	cpus := make([]scpu, len(stat.CPUStats))
	for i := 0; i < len(stat.CPUStats); i++ {
		var c scpu
		c.total = stat.CPUStats[i].User + stat.CPUStats[i].Nice + stat.CPUStats[i].System + stat.CPUStats[i].Idle
		c.idle = stat.CPUStats[i].Idle
		cpus[i] = c
	}

	if c.cache == nil {
		c.cache = cpus
		return "", nil
	}

	var rstr string
	for i := 0; i < len(cpus); i++ {
		/* rutime error without nil check of cache
		horrificly hard to read. so here's this:
		(1 - (cpu.idle-cachedCpu.idle) / (cpu.total - cachedCpu.total)) * 100 / 12.5 */
		rstr += bars[int((float64(1)-(float64(cpus[i].idle)-float64(c.cache[i].idle))/(float64(cpus[i].total)-float64(c.cache[i].total)))*float64(100)/float64(12.5))]
	}

	c.cache = cpus
	return rstr, nil
}

func (c *Cpubar) Int() int {
	return c.ii
}

func (c *Cpubar) Sig() int {
	return c.sig
}

func (c *Cpubar) SetPos(pos int) {
	c.pos = pos
}

func (c *Cpubar) Pos() int {
	return c.pos
}

func (_ Cpubar) Name() string {
	return "cpubars"
}
