package main

import (
	"errors"
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"time"

	//m "mvc_app/local_package/model"
	"sync"
)

// Step 1: In your main.go directory run: fo mod init examples/package
// Step 2: import ( examples/package/<YOUR PACKAGE NAME> )
// Step 3: Use it! with <YOUR PACKAGE NAME>.Something
// In order to export local package you have specify variable or function name as pascal case to export out

// To get around time.Sleep() we use below
var wg = sync.WaitGroup{}
var m = sync.RWMutex{}

func main() {
	//runtime.GOMAXPROCS(20)
	//m.PrintfModel()
	fmt.Println("Hi this is main")
	fmt.Println("Threads available checking", runtime.GOMAXPROCS(-1))
	res, err := GeneratePassword(12)
	if err != nil {
		fmt.Println("Error []", err)
	}
	fmt.Println(res)
	//wg.Add(1)
	//go func() {
	//	m.RLock()
	//	checkType(123123123)
	//	m.RUnlock()
	//	wg.Done()
	//}()
	//wg.Wait()
}

func checkType(values interface{}) {
	switch values.(type) {
	case int:
		fmt.Println("Type: int")
	case int16:
		fmt.Println("Type: int16")
	case int32:
		fmt.Println("Type: int32")
	case int64:
		fmt.Println("Type: int64")
	case string:
		fmt.Println("Type: string")
	case float32:
		fmt.Println("Type: float32")
	case float64:
		fmt.Println("Type: float64")
	case bool:
		fmt.Println("Type: bool")
	case byte:
		fmt.Println("Type: byte")
	case []rune:
		fmt.Println("Type: []rune")
	default:
		fmt.Println("No idea")
	}
}

func GeneratePassword(length int) (string, error) {
	var pw []string

	if length < 12 {
		return "", errors.New("password must longer than 12 character")
	}
	if length > 20 {
		return "", errors.New("password too long")
	}

	for i := 0; i < length; i++ {
		if i >= 10 {
			pw = append(pw, string(RandomChar(100000, i)))
		} else if i >= 7 {
			pw = append(pw, string(RandomSymbol(100, i)))
		} else if i >= 3 {
			pw = append(pw, strconv.Itoa(Base10Int(i)))
		} else {
			pw = append(pw, string(RandomChar(10000000, i)))
		}
	}
	return strings.Join(pw, ""), nil
}

func RandomInt(max int, randomNum int) (int) {
	rand.Seed(time.Now().Unix())
	return rand.Intn(randomNum + max)
}

func RandomChar(max int, randomNum int) (rune) {
	var text = []rune("abcdefghijklmnopqrstuvxyzABCDEFGHIJKLMNOPQRSTUVXYZ")
	char := text[RandomInt(max, randomNum) % len(text)]
	return char
}

func RandomSymbol(max int, randomNum int) (rune) {
	var symbol = []rune("!@#$%^&*()_+=-;:'|,./?><")
	char := symbol[RandomInt(max, randomNum) % len(symbol)]
	return char
}

func Base10Int(randomNum int) (int) {
	var Nums = []rune("0123456789")
	now := time.Now()
	num, err := strconv.Atoi(string(Nums[RandomInt(now.Minute() + randomNum, now.Nanosecond()) % len(Nums)]))
	if err != nil {
		fmt.Println("ERROR [*] while converting a rune to int -> Base10Int()", err)
	}
	return num 
}