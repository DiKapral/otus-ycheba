package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

func Top10(str string) []string {
	re := regexp.MustCompile(`[[:space:]]`)
	str = re.ReplaceAllString(str, " ")
	masStr := strings.Split(str, " ")

	stroca := map[string]int{}
	for _, val := range masStr {
		if val != "" {
			stroca[val]++
		}
	}

	type keyVal struct {
		Key string
		Val int
	}

	var sortStruct []keyVal

	for key, value := range stroca {
		sortStruct = append(sortStruct, keyVal{key, value})
	}

	sort.Slice(sortStruct, func(i, j int) bool {
		return sortStruct[i].Val > sortStruct[j].Val
	})

	if len(sortStruct) > 10 {
		sortStruct = sortStruct[0:10]
	} else {
		sortStruct = sortStruct[0:]
	}

	var sortReza []keyVal
	var sortReza2 []keyVal
	var sortReza3 []string
	for len(sortStruct) != 0 {
		for _, keyval := range sortStruct {
			for i, keyval2 := range sortStruct {
				if keyval2.Val == keyval.Val {
					sortReza = append(sortReza, keyval2)
				} else {
					sort.Slice(sortReza, func(i, j int) bool {
						return sortReza[i].Key < sortReza[j].Key
					})

					sortReza2 = append(sortReza2, sortReza...)
					sortStruct = sortStruct[len(sortReza):]
					sortReza = []keyVal{}
					break
				}
				if i == len(sortStruct)-1 {
					sort.Slice(sortReza, func(i, j int) bool {
						return sortReza[i].Key < sortReza[j].Key
					})

					sortReza2 = append(sortReza2, sortReza...)
					sortStruct = sortStruct[len(sortReza):]
					sortReza = []keyVal{}
				}
			}
		}
	}

	for _, v := range sortReza2 {
		sortReza3 = append(sortReza3, v.Key)
	}

	return sortReza3
}
