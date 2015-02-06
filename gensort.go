// Binary tmpl TODO
package main

import (
	"flag"
	"log"
	"os"
	"text/template"
)

var (
	output      = flag.String("output", "/dev/stdout", "The name of the file to write generated contents to")
	target      = flag.String("type", "int", "The type of the slice to generate sort function for")
	packageName = flag.String("package", "main", "The package name to put generated sort function under")
	funcName    = flag.String("name", "quickSort", "The name of the sort func")
	less        = flag.String("less", "x < y", "The 'less' expression to use")
	eq          = flag.String("eq", "x == y", "The 'equal' expression to use")

	tmpl = `/* generated by gensort */
package {{.package}}

import (
	"math/rand"
	"sync"
)

func {{.name}}(xs []{{.T}}) {
	const threshold = 1000

	partition := func(xs []{{.T}}) (low, high []{{.T}}) {
		var (
			chosen = rand.Intn(len(xs))
			y  = xs[chosen]
		)

		i, j := 0, 0
		for k, x := range xs {
			if {{.less}} {
				xs[k] = xs[j]
				xs[j] = xs[i]
				xs[i] = x
				i++
				j++
			} else if {{.eq}} {
				xs[k] = xs[j]
				xs[j] = x
				j++
			}
		}

		return xs[:i], xs[j:]
	}

	var wg sync.WaitGroup
	var {{.name}} func([]{{.T}})
	{{.name}} = func(xs []{{.T}}) {
	init:
		if len(xs) <= 1 {
			return
		}

		low, high := partition(xs)

		var short, long []{{.T}}
		if len(low) < len(high) {
			short, long = low, high
		} else {
			short, long = high, low
		}

    if len(short) > threshold {
      wg.Add(1)
      go func() {
        {{.name}}(short)
        wg.Done()
      }()
    } else {
      {{.name}}(short)
    }
    xs = long
    goto init
  }

  {{.name}}(xs)
  wg.Wait()
}
`
)

func main() {
	flag.Parse()

	file, err := os.Create(*output)
	if err != nil {
		log.Fatal(err)
	}

	model := map[string]string{
		"T":       *target,
		"package": *packageName,
		"less":    *less,
		"eq":      *eq,
		"name":    *funcName,
	}

	t := template.Must(template.New("Main").Parse(tmpl))
	if err := t.Execute(file, model); err != nil {
		log.Fatal(err)
	}
}