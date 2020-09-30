package main

import (
	"fmt"
	"sort"
)

func findItinerary(tickets [][]string) []string {
	from := make(map[string][]string)
	for _, m := range tickets {
		from[m[0]] = append(from[m[0]], m[1])
	}
	for k, v := range from {
		sort.Strings(v)
		from[k] = v
		fmt.Println(k, v)
	}
	t := "JFK"
	var result []string
	result = append(result, t)
	for  i := 0; i < 10; i ++{
		value, ok := from[t]
		if !ok {
			return result
		}
		if len(value) > 0 {
			new := value[0]
			from[t] = value[1:]
			t = new
			result = append(result, t)
		}else {
			return result
		}
	}
	return result
}

func main() {
	//tickets := [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}
	tickets :=  [][]string{{"JFK","KUL"},{"JFK","NRT"},{"NRT","JFK"}}
	fmt.Println(tickets)
	fmt.Println(findItinerary(tickets))
}
