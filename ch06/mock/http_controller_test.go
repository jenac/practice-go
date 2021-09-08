package mock

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	. "bou.ke/monkey"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetPageResponse(t *testing.T) {
	PatchInstanceMethod(reflect.TypeOf(http.DefaultClient), "Do", func( c *http.Client,  r *http.Request) (*http.Response, error) {
		return nil, fmt.Errorf("%s", "http fail")
	})
	tests := [2]struct {
		name  string
		url   string
		want1 int
		want2 []byte
		want3 error
	}{
		{
			name:  "not ok",
			url:   "http://www.baidu.com",
			want1: 400,
			want2: []byte{},
			want3: fmt.Errorf("%s", "http fail"),
		},
		{
			name:  "not ok",
			url:   "http://www.hao123.com",
			want1: 400,
			want2: []byte{},
			want3: fmt.Errorf("%s", "http fail"),
		},
	}
	Convey(tests[0].name, t, func() {
		code, _, err := GetPageResponse(tests[0].url)
		fmt.Println(code, err, )
		So(code, ShouldEqual, tests[0].want1)
		//So(result, ShouldEqual, tests[0].want2)
		So(err, ShouldNotBeNil)
	})
	Convey(tests[1].name, t, func() {
		code, _, err := GetPageResponse(tests[1].url)
		So(code, ShouldEqual, tests[1].want1)
		//So(result, ShouldEqual, tests[1].want2)
		So(err, ShouldNotBeNil)
	//
	})
}