package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/AlecAivazis/survey/v2"
	"golang.org/x/term"
)

func DoStringPrompt() {
	name := StringPrompt("what is your name?")
	fmt.Printf("Hello, %s!\n", name)
}

func DoPasswordPrompt() {
	password := PasswordPrompt("what is your password?")
	fmt.Printf("Hello, %s!\n", password)
}
func DoYesOrNo() {
	ok := YesNoPrompt("Dev.to is awesome!", true)
	if ok {
		fmt.Println("Agree!")
	} else {
		fmt.Println("Huh?")
	}
}

func DoCheckBoxes() {
	answers := Checkboxes(
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
	fmt.Println("Oh, I see! You like", s)

}
func Checkboxes(label string, opts []string) []string {
	res := []string{}
	prompt := &survey.MultiSelect{
		Message: label,
		Options: opts,
	}
	survey.AskOne(prompt, &res)

	return res
}

func YesNoPrompt(label string, def bool) bool {
	choices := "Y/n"
	if !def {
		choices = "y/N"
	}

	r := bufio.NewReader(os.Stdin)
	var s string

	for {
		fmt.Fprintf(os.Stderr, "%s (%s) ", label, choices)
		s, _ = r.ReadString('\n')
		s = strings.TrimSpace(s)
		if s == "" {
			return def
		}
		s = strings.ToLower(s)
		if s == "y" || s == "yes" {
			return true
		}
		if s == "n" || s == "no" {
			return false
		}
	}
}
func StringPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}
func PasswordPrompt(label string) string {
	var s string
	for {
		fmt.Fprint(os.Stderr, label+" ")
		b, _ := term.ReadPassword(int(syscall.Stdin))
		s = string(b)
		if s != "" {
			break
		}
	}
	fmt.Println()
	return s
}
