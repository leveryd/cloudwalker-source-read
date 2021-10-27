package main

func main(){
	x := make([]int, 1)
	println(cap(x))
	x = append(x, 1)
	println(cap(x))
	x = append(x, 2)
	println(cap(x))
	x = append(x, 3)
	println(cap(x))

	println(x)
}