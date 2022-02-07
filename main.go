package main

import "go-mastery/pkg/cli"

func main() {
	cli.DoStringPrompt()
	cli.DoPasswordPrompt()
	cli.DoYesOrNo()
	cli.DoCheckBoxes()

}
