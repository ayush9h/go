package main

import "fmt"

func main() {

	m := make(map[string]string)

	m["name"] = "golang"
	m["area"] = "backend"
	fmt.Println(m["name"], m["area"])

	// if does not exist, we get an empty value
	fmt.Println(m["fpe"])
	delete(m, "area")

	v, ok := m["name"]
	fmt.Println(v)
	fmt.Println(ok)
}
