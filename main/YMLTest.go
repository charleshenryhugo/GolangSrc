package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
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

func viperTest() {
	//name of config file(without extension)
	viper.SetConfigName("users")
	//paths to look for the config file in
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath(".")

	//tell the viper instance to watchConfig
	viper.WatchConfig()
	//provide a function for viper to run each time a config change occurs
	viper.OnConfigChange(
		func(e fsnotify.Event) {
			log.Println("config file changed:", e.Name)
		})

	//find and read the config file
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
	//read successfully
	viper.Set("students.student1.name", "Booooob Peeeeeeter")
	viper.WriteConfig()
	// for {
	// 	var stds students
	// 	viper.Unmarshal(&stds)
	// 	//images := viper.Sub("students").Sub("student2").GetStringSlice("imageUrls")
	// 	typeCheck(stds)
	// 	time.Sleep(time.Second)
	// }

	// [	map[selfIntro:map[long:I am Bob! short:Hi!]
	// 	 imageUrls:[01.jpg 02.jpg 03.jpg]
	// 	 shemale:false
	// 	 name:Bob Peter
	// 	 fullName:map[lastName:Peter firstName:Bob]
	// 	 sex:male birthday:1990-12-12
	// 	 ]
	// 	map[sex:female
	// 		birthday:1994-12-12
	// 		selfIntro:map[long:Iam Alice! short:Hi!]
	// 		imageUrls:[001.jpg 002.jpg]
	// 		shemale:true
	// 		name:Alice Lily
	// 		fullName:map[firstName:Alice lastName:Lily]
	// 	]
	// ]

}
