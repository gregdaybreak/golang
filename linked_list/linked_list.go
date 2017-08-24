package main


import "fmt"

type Student struct {  //we create a struct of Student
  age int
  weight int
  name string
  Next *Student   //this is the pointer to the next student if applicable
}

type Teacher struct { //we create a struct of Teacher
  age int
  weight int
  name string
  Next *Student  //this is the pointer to the student down the line
}

func main() {
  kyle := Student{7, 70, "Kyle", nil}  //kyle is a new Student but he is in the
                                       // back of line so he points to nil

  john := Student{10, 100, "John", &kyle}  //john is the next student and is in
                                           //front of kyle so he points to the
                                           //address where kyle is

  debra := Teacher{24, 185, "Debra", &john} //debra is the front of the line
                                            //she points to the address where
                                            //john is

  fmt.Println("The teacher is", debra.name)
  fmt.Println("The student directly behind Debra is", debra.Next.name) // the Teacher struct has a pointer to a student
                                                                       // called Next. by calling debra.Next.name we are going into
                                                                      // the struct for john and getting his name. (John)

  fmt.Println("The student directly behind John is", debra.Next.Next.name) //we can also get to the student behind john (Kyle) through debra.

  fmt.Println("The student directly behind John is", john.Next.name) //or get to Kyle through john

  fmt.Println("The student directly behind Kyle is", kyle.Next) //Kyle is the last kid so his Next is nil (no one is behind him)
  
}
