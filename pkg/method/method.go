package main

type Rect struct {
	width, height int
}


// Rect 의 area 메소드
func (r Rect) area() int {
	return r.width * r.height
}

func main() {

	rect := Rect{10,20}
	area := rect.area()
	print(area)
}
