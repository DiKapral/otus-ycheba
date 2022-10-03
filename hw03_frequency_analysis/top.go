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

	sortStruct := make([]keyVal, 0, len(stroca))

	for key, value := range stroca {
		sortStruct = append(sortStruct, keyVal{key, value})
	}

	sort.Slice(sortStruct, func(i, j int) bool {
		if sortStruct[i].Val == sortStruct[j].Val {
			return sortStruct[i].Key < sortStruct[j].Key
		}
		return sortStruct[i].Val > sortStruct[j].Val
	})

	if len(sortStruct) > 20 {
		sortStruct = sortStruct[0:20]
	} else {
		sortStruct = sortStruct[0:]
	}

	sortReza := make([]string, 0, len(sortStruct))

	for _, v := range sortStruct {
		sortReza = append(sortReza, v.Key)
	}

	return sortReza
}
