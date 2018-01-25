package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type students struct {
	Students []student `yaml:"students"`
}

type student struct {
	Name      string    `yaml:"name"`
	FullName  fullName  `yaml:"fullName"`
	Sex       string    `yaml:"sex"`
	SelfIntro selfIntro `yaml:"selfIntro"`
	ImageUrls []string  `yaml:"imageUrls"`
	Shemale   bool      `yaml:"shemale"`
}

type fullName struct {
	Firstname string `yaml:"firstName"`
	LastName  string `yaml:"lastName"`
}

type selfIntro struct {
	Long  string `yaml:"long"`
	Short string `yaml:"short"`
}

func ymlTest() {
	ymlFile, err := ioutil.ReadFile("users.yaml")
	if err != nil {
		log.Panic(err)
	}

	var stds students
	if err := yaml.Unmarshal(ymlFile, &stds); err != nil {
		log.Panic(err)
	}
	for _, std := range stds.Students {
		fmt.Println(std.Shemale)
	}

}
