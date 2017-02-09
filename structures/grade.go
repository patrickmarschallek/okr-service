package structures

import (
	"recipe/common"
	"reflect"
)

// Grade shows how good something was (0.3, 0.5, 0.7, 1.0)
type Grade float64

var grades = [5]Grade{
	0.0,
	0.3,
	0.5,
	0.7,
	1.0,
}

func (g Grade) Value() float64 {
	return reflect.ValueOf(g).Float()
}

func (g Grade) Up() {
	var pos = indexOf(g)
	if len(grades) < pos {
		g = grades[pos+1]
	}
}

func (g Grade) Down() {
	var pos = indexOf(g)
	if len(grades) != 0 {
		g = grades[pos-1]
	}
}

func indexOf(g Grade) int {
	var iSlice []interface{} = make([]interface{}, len(grades))
	for i, value := range grades {
		iSlice[i] = value
	}
	return common.IndexOf(iSlice, g)
}
