package common

import (
	"regexp"
	"strconv"
)

func GetNumberOfVersion(version string) []int64 {
	reg := regexp.MustCompile(`\d+`)
	strlist := reg.FindAllString(version, -1)

	numlist := make([]int64, 0)

	for _, strnum := range strlist {
		num, _ := strconv.ParseInt(strnum, 10, 64)
		numlist = append(numlist, num)
	}

	return numlist
}

//版本号比较 version 是否大于 version2
func IsVersionBigger(version1, version2 string) bool {

	//格式校验

	numlist1 := GetNumberOfVersion(version1)
	numlist2 := GetNumberOfVersion(version2)

	minlen := len(numlist1)
	if len(numlist2) < minlen {
		minlen = len(numlist2)
	}

	for i := 0; i < minlen; i++ {
		if numlist1[i] > numlist2[i] {
			return true
		} else if numlist1[i] < numlist2[i] {
			return false
		}
	}

	return false
}

//版本号比较 version 是否大于 version2
func IsVersionBiggerOrEq(version1, version2 string) bool {

	return !IsVersionBigger(version2, version1)
}
