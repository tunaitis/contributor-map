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

func Map(locations map[string]int) ([]byte, error) {
	template, err := readTemplate("map.svg")
	if err != nil {
		return nil, err
	}

	min := math.MaxInt
	max := 0

	for k := range locations {
		if locations[k] > max {
			max = locations[k]
		}
		if locations[k] < min {
			min = locations[k]
		}
	}

	var l []struct{string;int}
	for c := range locations {
		l = append(l, struct{string; int}{c, locations[c]})
	}
	sort.Slice(l, func(i, j int) bool {
		return l[i].int > l[j].int
	})

	//palette := []string{"#e5bb42", "#e7873f", "#df5f2c", "#ed3726", "#bd281c", "#861e11"}

	//palette := []string{"#EDF8E9", "#C8E8C2", "#A0D69B", "#74C476", "#41AB5D", "#238B45", "#005A32"}
	//palette := []string{"#E2EEE7", "#CAE3D6", "#B3D7C4", "#9BCCB2", "#84C1A1", "#6DB68F", "#55AA7D"}
	//palette := []string{"#e5f2eb", "#d8ebe1", "#cbe5d7", "#bedecd", "#b1d8c3", "#a3d1b9", "#96cbaf", "#89c4a5", "#7cbe9b", "#6fb791", "#62b187", "#55aa7d", "#4e9d73", "#48906a", "#418360", "#3b7657", "#34694d"}

	//palette := []string{"#d5e5ff", "#aaccff", "#80b3ff", "#5599ff", "#2a7fff", "#0066ff", "#0055d4", "#0044aa", "#003380"}
	//palette := []string{"#e5d5ff", "#ccaaff", "#b380ff", "#9955ff", "#7f2aff", "#6600ff", "#5500d4", "#4400aa", "#330080", "#220055"}
	//palette := []string{"#7a7a7a", "#8b7679", "#9b7076", "#a96973", "#b76171", "#c3596f", "#d04d6d", "#da416c", "#e62d6b", "#f0036a"}
	palette := []string{"#4cc9f0", "#4895ef", "#4361ee", "#3f37c9"}//, "#3a0ca3", "#480ca8", "#560bad", "#7209b7", "#b5179e", "#f72585"}
	//palette := []string{"#184e77", "#1e6091", "#1a759f", "#168aad", "#34a0a4", "#52b69a", "#76c893", "#99d98c", "#b5e48c", "#d9ed92"}
	//palette := []string{"#99e2b4", "#88d4ab", "#78c6a3", "#67b99a", "#56ab91", "#469d89", "#358f80", "#248277", "#14746f", "#036666"}
	//palette := []string{"#80ffdb", "#72efdd", "#64dfdf", "#56cfe1", "#48bfe3", "#4ea8de", "#5390d9", "#5e60ce", "#6930c3", "#7400b8"}
	//palette := []string{"#277da1", "#577590", "#4d908e", "#43aa8b", "#90be6d", "#f9c74f", "#f9844a", "#f8961e", "#f3722c", "#f94144"}
	//palette := []string{"#03071e", "#370617", "#6a040f", "#9d0208", "#d00000", "#dc2f02", "#e85d04", "#f48c06", "#faa307", "#ffba08"}

	/*
		start, _ := colorful.Hex("#d5e5ff")
		end, _ := colorful.Hex("#002255")
		palette := make([]string, 10)

		for i := range palette {
			c := start.BlendLuv(end, float64(i)/float64(len(palette) - 1))
			palette[i] = c.Hex()
		}
	*/

	style := strings.Builder{}
	style.WriteString("\n<style>\n")
	/*
		for i := range l {
			style.WriteString(fmt.Sprintf("\t.%s { fill: %s; }\n", strings.ToLower(l[i].string), palette[i]))
		}*/
	for i := range palette {
		style.WriteString(fmt.Sprintf("\t.palette-color-%d { fill: %s !important; }\n", i+1, palette[i]))
	}
	for k := range locations {
		color := int(math.Ceil(float64(locations[k]) / float64(max) * float64(len(palette)))) - 1
		style.WriteString(fmt.Sprintf("\t.%s { fill: %s; }\n", strings.ToLower(k), palette[color]))
	}
	style.WriteString("</style>\n")

	m := strings.Replace(string(template),
		"<!-- map_style -->",
		style.String(), 1)

	s := int(float64(max-min) / float64(len(palette)))
	for i := 0; i<len(palette) + 1; i++ {
		k := fmt.Sprintf("<!-- legend_value_%d -->", i + 1)
		v := float64(i * s + min)
		m = strings.Replace(m, k, util.NearestThousandFormat(v), 1)
	}

	for k := range locations {
		m = strings.Replace(m,
			fmt.Sprintf("<!-- %s_contributions -->", strings.ToLower(k)),
			fmt.Sprintf(" (%d contributions)", locations[k]), 1)
	}

	return []byte(m), nil
}
