package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Flags struct {
	useC *bool
	useD *bool
	useU *bool
	useI *bool
	useF *bool
	argF int
	/*	useF *string
		useS *string
		string*/
}

func equalStr(parametrI *bool, str1 string, str2 string) (result bool) {
	if *parametrI {
		result = strings.EqualFold(str1, str2)
		return result
	}
	result = str1 == str2
	return result
}

func readFile(path string, parametr *Flags) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Unable to open file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var previosString = ""
	for n := 0; n > -1; n++ {
		line, err := reader.ReadString('\n')

		if n == 0 {
			previosString = line
		}

		if *parametr.useF {
			strings.SplitAfterN(previosString, " ", parametr.argF)
			fmt.Print(parametr.argF)
		}

		if equalStr(parametr.useI, previosString, line) == false {
			if *parametr.useC {
				fmt.Print(n, " ", previosString)
			} else if *parametr.useD && n > 1 {
				fmt.Print(previosString)
			} else if *parametr.useU && n == 1 {
				fmt.Print(previosString)
			}
			previosString = line
			n = 0
		}

		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}
	}
}
func defaulFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Unable to open file:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}
		fmt.Print(line)
	}
}

func main() {
	var flags = Flags{useC: flag.Bool("c", false, "вывод кол-ва вхождений строки"),
		useD: flag.Bool("d", false, "вывод строк, которые повторились"),
		useU: flag.Bool("u", false, "вывод строк, которые не повторились"),
		useI: flag.Bool("i", false, "не учитывать регистр букв"),
		useF: flag.Bool("f", false, "не учитывать первые num_fields полей в строке")}
	flag.Parse()
	/*,
	useF: flag.String("f", "f", "не учитывать первые num_fields полей в строке"),
	useS: flag.String("s", "s", "не учитывать первые num_chars символов в строке"),
	}*/
	/*if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}
	path := os.Args[1]*/
	i := 0
	if *flags.useF {
		i = 1
		s := flag.Arg(0)
		flags.argF, _ = strconv.Atoi(s)
	}
	path := flag.Arg(i)

	//path := "/Users/itsumaden/go/src/Tests/text.txt"
	if *flags.useC || *flags.useD || *flags.useU {
		readFile(path, &flags)
		return
	}
	defaulFile(path)
}
