package main

func PointOne(n *int) {
	*n = *n + 1
}

func main() {
	m := 0
	PointOne(&m)
	println(m)
}
