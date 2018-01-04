package main 

import "fmt"

type contactInfo struct {
  email string
  twitter string
}

type person struct {
  firstName string
  lastName string
  homeTown string
}

type enhancedPerson struct {
  firstName string
  lastName string
  homeTown string
  contact contactInfo
}

func main() {
  jamal := person{"Jamal", "O'Garro", "NYC"}
  fmt.Printf("%+v", jamal)
  
  var friend person
  friend.firstName = "Cool"
  friend.lastName = "Person"
  friend.homeTown = "NYC"
  fmt.Printf("%+v", friend)
  
  mj := enhancedPerson{
    firstName: "Michael",
    lastName: "Jordan",
    homeTown: "NC",
    contact: contactInfo{
      email: "mike@bnike.com", 
      twitter: "jumpman23",
    },
  }
  fmt.Printf("%+v", mj)
}