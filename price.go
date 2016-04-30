package pricer

import (
    "strconv"
    "strings"
    "regexp"
    "fmt"
)

var (
    regexp_numbers *regexp.Regexp
)

func init() {
    regexp_numbers = regexp.MustCompile("[.0-9]+")
}


func NewPrice() (p *Price) {
    p = new(Price)
    p.default_type = "RUB"
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

func (p *Price) Get() string {
    return p.Price
}

func (p *Price) GetInt() int {
    price := p.Price
    c := strings.IndexRune(price, '.')
    if c > -1 {
        price = price[:c]
    }

    v, err := strconv.Atoi(price)
    if err != nil {
        fmt.Println(err)
        return 0
    }
    return v
}

func (p *Price) GetInt64() int64 {
    price := p.Price
    c := strings.IndexRune(price, '.')
    if c > -1 {
        price = price[:c]
    }

    v, err := strconv.ParseInt(price, 10, 64)
    if err != nil {
        return 0
    }
    return v
}

func (p *Price) GetFloat64() float64 {
    v, err := strconv.ParseFloat(p.Price, 64)
    if err != nil {
        return 0
    }
    return v
}

func (p *Price) GetType() string {
    return p.Price_type
}