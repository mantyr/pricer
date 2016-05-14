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

    price_rub_is      bool
    price_rub         string
    price_rub_int     int
    price_rub_int64   int64
    price_rub_float64 float64
}

var priceTests = []ParseTest{
    {"1 790  руб.",         "1790",    1790, 1790, 1790,    "RUB",          true,  "1790",    1790, 1790, 1790},
    {"1 790,20  руб.",      "1790.20", 1790, 1790, 1790.20, "RUB",          true,  "1790.20", 1790, 1790, 1790.20},
    {"1 790,20  руб.2",     "1790.20", 1790, 1790, 1790.20, "RUB",          true,  "1790.20", 1790, 1790, 1790.20},
    {"1 790,20",            "1790.20", 1790, 1790, 1790.20, "DEFAULT_TYPE", false, "1790.20", 1790, 1790, 1790.20},
    {"1 790,20€",           "1790.20", 1790, 1790, 1790.20, "EUR",          true, "134085.98", 134086, 134086, 134085.98},
    {"1 790,20$",           "1790.20", 1790, 1790, 1790.20, "USD",          true, "116363", 116363, 116363, 116363.00},
    {"1 790.20.123€",       "1790.20", 1790, 1790, 1790.20, "EUR",          true, "134085.98", 134086, 134086, 134085.98},
    {"1 790.20,123 EUR",    "1790.20", 1790, 1790, 1790.20, "EUR",          true, "134085.98", 134086, 134086, 134085.98},
    {"1 790.20,123 ₽",      "1790.20", 1790, 1790, 1790.20, "RUB",          true, "1790.20", 1790, 1790, 1790.20},
    {"1 790.20,123\u20BD'", "1790.20", 1790, 1790, 1790.20, "RUB",          true, "1790.20", 1790, 1790, 1790.20},
    {"",                    "0",          0,    0, 0,       "DEFAULT_TYPE", false,      "0",    0,    0,       0},
}

func TestPrices(t *testing.T) {
    p := NewPrice()
    p.SetDefaultType("DEFAULT_TYPE")

    SetCourseString("EUR", "RUB", "74.9")
    SetCourseString("USD", "RUB", "65")

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

        rub, err := p.SetConvert("RUB")
        if err != nil && test.price_type != "DEFAULT_TYPE" && rub.GetType() != "DEFAULT_TYPE" {
            t.Errorf("Error convert price to RUB, %q, %q, %q", test.source, test.price_type, err)
        }
        if rub.Get() != test.price_rub {
            t.Errorf("Error convert price to RUB price, %q, %q, %q, %q", test.source, test.price_rub, rub, rub.Get())
        }
        if rub.GetInt() != test.price_rub_int {
            t.Errorf("Error convert price to RUB prict_int, %q, %i, %i, %q", test.source, test.price_rub_int, rub.GetInt(), rub)
        }
        if rub.GetInt64() != test.price_rub_int64 {
            t.Errorf("Error convert price to RUB price_int64, %q, %i, %i", test.source, test.price_rub_int64, rub.GetInt64(), rub)
        }
        if rub.GetFloat64() != test.price_rub_float64 {
            t.Errorf("Error convert price to RUB price_float64, %q, %q", test.source, rub)
        }

        rub, err = p.SetConvertRUB()
        if err != nil && test.price_type != "DEFAULT_TYPE" && rub.GetType() != "DEFAULT_TYPE" {
            t.Errorf("Error convert price to RUB, %q, %q, %q", test.source, test.price_type, err)
        }
        if rub.Get() != test.price_rub {
            t.Errorf("Error convert price to RUB price, %q, %q, %q, %q", test.source, test.price_rub, rub, rub.Get())
        }
        if rub.GetInt() != test.price_rub_int {
            t.Errorf("Error convert price to RUB prict_int, %q, %i, %i, %q", test.source, test.price_rub_int, rub.GetInt(), rub)
        }
        if rub.GetInt64() != test.price_rub_int64 {
            t.Errorf("Error convert price to RUB price_int64, %q, %i, %i", test.source, test.price_rub_int64, rub.GetInt64(), rub)
        }
        if rub.GetFloat64() != test.price_rub_float64 {
            t.Errorf("Error convert price to RUB price_float64, %q, %q", test.source, rub)
        }

    }
}

func TestPricesModify(t *testing.T) {
    p := NewPrice()
    p.SetDefaultType("DEFAULT_TYPE")

    SetCourseString("EUR", "RUB", "74.9")
    SetCourseString("USD", "RUB", "65")
    p.Parse("120€")
    p.Plus("12")
    if p.Get() != "132" || p.GetType() != "EUR" {
        t.Errorf("Error Plus(), %q, %q", p.Get(), p.GetType())
    }
    p.PlusPercent("10")
    if p.Get() != "145.20" || p.GetType() != "EUR" {
        t.Errorf("Error PlusPercent(), %q, %q", p.Get(), p.GetType())
    }
    p.Plus("-2.15")
    if p.Get() != "143.05" || p.GetType() != "EUR" {
        t.Errorf("Error Minus, %q, %q", p.Get(), p.GetType())
    }

    p.PlusPercent("-10")
    if p.Get() != "128.75" || p.GetType() != "EUR" {
        t.Errorf("Error Minus, %q, %q", p.Get(), p.GetType())
    }
    p.PlusPercent("-10%")
    if p.Get() != "115.87" || p.GetType() != "EUR" {
        t.Errorf("Error Minus, %q, %q, %q", p.Get(), p.GetType(), p.Price_rat.FloatString(10))
    }
    p.Plus("-10%")
    if p.Get() != "104.28" || p.GetType() != "EUR" {
        t.Errorf("Error Minus, %q, %q, %q", p.Get(), p.GetType(), p.Price_rat.FloatString(10))
    }
}