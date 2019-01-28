package blackboard

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBlackboardString(t *testing.T) {

	Convey("when given a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("and retrieving a non-existing `string` with key `notexist`", func() {
			s := bb.StringP("notexist")
			Convey("should be nil", func() {
				So(s, ShouldBeNil)
			})
		})
	})

	Convey("when given a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("and setting a value with key `exist` and value `1`", func() {
			bb.SetValue("exist", 1)
			Convey("retrieving the value as a `string`", func() {
				s := bb.StringP("exist")
				Convey("should be nil", func() {
					So(s, ShouldBeNil)
				})
			})
		})
	})

	Convey("when using the blackboard global methods", t, func() {
		Convey("and setting a `string` with key `exist1` and value `test1`", func() {
			SetString("exist1", "test1")
			Convey("retrieving the value as a `string`", func() {
				s := StringP("exist1")
				Convey("should not be nil", func() {
					So(s, ShouldNotBeNil)
				})
				Convey("should equal `test1`", func() {
					So(*s, ShouldEqual, "test1")
				})
			})
		})
	})

	Convey("when using the blackboard global methods", t, func() {
		Convey("and setting a `string` with key `exist2` and value `test2` using a `string` pointer", func() {
			v := "test2"
			SetStringP("exist2", &v)
			Convey("retrieving the value as a `string`", func() {
				s := StringP("exist2")
				Convey("should not be nil", func() {
					So(s, ShouldNotBeNil)
				})
				Convey("should equal `test2`", func() {
					So(*s, ShouldEqual, "test2")
				})
			})
		})
	})

	Convey("when using the blackboard global methods", t, func() {
		Convey("and setting an existing `string` with key `exist3` and value `test3` using a `string` pointer", func() {
			v := "test3"
			SetStringP("exist3", &v)
			Convey("retrieving all `string` values", func() {
				ksl := AllString()
				Convey("the slice should not be nil", func() {
					So(ksl, ShouldNotBeNil)
				})
				Convey("the slice should have length of 3", func() {
					So(len(ksl), ShouldEqual, 3)
				})
				Convey("the slice should contain {`exist3`, correct pointer}", func() {
					So(ksl, ShouldContain, KS{"exist3", &v})
				})
			})
		})
	})
}
