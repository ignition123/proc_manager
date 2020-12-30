package lib

import(
	"log"
)

func Handlepanic() { 
 
    if err := recover(); err != nil {

       	log.Println(err)

   	}	

}

 