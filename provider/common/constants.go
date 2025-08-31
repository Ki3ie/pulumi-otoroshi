package common

var BoolTrue = GetBool(true)
var BoolFalse = GetBool(false)

func GetBool(b bool) *bool {
	return &b
}

var StringEmpty = GetString("")
var StringDefault = GetString("default")

func GetString(s string) *string {
	return &s
}
