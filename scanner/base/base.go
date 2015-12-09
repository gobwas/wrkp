package base

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/gobwas/wrkp/scanner"
	"io"
	"sync"
)

type Scanner struct {
	mu   sync.Mutex
	s    *bufio.Scanner
	buf  []word
	prev []byte
}

func New(rd io.Reader) *Scanner {
	s := bufio.NewScanner(rd)
	s.Split(bufio.ScanWords)

	return &Scanner{s: s}
}

func (self *Scanner) pushBuf(w ...word) {
	self.mu.Lock()
	defer self.mu.Unlock()
	self.buf = append(self.buf, w...)
}

func (self *Scanner) shiftBuf() (w word) {
	self.mu.Lock()
	defer self.mu.Unlock()
	w = self.buf[0]
	self.buf = self.buf[1:]

	return
}

func (self *Scanner) setPrev(w []byte) {
	self.mu.Lock()
	defer self.mu.Unlock()
	self.prev = w
}

func (self *Scanner) flushPrev() (w []byte) {
	self.mu.Lock()
	defer self.mu.Unlock()
	w = self.prev
	self.prev = nil

	return
}

func (self *Scanner) Scan() (scanner.Token, []byte, error) {
	if len(self.buf) > 0 {
		w := self.shiftBuf()
		return w.token, w.b, nil
	}

	if !self.s.Scan() {
		return scanner.EOF, nil, io.EOF
	}

	switch self.s.Text() {

	case "Running":
		words, err := self.readNext(4)
		if err != nil {
			return scanner.EOF, nil, io.EOF
		}

		return scanner.Url, words[3], nil

	case "threads":
		words, err := self.readNext(3)
		if err != nil {
			return scanner.EOF, nil, io.EOF
		}

		self.pushBuf(word{scanner.Connections, words[1]})

		return scanner.Threads, self.flushPrev(), nil

	case "Latency":
		words, err := self.readNext(4)
		if err != nil {
			return scanner.EOF, nil, io.EOF
		}

		b := []word{
			word{scanner.LatencyStdev, words[1]},
			word{scanner.LatencyMax, words[2]},
			word{scanner.LatencyDelta, words[3]},
		}

		self.pushBuf(b...)

		return scanner.LatencyAvg, words[0], nil

	case "Req/Sec":
		words, err := self.readNext(4)
		if err != nil {
			return scanner.EOF, nil, io.EOF
		}

		b := []word{
			word{scanner.RPSStdev, words[1]},
			word{scanner.RPSMax, words[2]},
			word{scanner.RPSDelta, words[3]},
		}

		self.pushBuf(b...)

		return scanner.RPSAvg, words[0], nil

	case "requests":
		words, err := self.readNext(4)
		if err != nil {
			return scanner.EOF, nil, io.EOF
		}

		b := []word{
			word{scanner.TotalDuration, bytes.Trim(words[1], ",")},
			word{scanner.TotalTransfer, bytes.Trim(words[2], ",")},
		}

		self.pushBuf(b...)

		return scanner.TotalRequests, self.flushPrev(), nil

	case "errors":
		words, err := self.readNext(8)
		if err != nil {
			return scanner.EOF, nil, io.EOF
		}

		b := []word{
			word{scanner.ErrorsRead, bytes.Trim(words[3], ",")},
			word{scanner.ErrorsWrite, bytes.Trim(words[5], ",")},
			word{scanner.ErrorsTimeout, bytes.Trim(words[7], ",")},
		}

		self.pushBuf(b...)

		return scanner.ErrorsConnect, words[1], nil

	case "Requests/sec:":
		words, err := self.readNext(1)
		if err != nil {
			return scanner.EOF, nil, io.EOF
		}

		return scanner.RequestsPerSec, words[0], nil

	case "Transfer/sec:":
		words, err := self.readNext(1)
		if err != nil {
			return scanner.EOF, nil, io.EOF
		}

		return scanner.TransferPerSec, words[0], nil
	}

	self.setPrev(self.s.Bytes())

	return self.Scan()
}

func (self *Scanner) readNext(l int) (map[int][]byte, error) {
	var index int

	words := make(map[int][]byte)

	for index < l {
		if !self.s.Scan() {
			return words, fmt.Errorf("unexpected end of input")
		}

		words[index] = self.s.Bytes()

		index++
	}

	return words, nil
}

type word struct {
	token scanner.Token
	b     []byte
}

func indexOf(l []int, n int) int {
	for i, v := range l {
		if v == n {
			return i
		}
	}

	return -1
}
