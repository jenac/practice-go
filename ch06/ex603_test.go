package ch06

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAdd_With_GoConvey(t *testing.T) {
	Convey("test add", t, func() {
		Convey("0+0", func() {
			So(Add(0, 0), ShouldEqual, 0)
		})

		Convey("1+3", func() {
			So(Add(1, 3), ShouldEqual, 4)
		})
	})

}
