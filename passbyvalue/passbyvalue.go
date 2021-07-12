package main
import (
	"fmt"
)
type User struct {
	Name string
	Pet  []string
}
func (p2 *User) newPet() {
	fmt.Println(*p2, "underlying value of p2 before")
	p2.Pet = append(p2.Pet, "Lucy")
	fmt.Println(*p2, "underlying value of p2 after")
}
func main() {
	u := User{Name: "Anna", Pet: []string{"Bailey"}} // this time we'll generate a pointer!
	fmt.Println(u, "u before")
	p := &u // Let's make a pointer to u!
	m := u
	fmt.Printf("%T, %v, %s \n", p, p, p)
	fmt.Println(p)
	fmt.Println(m)
	p.newPet()
	fmt.Println(u, "u after")
	// Does Anna have a new pet now? Yes!
}