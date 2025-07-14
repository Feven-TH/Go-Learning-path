package main

import(
	"testing"
	"reflect"
)

func Test_palindrome(t *testing.T){
	tests := []struct {
		input string
		want  bool
	}{
		{"madam", true},
		{"hello", false},
		{"", true},
		{"1p1p1/",true},
	}
	for _,test := range(tests){
		got := palindrome(test.input)
		if got != test.want{
			t.Errorf("Expected for palindrome(%q): %v, but got: %v", test.input, test.want, got)
		}
	}
}

func Test_wordcount(t *testing.T){
	tests := []struct{
		input string
		want map[string]int
	}{
        {"Hello world hello", map[string]int{"hello": 2, "world": 1}},
		{"Go Go Go language is fun", map[string]int{"go": 3, "language": 1, "is": 1, "fun": 1}},
		{"Code Code code test code", map[string]int{"code": 4, "test": 1}},
	}
	for _,test := range(tests){
		got, _ := wordCount(test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Expected wordCount(%q): %v, but got: %v", test.input, test.want, got)
		}
	}
}


	