package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var s = bufio.NewScanner(os.Stdin)

func main() {
	for {
		fmt.Print("\nテキストを変換するなら1、にーを変換するなら2を入力してください > ")
		s.Scan()
		input := s.Text()
		if input == "1" {
			fmt.Print("input > ")
			s.Scan()
			text := s.Text()
			morse, err := TextToMorse(text)
			if err != nil {
				fmt.Println(err)
			}
			ni, err := MorseToNI(morse)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("output >", ni)
		} else if input == "2" {
			fmt.Print("input >")
			s.Scan()
			ni := s.Text()
			morse, err := NIToMorse(ni)
			if err != nil {
				fmt.Println(err)
			}
			text, err := MorseToText(morse)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("output >", text)
		} else {
			fmt.Println("1か2を入力してください")
		}
	}
}

func TextToMorse(text string) (string, error) {
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

func MorseToText(morse string) (string, error) {
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
