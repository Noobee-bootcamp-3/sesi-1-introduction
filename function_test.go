package sesi1introdution

import (
	"fmt"
	"testing"
)

type M map[string]string

func combineString(name string, color string) (message string) {
	carString := fmt.Sprintf("Mobil %s berwarna %s", name, color)
	return carString
}

func Cetak(arg any) {
	fmt.Println(arg)
}

func TestFunction(t *testing.T) {
	var car = M{"name": "BMW", "color": "Black"}

	Cetak(combineString(car["name"], car["color"]))
}
