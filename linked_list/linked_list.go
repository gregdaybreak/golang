package main

import "fmt"

//Student is a structure with a pointer and a name
type Student struct {
	next *Student
	name string
}

//List represents our linked list
type List struct {
	head *Student
	tail *Student
}

//Insert something into the list
func (L *List) Insert(name string) {
	list := &Student{ //create a new variable list and put a new Student in it
		next: L.head, //assign the next student to the head of the list
		name: name,
	}

	L.head = list //put this new student as the head of the list

	l := L.head         //assign the head of the list to variable l
	for l.next != nil { //while the next is not nil (the end of the list)
		l = l.next //assign the next student to l
	}
	L.tail = l //assign the tail of the list
}

//Show the list of Students
func (L *List) Show() {
	list := L.head    //assign the head of the list
	for list != nil { //while we are not at the end of list
		fmt.Printf("%+v ->", list.name) //print out the name
		list = list.next                //go to the next student
	}
	fmt.Println() //newline
}

//Reverse the list
func (L *List) Reverse() {
	curr := L.head    //assign the head to a var curr
	var prev *Student //assign a Student to a var called prev
	L.tail = L.head   //the tail now equals the head

	for curr != nil { //while the list is not empty
		next := curr.next //next equals the current's next Student
		curr.next = prev  //now the current next is the previous
		prev = curr       //previous is now current
		curr = next       //current is now the next
	}
	L.head = prev //the head is now prev

}

func main() {
	l := List{}
	l.Insert("Jack")
	l.Insert("Sam")
	l.Insert("Fred")
	l.Insert("Greg")
	l.Insert("Frank")
	l.Insert("Jim")

	fmt.Printf("head: %v\n", l.head.name)
	fmt.Printf("tail: %v\n", l.tail.name)
	l.Show()
	fmt.Println("Reversing the List!")
	l.Reverse()
	l.Show()

	fmt.Printf("head: %v\n", l.head.name)
	fmt.Printf("tail: %v\n", l.tail.name)
}
