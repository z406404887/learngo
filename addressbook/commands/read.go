package commands

import (
	"bufio"
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/essem/learngo/addressbook/pb"
)

// Read a person information in address book
func Read(c pb.AddressBookServiceClient, reader *bufio.Reader) {
	fmt.Print("ID: ")
	idStr, _ := reader.ReadString('\n')

	id, err := strconv.ParseInt(strings.TrimSpace(idStr), 10, 64)
	if err != nil {
		fmt.Printf("Could not convert id: %v\n", err)
		return
	}

	r, err := c.Read(context.Background(), &pb.ReadRequest{Id: id})
	if err != nil {
		fmt.Printf("Could not read: %v\n", err)
		return
	}

	if r.Person == nil {
		fmt.Println("Not found")
		return
	}

	fmt.Printf("ID: %d\n", r.Person.Id)
	fmt.Printf("Name: %s\n", r.Person.Name)
	fmt.Printf("E-mail: %s\n", r.Person.Email)
}
