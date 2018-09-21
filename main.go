package main

import (
	"fmt"
	"github.com/nicolasvasquez/brandSelector/internal"
)

func main() {
	r := internal.Request{Brand: "RESERVOIR", Header: internal.Header{AccessToken: "qqvhB0nKX83z_l7eQCHm9w", Client: "futo4fniAExNw_Bdmxgz2w"}}
	if a, err := r.GetBrand(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(a)
	}
}
