//args: -Eexportloopref
package testdata

import "fmt"

func exportlooprefFunction() {
	var array [4]*int
	var slice []*int
	var ref *int
	var str struct{ x *int }

	fmt.Println("loop expecting 10, 11, 12, 13")
	for i, p := range []int{10, 11, 12, 13} {
		exportlooprefPrintp(&p)
		slice = append(slice, &p) // ERROR "exporting a pointer for the loop variable p"
		array[i] = &p             // ERROR "exporting a pointer for the loop variable p"
		if i%2 == 0 {
			ref = &p   // ERROR "exporting a pointer for the loop variable p"
			str.x = &p // ERROR "exporting a pointer for the loop variable p"
		}
		var vStr struct{ x *int }
		var vArray [4]*int
		var v *int
		if i%2 == 0 {
			v = &p
			vArray[1] = &p
			vStr.x = &p
		}
		_ = v
	}

	fmt.Println(`slice expecting "10, 11, 12, 13" but "13, 13, 13, 13"`)
	for _, p := range slice {
		exportlooprefPrintp(p)
	}
	fmt.Println(`array expecting "10, 11, 12, 13" but "13, 13, 13, 13"`)
	for _, p := range array {
		exportlooprefPrintp(p)
	}
	fmt.Println(`captured value expecting "12" but "13"`)
	exportlooprefPrintp(ref)
}

func exportlooprefPrintp(p *int) {
	fmt.Println(*p)
}
