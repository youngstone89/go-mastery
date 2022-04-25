package algorithms_test

import (
	"fmt"
	"strings"
	"testing"
)

func TestLongestSubstring(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
	}{
		{
			desc:  "abcabcbb",
			input: "abcabcbb",
		},
		{
			desc:  "bbbbb",
			input: "bbbbb",
		},
		{
			desc:  "pwwkew",
			input: "pwwkew",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			output, length := lengthOfLongestSubstring(tC.input)
			t.Log(output, length)
		})
	}
}

func lengthOfLongestSubstring(s string) (string, int) {
	output := ""
	length := 0

	//the longest substring
	// should know very case of substring of a string
	// "abcabcbb"
	// length = 8
	// index = 0, cases = 8
	// a, ab, abc, abca, abcab, abcabc, abacbcb, abcabcbb
	// a=arr[:0]=[:idx]
	// ab=arr[0:1]=[idx:+1]
	// abc=arr[0:2]=[idx:+2]
	// abcabcbb=arr[0:7]=[idx:len-1]
	// without repeating
	// dup check?
	// make char arr
	// compare every single char one by one
	// if ther is any dupe, filter out
	// this way, you can sort out substrings that don't have repeated characters.
	// and repeat this throughout all index's cases
	// and find the longest one among them.

	// b, bc, bca, bcab, babc, bvabcb, bacbcbb
	// index = 1, cases = 7
	// c, ca, cab, cabc, cabcb, cabcbb,
	// ....
	// b, bb
	// index = 6, cases = 2
	// b
	// index = 7, cases = 1

	charArr := []rune(s)
	for i := 0; i < len(charArr); i++ {
		substrArr := make([]string, len(charArr))
		for j := i; j < len(charArr)-1; j++ {
			subCharArr := charArr[0:j]
			if j == 0 {
				substrArr = append(substrArr, string(charArr[0:j]))
			}
			// if dup? continue
			compStr := ""
			if !isDup(subCharArr, compStr, substrArr) {
				substrArr = append(substrArr, string(charArr[0:j]))
			}

		}
		fmt.Println(substrArr)
	}

	return output, length
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
