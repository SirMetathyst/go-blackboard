package blackboard

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBlackboard(t *testing.T) {

	Convey("when givien a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("blackboard should not be nil", func() {
			So(bb, ShouldNotBeNil)
		})
	})

	Convey("when givien the blackboard singleton instance", t, func() {
		bb := Singleton()
		Convey("blackboard should not be nil", func() {
			So(bb, ShouldNotBeNil)
		})
	})

	Convey("when using the blackboard global methods", t, func() {
		Convey("and setting a value with key `test_key` and value `10`", func() {
			SetValue("test_key", 10)
			Convey("retrieving the value", func() {
				v, ok := Value("test_key")
				Convey("and `test_key` exists", func() {
					So(ok, ShouldBeTrue)
					Convey("the returned value is of type `int`", func() {
						i, ok := v.(int)
						So(ok, ShouldBeTrue)
						Convey("and the value is `10`", func() {
							So(i, ShouldEqual, 10)
						})
					})
				})
			})
		})

		Convey("and setting the existing value to `test_value`", func() {
			SetValue("test_key", "test_value")
			Convey("retrieving the value", func() {
				v, ok := Value("test_key")
				Convey("and `test_key` exists", func() {
					So(ok, ShouldBeTrue)
					Convey("the returned value is of type `string`", func() {
						s, ok := v.(string)
						So(ok, ShouldBeTrue)
						Convey("and the value is `test_value`", func() {
							So(s, ShouldEqual, "test_value")
						})
					})
				})
			})
		})
	})
}
