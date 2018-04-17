package main

import "fmt"

func main() {
	m := map[string]string{
		"a": "中国",
		"b": "美国",
		"c": "俄罗斯",
	}

	m2 := make(map[string]int)

	m3 := map[string]int{}

	fmt.Println(m, m2, m3)

	/*
	// 遍历map
	for k, v := range m {
		fmt.Println(k, v)
	}

	// update map
	var key = "a"
	if country, ok := m[key]; ok {
		fmt.Printf("m['%s'] = %s\n", key, country)
	} else {
		fmt.Printf("not found key '%s' of m\n", key)
	}

	*/
	a, ok := m["a"]
	fmt.Println(a, ok)

	d, ok := m["d"]
	fmt.Println(d, ok)
}
