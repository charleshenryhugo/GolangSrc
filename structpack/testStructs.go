package structpack

import (
	"fmt"
	"reflect"
	"unsafe"
)

type intAlias int
type intStruct struct {
	int
}

//You can not define a function for int
//but you can define a function for intAlias, which is a alias of int
//in the same package
func (pi *intAlias) add(para int) int {
	return int(*pi) + para
}

//You can not define a function for int
//but you can define a function for intStruct with int as a member
//in the same package
func (pi *intStruct) add(para int) int {
	return pi.int + para //or *pi.int + para
}

func IntInits() {
	iA := intAlias(5)
	piA := new(intAlias)
	*piA = 5
	iS := intStruct{5}
	piS := new(intStruct)
	(*piS).int = 5 //or piS.int = 5

	fmt.Println(iA.add(2))
	fmt.Println(iS.add(2))

	fmt.Println((*piA).add(2)) //or piA.add(2)
	fmt.Println(piS.add(2))
}

type dad struct {
	height int
	weight int
}

type mom struct {
	height int
	weight int
}

type son struct {
	dad
	mom
	height int
	weight int
}

func SonInits() {
	bob := son{dad{1, 2}, mom{7, 8}, 9, 10}
	fmt.Println(bob.height)
	fmt.Println(bob.dad.height)
	fmt.Println(bob.mom.height)
}

type innerS struct {
	int1 int
	int2 int
}
type outerS struct {
	i      int                                   //8
	c      float32                               //4
	int    "anonymous field int"                 //8
	innerS "anonymous field(Just like heritage)" //16
}

func OuterSInits() {
	outer1 := new(outerS)
	outer1.i = 5
	outer1.c = 5.0
	outer1.int = 50
	outer1.int1 = 1
	outer1.int2 = 2
	outer2 := outerS{5, 5.0, 50, innerS{1, 2}}
	fmt.Println("outer2.i:", &outer2.i, " ", unsafe.Sizeof(outer2.i))
	fmt.Println("outer2.c:", &outer2.c, " ", unsafe.Sizeof(outer2.c))
	fmt.Println("outer2.int:", &outer2.int, " ", unsafe.Sizeof(outer2.int))
	fmt.Println("outer2.int1:", &outer2.int1, " ", unsafe.Sizeof(outer2.int1))
	fmt.Println("outer2.int2:", &outer2.int2, " ", unsafe.Sizeof(outer2.int2))

	fmt.Printf("outer2.innerS: %p\n", &outer2.innerS)
	fmt.Println("outer2:", unsafe.Sizeof(outer2))
}

type interval struct {
	start int "time when an interval starts"
	end   int "time when an interval ends"
}

func IntervalInits() {
	Intv1 := new(interval)      //*interval
	Intv2 := interval{0, 5}     //interval
	Intv3 := &interval{0, 5}    //*interval
	Intv4 := interval{start: 0} //interval

	Intv5 := NewInterval(0, 5)
	Intv1.ShowIntv()
	Intv2.ShowIntv()
	Intv3.ShowIntv()
	Intv4.ShowIntv()
	Intv5.ShowIntv()
	fmt.Println(unsafe.Sizeof(*Intv5))
}

func NewInterval(start, end int) *interval {
	if start > end || start < 0 {
		return nil
	}
	return &interval{start, end}
}

func (intv *interval) ShowIntv() {
	fmt.Println(intv.end - intv.start)
}

//RefTags prints all tags of a struct-interval
func RefTags(t *interval) {
	tType := reflect.TypeOf(*t)
	fmt.Println(tType)
	nField := tType.NumField()
	for ix := 0; ix < nField; ix++ {
		ixField := tType.Field(ix)
		fmt.Println(ixField.Tag)
	}
}
