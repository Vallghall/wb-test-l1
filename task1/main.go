package main

import "fmt"

// Human is a type representing a person. Base type
type Human struct {
	name string // name field represents person's name
	age  int    // age field represents person's age
}

// String method makes Human type to implement Stringer interface
func (h Human) String() string {
	return fmt.Sprintf("Name: %s\nAge: %d\n Type:%T\n", h.name, h.age, h)
}

func (h Human) Introduce() {
	fmt.Printf("I am %s, %d years old\n", h.name, h.age)
}

// Action type is built upon Human type. Daughter type
type Action struct {
	Human // struct embedding
}

// String method shadows Human.String method
func (a Action) String() string {
	return "Shadowed Human.String()"
}

// Introduce method shadows Human.Introduce method
func (a Action) Introduce() {
	fmt.Println("Killer Queen!")
}

func main() {
	man := Human{"Kira Yoshikage", 33} // initializing value of type Human
	fmt.Println(man)                   // prints return value of a String() method

	a := Action{man} // initializing value of type Action
	fmt.Println(a)   // prints return value of a String() method !of type Human! until shadowed

	man.Introduce() // Calls Human::Introduce()
	a.Introduce()   // calls Human::Introduce() until shadowed

	if a.Human.age == a.age {
		fmt.Println("It is possible to straightforwardly refer to Human type fields on Action value")
	}
}
