package userdefinedtypes

import "fmt"

type HighScore Score
type Score int

func testScore() {

	var i int = 30
	var s Score = 100
	var hs HighScore = 200

	//hs = s
	//s = i
	s = Score(i)
	hs = HighScore(s)

	fmt.Println(hs)

}
