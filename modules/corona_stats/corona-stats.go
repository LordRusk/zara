/*
pulls and parses current coronavirus stats

written for modules/georona.go */
package corona_stats

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"jaytaylor.com/html2text"
)

const URL = "https://corona-stats.online"
const UsAddon = "/states/us"

/*
world, united states, error */
func GetStats() (string, string, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return "", "", err
	} else if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("http response from '%s' not ok: %d", URL, resp.StatusCode)
	}
	world, err := html2text.FromReader(resp.Body, html2text.Options{})
	if err != nil {
		return "", "", err
	}
	resp.Body.Close()

	resp, err = http.Get(URL + UsAddon)
	if err != nil {
		return "", "", err
	} else if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("http response from '%s' not ok: %d", URL+UsAddon, resp.StatusCode)
	}
	us, err := html2text.FromReader(resp.Body, html2text.Options{})
	if err != nil {
		return "", "", err
	}
	resp.Body.Close()

	return world, us, nil
}

type Stat struct {
	Name string
	Rank, TotalCases, NewCases, TotalDeaths,
	NewDeaths, Recovered, Active, Critical int
}

const geoSplitChar = "│"

/*
parses the output from GetReport()
into a useable structure */
func ExtractStats(str string) (map[string]*Stat, error) {
	stats := make(map[string]*Stat)
	strs := strings.Split(str, "\n")
	for i := 0; i < len(strs); i++ {
		if i%2 == 0 || i == 1 { /* if even */
			continue
		}

		var s Stat
		splitLine := strings.Split(strs[i], geoSplitChar)
		for n := 0; n < len(splitLine); n++ {
			ints := ExtractInts(splitLine[n]) /* modules/weather.go */
			// fmt.Printf("%s: %+v\n", splitLine[n], ints)
			switch n {
			case 0:
				if ints == nil {
					break
				}
				s.Rank = ints[0]
			case 1:
				s.Name = strings.TrimSpace(splitLine[n])
			case 2:
				if ints == nil {
					s.TotalCases = 0
				} else {
					s.TotalCases = ints[0]
				}
			case 3:
				if ints == nil {
					s.NewCases = 0
				} else {
					s.NewCases = ints[0]
				}
			case 4:
				if ints == nil {
					s.TotalDeaths = 0
				} else {
					s.TotalDeaths = ints[0]
				}
			case 5:
				if ints == nil {
					s.NewDeaths = 0
				} else {
					s.NewDeaths = ints[0]
				}
			case 6:
				if ints == nil {
					s.Recovered = 0
				} else {
					s.Recovered = ints[0]
				}
			case 7:
				if ints == nil {
					s.Critical = 0
				} else {
					s.Critical = ints[0]
				}
			default:
			}
		}
		if s.Name != "" {
			stats[s.Name] = &s
		}
	}

	return stats, nil
}

/*
from modules/weather.go

goes through a string and extracts ints | 0000 = 0 */
func ExtractInts(str string) []int {
	var rints []int
	var currentNum int
	var zero bool

	splitLine := strings.Split(str, "")
	for i := 0; i < len(splitLine); i++ {
		num, err := strconv.Atoi(splitLine[i])
		if err != nil {
			if splitLine[i] == "," {
				continue
			}

			if zero {
				rints = append(rints, currentNum)
			} else if currentNum != 0 {
				rints = append(rints, currentNum)
			}
			currentNum = 0
			zero = false
			continue
		}
		if num == 0 {
			zero = true
		}
		currentNum *= 10
		currentNum += num
	}

	return rints
}
