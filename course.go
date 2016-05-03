package pricer

import (
    "errors"
    "math/big"
)

func NewCourse() (c *ConvertPriceType) {
    c = new(ConvertPriceType)
    c.d = make(map[string]map[string]*big.Rat)
    return
}

func (c *ConvertPriceType) Set(source, destination string, ratio *big.Rat) *ConvertPriceType {
    c.Lock()
    defer c.Unlock()

    _, ok := c.d[source][destination]
    if !ok {
        c.d[source] = make(map[string]*big.Rat)
    }
    c.d[source][destination] = ratio

    return c
}

func (c *ConvertPriceType) Get(source, destination string) (ratio *big.Rat, ok bool) {
    c.RLock()
    defer c.RUnlock()

    ratio, ok = c.d[source][destination]
    return ratio, ok
}

var course_price *ConvertPriceType

func init() {
    course_price = NewCourse()
}

func SetCourse(source, destination string, ratio *big.Rat) *ConvertPriceType {
    return course_price.Set(source, destination, ratio)
}

func SetCourseString(source, destination string, ratio string) *ConvertPriceType {
    r := new(big.Rat)
    r.SetString(ratio)
    return course_price.Set(source, destination, r)
}

func GetCourse(source, destination string) (ratio *big.Rat, ok bool) {
    return course_price.Get(source, destination)
}

func (p *Price) SetConvert(destination string) (price *Price, err error) {
    if p.GetType() == destination {
        return p, nil
    }

    ratio, ok := course_price.Get(p.GetType(), destination)
    if !ok {
        return p, errors.New("Convert course no destination, "+destination)
    }

    price = NewPrice()
    price.Price_rat.Mul(p.Price_rat, ratio)

    price.Price        = price.Get()
    price.Price_source = price.Price+" "+destination
    price.Price_type   = destination

    return price, nil
}

func (p *Price) SetConvertRUB() (*Price, error) {
    return p.SetConvert("RUB")
}

func (p *Price) SetConvertUSD() (*Price, error) {
    return p.SetConvert("USD")
}

func (p *Price) SetConvertEUR() (*Price, error) {
    return p.SetConvert("EUR")
}

