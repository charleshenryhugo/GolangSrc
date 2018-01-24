package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Users struct {
	XMLName xml.Name `xml:"users"`
	Users   []User   `xml:"user"`
}

type User struct {
	XMLName xml.Name `xml:"user"`
	Type    string   `xml:"type,attr"`
	Name    string   `xml:"name"`
	Social  Social   `xml:"social"`
}

type Social struct {
	XMLName  xml.Name `xml:"social"`
	Facebook string   `xml:"facebook"`
	Twitter  string   `xml:"twitter"`
	Youtube  string   `xml:"youtube"`
}

func xmlTest() {
	//open our XMLFile
	xmlFile, err := os.Open("users.xml")
	//handle the error
	if err != nil {
		log.Panic(err)
	}

	//defer the closing of our XML file so that we can parser it later on
	defer xmlFile.Close()

	//read our xmlFile as a byte array
	byteValue, _ := ioutil.ReadAll(xmlFile)

	//initialize our Users array
	var users Users
	//unmarshall our byte array which contains our
	//XML file content into 'users' defined above
	xml.Unmarshal(byteValue, &users)

	//iterate through every user in users
	for _, u := range users.Users {
		fmt.Println(u.Type)
		fmt.Println(u.Name)
		fmt.Println(u.Social)
	}
}
