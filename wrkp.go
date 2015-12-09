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
	"os"
	"path/filepath"
	"strings"
)

var file = flag.String("f", "", "file to parse (allows glob)")
var report = &ReporterFlag{report_csv, []string{report_csv}}

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
	flag.Var(report, "resp", fmt.Sprintf("how should server response on message (%s)", strings.Join(report.e, ", ")))
	flag.Parse()

	var readers []io.Reader

	if *file != "" {
		p, err := filepath.Abs(*file)
		if err != nil {
			fmt.Println("resolve path error: ", err)
			os.Exit(1)
		}

		matches, err := filepath.Glob(p)
		if err != nil {
			fmt.Println("glob error: ", err)
			os.Exit(1)
		}

		for _, f := range matches {
			file, err := os.OpenFile(f, os.O_RDONLY, 0)
			if err != nil {
				fmt.Printf("open file %s error: %s\n", f, err)
				os.Exit(1)
			}

			readers = append(readers, file)
		}
	} else {
		readers = []io.Reader{os.Stdin}
	}

	var results []wrk.Result
	for _, rd := range readers {
		p := baseParser.New(baseScanner.New(rd))
		result, err := p.Parse()
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}

		results = append(results, *result)
	}

	var rep reporter.Reporter
	switch report.Get() {
	case report_csv:
		rep = csvReporter.New(';')
	default:
		fmt.Println("unknown report")
		os.Exit(1)
	}

	b, err := rep.Generate(results, reporter.AllFields)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(b))
}
