package go_call_my_name

import (
	"strconv"
)

func CallMyNameAndAge(name string, age int) string {
	return "Hello " + name + "with " +  strconv.Itoa(age) + " years old, have a nice day"
}