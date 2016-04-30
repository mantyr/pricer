# Golang Pricer - parser for money string

[![Build Status](https://travis-ci.org/mantyr/pricer.svg?branch=master)](https://travis-ci.org/mantyr/pricer) [![GoDoc](https://godoc.org/github.com/mantyr/pricer?status.png)](http://godoc.org/github.com/mantyr/pricer) [![Software License](https://img.shields.io/badge/license-The%20Not%20Free%20License,%20Commercial%20License-brightgreen.svg)](LICENSE.md)

This stable version

## Testing
```GO
type ParseTest struct {
    source        string
    price         string
    price_int     int
    price_int64   int64
    price_float64 float64
    price_type    string
}

var priceTests = []ParseTest{
    {"1 790  руб.",      "1790",    1790, 1790, 1790,    "RUB"},
    {"1 790,20  руб.",   "1790.20", 1790, 1790, 1790.20, "RUB"},
    {"1 790,20  руб.2",  "1790.20", 1790, 1790, 1790.20, "RUB"},
    {"1 790,20",         "1790.20", 1790, 1790, 1790.20, "DEFAULT_TYPE"},
    {"1 790,20€",        "1790.20", 1790, 1790, 1790.20, "EUR"},
    {"1 790.20.123€",    "1790.20", 1790, 1790, 1790.20, "EUR"},
    {"1 790.20,123 EUR", "1790.20", 1790, 1790, 1790.20, "EUR"},
}
```

## Installation

    $ go get github.com/mantyr/pricer

## Example
```GO
package main

import (
    "github.com/mantyr/pricer"
    "fmt"
)

func main() {
    price := pricer.NewPrice()
    price.SetDefaultType("DEFAULT_TYPE")

    price.Parse("1 179.20 $")

    fmt.Println(price.Get())        // print "1179.20"
    fmt.Println(price.GetInt())     // print "1179"
    fmt.Println(price.GetInt64())   // print "1179"
    fmt.Println(price.GetFloat64()) // print "1179.2"
    fmt.Println(price.GetType())    // print "USD"

    price.Parse("1 179.21 руб.")

    fmt.Println(price.Get())        // print "1179.21"
    fmt.Println(price.GetInt())     // print "1179"
    fmt.Println(price.GetInt64())   // print "1179"
    fmt.Println(price.GetFloat64()) // print "1179.21"
    fmt.Println(price.GetType())    // print "RUB"

    price.Parse("1 179.21")

    fmt.Println(price.Get())        // print "1179.21"
    fmt.Println(price.GetInt())     // print "1179"
    fmt.Println(price.GetInt64())   // print "1179"
    fmt.Println(price.GetFloat64()) // print "1179.21"
    fmt.Println(price.GetType())    // print "DEFAULT_TYPE"
}
```

## Author

[Oleg Shevelev][mantyr]

[mantyr]: https://github.com/mantyr
