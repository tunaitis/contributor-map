package render

import (
	"fmt"
	"github.com/tunaitis/contributor-map/internal/util"
	"math"
	"os"
	"path"
	"sort"
	"strings"
)

func readTemplate(name string) ([]byte, error) {
	template, err := os.ReadFile(path.Join("template", name))
	if err == nil {
		return template, nil
	}

	// If the data file is not found, look for it in the folder where the process started.
	e, err := os.Executable()
	if err != nil {
		return nil, err
	}

	template, err = os.ReadFile(path.Join(path.Dir(e), "template", name))
	if err != nil {
		return nil, err
	}

	return template, nil
}

type location struct {
	name  string
	value int
}

type fromTo struct {
	from int
	to   int
}

func sortValues(locations map[string]int) []int {
	var r []int
	for c := range locations {
		r = append(r, locations[c])
	}
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	return r
}

func makeUnique(items []int) []int {
	var r []int
	for i := range items {
		f := false
		for j := range r {
			if r[j] == items[i] {
				f = true
			}
		}
		if !f {
			r = append(r, items[i])
		}
	}

	return r
}

func classify(values []int, number int) map[int]fromTo {
	r := make(map[int]fromTo, number)

	if len(values) > number {
		classSize := math.Ceil(float64(len(values)) / float64(number))
		for i := range values {
			c := int(math.Round(float64(i) / classSize))
			if _, found := r[c]; found {
				if values[i] > r[c].to {
					r[c] = fromTo{from: r[c].from, to: values[i]}
				}
			} else {
				r[c] = fromTo{from: values[i], to: values[i]}
			}
		}
	} else {
		for i := range values {
			if _, found := r[i]; found {
				if values[i] > r[i].to {
					r[i] = fromTo{from: r[i].from, to: values[i]}
				}
			} else {
				r[i] = fromTo{from: values[i], to: values[i]}
			}
		}
	}

	return r
}

func Map(locations map[string]int, palette []string) ([]byte, error) {
	template, err := readTemplate("map.svg")
	if err != nil {
		return nil, err
	}

	//palette := []string{"#99e2b4", "#88d4ab", "#78c6a3", "#67b99a", "#56ab91", "#469d89", "#358f80", "#248277", "#14746f", "#036666"}

	values := sortValues(locations)
	values = makeUnique(values)
	classes := classify(values, len(palette))

	if len(classes) < len(palette) {
		palette = palette[0:len(classes)]
	}

	style := strings.Builder{}
	style.WriteString("\n<style>\n")

	for i := range palette {
		style.WriteString(fmt.Sprintf("\t.palette-color-%d { fill: %s !important; }\n", i+1, palette[i]))
	}

	for c := range classes {
		for k := range locations {
			if locations[k] >= classes[c].from && locations[k] <= classes[c].to {
				style.WriteString(fmt.Sprintf("\t.%s { fill: %s; }\n", strings.ToLower(k), palette[c]))
			}
		}
	}

	style.WriteString(fmt.Sprintf("#legend%d { display: inline !important; }", len(classes)))
	style.WriteString("</style>\n")

	m := strings.Replace(string(template),
		"<!-- map_style -->",
		style.String(), 1)

	for c := range classes {
		k := fmt.Sprintf(">%%%d<", c+1)
		r := fmt.Sprintf(">%s<", util.NearestThousandFormat(float64(classes[c].from)))
		m = strings.Replace(m, k, r, -1)
	}
	k := fmt.Sprintf(">%%%d<", len(classes)+1)
	r := fmt.Sprintf(">%s<", util.NearestThousandFormat(float64(classes[len(classes)-1].to)))
	m = strings.Replace(m, k, r, -1)

	for k := range locations {
		m = strings.Replace(m,
			fmt.Sprintf("<!-- %s_contributions -->", strings.ToLower(k)),
			fmt.Sprintf(" (%d contributions)", locations[k]), 1)
	}

	return []byte(m), nil
}
