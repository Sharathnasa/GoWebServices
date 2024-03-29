package loops

func main() {

	slice := []int{1, 2, 3}
	for i, v := range slice {
		println(i, v)
	}

	//looping over maps
	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for k, v := range wellKnownPorts {
		println(k, v)
	}

}
