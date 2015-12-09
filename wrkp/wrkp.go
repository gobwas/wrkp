package main

import (
	"flag"
	"fmt"
	baseParser "github.com/gobwas/wrkp/parser/base"
	"github.com/gobwas/wrkp/reporter"
	csvReporter "github.com/gobwas/wrkp/reporter/csv"
	baseScanner "github.com/gobwas/wrkp/scanner/base"
	"github.com/gobwas/wrkp/wrk"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var file = flag.String("f", "", "use this file pattern instead of stdin (allows glob)")
var report = &ReporterFlag{report_csv, []string{report_csv}}

//var delimiter = flag.String("d", ",", "delimiter for csv report")

const (
	report_csv = "csv"
)

type ReporterFlag struct {
	v string
	e []string
}

func (r *ReporterFlag) Set(s string) error {
	for _, e := range r.e {
		if e == s {
			r.v = s
			return nil
		}
	}

	return fmt.Errorf("expecting one of %s", r.e)
}
func (r ReporterFlag) String() string {
	return r.v
}
func (r ReporterFlag) Get() interface{} {
	return r.v
}

func main() {
	flag.Var(report, "resp", fmt.Sprintf("report type (available: %s)", strings.Join(report.e, ", ")))
	flag.Parse()
	log.SetFlags(0)

	var rep reporter.Reporter
	switch report.Get() {
	case report_csv:
		rep = csvReporter.New(',')
	default:
		log.Println("unknown report")
		os.Exit(1)
	}

	var (
		readers []io.Reader
		labels  []string
		fields  []wrk.Field
	)
	if *file != "" {
		p, err := filepath.Abs(*file)
		if err != nil {
			log.Println("resolve path error: ", err)
			os.Exit(1)
		}

		matches, err := filepath.Glob(p)
		if err != nil {
			log.Println("glob error: ", err)
			os.Exit(1)
		}

		for _, f := range matches {
			file, err := os.OpenFile(f, os.O_RDONLY, 0)
			if err != nil {
				log.Printf("open file %s error: %s\n", f, err)
				os.Exit(1)
			}

			labels = append(labels, filepath.Base(f))
			readers = append(readers, file)
		}

		fields = reporter.AllFields
	} else {
		readers = []io.Reader{os.Stdin}
		fields = reporter.AllExcept(wrk.Label)
	}

	var results []wrk.Result
	for i, rd := range readers {
		p := baseParser.New(baseScanner.New(rd))
		result, err := p.Parse()
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}

		if label, ok := labels[i]; ok {
			result.Label = label
		}

		results = append(results, *result)
	}

	b, err := rep.Generate(results, fields)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Fprint(os.Stdout, string(b))
	os.Exit(0)
}
