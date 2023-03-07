package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"html"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"
)

var rander *rand.Rand

func init() {
	seed := os.Getpid()
	fmt.Printf("{utilsInit seed}|%d", seed)
	rander = rand.New(rand.NewSource(int64(seed)))
}

/*
//时间格式化模板
const TimeFormart string = "2006-01-02 15:04:05"

//格式化时间
var formatMap = map[string]string{"YY": "06", "YYYY": "2006", "MM": "01", "DD": "02", "HH": "15", "mm": "04", "ss": "05"}

func GetFormatTimeStr(t time.Time, format string) string {
	realFormat := TimeFormart
	if len(format) > 0 {
		realFormat = format
		for k, v := range formatMap {
			realFormat = strings.ReplaceAll(realFormat, k, v)
		}
	}
	return t.Format(realFormat)
}
*/

//获取简单的uid 使用时间戳
func GetSimpleRandId() string {
	return time.Now().Format("060102150405") + strconv.Itoa(rander.Intn(90)+10)
}

// 数组去重，快慢指针
func RemoveDuplicates(str []string) (strWithoutDup []string) {
	strMap := make(map[string]string)
	for _, v := range str {
		strMap[v] = v
	}
	for _, v := range strMap {
		strWithoutDup = append(strWithoutDup, v)
	}
	return
}

// 字符串转md5
func GetMd5FromStr(str *string) string {
	buffer := []byte(*str)
	return GetMd5FromBytes(&buffer)
}

//bytes转md5
func GetMd5FromBytes(buffer *[]byte) string {
	h := md5.New()
	h.Write(*buffer)
	return hex.EncodeToString(h.Sum(nil))
}

//还原被转义的Hmtl中的< > &
func DecodeHtml(content *string) {
	*content = html.UnescapeString(*content)
}

//Html转纯文本(使用正则表达式)
func Html2text(html string) string {
	text := html

	//删除\r\t
	text = regexp.MustCompile(`[\r|\t|\n]`).ReplaceAllString(text, "")
	//样式以及脚本标签
	text = regexp.MustCompile(`<(style|script)[^<]*</(style|script)>`).ReplaceAllString(text, "")
	//删除html标签
	text = regexp.MustCompile(`<\s+[^<]+>`).ReplaceAllString(text, "")
	//清除一些不需要换行的标签（内联标签）
	text = regexp.MustCompile(`<(a[^a-zA-Z]|b[^r]|em|strong|span)[^<]*>`).ReplaceAllString(text, " ")
	text = regexp.MustCompile(`</(a|b[^r]|em|strong|span)[^<]*>`).ReplaceAllString(text, " ")

	//将所有的剩下的非内敛标签换成p标签
	text = regexp.MustCompile(`<\s*[^/][^<]*>`).ReplaceAllString(text, "<p>")
	text = regexp.MustCompile(`<\s*[/][^<]*>`).ReplaceAllString(text, "</p>")

	//将p标签换成\n以保持换行
	text = regexp.MustCompile(`(<p>|</p>)`).ReplaceAllString(text, "\n")
	text = regexp.MustCompile(`(\n\s*)+`).ReplaceAllString(text, "\n")
	text = regexp.MustCompile(`^\n`).ReplaceAllString(text, "")
	text = regexp.MustCompile(` +`).ReplaceAllString(text, "")

	//删除unicode字符
	//todo 这里有问题
	text = regexp.MustCompile(`u[0-9a-f]{4}`).ReplaceAllString(text, "")

	return text
}
