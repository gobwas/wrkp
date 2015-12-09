# :dizzy_face: [wrk](https://github.com/wg/wrk)p

[![GoDoc](https://godoc.org/github.com/gobwas/wrkp?status.svg)](https://godoc.org/github.com/gobwas/wrkp)

> wrk result go parser

## Install

```shell
go get github.com/gobwas/wrkp...
```

## Usage

```shell
wrk example.com | wrkp > report.csv
```

Or

```shell
wrk example.com > example.com.wrk
...

wrkp -f="$HOME/*.wrk" -r=csv
```