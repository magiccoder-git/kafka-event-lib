package main

import (
	"fmt"
)



// The KEL (KafkaEventLib) is the main config for the Producer / Coustumer
type kel struct {

	//Channels :: a list with alle channels, who the service / endpoints are used.
	//The service can only use the channel that are  listen for the producer.
	Channels []string
	Adress string
	EName string
}
//INIT | Define the main config for the lib
func (k *kel) Init(address string, ename string) {

	if (address == "" || ename == "" ){
		panic("Error, missing init");
	}
	k.Adress = address;
	k.EName = ename;
	fmt.Println("[I] Channel permission for topics :: ")
	for i := 0; i < len(k.Channels); i++ {
		fmt.Println( k.Channels[i] + "  ")
	}
	k.Sub();
}