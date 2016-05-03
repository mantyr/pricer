package pricer

import (
    "strconv"
    "strings"
    "regexp"
    "math/big"
)

var regexp_numbers = regexp.MustCompile("[.,0-9]+")

func NewPrice() (p *Price) {
    p = new(Price)
    p.default_type = "RUB"
    p.Price_rat    = new(big.Rat)
    return p
}

func (p *Price) SetDefaultType(name string) *Price {
    p.default_type = name
    return p
}

func (p *Price) Parse(price string) *Price {
    p.Price_source = price
    p.Price        = p.parse_value()
    p.Price_type   = p.parse_type()

    p.Price_rat.SetString(p.Price)
    return p
}

func (p *Price) SetFloat64(price float64, price_type ...string) *Price {
    p.Price_rat.SetFloat64(price)
    if len(price_type) > 0 {
        p.Price_type = price_type[0]
    }
    p.Price_source = p.String()+" "+p.Price_type
    p.Price        = p.Price_source
    return p
}

func init() {
    regexp_numbers = regexp.MustCompile("[,.0-9]+")
}

func (p *Price) parse_value() string {
    v := strings.Join(regexp_numbers.FindAllString(p.Price_source, -1), "")
    v  = strings.Replace(v, ",", ".", -1)

    c := strings.IndexRune(v, '.')
    if c > -1 {
        c2 := strings.IndexRune(v[c+1:], '.')
        if c2 > -1 {
            v = v[:c+c2+2]
        }
    }

    if v[len(v)-1:] == "." {
        v = v[:len(v)-1]
    }
    return v
}

func (p *Price) parse_type() string {
    for _, price := range PriceTypes {
        if strings.Index(p.Price_source, price.Search) > -1 {
            return price.Result
        }
    }
    return p.default_type
}

// todo: Add a standard format for national currencies
func (p *Price) String() string {
    return p.Get()+" "+p.Price_type
}

func (p *Price) Get() string {
    price := p.Price_rat.FloatString(2)
    if price[len(price)-2:] == "00" {
        return price[:len(price)-3]
    }
    return price
}

func (p *Price) GetFloatString(prec int) string {
    return p.Price_rat.FloatString(prec)
}

func (p *Price) GetInt() int {
    price := p.GetFloatString(0)

    v, err := strconv.Atoi(price)
    if err != nil {
        return 0
    }
    return v
}

func (p *Price) GetInt64() int64 {
    price := p.GetFloatString(0)

    v, err := strconv.ParseInt(price, 10, 64)
    if err != nil {
        return 0
    }
    return v
}

func (p *Price) GetFloat64() float64 {
    v, _ := p.Price_rat.Float64()
    return v
}

func (p *Price) GetType() string {
    return p.Price_type
}