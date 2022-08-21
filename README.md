# Time Parser for Golang

[![GitHub Releases](https://img.shields.io/github/v/release/nhatthm/timeparser)](https://github.com/nhatthm/timeparser/releases/latest)
[![Build Status](https://github.com/nhatthm/timeparser/actions/workflows/test.yaml/badge.svg)](https://github.com/nhatthm/{}name/actions/workflows/test.yaml)
[![codecov](https://codecov.io/gh/nhatthm/timeparser/branch/master/graph/badge.svg?token=eTdAgDE2vR)](https://codecov.io/gh/nhatthm/timeparser)
[![Go Report Card](https://goreportcard.com/badge/go.nhat.io/timeparser)](https://goreportcard.com/report/go.nhat.io/timeparser)
[![GoDevDoc](https://img.shields.io/badge/dev-doc-00ADD8?logo=go)](https://pkg.go.dev/go.nhat.io/timeparser)
[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.com/donate/?hosted_button_id=PJZSGJN57TDJY)

`timeparser` provides flexibility in parsing time from string for Golang. It allows either `RFC3339` or `YMD`.

## Prerequisites

- `Go >= 1.17`

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

If this project help you reduce time to develop, you can give me a cup of coffee :)

### Paypal donation

[![paypal](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/donate/?hosted_button_id=PJZSGJN57TDJY)

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;or scan this

<img src="https://user-images.githubusercontent.com/1154587/113494222-ad8cb200-94e6-11eb-9ef3-eb883ada222a.png" width="147px" />
