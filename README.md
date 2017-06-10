# go-lol
Go package client for League of Legends


go-lol is a Go client library for accessing the [GitHub API][].

**Documentation:** [![GoDoc](https://godoc.org/github.com/doguti/go-lol/github?status.svg)](https://godoc.org/github.com/doguti/go-lol/github)

**Build Status:** [![Build Status](https://travis-ci.org/doguti/go-lol.svg?branch=master)](https://travis-ci.org/doguti/go-lol)

**Test Coverage:** [![Test Coverage](https://coveralls.io/repos/github/doguti/go-lol/badge.svg?branch=master)](https://coveralls.io/github/doguti/go-lol?branch=master)

**Build History**

[![Build history for master branch](https://buildstats.info/travisci/chart/doguti/go-lol?branch=master&buildCount=50)](https://travis-ci.org/doguti/go-lol/branches)

go-lol requires Go version 1.7 or greater.


# MANUAL


import "github.com/doguti/go-lol"

```go
package main

import (
	"context"
	"github.com/doguti/go-lol"
	"fmt"
)



func main(){
	api_key := "<KEY API>"
	ctx := context.Background()
	c := lol.NewClient(nil, api_key)
	ch,res,_ := c.Champions.Get(ctx,"<USER>","<Method Name by default>")
	fmt.Printf("%s   %+v",ch.Name,res)

}
```