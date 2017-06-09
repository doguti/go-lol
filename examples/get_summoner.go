package main

import (
	"context"
	"github.com/doguti/go-lol/lol"
	"fmt"
	"os"
)


func main(){
	api_key := os.Getenv("LOL_KEY")

	c := lol.NewClient(nil, api_key)

	ch,_, err := c.Summoners.Get(context.Background(),"WolfKon")

	if err != nil{
		fmt.Printf("%+v", err)
	}else {
		fmt.Printf("Name of user %s", *ch.Name)
	}

}