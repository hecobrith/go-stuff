/* This blueprint defines an expected set of actions that 
any concrete implementation must fulfill in order to be 
considered a type of interface. 
To define an interface, you define a set of methods; 
any data that contains those methods with the correct 
signatures fulfills the contract and is considered 
a type of that interface 
*/

package main

import (
	"log"
)

type Person struct {
	Name string
	Age int
}

type Dog struct { }

func (d *Dog) SayHello() {
	log.Println("Wuuuuf")
}

func (p *Person) SayHello(){
	log.Println("Hello", p.Name)
}

type Friend interface {
	SayHello()
}

func Greet (f Friend) {
	f.SayHello()
}

func main() {
	guy := new(Person)
	guy.Name = "Héctor"
	Greet(guy)
	dog := new(Dog)
	Greet(dog)
}

func foo(i interface{}){
	switch v := i.(type) {
		case int:
			log.Println("a numnber", v)
		case string:
			log.Println("its a string", v)
		default:
			log.Println("some typee", v)
		}
}