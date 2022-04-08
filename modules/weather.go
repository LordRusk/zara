package modules

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"jaytaylor.com/html2text"
)

const weatherLink = "https://wttr.in/"

/*
shows precipitation, daily low and daily high
with varied weather report update interval
separate from modules update interval.
*/
type Weather struct {
	barstr, location string
	upInterv         int
	lastUpdate       time.Time

	ii, sig, pos int
}

/*
returns a weather module that shows precipitation,
daily low and daily high. updateWeatherReport is
how many minutes between weather report updates.
this does mean the weather report itself is not
pulled every time zara calls Run().

location is inserted here: wttr.in/'location'.
see wttr.in for details

if no location is given, it use's your ip,
this can be inaccurate like in my case.
*/
func NewWeather(interv, sig, updateWeatherReport int, location string) *Weather {
	return &Weather{
		location: location,
		upInterv: updateWeatherReport,
		ii:       interv,
		sig:      sig,
	}
}

func getReport(location string) (string, error) {
	resp, err := http.Get(weatherLink + location)
	if err != nil {
		return "", err
	} else if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("http response from '%s' not ok: %d", weatherLink+location, resp.StatusCode)
	}
	defer resp.Body.Close()

	str, err := html2text.FromReader(resp.Body, html2text.Options{})
	if err != nil {
		return "", fmt.Errorf("Failed to convert html to plaintext: %s", err)
	}

	return str, nil
}

/* goes through a string and extracts ints | 0000 = 0 */
func extractInts(str string) []int {
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

const ( /* significant line numbers */
	current = 3
	day     = 12
	percip  = 15
)

func (w *Weather) Run() (string, error) {
	if w.lastUpdate.Add(time.Duration(w.upInterv) * time.Minute).Before(time.Now()) {
		w.lastUpdate = time.Now()
		report, err := getReport(w.location)
		if err != nil {
			return "", fmt.Errorf("Failed to get weather report: %s", err)
		}
		log := strings.Split(report, "\n") /* split essentially by line number */

		if len(log) < percip+1 {
			return "", fmt.Errorf("No weather report?\n\nreport: '%s'", report)
		}
		temps := append(extractInts(log[current]), extractInts(log[day])...)
		percips := extractInts(log[percip])

		sort.Ints(temps)
		sort.Ints(percips)

		w.barstr = fmt.Sprintf("☔%d%% ❄%d° ☀%d°", percips[len(percips)-1], temps[0], temps[len(temps)-1])
	}

	return w.barstr, nil
}

func (w *Weather) Int() int {
	return w.ii
}

func (w *Weather) Sig() int {
	return w.sig
}

func (w *Weather) SetPos(pos int) {
	w.pos = pos
}

func (w *Weather) Pos() int {
	return w.pos
}

func (_ Weather) Name() string {
	return "Weather"
}
