package main

import (
	"fmt"
	"regexp"
)

func main() {
	myStr := `%^&**^&^*7*&%$#!@#$%^&*(){}?":992ncj sd2091'(1)`
	re := regexp.MustCompile(`[\s+.!/_,$%^*("')]+|[+—()?【】“”！，。？、~@#￥%…&*（）]+|{+|}+|:`)
	myStr = re.ReplaceAllString(myStr, "")
	fmt.Println(myStr)
	fmt.Println(10 * 3 / 4)
}
