package main

import (
"fmt"
pb "./addressbook/generated_code"
proto "github.com/golang/protobuf/proto"
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



  // The whole purpose of using protocol buffers is to serialize your data so that it can be parsed elsewhere.
  book := pb.AddressBook{}
  // or do &pb.AddressBook{} and then use book
  book.People = []*pb.Person{}
  book.People = append(book.People, &p)
  fmt.Println(book)
  out, _ := proto.Marshal(&book)
  fmt.Println(out) // out is in byte format

  err := proto.Unmarshal(out,  &book)
    if err != nil {
    fmt.Println(err)
  }
  fmt.Println(&book) // converting byte to str


}





// go run main.go