package reporter

import (
	"fmt"
	"github.com/gobwas/wrkp/wrk"
)

var AllFields = []wrk.Field{
	wrk.Label,
	wrk.Url,
	wrk.Threads,
	wrk.Connections,
	wrk.LatencyAvg,
	wrk.LatencyStdev,
	wrk.LatencyMax,
	wrk.LatencyDelta,
	wrk.RPSAvg,
	wrk.RPSStdev,
	wrk.RPSMax,
	wrk.RPSDelta,
	wrk.TotalRequests,
	wrk.TotalDuration,
	wrk.TotalTransfer,
	wrk.ErrorsConnect,
	wrk.ErrorsRead,
	wrk.ErrorsWrite,
	wrk.ErrorsTimeout,
	wrk.RequestsPerSec,
	wrk.TransferPerSec,
}

func AllExcept(exc ...wrk.Field) (result []wrk.Field) {
All:
	for _, f := range AllFields {
		for _, e := range exc {
			if f == e {
				continue All
			}
		}

		result = append(result, f)
	}

	return
}

func GetFieldValue(r wrk.Result, f wrk.Field) string {
	switch f {
	case wrk.Url:
		return fmt.Sprintf("%v", r.Url)

	case wrk.Threads:
		return fmt.Sprintf("%v", r.Threads)

	case wrk.Connections:
		return fmt.Sprintf("%v", r.Connections)

	case wrk.LatencyAvg:
		return fmt.Sprintf("%v", r.Latency.Average.Nanoseconds())

	case wrk.LatencyStdev:
		return fmt.Sprintf("%v", r.Latency.Stdev.Nanoseconds())

	case wrk.LatencyMax:
		return fmt.Sprintf("%v", r.Latency.Max.Nanoseconds())

	case wrk.LatencyDelta:
		return fmt.Sprintf("%v", r.Latency.Delta)

	case wrk.RPSAvg:
		return fmt.Sprintf("%v", r.RPS.Average)

	case wrk.RPSStdev:
		return fmt.Sprintf("%v", r.RPS.Stdev)

	case wrk.RPSMax:
		return fmt.Sprintf("%v", r.RPS.Max)

	case wrk.RPSDelta:
		return fmt.Sprintf("%v", r.RPS.Delta)

	case wrk.TotalRequests:
		return fmt.Sprintf("%v", r.Total.Requests)

	case wrk.TotalDuration:
		return fmt.Sprintf("%v", r.Total.Duration.Nanoseconds())

	case wrk.TotalTransfer:
		return fmt.Sprintf("%v", r.Total.Transfer)

	case wrk.ErrorsConnect:
		return fmt.Sprintf("%v", r.Errors.Connect)

	case wrk.ErrorsRead:
		return fmt.Sprintf("%v", r.Errors.Read)

	case wrk.ErrorsWrite:
		return fmt.Sprintf("%v", r.Errors.Write)

	case wrk.ErrorsTimeout:
		return fmt.Sprintf("%v", r.Errors.Timeout)

	case wrk.RequestsPerSec:
		return fmt.Sprintf("%v", r.RequestsPerSec)

	case wrk.TransferPerSec:
		return fmt.Sprintf("%v", r.TransferPerSec)

	case wrk.Label:
		return fmt.Sprintf("%v", r.Label)
	}

	return ""
}

type Reporter interface {
	Generate([]wrk.Result, []wrk.Field) ([]byte, error)
}
