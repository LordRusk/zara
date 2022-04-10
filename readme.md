Dwm statusbar modeled after [gocaudices](https://github.com/lordrusk/gocaudices) expandable in Go through an [interface](https://github.com/LordRusk/zara/blob/master/modules/modules.go).

### modules
I've written many modules for zara to replace many of scripts from my [dotfiles](https://github.com/LordRusk/artixdwm/tree/master/.local/bin/statusbar) - most luke smith scripts from over 2 years ago. Heres a list:

- internet
- time
- audio system independent volume
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

### Zara vs Gocaudices & Dwmblocks
One problem I've encountered while writing scripts and programs for dwmblocks is that the script is just ran. Any information shared between runs must be stored in a file, and first run checks are checked every time the script is ran.

This is significantly different when writing modules for zara. Most of the first checks, along with things like parsing `go-awk` strings into a program for re-use, can be pushed to the `NewModule()` function and never need to be checked in `Module.Run()`. Further, any information shared between run's can can be stored in the structure that satisfies the `Module` interface. No writing to files and reading files every run.

You can also expect a lot better performance because it isn't calling an external program every time a module is updated, only running a function.

### TODO
Add a battery module. I don't have a laptop. This is one someone should create a pull request for.

### other modules
Here's a link list to other peoples modules:
