/* where the modules are configured for zars */
package main

import "github.com/lordrusk/zara/modules"

var delim = " "

var activeStats = &modules.WhichCovidStats{
	Rank:      true,
	TotalSick: true,
	SickToday: true,
	Dead:      true,
	DeadToday: true,
	Recovered: true,
	Active:    true,
	Critical:  true,
}

var mods = []modules.Module{
	modules.NewPacPackages(0, 8),
	/* modules.NewGeorona(60, 19, "USA (US)", "", activeStats), /* us */
	/* modules.NewGeorona(60, 19, "", "", activeStats), /* world */
	modules.NewMemory(6, 14, true),
	modules.NewMemory(6, 14, false),
	modules.NewCpu(3, 13),
	modules.NewCpubar(1, 22),
	modules.NewDisk(7, 15, "/mnt"),
	modules.NewDisk(7, 15, "/home"),
	modules.NewDisk(7, 15, ""),
	modules.NewMoonphase(12*60*60, 18, "kennewick"),
	modules.NewWeather(60, 5, 10, "kennewick"),
	modules.NewNettraf(16),
	modules.NewBat(10, 3),
	modules.NewAudio(0, 10),
	modules.NewTime(1, 1),
	modules.NewInternet(5, 4),
}

/* For my dotfiles
1 clock | time
2 sip (unimplemented)
3 battery
4 internet
5 weather
6 news (unimplemented)
7 torrent (unimplemented)
8 pacpackages
9 recicon (unimplemented)
10 volume
11 music (unimplemented)
12 mailbox (unimplemented)
13 cpu
14 memory
15 disk
16 nettraf
17 crypto (unimplemented)
18 astrological
19 georona
20 help-icon (unimplemented)
21 vpnstat (unimplemented)
22 cpubar
*/
