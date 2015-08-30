package common

import (
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

//获取随机数
func GetRandom(Number int) int {
	return rand.Intn(Number)
}

//获取随机数，指定区间
func GetRandByRange(min int, max int) int {
	return min + rand.Intn(max-min)
}

//获取一定比例的随机数
func GetRandByPercent(MaxNum int, percent float64) []int {
	mVal := make(map[int]bool)

	//有效性范围限制
	if percent < 0.0 {
		percent = 0.0
	} else if percent > 1.0 {
		percent = 1.0
	}

	count := int(float64(MaxNum) * percent)

	for len(mVal) < count {
		key := rand.Intn(MaxNum)
		_, isOK := mVal[key]
		if !isOK {
			mVal[key] = true
		}
	}

	list := make([]int, 0)
	for k, _ := range mVal {
		list = append(list, k)
	}

	return list
}

func GetRandByUnduplicated(min int, max int, randed map[string]interface{}) int {
	if randed == nil || len(randed) == 0 {
		return GetRandByRange(min, max)
	}

	var r []int
	for i := min; i < max; i++ {
		str := strconv.Itoa(i)
		if _, ok := randed[str]; !ok {
			r = append(r, i)
		}
	}

	index := GetRandom(len(r))
	return r[index]
}
