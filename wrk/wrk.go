package wrk

import (
	"net/url"
	"time"
)

type Field int

const (
	Url Field = iota
	Threads
	Connections
	LatencyAvg
	LatencyStdev
	LatencyMax
	LatencyDelta
	RPSAvg
	RPSStdev
	RPSMax
	RPSDelta
	TotalRequests
	TotalDuration
	TotalTransfer
	ErrorsConnect
	ErrorsRead
	ErrorsWrite
	ErrorsTimeout
	RequestsPerSec
	TransferPerSec
)

func (f Field) String() string {
	switch f {
	case Url:
		return "Url"

	case Threads:
		return "Threads"

	case Connections:
		return "Connections"

	case LatencyAvg:
		return "LatencyAvg"

	case LatencyStdev:
		return "LatencyStdev"

	case LatencyMax:
		return "LatencyMax"

	case LatencyDelta:
		return "LatencyDelta"

	case RPSAvg:
		return "RPSAvg"

	case RPSStdev:
		return "RPSStdev"

	case RPSMax:
		return "RPSMax"

	case RPSDelta:
		return "RPSDelta"

	case TotalRequests:
		return "TotalRequests"

	case TotalDuration:
		return "TotalDuration"

	case TotalTransfer:
		return "TotalTransfer"

	case ErrorsConnect:
		return "ErrorsConnect"

	case ErrorsRead:
		return "ErrorsRead"

	case ErrorsWrite:
		return "ErrorsWrite"

	case ErrorsTimeout:
		return "ErrorsTimeout"

	case RequestsPerSec:
		return "RequestsPerSec"

	case TransferPerSec:
		return "TransferPerSec"

	default:
		return ""
	}
}

type Latency struct {
	Average, Stdev, Max time.Duration
	Delta               float64
}

type RPS struct {
	Average, Stdev, Max, Delta float64
}

type Total struct {
	Requests int64
	Duration time.Duration
	Transfer float64
}

type Errors struct {
	Connect, Read, Write, Timeout int64
}

type Result struct {
	Url            *url.URL
	Threads        int64
	Connections    int64
	RequestsPerSec float64
	TransferPerSec float64
	Latency        Latency
	RPS            RPS
	Total          Total
	Errors         Errors
}
