package main

import "fmt"

type pair struct {
	timestamp int
	value     string
}

type TimeMap struct {
	data map[string][]pair
}

func Constructor() TimeMap {
	return TimeMap{
		data: make(map[string][]pair),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.data[key] = append(this.data[key], pair{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	values, ok := this.data[key]
	if !ok {
		return ""
	}

	left, right := 0, len(values)-1
	res := ""

	for left < right {
		mid := (left + right) / 2

		if values[mid].timestamp <= timestamp {
			res = values[mid].value
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return res
}

func main() {

	timeMap := Constructor()

	timeMap.Set("love", "high", 10)
	timeMap.Set("love", "low", 20)

	fmt.Println(timeMap.Get("love", 5))  // ""
	fmt.Println(timeMap.Get("love", 10)) // "high"
	fmt.Println(timeMap.Get("love", 15)) // "high"
	fmt.Println(timeMap.Get("love", 20)) // "low"
	fmt.Println(timeMap.Get("love", 25)) // "low"
}
