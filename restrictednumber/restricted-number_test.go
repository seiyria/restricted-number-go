package restrictednumber

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRestrictedNumber(t *testing.T) {
	var rn *RestrictedNumber

	Convey("Instantiation works", t, func() {
		rn = New()
		So(rn.Max(), ShouldEqual, 0)
	})
}
