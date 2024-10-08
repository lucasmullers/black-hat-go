package main

// Go contains standard packages for data encoding. The most common packages you`re likely to use are JSON and XML
import (
	"encoding/json"
	"fmt"
)

type Foo struct {
	Bar string
	Baz string
}

func main() {
	f := Foo{"Joe Junior", "Hello Shabado"}
	b, _ := json.Marshal(f)
	fmt.Println(string(b))
	json.Unmarshal(b, &f)
}
