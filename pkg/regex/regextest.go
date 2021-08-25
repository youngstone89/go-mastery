package regex

import (
	"fmt"
	"regexp"
)

func DoRegexTest()  {

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	//https://log.api.nuance.com/producers
	host_domain := "nuance.com"
	endpoint_subpath := "/producers"
	regex_pattern := fmt.Sprintf(`^http(s*)://(.+)%s%s$`,host_domain,endpoint_subpath)
	test_url := "https://log.api.nuance.com/producers"
	//test_url := "http://log.api.nuance.com/producers"
	//test_url := "http://crd4.dev.log.api.nuance.com/producers"
	//test_url := "https://crd4.dev.log.api.nuance.com/producers"
	//test_url := "https://crd4.dev.log.api.nuancex.com/producer"
	fmt.Println(regex_pattern)
	isMatch, _ := regexp.MatchString("^http(s*)://(.+)nuance.com/producers$", test_url)
	fmt.Println(isMatch)
}
