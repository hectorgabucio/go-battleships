package stuff

import (
	"bufio"
	"fmt"
	"os"
	"text/template"

	"github.com/rs/zerolog/log"
)

func decInt(i int, by int) string {
	return fmt.Sprintf("%d", i-by)
}
func incInt(i int, by int) string {
	return fmt.Sprintf("%d", i+by)
}

type shipDesc struct {
	St   string
	Name string
}

type templateStruct struct {
	Letters []byte
	Numbers []int
	Names   []shipDesc
}

var ts = templateStruct{
	Letters: []byte{'/', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J'},
	Numbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	Names: []shipDesc{
		{"A", "Carrier (5)"},
		{"B", "Submarine (4)"},
		{"C", "Destroyer (3)"},
		{"D", "Cruiser (3)"},
		{"E", "Patrol (2)"},
	},
}

// CreateFrontpage is
func CreateFrontpage() {
	log.Debug().Msg("Creating file")

	var fm = template.FuncMap{
		"decInt": decInt,
		"incInt": incInt,
	}

	shortnerTemplate := template.Must(template.New("index.html").Funcs(fm).ParseFiles("./index.html"))

	fo, err := os.Create("static/index.html")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	w := bufio.NewWriter(fo)

	shortnerTemplate.Execute(w, ts)

	if err = w.Flush(); err != nil {
		panic(err)
	}

	log.Debug().Msg("File created")
}