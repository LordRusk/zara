package modules

import (
	"fmt"
	// "io"
	"net/http"
	"strings"

	"jaytaylor.com/html2text"
)

type Georona struct {
	country, state string

	ii, sig, pos int
}

/*
returns corona virus stats

if country is "" it will return global stats
`
assuming you set country to `US` if state isn't ""
it will return state specific results. */
func NewGeorona(interv, sig int, country, state string) *Georona {
	return &Georona{
		country: country,
		state:   state,
		ii:      interv,
		sig:     sig,
	}
}

const georonaURL = "https://corona-stats.online"
const georonaUsAddon = "/states/us"

/*
world, united states, error */
func getStats() (string, string, error) {
	resp, err := http.Get(georonaURL)
	if err != nil {
		return "", "", err
	} else if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("http response from '%s' not ok: %d", georonaURL, resp.StatusCode)
	}
	world, err := html2text.FromReader(resp.Body, html2text.Options{})
	if err != nil {
		return "", "", err
	}
	resp.Body.Close()

	resp, err = http.Get(georonaURL + georonaUsAddon)
	if err != nil {
		return "", "", err
	} else if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("http response from '%s' not ok: %d", georonaURL, resp.StatusCode)
	}
	us, err := html2text.FromReader(resp.Body, html2text.Options{})
	if err != nil {
		return "", "", err
	}
	resp.Body.Close()

	return world, us, nil
}

type stat struct {
	name string
	rank, totalCases, newCases, totalDeaths,
	newDeaths, recovered, active, critical int
}

const geoSplitChar = "│"

func extractStats(str string) (map[string]stat, error) {
	stats := make(map[string]stat)
	strs := strings.Split(str, "\n")
	for i := 0; i < len(strs); i++ {
		if i%2 == 0 || i == 1 { /* if even */
			continue
		}

		var s stat
		splitLine := strings.Split(strs[i], geoSplitChar)
		for n := 0; n < len(splitLine); n++ {
			ints := extractInts(splitLine[n]) /* modules/weather.go */
			// fmt.Printf("%s: %+v\n", splitLine[n], ints)
			switch n {
			case 0:
				if ints == nil {
					break
				}
				s.rank = ints[0]
			case 1:
				s.name = strings.TrimSpace(splitLine[n])
			case 2:
				if ints == nil {
					s.totalCases = 0
				} else {
					s.totalCases = ints[0]
				}
			case 3:
				if ints == nil {
					s.newCases = 0
				} else {
					s.newCases = ints[0]
				}
			case 4:
				if ints == nil {
					s.totalDeaths = 0
				} else {
					s.totalDeaths = ints[0]
				}
			case 5:
				if ints == nil {
					s.newDeaths = 0
				} else {
					s.newDeaths = ints[0]
				}
			case 6:
				if ints == nil {
					s.recovered = 0
				} else {
					s.recovered = ints[0]
				}
			case 7:
				if ints == nil {
					s.critical = 0
				} else {
					s.critical = ints[0]
				}
			default:
			}
		}
		if s.name != "" {
			stats[s.name] = s
		}
	}

	return stats, nil
}

func (_ Georona) Run() (string, error) {
	worldReport, usReport, err := getStats()
	if err != nil {
		return "", err
	}

	worldStats, err := extractStats(worldReport)
	if err != nil {
		return "", err
	}

	usStats, err := extractStats(usReport)
	if err != nil {
		return "", err
	}

	fmt.Printf("%+v\n\nus:\n%+v\n", worldStats, usStats)

	return "", nil
}

func (g *Georona) Int() int {
	return g.ii
}

func (g *Georona) Sig() int {
	return g.sig
}

func (g *Georona) SetPos(pos int) {
	g.pos = pos
}

func (g *Georona) Pos() int {
	return g.pos
}

func (_ Georona) Name() string {
	return "Georona"
}
