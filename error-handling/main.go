package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type GradeError struct {
	LineNumber int
	Value      string
	Operation  string
	Err        error
}

// When any error needs to be printer or converted to string (fmt.Printf)
// Go automatically calls the Error() method
func (e *GradeError) Error() string {
	return fmt.Sprintf("error during %s at line %d (value: %s): %v", e.Operation, e.LineNumber, e.Value, e.Err)
}

func calculateAverage(filename string) (float64, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, fmt.Errorf("failed to open file: %w", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum float64
	var count int

	for lineNum := 1; scanner.Scan(); lineNum++ {
		gradeStr := scanner.Text()

		if gradeStr == "" {
			continue
		}

		grade, err := strconv.ParseFloat(gradeStr, 64)

		if err != nil {
			return 0, &GradeError{
				LineNumber: lineNum,
				Value:      gradeStr,
				Operation:  "grade conversion",
				Err:        err,
			}
		}

		if grade > 100 || grade <= 0 {
			return 0, &GradeError{
				LineNumber: lineNum,
				Value:      gradeStr,
				Operation:  "grade validation",
				Err:        fmt.Errorf("grade out of valid range (0-100)"),
			}
		}

		sum += grade
		count++
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file %w", err)
	}

	if count == 0 {
		return 0, fmt.Errorf("no valid grades found in file")
	}

	return sum / float64(count), nil
}

func main() {

	average, err := calculateAverage("grades.txt")

	if err != nil {
		// Attempts to convert err to our specific *GradeError type
		// Tells us whether this conversion was successful
		// Think of it like trying to identify a specific breed of dog. You might have a general "animal" (our error interface),
		// but you want to know if it's specifically a "golden retriever" (our *GradeError).
		if gradeErr, ok := err.(*GradeError); ok {
			fmt.Printf("Grade process error: %v\n", gradeErr)
		} else {
			fmt.Println("Error: %v\n", err)
		}

		os.Exit(1)
	}

	fmt.Println(average)
}
