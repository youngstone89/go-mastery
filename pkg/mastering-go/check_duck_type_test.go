package masteringgo_test

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"testing"
	"time"
)

type Shape2D interface {
	Perimeter() float64
}
type circle struct {
	R float64
}

func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.R
}
func TestDuckType(t *testing.T) {
	a := circle{R: 1.5}
	fmt.Printf("R %.2f -> Perimeter %.3f \n", a.R, a.Perimeter())
	_, ok := interface{}(a).(Shape2D)
	if ok {
		fmt.Println("a is a Shape2D!")
	}
}

const min = 1
const max = 5

func rF64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

type Shape3D interface {
	Vol() float64
}
type Cube struct {
	x float64
}

func (c Cube) Vol() float64 {
	return c.x * c.x * c.x
}

type Cuboid struct {
	x float64
	y float64
	z float64
}

func (c Cuboid) Vol() float64 {
	return c.x * c.y * c.z
}

type Sphere struct {
	r float64
}

func (c Sphere) Vol() float64 {
	return 4 / 3 * math.Pi * c.r * c.r * c.r
}

type shapes []Shape3D

// Implementing sort.Interface
func (a shapes) Len() int {
	return len(a)
}
func (a shapes) Less(i, j int) bool {
	return a[i].Vol() < a[j].Vol()
}
func (a shapes) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func PrintShapes(a shapes) {
	for _, v := range a {
		switch v.(type) {
		case Cube:
			fmt.Printf("Cube: volume %.2f\n", v.Vol())
		case Cuboid:
			fmt.Printf("Cuboid: volume %.2f\n", v.Vol())
		case Sphere:
			fmt.Printf("Sphere: volume %.2f\n", v.Vol())
		default:
			fmt.Println("Unknown data type!")
		}
	}
	fmt.Println()
}
func TestSorting(t *testing.T) {
	data := shapes{}
	rand.Seed(time.Now().Unix())
	for i := 0; i < 3; i++ {
		cube := Cube{rF64(min, max)}
		cuboid := Cuboid{rF64(min, max), rF64(min, max), rF64(min, max)}
		sphere := Sphere{rF64(min, max)}
		data = append(data, cube)
		data = append(data, cuboid)
		data = append(data, sphere)
	}
	PrintShapes(data)
	// Sorting
	sort.Sort(shapes(data))
	PrintShapes(data)
	// Reverse sorting
	sort.Sort(sort.Reverse(shapes(data)))
	PrintShapes(data)
}
