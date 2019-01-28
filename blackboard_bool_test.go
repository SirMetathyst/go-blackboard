package blackboard

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBlackboardBool(t *testing.T) {

	Convey("when given a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("and retrieving a non-existing `bool` with key `notexist`", func() {
			b := bb.BoolP("notexist")
			Convey("should be nil", func() {
				So(b, ShouldBeNil)
			})
		})
	})

	Convey("when given a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("and setting a value with key `exist` and value `10`", func() {
			bb.SetValue("exist", 10)
			Convey("retrieving the value as a `bool`", func() {
				i := bb.BoolP("exist")
				Convey("should be nil", func() {
					So(i, ShouldBeNil)
				})
			})
		})
	})

	Convey("when using the blackboard global methods", t, func() {
		Convey("and setting a `bool` with key `exist1` and value `true`", func() {
			SetBool("exist1", true)
			Convey("retrieving the value as a `bool`", func() {
				b := BoolP("exist1")
				Convey("should not be nil", func() {
					So(b, ShouldNotBeNil)
				})
				Convey("should equal `true`", func() {
					So(*b, ShouldBeTrue)
				})
			})
		})
	})

	Convey("when using the blackboard global methods", t, func() {
		Convey("and setting a `bool` with key `exist2` and value `true` using a `bool` pointer", func() {
			v := true
			SetBoolP("exist2", &v)
			Convey("retrieving the value as a `bool`", func() {
				b := BoolP("exist2")
				Convey("should not be nil", func() {
					So(b, ShouldNotBeNil)
				})
				Convey("should equal `true`", func() {
					So(*b, ShouldBeTrue)
				})
			})
		})
	})

	Convey("when using the blackboard global methods", t, func() {
		Convey("and setting an existing `bool` with key `exist3` and value `true` using a `bool` pointer", func() {
			v := true
			SetBoolP("exist3", &v)
			Convey("retrieving all `bool` values", func() {
				kbl := AllBool()
				Convey("the slice should not be nil", func() {
					So(kbl, ShouldNotBeNil)
				})
				Convey("the slice should have length of 3", func() {
					So(len(kbl), ShouldEqual, 3)
				})
				Convey("the slice should contain {`exist3`, correct pointer}", func() {
					So(kbl, ShouldContain, KB{"exist3", &v})
				})
			})
		})
	})
}
