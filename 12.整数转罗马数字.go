/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */

// @lc code=start
func intToRoman(num int) string {
	value := []struct {
		num   int
		roman string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := []byte{}
	for _, v := range value {
		for num >= v.num {
			num -= v.num
			roman = append(roman, v.roman...)
		}
		if num == 0 {
			break
		}
	}
	return string(roman)
}

// @lc code=end

