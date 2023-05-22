package main

import (
	"fmt"

	"github.com/asackan/study-golang/example3_dictionary/dict"
)

func main() {
	dictionary := dict.Dictionary{"first": "Apple"}

	// first 찾기 (존재함)
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("  definition is", definition)
	}

	// second 찾기 (없음)
	definition, err = dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("  definition is", definition)
	}

	// hello 추가
	word := "hello"
	definition = "Cloud"
	err = dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(" [NEW]  ", word, "  definition is", definition)
	}

	// hello 재추가 (안됨)
	err = dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word, "'s definition is ", dictionary[word])
	}

	// blabla 변경 (안됨)
	err = dictionary.Update("blabla", "Greeting")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word, "'s definition is ", dictionary[word])
	}

	// hello 업데이트
	err = dictionary.Update(word, "Greeting")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word, "'s definition is ", dictionary[word])
	}

	// first 삭제
	word = "first"
	err = dictionary.Delete(word)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Delete", word, "complete")
	}

	err = dictionary.Delete(word)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Delete", word, "complete")
	}

}
