package main

import (
	"fmt"

	"github.com/botviet/vibo/chatbot/nlp"
	"github.com/botviet/vibo/utility"
)

func main() {

	// transform()
	// spellCorrect()
	similarity()

}

func similarity() {
	var dic nlp.Dictionary
	dic.Load("./storage/big-text")

	score, _ := dic.Similarity("tôi dắt chó đi dạo", "tôi đi dạo cùng chó")
	fmt.Println(score)
	// output: 0.883454233049763

	score, _ = dic.Similarity("tôi dắt chó đi dạo", "tôi đi dạo cùng mèo")
	fmt.Println(score)
	// output: 0.510749824206131

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
