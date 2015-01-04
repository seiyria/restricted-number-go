package restrictednumber

import (
	resn "github.com/seiyria/restricted-number-go/restrictednumber"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRestrictedNumber(t *testing.T) {
	var rn *resn.RestrictedNumber

	Convey("Instantiation works", t, func() {
		rn = resn.New()
		So(rn.Max(), ShouldEqual, 0)
	})
}
