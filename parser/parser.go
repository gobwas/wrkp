package parser

import (
	"github.com/gobwas/wrkp/wrk"
)

type Parser interface {
	Parse() ([]wrk.Result, error)
}
