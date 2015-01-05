package restrictednumber

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRestrictedNumber(t *testing.T) {
	var rn *RestrictedNumber

	Convey("Instantiation works and sets values correctly with default parameters", t, func() {
		rn = New()
		So(rn.Min(), ShouldEqual, 0)
		So(rn.Max(), ShouldEqual, 0)
		So(rn.Val(), ShouldEqual, 0)
	})

	Convey("Instantiation works and sets values correctly with specifier function", t, func() {
		rn = NewSet(0, 100, 50)
		So(rn.Min(), ShouldEqual, 0)
		So(rn.Max(), ShouldEqual, 100)
		So(rn.Val(), ShouldEqual, 50)
	})

	Convey("Clamping works as expected", t, func() {
		Convey("For specifier", func() {
			rn = NewSet(0, 100, 101)
			So(rn.Val(), ShouldEqual, 100)

			rn = NewSet(0, 100, -1)
			So(rn.Val(), ShouldEqual, 0)
		})

		Convey("For default", func() {
			rn = NewSet(0, 100, 0)
			rn.SetVal(101)
			So(rn.Val(), ShouldEqual, 100)
			rn.SetVal(-1)
			So(rn.Val(), ShouldEqual, 0)
		})
	})

	Convey("Automatically setting to min or max should work", t, func() {
		Convey("ToMin", func() {
			rn = NewSet(0, 100, 50)
			rn.ToMin()
			So(rn.Val(), ShouldEqual, 0)
		})

		Convey("ToMax", func() {
			rn = NewSet(0, 100, 50)
			rn.ToMax()
			So(rn.Val(), ShouldEqual, 100)
		})
	})

	Convey("Changing boundaries should work as expected", t, func() {
		Convey("Min should change", func() {
			rn = NewSet(0, 100, 0)
			rn.SetMin(50)
			So(rn.Min(), ShouldEqual, 50)
		})

		Convey("Max should change", func() {
			rn = NewSet(0, 100, 0)
			rn.SetMax(50)
			So(rn.Max(), ShouldEqual, 50)
		})

		Convey("Min should not go above max", func() {
			rn = NewSet(0, 100, 0)
			err := rn.SetMin(150)
			So(rn.Min(), ShouldEqual, 0)
			So(err, ShouldEqual, ErrMinGreaterThanMax)
		})

		Convey("Max should not go below min", func() {
			rn = NewSet(0, 100, 0)
			err := rn.SetMax(-50)
			So(rn.Max(), ShouldEqual, 100)
			So(err, ShouldEqual, ErrMaxSmallerThanMin)
		})

		Convey("Min should appropriately re-scale current", func() {
			rn = NewSet(0, 100, 0)
			rn.SetMin(50)
			So(rn.Min(), ShouldEqual, 50)
			So(rn.Val(), ShouldEqual, 50)
			So(rn.AtMin(), ShouldEqual, true)
		})

		Convey("Max should appropriately re-scale current", func() {
			rn = NewSet(0, 100, 100)
			rn.SetMax(50)
			So(rn.Max(), ShouldEqual, 50)
			So(rn.Val(), ShouldEqual, 50)
			So(rn.AtMax(), ShouldEqual, true)
		})
	})

	Convey("Add and subtract should work correctly", t, func() {
		rn = NewSet(0, 100, 50)
		rn.Add(20)
		So(rn.Val(), ShouldEqual, 70)

		rn.Sub(30)
		So(rn.Val(), ShouldEqual, 40)
	})

	Convey("Percent should be correctly calculated", t, func() {
		rn = NewSet(0, 100, 100)
		So(rn.AsPercent(), ShouldEqual, 100)
		So(rn.IsPercent(100), ShouldEqual, true)

		rn = NewSet(0, 200, 100)
		So(rn.AsPercent(), ShouldEqual, 50)
		So(rn.IsPercent(50), ShouldEqual, true)
	})

	Convey("Modifying by percent should work correctly", t, func() {
		rn = NewSet(0, 200, 0)
		rn.SetToPercent(50)
		So(rn.AsPercent(), ShouldEqual, 50)
		So(rn.Val(), ShouldEqual, 100)

		rn.AddByPercent(20)
		So(rn.AsPercent(), ShouldEqual, 70)
		So(rn.Val(), ShouldEqual, 140)

		rn.SubByPercent(40)
		So(rn.AsPercent(), ShouldEqual, 30)
		So(rn.Val(), ShouldEqual, 60)
	})

	Convey("Checking percentages should work correctly", t, func() {
		rn = NewSet(0, 100, 66)

		Convey("LessThanPercent", func() {
			So(rn.LessThanPercent(67), ShouldEqual, true)
			So(rn.LessThanPercent(66), ShouldEqual, false)
			So(rn.LessThanPercent(65), ShouldEqual, false)
		})

		Convey("LessThanEqualsPercent", func() {
			So(rn.LessThanEqualsPercent(67), ShouldEqual, true)
			So(rn.LessThanEqualsPercent(66), ShouldEqual, true)
			So(rn.LessThanEqualsPercent(65), ShouldEqual, false)
		})

		Convey("GreaterThanPercent", func() {
			So(rn.GreaterThanPercent(67), ShouldEqual, false)
			So(rn.GreaterThanPercent(66), ShouldEqual, false)
			So(rn.GreaterThanPercent(65), ShouldEqual, true)
		})

		Convey("GreaterThanEqualsPercent", func() {
			So(rn.GreaterThanEqualsPercent(67), ShouldEqual, false)
			So(rn.GreaterThanEqualsPercent(66), ShouldEqual, true)
			So(rn.GreaterThanEqualsPercent(65), ShouldEqual, true)
		})
	})
}
