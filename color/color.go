/**
  打印字符颜色
  Use:
  color.Green("SSH key %s created!")
*/
package color

import (
	"fmt"
)

const (
	RED     = uint8(iota + 91) //红色
	GREEN                      //绿色
	YELLOW                     //黄色
	BLUE                       //蓝色
	MAGENTA                    //洋红
	BLUE2                      //湖蓝
)

func Color(s interface{}, col uint8) string {
	str := fmt.Sprintf("\x1b[%dm%v\x1b[0m", col, s)
	fmt.Println(str)
	return str
}

func None(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

func Red(v interface{}) string {
	return Color(v, RED)
}

func Green(v ...interface{}) string {
	// return Color(strings.Join(v, ''))
	return Color(v, GREEN)
}

func Yellow(v interface{}) string {
	return Color(v, YELLOW)
}

func Blue(v interface{}) string {
	return Color(v, BLUE)
}

func Magenta(v interface{}) string {
	return Color(v, MAGENTA)
}

func Blue2(v interface{}) string {
	return Color(v, BLUE2)
}
