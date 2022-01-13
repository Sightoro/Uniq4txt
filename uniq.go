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
	if parametr.argS > len(str1) {
		str1 = ""
	}
	if parametr.argS > len(str2) {
		str2 = ""
	}
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

func readFile(pathIn string, pathOut string, parametr *Flags) {
	fileIn, err := os.Open(pathIn)
	if err != nil {
		fmt.Println("Unable to open input_file:", err)
		return
	}

	defer fileIn.Close()

	reader := bufio.NewReader(fileIn)
	var previosString = ""
	for n := 0; n > -1; n++ {

		line, err := reader.ReadString('\n')

		if n == 0 {
			previosString = line
		}
		if equalStr(parametr, previosString, line) == false {
			if *parametr.useC {
				str2Output(strconv.Itoa(n)+" "+previosString, pathOut)
				fmt.Print(n, " ", previosString)
			} else if *parametr.useD && n > 1 {
				str2Output(previosString, pathOut)
				fmt.Print(previosString)
			} else if *parametr.useU && n == 1 {
				str2Output(previosString, pathOut)
				fmt.Print(previosString)
			} else if *parametr.useS {
				str2Output(previosString, pathOut)
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

func defaulFile(pathIn string, pathOut string) {
	fileIn, err := os.Open(pathIn)
	fileOut, err1 := os.Create(pathOut)

	if err != nil {
		fmt.Println("Unable to open file:", err)
		return
	}
	if err1 != nil {
		fmt.Println("Unable to create file:", err1)
		return
	}
	defer fileIn.Close()
	defer fileOut.Close()
	reader := bufio.NewReader(fileIn)
	writer := bufio.NewWriter(fileOut)
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
		writer.WriteString(line)
		writer.Flush()
		fmt.Print(line)
	}
}

func str2Output(str string, output string) {
	f, err := os.OpenFile(output, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString(str); err != nil {
		panic(err)
	}
}

/*func CheckFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		panic("file")
		return false
	}
	return true
}*/

func main() {
	var flags = Flags{useC: flag.Bool("c", false, "вывод кол-ва вхождений строки"),
		useD: flag.Bool("d", false, "вывод строк, которые повторились"),
		useU: flag.Bool("u", false, "вывод строк, которые не повторились"),
		useI: flag.Bool("i", false, "не учитывать регистр букв"),
		useF: flag.Bool("f", false, "не учитывать первые num_fields полей в строке"),
		useS: flag.Bool("s", false, "не учитывать первые num_chars символов в строке")}
	flag.Parse()

	in := 0
	out := 1
	pathIn := ""
	pathOut := ""
	if *flags.useF {
		if *flags.useS {
			in = 2
			out = 3
			s := flag.Arg(0)
			flags.argS, _ = strconv.Atoi(s)
			f := flag.Arg(1)
			flags.argF, _ = strconv.Atoi(f)
		} else {
			in = 1
			out = 2
			f := flag.Arg(0)
			flags.argF, _ = strconv.Atoi(f)
		}
	} else if *flags.useS {
		in = 1
		out = 2
		s := flag.Arg(0)
		flags.argS, _ = strconv.Atoi(s)
	}
	pathIn = flag.Arg(in)
	pathOut = flag.Arg(out)
	for {
		if pathIn == "" {
			fmt.Println("Type input_file name: ")
			fmt.Fscan(os.Stdin, &pathIn)
		} else {
			break
		}
	}

	for {
		if pathOut == "" {
			fmt.Println("Type out_file name: ")
			fmt.Fscan(os.Stdin, &pathOut)
		} else {
			break
		}
	}

	//CheckFileExist(pathOut)

	if *flags.useC || *flags.useD || *flags.useU || *flags.useF || *flags.useS {
		readFile(pathIn, pathOut, &flags)
		return
	}
	defaulFile(pathIn, pathOut)
}
