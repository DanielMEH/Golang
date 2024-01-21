package main

import "fmt"

func main() {

	// types for variables and constants in golang
	d := 1
	var e int = 2
	var (
		f = 3
	)
	const g = 4
	const (
		h = 5
	)

	fmt.Println(d, e, f, g, h)

	conversion()

	typesNumeric()

}

func conversion() {
	// fmt all valores
	// %v valor
	// %T tipo
	// %t boolean
	// %d int
	// %b binario
	// %c char
	// %x hexadecimal
	// %f float
	// %s string
	// %p ponteiro
	// %q aspas duplas
	// %#U unicode
	// %#x hexadecimal
	// %#o octal

}

func typesNumeric() {
	// int8, int16, int32, int64
	// uint8, uint16, uint32, uint64
	// byte = uint8
	// rune = int32
	// float32, float64
	// complex64, complex128
	IotalType()
}

func IotalType() {
	const (
		a = iota
		b
		c = iota
		d
	)

	fmt.Println(a, b, c, d)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)

	// bit shifting

	const (
		_  = iota
		kb = 1 << (iota * 10)
		mb = 2 << (iota * 10)
		gb = 3 << (iota * 10)
	)

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)

	x := 1
	for {
		x++
		if x > 100 {
			break
		}

		if x%2 != 0 {
			continue
		}

	}

	// switch
	switch {
	case false:
		fmt.Println("false")
	case true:
		fmt.Println("true")
		fallthrough
	default:
		fmt.Println("default")

	}

	sliceArray()

}

func sliceArray() {

	x := []int{1, 2, 3, 4, 5}

	y := []int{6, 7, 8, 9, 10}
	x = append(x, y...)

	x = append((x[:2]), x[4:]...)
	fmt.Println(x)
	fmt.Println(x[1:3])

	for i, v := range x {

		fmt.Println(i, " ", v)
	}

	// make
	z := make([]int, 5, 10)
	z[0] = 1
	z[3] = 4
	z[4] = 5
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	//maps

	mapData()

}

func mapData() {

	fmt.Println("maps")

	m := map[string]int{
		"james": 32,
		"jenny": 27,
	}
	delete(m, "james")

	for k, v := range m {
		fmt.Println(k, v)
	}

	structData()
}

func structData() {

	fmt.Println("structs")

	type person []struct {
		name string
		age  int
	}

	var p person

	for i := 0; i < 7; i++ {
		p = append(p, struct {
			name string
			age  int
		}{name: "james", age: 32})
	}
	fmt.Println(p)



 fmt.Println("POO")


type foo int



// func (r receiver) identifier(parameters) (return(s)) { ... }
fbo(1,2,3,4,5,6,7,8,9,10)






}
func fbo(x ...int){
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	fmt.Println("sum", sum)
}
