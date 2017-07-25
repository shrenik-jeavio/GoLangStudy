package main
import "fmt"

func main(){
	ele := map[string]map[string]string{
		"H": map[string]string{
			"name": "shrenik",
			"state": "Gujarat",
		},
		"G": map[string]string{
			"name": "rajesh",
			"state": "rajasthan",
		},
	}

	if el, ok := ele["G"]; ok {
		fmt.Println(el["name"], el["state"])
	}
	for k, v := range ele {
		fmt.Println(k, v)
	}
}
