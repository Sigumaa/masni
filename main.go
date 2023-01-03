package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func TextToMorse(text string, morseCode map[rune]string) (string, error) {

	runes := []rune(text)
	result := ""

	for _, r := range runes {
		result += morseCode[r]
		result += " "
	}
	return result, nil
}

func MorseToNI(morse string) (string, error) {
	runes := []rune(morse)
	result := ""

	for _, r := range runes {
		switch r {
		case '・':
			result += "に"
		case '－':
			result += "にー"
		case ' ':
			result += " "
		}
	}
	return result, nil
}

func NIToMorse(ni string) (string, error) {
	runes := []rune(ni)
	result := ""

	// "にー"という二文字は一文字の-になる必要がある

	for i, r := range runes {
		if r == 'に' {
			if i+1 < len(runes) && runes[i+1] == 'ー' {
				result += "－"
			} else {
				result += "・"
			}
		} else if r == 'ー' {
			// 何もしない
		} else if r == ' ' {
			result += " "
		}
	}
	return result, nil
}

func MorseToText(morse string, morseCode map[rune]string) (string, error) {
	result := ""
	word := strings.Split(morse, " ")
	for _, w := range word {
		for hiragana, m := range morseCode {
			if m == w {
				result += string(hiragana)
				break
			}
		}
	}
	return result, nil
}

var s = bufio.NewScanner(os.Stdin)

func main() {
	morseCode := map[rune]string{
		'あ': "－－・－－",
		'い': "・－",
		'う': "・・－",
		'え': "－・－－－",
		'お': "・－・・・",
		'か': "・－・・",
		'き': "－・－・・",
		'く': "・・・－",
		'け': "－・－－",
		'こ': "－－－－",
		'さ': "－・－・－",
		'し': "－－・－・",
		'す': "－－－・－",
		'せ': "・－－－・",
		'そ': "－－－・",
		'た': "－・",
		'ち': "・・－・",
		'つ': "・－－・",
		'て': "・－・－－",
		'と': "・・－・・",
		'な': "・－・",
		'に': "－・－・",
		'ぬ': "・・・・",
		'ね': "－－・－",
		'の': "・・－－",
		'は': "－・・・",
		'ひ': "－－・・－",
		'ふ': "－－・・",
		'へ': "・",
		'ほ': "－・・",
		'ま': "－・・－",
		'み': "・・－・－",
		'む': "－",
		'め': "－・・・－",
		'も': "－・・－・",
		'や': "・－－",
		'ゆ': "－・・－－",
		'よ': "－－",
		'ら': "・・・",
		'り': "－－・",
		'る': "－・－－・",
		'れ': "－－－",
		'ろ': "・－・－",
		'わ': "－・－",
		'を': "・－－－",
		'ん': "・－・－・",
	}
	for {
		fmt.Print("\nテキストを変換するなら1、にーを変換するなら2を入力してください > ")
		s.Scan()
		input := s.Text()

		if input != "1" && input != "2" {
			fmt.Println("1か2を入力してください")
			continue
		}

		fmt.Print("input > ")
		s.Scan()
		text := s.Text()
		var output string
		if input == "1" {
			result, err := TextToMorse(text, morseCode)
			if err != nil {
				fmt.Println(err)
				continue
			}

			result, err = MorseToNI(result)
			if err != nil {
				fmt.Println(err)
				continue
			}
			output = result
		} else {
			result, err := NIToMorse(text)
			if err != nil {
				fmt.Println(err)
				continue
			}

			result, err = MorseToText(result, morseCode)
			if err != nil {
				fmt.Println(err)
				continue
			}
			output = result
		}

		fmt.Println("output >", output)
	}
}
