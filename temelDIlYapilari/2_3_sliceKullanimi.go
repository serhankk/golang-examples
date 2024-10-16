package main

import "fmt"

func main() {

	// Initialize slice
	programmingLanguages := make([]string, 0, 10)
	programmingLanguages = append(programmingLanguages, "go", "javascript", "python", "c++", "c#", "ruby", "java", "react")
	fmt.Printf("programming languages: %v\n", programmingLanguages)

	backendLanguages := make([]string, 0, 10)
	backendLanguages = append(backendLanguages, programmingLanguages[0], programmingLanguages[2],
		programmingLanguages[3], programmingLanguages[4], programmingLanguages[5], programmingLanguages[6])

	fmt.Printf("backend languages: %v\n", backendLanguages)

	frontendLanguages := make([]string, 0, 10)
	frontendLanguages = append(frontendLanguages, programmingLanguages[1], programmingLanguages[7])
	fmt.Printf("frontend languages: %v\n", frontendLanguages)
}
