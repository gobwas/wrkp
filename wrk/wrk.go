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
	Label
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
		return "Thread Latency Avg"

	case LatencyStdev:
		return "Thread Latency Stdev"

	case LatencyMax:
		return "Thread Latency Max"

	case LatencyDelta:
		return "Thread Latency Delta"

	case RPSAvg:
		return "Thread RPS Avg"

	case RPSStdev:
		return "Thread RPS Stdev"

	case RPSMax:
		return "Thread RPS Max"

	case RPSDelta:
		return "Thread RPS Delta"

	case TotalRequests:
		return "Total Requests"

	case TotalDuration:
		return "Total Duration"

	case TotalTransfer:
		return "Total Transfer"

	case ErrorsConnect:
		return "Connect Errors"

	case ErrorsRead:
		return "Read Errors"

	case ErrorsWrite:
		return "Write Errors"

	case ErrorsTimeout:
		return "Timeout Errors"

	case RequestsPerSec:
		return "Requests Per Second"

	case TransferPerSec:
		return "Transfer Per Second"

	case Label:
		return "Label"

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
	Label          string
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
