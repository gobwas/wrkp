package csv

import (
	"bytes"
	"encoding/csv"
	"github.com/gobwas/wrkp/reporter"
	"github.com/gobwas/wrkp/wrk"
)

type Reporter struct {
	delimiter rune
}

func New(d rune) *Reporter {
	return &Reporter{d}
}

//name;latency avg;latency stdev;latency max;latency +/-;rps avg;rps stdev;rps max;rps +/-;total req;total time;total traffic;errors connect;errors read;errors write;errors timeout;rps;transfersec;
func (self *Reporter) Generate(r []wrk.Result, f []wrk.Field) (b []byte, err error) {
	lines := make(map[int][]string)

	for _, field := range f {
		lines[0] = append(lines[0], field.String())

		for line, result := range r {
			lines[line+1] = append(lines[line+1], reporter.GetFieldValue(result, field))
		}
	}

	buf := bytes.Buffer{}
	c := csv.NewWriter(&buf)
	c.Comma = self.delimiter

	for _, l := range lines {
		c.Write(l)
	}

	c.Flush()

	return buf.Bytes(), nil
}
