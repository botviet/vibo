# vibo

Project Chatbot cho tiếng việt

## Install

```cmd
go get github.com/botviet/vibo
```

## Usage

### Transform between UNICODE and VIQR

```go
import (
    "fmt"

    "github.com/botviet/vibo/utility"
)

func main() {

    fmt.Println(utility.Transform("nghiêng", utility.UNICODE, utility.VIQR))
    // → nghie^ng
    fmt.Println(utility.Transform("ngu+o+`i", utility.VIQR, utility.UNICODE))
    // → người
}
```

### Spelling correction

```go
import (
    "fmt"

    "github.com/botviet/vibo/chatbot/nlp"
)

func main() {
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
```

### Text similarity

```go
import (
    "fmt"

    "github.com/botviet/vibo/chatbot/nlp"
)

func main() {
    var dic nlp.Dictionary
    dic.Load("./storage/big-text")

    score, _ := dic.Similarity("tôi dắt chó đi dạo", "tôi đi dạo cùng chó")
    fmt.Println(score)
    // output: 0.883454233049763

    score, _ = dic.Similarity("tôi dắt chó đi dạo", "tôi đi dạo cùng mèo")
    fmt.Println(score)
    // output: 0.510749824206131
}
```
