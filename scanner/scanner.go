package scanner

type Token int

const (
	EOF Token = iota
	Url
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

func (t Token) String() string {
	switch t {
	case EOF:
		return "EOF"

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
	}

	return ""
}

type Scanner interface {
	Scan() (Token, []byte, error)
}
