package config

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestConfig(t *testing.T){
	data, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		t.Error("so such file or director")
	}
	fmt.Println(string(data))
	fileData := Conf.getConfig()
	fmt.Printf("%T \n%v\n" ,fileData,fileData)
}