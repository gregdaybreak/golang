package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

//Student struct with pointer to next Student
type Student struct {
	name string
	age  int
	ssn  string
	Next *Student
}

func main() {

	var students *Student //create a new pointer variable called students

	students = nil //assign the pointer to nil

	if len(os.Args) < 2 { //if length of arguments is less than two
		fmt.Println("There is no file")
		return
	}

	filename := os.Args[1] //filename is first argument

	data, _ := ioutil.ReadFile(filename) //read the file in

	for _, line := range strings.Split(string(data), "\n") { //range over the file contents and split it on newline

		if line == "" { //if at the end of the file skip the below stuff
			continue
		}

		tempData := strings.Split(string(line), " ") //split each item by whitespace and put into tempData

		newStudent := &Student{ //populate our struct with the data
			name: tempData[0],
			ssn:  tempData[2],
			Next: students,
		}
		newStudent.age, _ = strconv.Atoi(tempData[1]) //convert the age to an integer and put it into the new students age

		students = newStudent //our empty student now points to this new student created from the file
	}
	fmt.Println("Our list of students")
	for s := students; s != nil; s = s.Next { //s equals students.Next (the pointer) as long as s is not nil (at the end of the list)
		//print out the info and then go to the next student.
		fmt.Println(s.name, s.age, s.ssn, s.Next)
	}

}
