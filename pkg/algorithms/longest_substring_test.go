package algorithms_test

import (
	"fmt"
	"strings"
	"testing"
)

// find the longest substring without repeating characters
func TestLongestSubstring(t *testing.T) {
	testCases := []struct {
		desc   string
		input  string
		expect string
		output int
	}{
		{
			desc:   "abcabcbb",
			input:  "abcabcbb",
			expect: "abc",
			output: 3,
		},
		{
			desc:   "bbbbb",
			input:  "bbbbb",
			expect: "b",
			output: 1,
		},
		{
			desc:   "pwwkew",
			input:  "pwwkew",
			expect: "wke",
			output: 3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			output, length := lengthOfLongestSubstring(tC.input)
			if tC.output != length {
				t.Fail()
			}
			if tC.expect != output {
				t.Fail()
			}
			t.Logf("passed %s, %d", output, length)
		})
	}
}

func lengthOfLongestSubstring(s string) (string, int) {

	var outputList []string
	charArr := []rune(s)
	// outer loop
	for i := 0; i < len(charArr); i++ {

		var outerLoopOutput []string

		// inner loop
	INNER:
		for j := i + 1; j < len(charArr); j++ {

			// map for checking repetitiveness of character
			distinctMap := make(map[string]int)

			s := charArr[:j]
			substrCharArr := []rune(s)

			// check if idstinctMap has the character, if not put it with # of appearance.
			for _, v := range substrCharArr {
				s := string(v)
				if _, exist := distinctMap[s]; exist {
					// increase by 1
					distinctMap[s] += 1
				} else {
					// set to 1
					distinctMap[s] = 1
				}
			}

			// check repetition
			for _, v := range distinctMap {
				if v > 1 {
					fmt.Printf("\n %s has repetitive characters", string(s))
					continue INNER
				}
			}
			fmt.Printf("\n %s has passed repetitiveness check", string(s))
			outerLoopOutput = append(outerLoopOutput, string(s))

		}
		longest := 0
		longestStr := ""
		// find the longest one in ordered list
		for _, s := range outerLoopOutput {
			cArr := []rune(s)
			length := len(cArr)
			if longest < length {
				longest = length
				longestStr = s
			}
		}
		outputList = append(outputList, longestStr)
	}

	var output string
	var size int
	for _, s := range outputList {
		b := []byte(s)
		if size < len(b) {
			size = len(b)
			output = s
		}
	}

	return output, size
}

func isDup(subCharArr []rune, compStr string, substrArr []string) bool {
	for k := 0; k < len(subCharArr); k++ {
		if k == 0 {
			compStr = substrArr[k]
			continue
		}
		if strings.Contains(compStr, string(substrArr[k])) {
			return true
		}
	}
	return false
}
