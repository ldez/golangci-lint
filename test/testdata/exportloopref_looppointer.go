//args: -Eexportloopref
//config_path: testdata/configs/looppointer.yml
package testdata

import "fmt"

func looppointerFunction() {
	var array [4]*int
	var slice []*int
	var ref *int
	var str struct{ x *int }

	fmt.Println("loop expecting 10, 11, 12, 13")
	for i, p := range []int{10, 11, 12, 13} {
		looppointerPrintp(&p)     // ERROR "taking a pointer for the loop variable p"
		slice = append(slice, &p) // ERROR "taking a pointer for the loop variable p"
		array[i] = &p             // ERROR "taking a pointer for the loop variable p"
		if i%2 == 0 {
			ref = &p   // ERROR "taking a pointer for the loop variable p"
			str.x = &p // ERROR "taking a pointer for the loop variable p"
		}
		var vStr struct{ x *int }
		var vArray [4]*int
		var v *int
		if i%2 == 0 {
			v = &p         // ERROR "taking a pointer for the loop variable p"
			vArray[1] = &p // ERROR "taking a pointer for the loop variable p"
			vStr.x = &p    // ERROR "taking a pointer for the loop variable p"
		}
		_ = v
	}

	fmt.Println(`slice expecting "10, 11, 12, 13" but "13, 13, 13, 13"`)
	for _, p := range slice {
		looppointerPrintp(p)
	}
	fmt.Println(`array expecting "10, 11, 12, 13" but "13, 13, 13, 13"`)
	for _, p := range array {
		looppointerPrintp(p)
	}
	fmt.Println(`captured value expecting "12" but "13"`)
	looppointerPrintp(ref)
}

func looppointerPrintp(p *int) {
	fmt.Println(*p)
}
