package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"

	"../config"
)

func GetArg() string {
	args := os.Args
	if len(os.Args) <= 1 {
		fmt.Println("\033[0;31m not enoght args")
		fmt.Println("\033[37m try :\n\t  \033[33m <app-name> <arg>")
		fmt.Println("\033[37m for example:\n\t \033[33m  ./botline 123")
		fmt.Println("\033[37m or\n\t \033[33m go run *go 123")
		os.Exit(0)
	}
	return args[1]
}

func InArrayInt64(arr []int64, str int64) bool {
	for _, tar := range arr {
		if tar == str {
			return true
		}
	}
	return false
}

func InArrayString(ArrList []string, rstr string) bool {
	for _, x := range ArrList {
		if x == rstr {
			return true
		}
	}
	return false
}

func RemoveString(s []string, r string) []string {
	new := make([]string, len(s))
	copy(new, s)
	for i, v := range new {
		if v == r {
			return append(new[:i], new[i+1:]...)
		}
	}
	return s
}

func PanicHandle(s string) {
	if r := recover(); r != nil {
		Ides := fmt.Sprintf("\nEror: %s \nFunc: %v", r, s)
		println(Ides)
	}
}

func GetMentionData(data string) []string {
	var midmen []string
	var midbefore []string
	res := config.Mentions{}
	json.Unmarshal([]byte(data), &res)
	for _, v := range res.MENTIONEES {
		if InArrayString(midbefore, v.Mid) == false {
			midbefore = append(midbefore, v.Mid)
			midmen = append(midmen, v.Mid)
		}
	}

	return midmen
}

func GetIP() net.IP {
	conn, err := net.Dial("udp", "0.0.0.0:80")
	if err != nil {
		return nil
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}

func IndexOf(data []string, element string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

func Checkmulti(list1 []string, list2 []string) bool {
	for _, v := range list1 {
		if InArrayString(list2, v) {
			return true
		}
	}
	return false
}

func StripOut(kata string) string {
	kata = strings.TrimSpace(kata)
	return kata
}

var letters = []rune("123456789")

func RandomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func Fancy(text string) string {
	return text
}

func AppendLastD(s [][]string, e []string) [][]string {
	s = append(s, e)
	if len(s) >= 1000 {
		s = s[100:]
		return s
	}
	return s
}

func AppendLast(s []string, e string) []string {
	s = RemoveString(s, e)
	s = append(s, e)
	if len(s) >= 1000 {
		s = s[100:]
		return s
	}
	return s
}
