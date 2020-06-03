package main

import (
	"fmt"
	"reflect"
)

//"context"
//"github.com/docker/docker/api/types"
//"github.com/docker/docker/api/types/container"
//"github.com/docker/docker/client"
//"github.com/docker/docker/pkg/stdcopy"
//"io"
//"os"
// "go-docker/docker"

type Test struct {
	Name string
}

func main() {
	// docker.Docker_api()
	var a Test
	a.Name = "test"
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))
}
