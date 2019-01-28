package blackboard

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBlackboardInt(t *testing.T) {

	Convey("when given a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("and retrieving a non-existing `int` with key `notexist`", func() {
			i := bb.IntP("notexist")
			Convey("should be nil", func() {
				So(i, ShouldBeNil)
			})
		})
	})

	Convey("when given a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("and setting a value with key `exist` and value `test`", func() {
			bb.SetValue("exist", "test")
			Convey("retrieving the value as an `int`", func() {
				i := bb.IntP("exist")
				Convey("should be nil", func() {
					So(i, ShouldBeNil)
				})
			})
		})
	})

	Convey("when using the blackboard global methods", t, func() {
		Convey("and setting an `int` with key `exist1` and value `20`", func() {
			SetInt("exist1", 20)
			Convey("retrieving the value as an `int`", func() {
				i := IntP("exist1")
				Convey("should not be nil", func() {
					So(i, ShouldNotBeNil)
				})
				Convey("should equal `20`", func() {
					So(*i, ShouldEqual, 20)
				})
			})
		})
	})

	Convey("when using the blackboard global methods", t, func() {
		Convey("and setting an `int` with key `exist2` and value `30` using an `int` pointer", func() {
			v := 30
			SetIntP("exist2", &v)
			Convey("retrieving the value as an `int`", func() {
				i := IntP("exist2")
				Convey("should not be nil", func() {
					So(i, ShouldNotBeNil)
				})
				Convey("should equal `30`", func() {
					So(*i, ShouldEqual, 30)
				})
			})
		})
	})

	Convey("when using the blackboard global methods", t, func() {
		Convey("and setting an existing `int` with key `exist3` and value `40` using an `int` pointer", func() {
			v := 30
			SetIntP("exist3", &v)
			Convey("retrieving all `int` values", func() {
				kil := AllInt()
				Convey("the slice should not be nil", func() {
					So(kil, ShouldNotBeNil)
				})
				Convey("the slice should have length of 3", func() {
					So(len(kil), ShouldEqual, 3)
				})
				Convey("the slice should contain {`exist3`, correct pointer}", func() {
					So(kil, ShouldContain, KI{"exist3", &v})
				})
			})
		})
	})
}
