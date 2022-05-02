package modules

import (
	"fmt"
	"strings"

	humanize "github.com/dustin/go-humanize"
	rona "github.com/lordrusk/zara/modules/corona_stats"
)

/*
struct used for toggling georona modules. */
type WhichCovidStats = struct {
	Rank, TotalSick, SickToday, Dead, DeadToday, Recovered, Active, Critical bool
}

type Georona struct {
	country, state string
	as             *WhichCovidStats

	ii, sig, pos int
}

/*
returns corona virus stats

if country is "" it will return global stats
if state isn't "", regardless of country, it will
pull stats for that state in the US.

Find names on corona-stats.online */
func NewGeorona(interv, sig int, country, state string, as *WhichCovidStats) *Georona {
	if country == "" {
		country = "World"
	}

	return &Georona{
		country: country,
		state:   state,
		as:      as,
		ii:      interv,
		sig:     sig,
	}
}

var (
	grank      = "🏆"
	gtotalsick = "😷"
	gsicktoday = "🤢"
	gdead      = "☠"
	gdeadtoday = "💀"
	grecovered = "😓"
	gactive    = "🤮"
	gcritical  = "😵"
)

func statToStr(s *rona.Stat, as *WhichCovidStats) string {
	var rstr string
	if as.Rank == true {
		if s.Rank != 0 {
			rstr += fmt.Sprintf("%s%s ", grank, humanize.Comma(int64(s.Rank)))
		}
	}
	if as.TotalSick == true {
		if s.TotalCases != 0 {
			rstr += fmt.Sprintf("%s%s ", gtotalsick, humanize.Comma(int64(s.TotalCases)))
		}
	}
	if as.SickToday == true {
		if s.NewCases != 0 {
			rstr += fmt.Sprintf("%s%s▲ ", gsicktoday, humanize.Comma(int64(s.NewCases)))
		}
	}
	if as.Dead == true {
		if s.TotalDeaths != 0 {
			rstr += fmt.Sprintf("%s%s ", gdead, humanize.Comma(int64(s.TotalDeaths)))
		}
	}
	if as.DeadToday == true {
		if s.NewDeaths != 0 {
			rstr += fmt.Sprintf("%s%s▲ ", gdeadtoday, humanize.Comma(int64(s.NewDeaths)))
		}
	}
	if as.Recovered == true {
		if s.Recovered != 0 {
			rstr += fmt.Sprintf("%s%s ", grecovered, humanize.Comma(int64(s.Recovered)))
		}
	}
	if as.Active == true {
		if s.Active != 0 {
			rstr += fmt.Sprintf("%s%s ", gactive, humanize.Comma(int64(s.Active)))
		}
	}
	if as.Critical == true {
		if s.Critical != 0 {
			rstr += fmt.Sprintf("%s%s ", gcritical, humanize.Comma(int64(s.Critical)))
		}
	}

	return strings.TrimSpace(rstr)
}

func (g *Georona) Run() (string, error) {
	worldReport, usReport, err := rona.GetStats()
	if err != nil {
		return "", err
	}

	if g.state != "" {
		stats, err := rona.ExtractStats(usReport)
		if err != nil {
			return "", err
		}

		if stats[g.state] == nil {
			return "", fmt.Errorf("%s not a state...check corona-stats.online/states/us", g.state)
		}
		return statToStr(stats[g.state], g.as), nil
	} else {
		stats, err := rona.ExtractStats(worldReport)
		if err != nil {
			return "", err
		}

		if stats[g.country] == nil {
			return "", fmt.Errorf("%s not a country...check corona-stats.online", g.country)
		}
		return statToStr(stats[g.country], g.as), nil
	}
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
