# vibo

Project Chatbot cho tiếng việt

## Install

```cmd
go get github.com/botviet/vibo
```

## Usage

### Word tokenize

```go
import (
    "fmt"

    "github.com/botviet/vibo/chatbot/nlp"
)

func main() {
    var dic nlp.Dictionary
    dic.DefaultLoad()

    words, cant := dic.WordTokenize([]string{"ngày mai Hà nội có mưa không"})

    for typ, w := range words {
        fmt.Println(typ, w)
    }

    if len(cant) > 0 {

        fmt.Println("Can't Readable")
        for _, w := range cant {
            fmt.Println(w)
        }
    }

    /*
        output:
            ngày mai [date]
            Hà nội [location_vietnam]
            có mưa [weather]
            không [word_common]
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
    dic.DefaultLoad()

    score, _ := dic.Similarity("mai mưa không", "ngày mai trời mưa không nhỉ")
    fmt.Println(score)
    // output: 0.640367228946234

    score, _ = dic.Similarity("tôi hàng ngày đều dắt chó đi dạo", "đi dạo với chó thú lắm")
    fmt.Println(score)
    // output: 0.5133125853160265
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
    dic.DefaultLoad()

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
