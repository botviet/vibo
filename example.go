package main

import (
	"fmt"

	"github.com/botviet/vibo/chatbot/nlp"
	"github.com/botviet/vibo/utility"
)

func main() {

	// transform()
	spellCorrect()

}

func spellCorrect() {
	var dic nlp.Dictionary
	dic.Load("./storage/big-text")

	fmt.Println(dic.Correction("lam", []string{"gì"}, []string{"đang"}))
	fmt.Println(dic.Correction("an", []string{"cơm"}, []string{"đi"}))
	fmt.Println(dic.Correction("oi", []string{"bạn"}, []string{}))
	fmt.Println(dic.Correction("1nghienga", []string{"ngả"}, []string{"đi"}))
	fmt.Println(dic.Correction("chan", []string{"quá", "đi"}, []string{"haizz"}))
	fmt.Println(dic.Correction("di", []string{"chơi", "không"}, []string{"ê"}))
	/*
		output:
			làm
			ăn
			ơi
			nghiêng
			chán
			đi
	*/
}

func transform() {
	fmt.Println(utility.Transform("nghiêng", utility.UNICODE, utility.VIQR))
	fmt.Println(utility.Transform("ngu+o+`i", utility.VIQR, utility.UNICODE))
}
