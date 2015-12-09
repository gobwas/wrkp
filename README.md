# :dizzy_face: [wrk](https://github.com/wg/wrk)p

> wrk result go parser

## Install

```shell
go get github.com/gobwas/wrkp
```

## Usage

```shell
wrk example.com | wrkp > report.csv
```

Or

```shell
wrk example.com > example.com.wrk
...

wrkp -f *.wrk -r csv
```

## Docs

Are [here](https://godoc.org/github.com/gobwas/wrkp).

