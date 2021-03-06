package main

import (
	types "coursera/hw3_bench/types"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {}

var r *regexp.Regexp

func init() {
	r = regexp.MustCompile("@")
}

// вам надо написать более быструю оптимальную этой функции
func FastSearch(out io.Writer) {

	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	seenBrowsers := map[string]bool{}
	uniqueBrowsers := 0
	foundUsers := ""

	lines := strings.Split(string(fileContents), "\n")

	for i, line := range lines {
		user := types.UserInfo{} 
		// fmt.Printf("%v %v\n", err, line)

		err := user.UnmarshalJSON([]byte(line))
		if err != nil {
			panic(err)
		}

		// data logic
		isAndroid := false
		isMSIE := false

		for _, browser := range user.Browsers {
			if strings.Contains(browser, "Android") {
				isAndroid = true

				if seen := seenBrowsers[browser]; !seen {
					// log.Printf("SLOW New browser: %s, first seen: %s", browser, user["name"])
					seenBrowsers[browser] = true
					uniqueBrowsers++
				}

			} else if strings.Contains(browser, "MSIE") {
				isMSIE = true
				if seen := seenBrowsers[browser]; !seen {
					// log.Printf("SLOW New browser: %s, first seen: %s", browser, user["name"])
					seenBrowsers[browser] = true
					uniqueBrowsers++
				}
			}
		}

		if !(isAndroid && isMSIE) {
			continue
		}

		// log.Println("Android and MSIE user:", user["name"], user["email"])
		email := r.ReplaceAllString(user.Email, " [at] ")
		foundUsers += fmt.Sprintf("[%d] %s <%s>\n", i, user.Name, email)
	}

	fmt.Fprintln(out, "found users:\n"+foundUsers)
	fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers))
}
