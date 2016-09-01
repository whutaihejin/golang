package main

import "fmt"

type ResultPOI struct {
	Uid    string
	Srctag string
}

func (r ResultPOI) String() string {
	return fmt.Sprintf("%s-%s", r.Srctag, r.Uid)
}

// clean or filter the sensitive record for the business
func resultClean(resultArray []ResultPOI, blackDict map[string]bool) []ResultPOI {
	// we will deal with the result array in place,
	// or shrinke the array in other words
	k := 0
	for _, record := range resultArray {
		key := fmt.Sprintf("%s-%s", record.Srctag, record.Uid)
		if blackDict[key] {
			continue
		}
		resultArray[k] = record
		k++
	}
	return resultArray[:k]
}

func main() {
	dict := make(map[string]bool)
	dict["penguin-123"] = true
	dict["penguin-1234"] = true
	dict["penguin-1236"] = true

	array := []ResultPOI{
		{Uid: "22", Srctag: "penguin"},
		{Uid: "1236", Srctag: "penguin"},
		{Uid: "33", Srctag: "penguin"},
	}
	for _, value := range array {
		fmt.Printf("%s\n", value)
	}
	filteredArray := resultClean(array, nil)
	fmt.Printf("len=%d\n", len(filteredArray))
	for _, value := range filteredArray {
		fmt.Printf("%s\n", value)
	}

	dealed := resultClean(nil, dict)
	fmt.Println(len(dealed))
}
