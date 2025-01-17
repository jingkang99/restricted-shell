package main

import (
	"fmt"
	"time"
	"bufio"
	"errors"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var allowed = `\b(?:awk|cat|cd|cp|curl|echo|exit|file|find|grep|head|ls|ll|ldd|locate|more|netstat|ping|ps|pwd|sed|sort|tail|tar|uniq|w|wc|help|plus)\b`
var helpStr =      `awk cat cd cp curl echo exit file find grep head ls ll ldd locate more netstat ping ps pwd sed sort tail tar uniq w wc help plus`
var prohibt = `\s*(?:/proc|/var|/etc|/boot|/dev|/root|/bin|/sbin|/lib|/usr|/sys)\b`

func main() {
	regx := regexp.MustCompile(allowed)
	regp := regexp.MustCompile(prohibt)

	f, err := os.CreateTemp(".", "rshell-")
	if err != nil {
        	fmt.Fprintln(os.Stderr, err)
	}
	defer f.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		now := fmt.Sprint(time.Now())
		if _, err := f.Write( []byte( ( now + "\t" + cmdString) ) ); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if len(cmdString) < 2 {
			continue
		}

		if regp.MatchString(cmdString) {
			fmt.Fprintln(os.Stderr, "path not allowed")
			continue
		}

		if regx.MatchString(cmdString) {
			err = runCommand(cmdString)
		} else {
			fmt.Fprintln(os.Stderr, strings.TrimSpace(cmdString)+" not allowed")
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func runCommand(commandStr string) error {
	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommandStr := strings.Fields(commandStr)
	switch arrCommandStr[0] {
	case "exit":
		os.Exit(0)
	case "help":
		fmt.Fprintln(os.Stdout, "allowed cmd: "+helpStr)
		return nil
	case "ll":
		strArr := []string {"ls", "-l"}
		arrCommandStr = strArr
	case "plus":
		if len(arrCommandStr) < 3 {
			return errors.New("2 or more arguments required")
		}
		arrNum := []int64{}
		for i, arg := range arrCommandStr {
			if i == 0 {
				continue
			}
			n, _ := strconv.ParseInt(arg, 10, 64)
			arrNum = append(arrNum, n)
		}
		fmt.Fprintln(os.Stdout, sum(arrNum...))
		return nil
	}
	cmd := exec.Command(arrCommandStr[0], arrCommandStr[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func sum(numbers ...int64) int64 {
	res := int64(0)
	for _, num := range numbers {
		res += num
	}
	return res
}

// https://hackernoon.com/today-i-learned-making-a-simple-interactive-shell-application-in-golang-aa83adcb266a
