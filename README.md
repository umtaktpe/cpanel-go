# cpanel-go

# Installation
```
go get github.com/umtaktpe/cpanel-go
```

# Usage
```go
package main

import (
	"fmt"
	"log"

	"github.com/umtaktpe/cpanel-go"
	"github.com/umtaktpe/cpanel-go/client"
	"github.com/umtaktpe/cpanel-go/model"
)

func main() {
	config := cpanel.NewConfig("username", "password")
    client := client.NewClient(config)


	params := &model.LicenseInfoParams{
		Expired: "0",
		GroupId: "groupid",
		Output:  "json",
	}

	response, err := client.LicenseInfo(params)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(response)
}
```