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
	useS *bool
	argF int
	argS int
}

func equalStr(parametr *Flags, str1 string, str2 string) (result bool) {
	if *parametr.useI {
		if *parametr.useS {
			if (str1 == "") && (str2 == "") {
				return true
			}
			if str1 == "" {
				result = strings.EqualFold(str1, str2[parametr.argS:len(str2)])
				return result
			}
			if str2 == "" {
				result = strings.EqualFold(str1[parametr.argS:len(str1)], str2)
				return result
			}
			result = strings.EqualFold(str1[parametr.argS:len(str1)], str2[parametr.argS:len(str2)])
			return result
		}
		result = strings.EqualFold(str1, str2)
		return result
	}
	if *parametr.useS {
		if (str1 == "") && (str2 == "") {
			return true
		}
		if str1 == "" {
			result = str1 == str2[parametr.argS:len(str2)]
			return result
		}
		if str2 == "" {
			result = str1[parametr.argS:len(str1)] == str2
			return result
		}
		result = str1[parametr.argS:len(str1)] == str2[parametr.argS:len(str2)]
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

		if equalStr(parametr, previosString, line) == false {
			if *parametr.useC {
				fmt.Print(n, " ", previosString)
			} else if *parametr.useD && n > 1 {
				fmt.Print(previosString)
			} else if *parametr.useU && n == 1 {
				fmt.Print(previosString)
			} else if *parametr.useS {
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
		useF: flag.Bool("f", false, "не учитывать первые num_fields полей в строке"),
		useS: flag.Bool("s", false, "не учитывать первые num_chars символов в строке")}
	flag.Parse()

	i := 0
	if *flags.useF {
		if *flags.useS {
			i = 2
			s := flag.Arg(0)
			flags.argS, _ = strconv.Atoi(s)
			f := flag.Arg(1)
			flags.argF, _ = strconv.Atoi(f)
		} else {
			i = 1
			f := flag.Arg(0)
			flags.argF, _ = strconv.Atoi(f)
		}
	} else if *flags.useS {
		i = 1
		s := flag.Arg(0)
		flags.argS, _ = strconv.Atoi(s)
	}
	path := flag.Arg(i)

	if *flags.useC || *flags.useD || *flags.useU || *flags.useF || *flags.useS {
		readFile(path, &flags)
		return
	}
	defaulFile(path)
}
