package blackboard

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBlackboardString(t *testing.T) {

	Convey("when given a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("when retrieving a non-existing string pointer with key `notexist`", func() {
			s := bb.StringP("notexist")
			Convey("bool pointer should be nil", func() {
				So(s, ShouldBeNil)
			})
		})
	})

	Convey("when given a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("when using SetString with key `exist` and value `test`", func() {
			bb.SetString("exist", "test")
			Convey("when retrieving an existing string pointer with key `exist`", func() {
				s := bb.StringP("exist")
				Convey("string pointer should not be nil", func() {
					So(s, ShouldNotBeNil)
				})
				Convey("string should equal `true`", func() {
					So(*s, ShouldEqual, "test")
				})
			})
		})
	})

	Convey("when given a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("when using SetStringP with key `exist` and value `test`", func() {
			v := "test"
			bb.SetStringP("exist", &v)
			Convey("when retrieving an existing string pointer with key `exist`", func() {
				s := bb.StringP("exist")
				Convey("string pointer should not be nil", func() {
					So(s, ShouldNotBeNil)
				})
				Convey("string should equal `test`", func() {
					So(*s, ShouldEqual, "test")
				})
			})
		})
	})

	Convey("when given a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("when using SetStringP with key `exist` and value `test`", func() {
			v := "test"
			bb.SetStringP("exist", &v)
			Convey("when retrieving existing string pointers with AllString", func() {
				ksl := bb.AllString()
				Convey("KeyString slice should not be nil", func() {
					So(ksl, ShouldNotBeNil)
				})
				Convey("KeyString slice should have length of 1", func() {
					So(len(ksl), ShouldEqual, 1)
				})
				Convey("KeyString slice should contain {`exist`, correct string pointer}", func() {
					So(ksl, ShouldContain, KS{"exist", &v})
				})
			})
		})
	})

	Convey("when given a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("when using SetValue with key `exist` and value `1`", func() {
			bb.SetValue("exist", 1)
			Convey("when retrieving value as string pointer with key `exist`", func() {
				s := bb.StringP("exist")
				Convey("string pointer should be nil", func() {
					So(s, ShouldBeNil)
				})
			})
		})
	})
}
