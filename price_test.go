package pricer

import (
    "testing"
)

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
    {"1 790,20$",        "1790.20", 1790, 1790, 1790.20, "USD"},
    {"1 790.20.123€",    "1790.20", 1790, 1790, 1790.20, "EUR"},
    {"1 790.20,123 EUR", "1790.20", 1790, 1790, 1790.20, "EUR"},
}

func TestPrices(t *testing.T) {
    p := NewPrice()
    p.SetDefaultType("DEFAULT_TYPE")

    for _, test := range priceTests {
        p.Parse(test.source)
        if p.Get() != test.price {
            t.Errorf("Error string price, %q, %q", test.source, p.Get())
        }
        if p.GetInt() != test.price_int {
            t.Errorf("Error string price_int, %q, %q", test.source, p.GetInt())
        }
        if p.GetInt64() != test.price_int64 {
            t.Errorf("Error string price_int64, %q, %q", test.source, p.GetInt64())
        }
        if p.GetFloat64() != test.price_float64 {
            t.Errorf("Error string price_float64, %q, %q", test.source, p.GetFloat64())
        }
        if p.GetType() != test.price_type {
            t.Errorf("Error string price_type, %q, %q", test.source, p.GetType())
        }
    }
}

