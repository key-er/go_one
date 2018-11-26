package main

import (
"fmt"
pb "./addressbook/generated_code"
)

func main() {
  p := pb.Person{
    Id:    1234,
    Name:  "John Doe",
    Email: "jdoe@example.com",
    Phones: []*pb.Person_PhoneNumber{
      {Number: "555-4321", Type: pb.Person_HOME},
    },
  }

  fmt.Println(p)
}