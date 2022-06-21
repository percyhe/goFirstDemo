package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

/**
帮助文本
*/
func flagUsage() {
	usageText := `main is an example cli tool.
  Usage:
  main command [arguments]
  The commands are:
  uppercase uppercase a string 
  lowercase lowercase a string
Use "main [command] --help" for more information about a command.`
	fmt.Fprintf(os.Stderr, "%s\n\n", usageText)
}
func TestFlag() {
	flag.Usage = flagUsage
	uppercaseCmd := flag.NewFlagSet("uppercase", flag.ExitOnError)
	lowercaseCmd := flag.NewFlagSet("lowercase", flag.ExitOnError)

	if len(os.Args) == 1 {
		flag.Usage()
		return
	}
	switch os.Args[1] {

	case "uppercase":
		s := uppercaseCmd.String("s", "", "A string of text to be uppercased")
		uppercaseCmd.Parse(os.Args[2:])
		fmt.Println(strings.ToUpper(*s))
	case "lowercase":
		s := lowercaseCmd.String("s", "", "A string of text to be lowercased")
		lowercaseCmd.Parse(os.Args[2:])
		fmt.Println(strings.ToLower(*s))
	default:
		flag.Usage()

	}
}

func main() {
	TestFlag()
	/*	flag.Usage = func() {
				usageText := `Usage example04 [OPTION]
		  An example of customizing usage output

		  -s --s  example string argument,default:  String help text
		  -i --i  example integer argument,default: Int help text
		  -b --b  example boolean argument,default: Bool help text`
				fmt.Fprintf(os.Stderr,"%s\n",usageText)
			}
			s := flag.String("s","Hello world","String help text")
			i := flag.Int("i",1,"Int help text")
			b := flag.Bool("b",false,"Bool help text")
			flag.Parse()
			fmt.Println("value of s:",*s)
			fmt.Println("value of i:",*i)
			fmt.Println("value of b:",*b)*/
}
