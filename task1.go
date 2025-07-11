package main

import (
	"fmt"
	"errors"
	"bufio"         
	"os"
	"strings"        
)

func validate(grade int) error {
	if grade < 0 || grade > 100 {
		return errors.New("Enter a valid grade")
	}
	return nil
}

func average(grades map[string]int, n int) float64 {
	sum := 0
	for _, grade := range grades {
		sum += grade
	}
	return float64(sum) / float64(n)
}

func main() {
	reader := bufio.NewReader(os.Stdin) 
	var name string
	var n int
	var grades = make(map[string]int)
	var subject string
	var grade int

	fmt.Println("What is your name?")
	fmt.Scan(&name)
	
	fmt.Println("How many subjects do you have?")
	fmt.Scan(&n)
	reader.ReadString('\n') 

	for i := 0; i < n; i++ {
		fmt.Printf("\nEnter subject #%d name:\n", i+1)
		subjectLine, _ := reader.ReadString('\n')  
		subject = strings.TrimSpace(subjectLine)  

		fmt.Printf("Enter grade for %s:\n", subject)
		fmt.Scan(&grade)
		reader.ReadString('\n') 

		err := validate(grade)
		if err != nil {
			fmt.Println(err)
			i-- 
			continue
		}
		grades[subject] = grade
	}

	mean := average(grades, n)
	fmt.Printf("Name: %s\n", name)
	fmt.Println("Your grades are:")
	for sub, grade := range grades {
		fmt.Printf("- %s: %d\n", sub, grade)
	}
	fmt.Printf("Your average grade is: %.2f\n", mean)
}
