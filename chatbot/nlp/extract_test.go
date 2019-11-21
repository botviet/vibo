package nlp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractNumber(t *testing.T) {

	var tests = []struct {
		input string
		want  []string
	}{
		{
			"9,316,420 回視聴",
			[]string{"9,316,420"},
		},
		{
			"897 lượt xem",
			[]string{"897"},
		},
		{
			"1,235,915 views",
			[]string{"1,235,915"},
		},
		{
			"1,235 và 915 views",
			[]string{"1,235", "915"},
		},
		{
			". , 1,2",
			[]string{"1,2"},
		},
		{
			"3 .2",
			[]string{"3"},
		},
		{
			"2 botviet.198 23a malong198@gmail.com",
			[]string{"2"},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, extractNumber(test.input))
	}
}

func TestExtractDate(t *testing.T) {

	var tests = []struct {
		input string
		want  []string
	}{
		{
			"2019/03/22 date",
			[]string{"2019/03/22"},
		},
		{
			"2019-03-22 日付",
			[]string{"2019-03-22"},
		},
		{
			"(Cập nhật: 17/07/2012 00:00)",
			[]string{"17/07/2012"},
		},
		{
			"(Cập nhật: 17-07-2012 00:00)",
			[]string{"17-07-2012"},
		},
		{
			"03-07-2012",
			[]string{"03-07-2012"},
		},
		{
			"06:00 02/04/2019",
			[]string{"02/04/2019"},
		},
		{
			"Mar 11th, 2015 4:32 PM",
			[]string{"Mar 11th, 2015"},
		},
		{
			"Thứ bảy, 13/7/2019, 15:52 (GMT+7)",
			[]string{"13/7/2019"},
		},
		{
			"Thứ bảy, 13/7/2019 và 15/8/2019",
			[]string{"13/7/2019", "15/8/2019"},
		},
		{
			"mai 13/7/2019 và 15-8-2019 , 09 8 2022",
			[]string{"13/7/2019", "15-8-2019", "09 8 2022"},
		},
		{
			"mai 13/7 và 15-8 , 09-8 nhé",
			[]string{"13/7", "15-8", "09-8"},
		},
		{
			"In 1991 and 1850 1500",
			[]string{"1991", "1850"},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, extractDate(test.input))
	}
}

func TestExtractTime(t *testing.T) {

	var tests = []struct {
		input string
		want  []string
	}{
		{
			"1h nhé",
			[]string{"1h"},
		},
		{
			"01 g ấy",
			[]string{"01 g"},
		},
		{
			"01 giờ á",
			[]string{"01 giờ"},
		},
		{
			"1h or 2 h",
			[]string{"1h", "2 h"},
		},
		{
			"24h nhé",
			[]string{"24h"},
		},
		{
			"23h59p cơ",
			[]string{"23h59p"},
		},
		{
			"lại 23h59' :(",
			[]string{"23h59'"},
		},
		{
			"là 23 h  59 ' và 23h 59p",
			[]string{"23 h 59", "23h 59p"},
		},
		{
			"23h59p50s 98@g",
			[]string{"23h59p"},
		},
		{
			"6:17PM (Việt Nam) 6:22AM (America) 1 AM hoặc 13PM",
			[]string{"6:17PM", "6:22AM", "1 AM", "13PM"},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, extractTime(test.input))
	}
}

func TestExtractURL(t *testing.T) {

	var tests = []struct {
		input string
		want  []string
	}{
		{
			"golang.com",
			[]string{"golang.com"},
		},
		{
			"go trang chủ là golang.com nhé",
			[]string{"golang.com"},
		},
		{
			"https://www.youtube.com/ và malong199@gmail.",
			[]string{"https://www.youtube.com/"},
		},
		{
			"xem phim tại youtube.com, đọc báo trên https://vnexpress.net/",
			[]string{"youtube.com", "https://vnexpress.net/"},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, extractURL(test.input))
	}
}

func TestExtractTemperature(t *testing.T) {

	var tests = []struct {
		input string
		want  []string
	}{
		{
			"13.25°C 3.25 ° f",
			[]string{"13.25°C", "3.25 ° f"},
		},
		{
			"Nhiệt độ: 13.25°C Nhiệt độ cao nhất: 15.2°C 15@°C",
			[]string{"13.25°C", "15.2°C"},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, extractTemperature(test.input))
	}
}

func TestExtractProbability(t *testing.T) {

	var tests = []struct {
		input string
		want  []string
	}{
		{
			"13.25% và 20,2 %",
			[]string{"13.25%", "20,2 %"},
		},
		{
			"Độ ẩm: 54 % - 55 ％ 15@%",
			[]string{"54 %", "55 ％"},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, extractProbability(test.input))
	}
}

func TestExtractEmail(t *testing.T) {

	var tests = []struct {
		input string
		want  []string
	}{
		{
			"bạn liên hệ tới botviet.asia@gmail.com nhé",
			[]string{"botviet.asia@gmail.com"},
		},
		{
			"email: botviet.asia@gmail.com hoặc botviet_asia@abc.com. s@ as@@.com",
			[]string{"botviet.asia@gmail.com", "botviet_asia@abc.com"},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, extractEmail(test.input))
	}
}

func TestExtractPhone(t *testing.T) {

	var tests = []struct {
		input string
		want  []string
	}{
		{
			`
			(021)1234567
			(123) 456 7899 sao nữa
			(123).456.7899 32222ss2d
			(123)-456-7899 e
			123-456-7899
			123 456 7899
			1234567899
			+6020863855
			as
			12345
			2 + 3
			123422.2+ 2
			`,
			[]string{"(021)1234567", "(123) 456 7899", "(123).456.7899", "(123)-456-7899", "123-456-7899", "123 456 7899", "1234567899", "+6020863855"},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, extractPhone(test.input))
	}
}

func TestExtractAnnotation(t *testing.T) {

	var tests = []struct {
		input string
		want  []string
	}{
		{
			"{{.person@123}} là anotation",
			[]string{"{{.person@123}}"},
		},
		{
			"{{.img@123}} và {{.buy@no}}",
			[]string{"{{.img@123}}", "{{.buy@no}}"},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, extractAnnotation(test.input))
	}
}
