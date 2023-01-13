package urlstt

import (
	"fmt"
	"net/url"
	"reflect"
)

func GetURLParamsString(urlParamsStruct any) string {
	beforeEncode := url.Values{}

	t := reflect.TypeOf(urlParamsStruct)
	v := reflect.ValueOf(urlParamsStruct)

	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("url")
		if tag == "" {
			continue
		}

		beforeEncode.Add(tag, fmt.Sprintf("%v", v.Field(i)))
	}

	return beforeEncode.Encode()
}
