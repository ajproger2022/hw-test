package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	if text != "" {
		type item struct {
			word string
			cnt  int
		}

		var result []string
		text = strings.Replace(text, "\r", " ", -1)
		text = strings.Replace(text, "\n", " ", -1)
		listWords := strings.Split(text, " ")

		if len(listWords) != 0 {
			listWordsCount := make(map[string]int)
			for i := 0; i < len(listWords); i++ {
				listWords[i] = strings.TrimSpace(listWords[i])
				if listWords[i] != "" {
					listWordsCount[listWords[i]]++
				}
			}

			var topList []item
			for k, v := range listWordsCount {
				topList = append(topList, item{word: k, cnt: v})
			}

			sort.SliceStable(topList, func(i, j int) bool {
				return topList[i].word < topList[j].word
			})
			sort.SliceStable(topList, func(i, j int) bool {
				return topList[i].cnt > topList[j].cnt
			})
			//fmt.Println(topList)

			for k, v := range topList {
				if k == 10 {
					break
				}
				result = append(result, v.word)
			}

			return result
		}
	}

	return nil
}
