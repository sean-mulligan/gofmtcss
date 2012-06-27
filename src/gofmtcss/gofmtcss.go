package gofmtcss

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"sort"
)

func main() {
	//Check that we have the right number of arguments.
	if len(os.Args) < 2 {
		log.Fatal("Please pass the file path in the command line.")
		return
	}
	path := os.Args[1]
	file, err := ioutil.ReadFile(path);
	if  err != nil {
		log.Fatal(err)
		return
	}
	input := string(file)
	
	//Create a regex that picks out the bracket sets {...} out of the css files.
	cssRegex, err := regexp.Compile(`\{[^\}]+\}`)
	if  err != nil {
		log.Fatal(err)
		return
	}
	
	//Iterate through the sections, replacing the found sections with the
	//ordered sections.
	//Issue: This has the possibility at the moment for replacing the
	//	wrong section if there are equivalent sections in the document.
	//	The matching patterns used are somewhat naive.
	sections := cssRegex.FindAllString(input,-1)
	output := input
	for _, exp := range sections {
		//Skip empty iterations
		if strings.Trim(exp, "{} \n\t\r") != "" {
			//Remove the starting and ending whitespace and brackets
			temp := strings.Trim(fmt.Sprintf("%s", exp), "{} \n\t\r")
			//Split the rules. We assume we are passed valid CSS.
			rules := strings.SplitAfter(temp,";")
			//Format the rules
			for pos, r := range rules {
				rules[pos] = fmt.Sprintf("\t%s", strings.Trim(r, " \n\t\r"))
			}
			//Sort the rules
			sort.Strings(rules)
			//Join the rules back together and put back in the brackets
			rule := strings.Join(rules, "\n")
			rule = fmt.Sprintf("{%s\n}", rule)
			//Find the original ruleset and replace it with the formatted one
			output = strings.Replace(output, exp, rule, 1)
		}
	}
	ioutil.WriteFile(path, []byte(output), 0)
}