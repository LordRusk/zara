[![Go Report Card](https://goreportcard.com/badge/github.com/lordrusk/zara)](https://goreportcard.com/report/github.com/lordrusk/zara)

Dwm statusbar modeled after [gocaudices](https://github.com/lordrusk/gocaudices) expandable in Go through an [interface](https://github.com/LordRusk/zara/blob/master/modules/modules.go).

### modules
I've written many modules for zara to replace many of scripts from my [dotfiles](https://github.com/LordRusk/artixdwm/tree/master/.local/bin/statusbar) - most luke smith scripts from over 2 years ago. Heres a list:

- internet
- time
- audio system independent volume
- OS independent batteries
- network traffic
- weather
- moonphase
- disk
- cpu
- cpubars
- memory
- georona
- pacpackages

### How To?
`git clone https://github.com/lordrusk/zara` then edit `mods.go` to your liking. `go mod tidy` to make sure you have all dependencies installed. `go install` to (re)install.

### Signals
The module definition of `modules.NewAudio(0, 10),` would be updated like `kill -44 $(pidof zara)` A dwm volume mute keybind might look like `{ 0, XF86XK_AudioMute, spawn, SHCMD("pamixer -t; kill -44 $(pidof gocaudices)") },`. NOTE: You're updating sig 44 because the first 34 signals are in use.

### Zara vs Gocaudices & Dwmblocks
One problem I've encountered while writing scripts and programs for dwmblocks is that the script is just ran. Any information shared between runs must be stored in a file, and first run checks are checked every time the script is ran.

This is significantly different when writing modules for zara. Most of the first checks, along with things like parsing `go-awk` strings into a program for re-use, can be pushed to the `NewModule()` function and never need to be checked in `Module.Run()`. Further, any information shared between run's can can be stored in the structure that satisfies the `Module` interface. No writing to files and reading files every run.

You can also expect a lot better performance because it isn't calling an external program every time a module is updated, only running a function.

### other modules
Here's a link list to other peoples modules:


## AWESOME BARS
dwm bars that I think are awesome! check them out and give them a star!

• [sysmon](https://github.com/blmayer/sysmon/tree/main) I would use this if I hadn't made zara

• [spoon](https://git.2f30.org/spoon/) I don't know much C but this is great

• [rsblocks](https://github.com/MustafaSalih1993/rsblocks) I don't know much Rust, but featureful and well starred, makes me wanna get my status emoji game up to par

• [mblocks](https://gitlab.com/mhdy/mblocks) another great rusty bar

• [integrated-status-text](https://dwm.suckless.org/patches/integrated-status-text) the way god intended
  
• [gods](https://github.com/schachmat/gods) ICONIC

• [dwmblocks-async](https://github.com/UtkarshVerma/dwmblocks-async) Awesome! I wrote this project because dwmblocks wasn't async...and I've lived without bar clickability since...maybe should have gone with this and learned C!

•[Luke Smith's Dwmblocks](https://github.com/LukeSmithxyz/dwmblocks) how could I forget where it all began?
