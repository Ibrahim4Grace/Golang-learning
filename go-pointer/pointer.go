package main 
import "fmt"


//we can use pointer when dealing with higher numbers
func main(){
 age := 32
  
 //create pointer
 var agePointer = &age

//dereference pointer
 fmt.Println("Age:", *agePointer)

 adultyears := getAdultYears(agePointer)
 fmt.Println("adult year:", adultyears)
}

func getAdultYears(age *int) int {

    return *age -18
}