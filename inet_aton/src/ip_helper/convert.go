package ip_helper

import (
	"errors"
	"unicode"
)

//the state of reading ip string
type STATECODE int

const (
	ERROR 		STATECODE = -1
	READING_DOT STATECODE = 1
	READING_NUM	STATECODE = 2
)

func IpConvert(ip string) (int, error) {
	digit := 0
	dotNum := 0         // num of dot
	input := 0         // current reading
	state := READING_NUM // reading state
	result := 0
	i := 0	//index of ip string

	if unicode.IsSpace(rune(ip[i])) {
		i = 1
	}

	for ; i < len(ip) ; i++ {
		input = int(ip[i])
		//reading digit
		if state == READING_NUM {
			if input >= '0' && input <= '9' {
				state = READING_DOT
				digit = input - '0'
			} else {
				//fmt.Println("error 000. index = ", i)
				return -1, errors.New("invalid ip string")
			}
		//reading dot or error
		} else {
			if input >= '0' && input <= '9' {
				digit = digit * 10 + (input-'0')
				if digit > 255 || digit < 0 {
					//fmt.Println("error 001. index = ", i)
					return -1, errors.New("invalid ip string")
				}
			} else if input == '.' {
				if dotNum == 3 {
					//fmt.Println("error 002. index = ", i)
					return -1, errors.New("invalid ip string")
				}
				dotNum++
				result = (result<<8) + digit
				digit = 0
				if unicode.IsSpace(rune(ip[i+1])) {
					//如果dot后面跟着space，下一次循环不能直接进去READING_NUM状态
					state = READING_DOT
				} else {
					state = READING_NUM
				}
			} else if unicode.IsSpace(rune(input)) {
				if unicode.IsDigit(rune(ip[i-1])) && unicode.IsDigit(rune(ip[i+1])) {
					//空格前后都是数字
					//fmt.Println("error 00x. index = ", i)
					return -1, errors.New("invalid ip string")
				} else {
					continue
				}
			} else if i == len(ip) {
				if dotNum != 3 {
					//fmt.Println("error 003. index = ", i)
					return -1, errors.New("invalid ip string")
				}
				result = (result<<8) + digit
				break
			} else {
				//fmt.Println("error 004. index = ", i)
				return -1, errors.New("invalid ip string")
			}
		}
	}

	if i == len(ip) {
		if dotNum != 3 {
			//fmt.Println("error 005")
			return -1, errors.New("invalid ip string")
		}
		result = (result<<8) + digit
	}

	return result, nil
}
