package main

import (
	"fmt"
	"bufio"         
	"os"
	"strings"  
	"unicode"
)

func palindrome(text string) bool {
	var clean [] rune
	
	for _,w := range(text){
		if unicode.IsLetter(w) || unicode.IsSpace(w){
			clean = append(clean,w)
		}
	}
	i,j := 0, len(clean)-1
	for i < j{
		if clean[i] != clean[j]{
			return false
		}
		i++
    	j--
	}
	return true
}

func wordCount(text string) (map[string]int,error){
	if strings.TrimSpace(text) == ""{
		return nil, fmt.Errorf("Input text is empty")
	}
	counts := make(map[string]int)
	var clean [] rune
	
	for _,w := range(text){
		if unicode.IsLetter(w) || unicode.IsSpace(w){
			clean = append(clean,w)
		}
	}
	words := strings.Fields(strings.ToLower(string(clean)))
	for _,word := range(words){
		counts[word]++
	}
	return counts,nil
}

func main(){
	fmt.Println("Welcome")
	fmt.Println("Please choose a service to continue:")
	fmt.Println("1 - Word Frequency Count" )
	fmt.Println("2 - Palindrome Check")
	fmt.Println("Enter your choice (1 or 2):")
	
	var choice string
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(line)

	
	switch choice {
	case "1":
		fmt.Println("Please enter your text for word count check: ")
		line,_ := reader.ReadString('\n')
		input := strings.TrimSpace(line)
		
		counts,err := wordCount(input)
		if err != nil{
			fmt.Println(err)
			return
		}
		for key, value := range counts {
    		fmt.Printf("Word: %-10s | Count: %d\n", key, value)
		}

	case "2":
		fmt.Println("Please enter the yout text for palindrome check: ")
		line,_ := reader.ReadString('\n')
		input := strings.TrimSpace(line)
		res := palindrome(input)

		if res == true{
			fmt.Println("Yes it is a palindrome")
		}else{
			fmt.Println("No it's not a palindrome")
		}
	default:
		fmt.Println("Invalid input. Choose 1 for Word Count or 2 for Palindrome.")
	}
}