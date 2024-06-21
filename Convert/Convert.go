package convert

import "strconv"

// PrsToInt تبدیل استرینگ به اعداد و برعکس
func PrsToInt(str string) int {
	test,_ :=strconv.Atoi(str)
	return int(test)
}
// PrsToInt تبدیل عدد به استرینگ
func ToString(number int) string {
	res :=strconv.Itoa(number)
	return res
}