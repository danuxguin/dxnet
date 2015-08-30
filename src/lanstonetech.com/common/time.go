package common

import (
	"time"
)

//年月日转成时间戳
func Timestamp(y, m, d, hour, min, sec int) int64 {
	t := time.Date(y, time.Month(m), d, hour, min, sec, 0, time.UTC)
	//t, _ := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", y, m, d, hour, min, sec))
	return t.Unix()
}

//两个时间戳判断是不是同一天
func IsSameDay(time1, time2 int64) bool {

	t1 := time.Unix(time1, 0)
	t2 := time.Unix(time2, 0)

	if t1.Year() != t2.Year() {
		return false
	}

	if t1.YearDay() == t2.YearDay() {
		return true
	}

	return false
}

//判断两个时间戳相差几天
func DiffDay(time1, time2 int64) (int, error) {
	ft1 := time.Unix(time1, 0).Format("2006-01-02")
	ft2 := time.Unix(time2, 0).Format("2006-01-02")
	t1, err := time.Parse("2006-01-02", ft1)
	if err != nil {
		return 0, err
	}
	t2, err2 := time.Parse("2006-01-02", ft2)
	if err2 != nil {
		return 0, err2
	}

	return Abs(int(t1.Sub(t2).Seconds() / 86400)), nil
}

func GetYear() int {
	now := time.Now().Unix()
	t := time.Unix(now, 0)
	return t.Year()
}

func GetYearMonthDay() (int, int, int) {
	now := time.Now().Unix()
	t := time.Unix(now, 0)
	return int(t.Year()), int(t.Month()), int(t.Day())
}

//返回格式如下 2015001 2015表示年份，后面的001表示是这一年的第几天
func GetYearDay() int {
	t := time.Now()
	ret := t.Year()*1000 + t.YearDay()
	return ret
}

func GetTimeFromStrDate(date string) (int, int, int) {
	d, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0, 0, 0
	}

	year := d.Year()
	month := int(d.Month())
	day := d.Day()
	return year, month, day
}

func GetZodiac(year int) (zodiac string) {
	if year <= 0 {
		zodiac = "-1"
	}
	start := 1901
	x := (start - year) % 12
	if x == 1 || x == -11 {
		zodiac = "鼠"
	}
	if x == 0 {
		zodiac = "牛"
	}
	if x == 11 || x == -1 {
		zodiac = "虎"
	}
	if x == 10 || x == -2 {
		zodiac = "兔"
	}
	if x == 9 || x == -3 {
		zodiac = "龙"
	}
	if x == 8 || x == -4 {
		zodiac = "蛇"
	}
	if x == 7 || x == -5 {
		zodiac = "马"
	}
	if x == 6 || x == -6 {
		zodiac = "羊"
	}
	if x == 5 || x == -7 {
		zodiac = "猴"
	}
	if x == 4 || x == -8 {
		zodiac = "鸡"
	}
	if x == 3 || x == -9 {
		zodiac = "狗"
	}
	if x == 2 || x == -10 {
		zodiac = "猪"
	}
	return
}

func GetAge(year int) (age int) {
	if year <= 0 {
		age = -1
	}
	nowyear := time.Now().Year()
	age = nowyear - year
	return
}

func GetConstellation(month, day int) (star string) {
	if month <= 0 || month >= 13 {
		star = "-1"
	}
	if day <= 0 || day >= 32 {
		star = "-1"
	}
	if (month == 1 && day >= 20) || (month == 2 && day <= 18) {
		star = "水瓶座"
	}
	if (month == 2 && day >= 19) || (month == 3 && day <= 20) {
		star = "双鱼座"
	}
	if (month == 3 && day >= 21) || (month == 4 && day <= 19) {
		star = "白羊座"
	}
	if (month == 4 && day >= 20) || (month == 5 && day <= 20) {
		star = "金牛座"
	}
	if (month == 5 && day >= 21) || (month == 6 && day <= 21) {
		star = "双子座"
	}
	if (month == 6 && day >= 22) || (month == 7 && day <= 22) {
		star = "巨蟹座"
	}
	if (month == 7 && day >= 23) || (month == 8 && day <= 22) {
		star = "狮子座"
	}
	if (month == 8 && day >= 23) || (month == 9 && day <= 22) {
		star = "处女座"
	}
	if (month == 9 && day >= 23) || (month == 10 && day <= 22) {
		star = "天秤座"
	}
	if (month == 10 && day >= 23) || (month == 11 && day <= 21) {
		star = "天蝎座"
	}
	if (month == 11 && day >= 22) || (month == 12 && day <= 21) {
		star = "射手座"
	}
	if (month == 12 && day >= 22) || (month == 1 && day <= 19) {
		star = "摩羯座"
	}

	return star
}
