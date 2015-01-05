package restrictednumber

import (
	"errors"
	"math"
)

var (
	ErrMinGreaterThanMax = errors.New("RestrictedNumber: min cannot be larger than max")
	ErrMaxSmallerThanMin = errors.New("RestrictedNumber: max cannot be lower than min")
)

type RestrictedNumber struct {
	min, max, cur int
	Name          string
}

func New() *RestrictedNumber {
	rs := new(RestrictedNumber)
	return rs
}

func NewSet(min, max, cur int) *RestrictedNumber {
	rs := new(RestrictedNumber)
	rs.min = min
	rs.max = max
	rs.SetVal(cur)
	return rs
}

func (rn *RestrictedNumber) Val() int {
	return rn.cur
}

func (rn *RestrictedNumber) Min() int {
	return rn.min
}

func (rn *RestrictedNumber) Max() int {
	return rn.max
}

func (rn *RestrictedNumber) SetMin(num int) error {

	if num > rn.max {
		return ErrMinGreaterThanMax
	}

	rn.min = num
	rn.SetVal(rn.cur)

	return nil
}

func (rn *RestrictedNumber) SetMax(num int) error {

	if num < rn.min {
		return ErrMaxSmallerThanMin
	}

	rn.max = num
	rn.SetVal(rn.cur)

	return nil
}

func (rn *RestrictedNumber) SetVal(num int) {
	// Math.min
	if num < rn.min {
		num = rn.min
	}

	// Math.max
	if num > rn.max {
		num = rn.max
	}

	rn.cur = num
}

// Setter functions
func (rn *RestrictedNumber) Add(num int) {
	rn.SetVal(rn.cur + num)
}

func (rn *RestrictedNumber) Sub(num int) {
	rn.Add(-num)
}

func (rn *RestrictedNumber) ToMax() {
	rn.SetVal(rn.max)
}

func (rn *RestrictedNumber) ToMin() {
	rn.SetVal(rn.min)
}

// Percentifying functions
func (rn *RestrictedNumber) AsPercent() int {
	return int(math.Floor(float64(rn.cur) / float64(rn.max) * 100))
}

func (rn *RestrictedNumber) getPercent(perc int) int {
	return int(math.Floor(float64(perc) * float64(rn.max) / 100))
}

func (rn *RestrictedNumber) SetToPercent(perc int) {
	rn.SetVal(rn.getPercent(perc))
}

func (rn *RestrictedNumber) AddByPercent(perc int) {
	rn.Add(rn.getPercent(perc))
}

func (rn *RestrictedNumber) SubByPercent(perc int) {
	rn.AddByPercent(-perc)
}

// Checker functions

func (rn *RestrictedNumber) AtMin() bool {
	return rn.min == rn.cur
}

func (rn *RestrictedNumber) AtMax() bool {
	return rn.max == rn.cur
}

func (rn *RestrictedNumber) IsPercent(perc int) bool {
	return rn.cur == rn.getPercent(perc)
}

func (rn *RestrictedNumber) LessThanPercent(perc int) bool {
	return !rn.GreaterThanEqualsPercent(perc)
}

func (rn *RestrictedNumber) LessThanEqualsPercent(perc int) bool {
	return !rn.GreaterThanPercent(perc)
}

func (rn *RestrictedNumber) GreaterThanPercent(perc int) bool {
	return rn.cur > rn.getPercent(perc)
}

func (rn *RestrictedNumber) GreaterThanEqualsPercent(perc int) bool {
	return rn.cur >= rn.getPercent(perc)
}
