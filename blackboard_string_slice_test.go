package blackboard

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBlackboardStringSlice(t *testing.T) {

	Convey("when given a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("and retrieving a non-existing `string` slice with key `notexist`", func() {
			ss := bb.StringSliceP("notexist")
			Convey("should be nil", func() {
				So(ss, ShouldBeNil)
			})
		})
	})

	Convey("when given a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("and setting a value with key `exist` and value `1`", func() {
			bb.SetValue("exist", 1)
			Convey("retrieving the value as `string` slice", func() {
				ss := bb.StringSliceP("exist")
				Convey("should be nil", func() {
					So(ss, ShouldBeNil)
				})
			})
		})
	})

	Convey("when given a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("and setting a `string` slice with key `exist` and value `test1`, `test2`", func() {
			ssv := []string{"test1", "test2"}
			bb.SetStringSlice("exist", ssv)
			Convey("retrieving the existing string slice", func() {
				ss := bb.StringSliceP("exist")
				Convey("should not be nil", func() {
					So(ss, ShouldNotBeNil)
				})
				Convey("the string slice should equal `test1`, `test2`", func() {
					So((*ss)[0], ShouldEqual, ssv[0])
					So((*ss)[1], ShouldEqual, ssv[1])
				})
			})
		})
	})

	Convey("when given a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("and setting a `string` slice with key `exist` and value `test1`, `test2 using a `string` slice pointer", func() {
			ssv := []string{"test1", "test2"}
			bb.SetStringSliceP("exist", &ssv)
			Convey("retrieving an existing `string` slice with key `exist`", func() {
				ss := bb.StringSliceP("exist")
				Convey("should not be nil", func() {
					So(ss, ShouldNotBeNil)
				})
				Convey("the `string` slice should equal `test1`, `test2`", func() {
					So((*ss)[0], ShouldEqual, ssv[0])
					So((*ss)[1], ShouldEqual, ssv[1])
				})
			})
		})
	})

	Convey("when given a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("and setting a `string` slice with key `exist` and value `test1`, `test2`", func() {
			ssv := []string{"test1", "test2"}
			bb.SetStringSliceP("exist", &ssv)
			Convey("retrieving all `string` slice values", func() {
				kssl := bb.AllStringSlice()
				Convey("the slice should not be nil", func() {
					So(kssl, ShouldNotBeNil)
				})
				Convey("the slice should have length of 1", func() {
					So(len(kssl), ShouldEqual, 1)
				})
				Convey("the slice should contain exist{`test1`, `test2`}", func() {
					So(kssl, ShouldContain, KSS{"exist", &ssv})
				})
			})
		})
	})

	Convey("when using the blackboard global methods", t, func() {
		Convey("and setting a `string` slice with key `exist1` and value `test1`, `test2`", func() {
			ssv := []string{"test1", "test2"}
			SetStringSlice("exist1", ssv)
			Convey("retrieving the existing string slice", func() {
				ss := StringSliceP("exist1")
				Convey("should not be nil", func() {
					So(ss, ShouldNotBeNil)
				})
				Convey("the string slice should equal `test1`, `test2`", func() {
					So((*ss)[0], ShouldEqual, ssv[0])
					So((*ss)[1], ShouldEqual, ssv[1])
				})
			})
		})
	})

	Convey("when using the blackboard global methods", t, func() {
		Convey("and setting a `string` slice with key `exist2` and value `test2`, `test3 using a `string` slice pointer", func() {
			ssv := []string{"test3", "test4"}
			SetStringSliceP("exist2", &ssv)
			Convey("retrieving an existing `string` slice with key `exist2`", func() {
				ss := StringSliceP("exist2")
				Convey("should not be nil", func() {
					So(ss, ShouldNotBeNil)
				})
				Convey("the `string` slice should equal `test2`, `test3`", func() {
					So((*ss)[0], ShouldEqual, ssv[0])
					So((*ss)[1], ShouldEqual, ssv[1])
				})
			})
		})
	})

	Convey("when using the blackboard global methods", t, func() {
		Convey("and setting a `string` slice with key `exist3` and value `test5`, `test6`", func() {
			ssv1 := []string{"test1", "test2"}
			ssv2 := []string{"test3", "test4"}
			ssv3 := []string{"test5", "test6"}
			SetStringSliceP("exist3", &ssv3)
			Convey("retrieving all `string` slice values", func() {
				kssl := AllStringSlice()
				Convey("the slice should not be nil", func() {
					So(kssl, ShouldNotBeNil)
				})
				Convey("the slice should have length of 3", func() {
					So(len(kssl), ShouldEqual, 3)
				})
				Convey("the slice should contain exist1{`test1`, `test2`}", func() {
					So(kssl, ShouldContain, KSS{"exist1", &ssv1})
				})

				Convey("the slice should contain exist2{`test3`, `test4`}", func() {
					So(kssl, ShouldContain, KSS{"exist2", &ssv2})
				})

				Convey("the slice should contain exist3{`test5`, `test6`}", func() {
					So(kssl, ShouldContain, KSS{"exist3", &ssv3})
				})
			})
		})
	})
}
