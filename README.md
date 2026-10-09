# Time Parser for Golang

[![GitHub Releases](https://img.shields.io/github/v/release/nhatthm/timeparser)](https://github.com/nhatthm/timeparser/releases/latest)
[![Build Status](https://github.com/nhatthm/timeparser/actions/workflows/test.yaml/badge.svg)](https://github.com/nhatthm/timeparser/actions/workflows/test.yaml)
[![codecov](https://codecov.io/gh/nhatthm/timeparser/branch/master/graph/badge.svg?token=eTdAgDE2vR)](https://codecov.io/gh/nhatthm/timeparser)
[![GoDevDoc](https://img.shields.io/badge/dev-doc-00ADD8?logo=go)](https://pkg.go.dev/go.nhat.io/timeparser)
[![Donate](https://img.shields.io/badge/%20-Donate-%20?style=flat&logo=githubsponsors&color=E5E4E2)](http://donate.nhat.me)

`timeparser` provides flexibility in parsing time from string for Golang. It allows either `RFC3339` or `YMD`.

## Prerequisites

- `Go >= 1.23`

## Install

```bash
go get go.nhat.io/timeparser
```

## Usage

### `func Parse(s string) (time.Time, error)`

Parse a time in `string` to `time.Time`. `s` could be `RFC3339` or `YMD`.

### `func ParsePeriod(from, to string) (start *time.Time, end *time.Time, err error)`

Parse a time period from `string` to `time.Time`.

`from` and `to` could be `RFC3339` or `YMD`. It is `nil` if the string is empty.

## Donation

If this project saved you some development time, buy me a cup of coffee :)

[![donate](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](http://donate.nhat.me)

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;or scan this

<img src="https://github.com/nhatthm/donate.nhat.me/blob/master/images/qr_sponsor.png" width="147px" />

[<sub><sup>[table of contents]</sup></sub>](#table-of-contents)
