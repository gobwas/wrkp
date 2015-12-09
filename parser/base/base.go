package base

import (
	"bytes"
	"fmt"
	"github.com/gobwas/wrkp/scanner"
	"github.com/gobwas/wrkp/wrk"
	"io"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

type Parser struct {
	s scanner.Scanner
}

func New(s scanner.Scanner) *Parser {
	return &Parser{s}
}

func (self *Parser) Parse() (*wrk.Result, error) {
	var result wrk.Result

	for {
		t, b, err := self.s.Scan()
		if err != nil {
			if err == io.EOF {
				return &result, nil
			}

			return nil, err
		}

		switch t {
		case scanner.Url:
			u, err := url.Parse(string(b))
			if err != nil {
				return nil, err
			}

			result.Url = u

		case scanner.Threads:
			v, err := strconv.ParseInt(string(b), 10, 64)
			if err != nil {
				return nil, err
			}

			result.Threads = v

		case scanner.Connections:
			v, err := strconv.ParseInt(string(b), 10, 64)
			if err != nil {
				return nil, err
			}

			result.Connections = v

		case scanner.LatencyAvg:
			v, err := time.ParseDuration(string(b))
			if err != nil {
				return nil, err
			}

			result.Latency.Average = v

		case scanner.LatencyStdev:
			v, err := time.ParseDuration(string(b))
			if err != nil {
				return nil, err
			}

			result.Latency.Stdev = v

		case scanner.LatencyMax:
			v, err := time.ParseDuration(string(b))
			if err != nil {
				return nil, err
			}

			result.Latency.Max = v

		case scanner.LatencyDelta:
			v, err := strconv.ParseFloat(string(bytes.TrimSuffix(b, []byte{'%'})), 64)
			if err != nil {
				return nil, err
			}

			result.Latency.Delta = v

		case scanner.RPSAvg:
			v, err := parse(string(b), siMap)
			if err != nil {
				return nil, err
			}

			result.RPS.Average = v

		case scanner.RPSStdev:
			v, err := parse(string(b), siMap)
			if err != nil {
				return nil, err
			}

			result.RPS.Stdev = v

		case scanner.RPSMax:
			v, err := parse(string(b), siMap)
			if err != nil {
				return nil, err
			}

			result.RPS.Max = v

		case scanner.RPSDelta:
			v, err := strconv.ParseFloat(string(bytes.TrimSuffix(b, []byte{'%'})), 64)
			if err != nil {
				return nil, err
			}

			result.RPS.Delta = v

		case scanner.TotalRequests:
			v, err := strconv.ParseInt(string(b), 10, 64)
			if err != nil {
				return nil, err
			}

			result.Total.Requests = v

		case scanner.TotalDuration:
			v, err := time.ParseDuration(string(b))
			if err != nil {
				return nil, err
			}

			result.Total.Duration = v

		case scanner.TotalTransfer:
			v, err := parse(string(b), fileSizeMap)
			if err != nil {
				return nil, err
			}

			result.Total.Transfer = v

		case scanner.ErrorsConnect:
			v, err := strconv.ParseInt(string(b), 10, 64)
			if err != nil {
				return nil, err
			}

			result.Errors.Connect = v

		case scanner.ErrorsRead:
			v, err := strconv.ParseInt(string(b), 10, 64)
			if err != nil {
				return nil, err
			}

			result.Errors.Read = v

		case scanner.ErrorsWrite:
			v, err := strconv.ParseInt(string(b), 10, 64)
			if err != nil {
				return nil, err
			}

			result.Errors.Write = v

		case scanner.ErrorsTimeout:
			v, err := strconv.ParseInt(string(b), 10, 64)
			if err != nil {
				return nil, err
			}

			result.Errors.Timeout = v

		case scanner.RequestsPerSec:
			v, err := strconv.ParseFloat(string(b), 64)
			if err != nil {
				return nil, err
			}

			result.RequestsPerSec = v

		case scanner.TransferPerSec:
			v, err := parse(string(b), fileSizeMap)
			if err != nil {
				return nil, err
			}

			result.TransferPerSec = v
		}
	}

	return nil, fmt.Errorf("unexpected error")
}

const (
	_          = iota // ignore first value by assigning to blank identifier
	KB float64 = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

const (
	Kilo float64 = 1000
	Mega         = Kilo * 1000
	Giga         = Mega * 1000
	Tera         = Giga * 1000
	Peta         = Tera * 1000
	Exa          = Peta * 1000
)

var fileSizeMap map[string]float64 = map[string]float64{
	"KB": KB,
	"MB": MB,
	"GB": GB,
	"TB": TB,
	"PB": PB,
	"EB": EB,
	"ZB": ZB,
	"YB": YB,
}

var siMap map[string]float64 = map[string]float64{
	"k": Kilo,
	"M": Mega,
	"G": Giga,
	"T": Tera,
	"P": Peta,
	"E": Exa,
}

var numeric = regexp.MustCompile(`^([0-9\.]+)([a-zA-Z]*)$`)

func parse(s string, m map[string]float64) (float64, error) {
	matches := numeric.FindStringSubmatch(s)
	if len(matches) != 3 {
		return 0, fmt.Errorf("could not parse file size: %s", s)
	}

	val, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return 0, err
	}

	if matches[2] == "" {
		return val, nil
	}

	for s, k := range m {
		if s == matches[2] {
			return val * k, nil
		}
	}

	return 0, fmt.Errorf("Unknown size: %s", matches[2])
}
