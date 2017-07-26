package main

import (
	"bufio"
	"strings"
	"log"
	"github.com/gogo/protobuf/proto"
	"os"
	"io"
	"io/ioutil"
	"fmt"
)

func promptForAddress(r io.Reader) (*Person, error) {
	// A protocol buffer can be created like any struct.
	p := Person{}

	rd := bufio.NewReader(r)
	p.Id = new(int64)

	fmt.Print("Enter person ID number: ")
	// An int32 field in the .proto file is represented as an int32 field
	// in the generated Go struct.

		if _, err := fmt.Fscanf(rd, "%d\n", p.Id); err != nil {
			fmt.Printf("Error: %+v",err )
		return &p, err
	}


	fmt.Print("Enter name: ")
	name, err := rd.ReadString('\n')
	if err != nil {
		return &p, err
	}

	// A string field in the .proto file results in a string field in Go.
	// We trim the whitespace because rd.ReadString includes the trailing
	// newline character in its output.
	p.Name = &name //strings.TrimSpace()

	fmt.Print("Enter email address (blank for none): ")
	email, err := rd.ReadString('\n')
	if err != nil {
		return &p, err
	}
	p.Email = &email //strings.TrimSpace(email)

	for {
		fmt.Print("Enter a phone number (or leave blank to finish): ")
		phone, err := rd.ReadString('\n')
		if err != nil {
			return &p, err
		}
		phone = strings.TrimSpace(phone)
		if phone == "" {
			break
		}

		// The PhoneNumber message type is nested within the Person
		// message in the .proto file.  This results in a Go struct
		// named using the name of the parent prefixed to the name of
		// the nested message.  Just as with pb.Person, it can be
		// created like any other struct.
		pn := &Person_PhoneNumber{
			Number: &phone,
		}

		fmt.Print("Is this a mobile, home, or work phone? ")
		ptype, err := rd.ReadString('\n')
		if err != nil {
			return &p, err
		}
		ptype = strings.TrimSpace(ptype)

		// A proto enum results in a Go constant for each enum value.
		/*switch ptype {
			case "mobile":
				pn.Type = Person_MOBILE
			case "home":
				pn.Type = Person_HOME
			case "work":
				pn.Type = Person_WORK
			default:
				fmt.Printf("Unknown phone type %q.  Using default.\n", ptype)
		}*/

		// A repeated proto field maps to a slice field in Go.  We can
		// append to it like any other slice.
		p.Phones = append(p.Phones, pn)
	}

	return &p, nil
}

func main(){
	/*
	person := &Person{
		Name: proto.String("Shrenik"),
		Id:  proto.Int32(1),
		Email: proto.String("shrenik@jeavio.com"),
	}

	data, err := proto.Marshal(person)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newPerson := &Person{
		Name: proto.String("Rajesh"),
		Id:  proto.Int32(2),
		Email: proto.String("rajesh@jeavio.com"),
	}
	err = proto.Unmarshal(data, newPerson)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	// Now test and newTest contain the same data.
	if person.GetName() != newPerson.GetName() {
		log.Fatalf("data mismatch %q != %q", person.GetName(), newPerson.GetName())
	}

	log.Printf("Unmarshalled to: %+v", newPerson)
	*/

	if len(os.Args) != 2 {
		log.Fatalf("Usage:  %s ADDRESS_BOOK_FILE\n", os.Args[0])
	}
	fname := os.Args[1]

	// Read the existing address book.
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s: File not found.  Creating new file.\n", fname)
		} else {
			log.Fatalln("Error reading file:", err)
		}
	}

	book := &AddressBook{}

	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	// Add an address.
	addr, err := promptForAddress(os.Stdin)
	if err != nil {
		log.Fatalln("Error with address:", err)
	}
	book.People = append(book.People, addr)

	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}
