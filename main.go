package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"strings"
)

func CheckBoxes(label string, opts []string) []string {
	res := []string{}
	prompt := &survey.MultiSelect{
		Message: label,
		Options: opts,
	}
	survey.AskOne(prompt, &res)

	return res
}

func main() {
	answers := CheckBoxes(
		"Which are your favourite programming languages?",
		[]string{
			"C",
			"Python",
			"Java",
			"C++",
			"C#",
			"Visual Basic",
			"JavaScript",
			"PHP",
			"Assembly Language",
			"SQL",
			"Groovy",
			"Classic Visual Basic",
			"Fortran",
			"R",
			"Ruby",
			"Swift",
			"MATLAB",
			"Go",
			"Prolog",
			"Perl",
		},
	)
	s := strings.Join(answers, ", ")
	fmt.Printf("You chose: %s\n", s)
}
