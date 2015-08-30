package common

import (
	"fmt"
	"strconv"
)

func MapValueToString(m map[string]string, k string) string {

	v, found := m[k]
	if !found {
		panic(fmt.Errorf("MapValueToString failed! key=%s not found", k))
	}

	return v
}

func MapValueToBool(m map[string]string, k string) bool {

	v, found := m[k]
	if !found {
		panic(fmt.Errorf("MapValueToInt8 failed! key=%s not found", k))
	}

	r, err := strconv.ParseBool(v)
	if err != nil {
		panic(err)
	}

	return r
}

func MapValueToInt(m map[string]string, k string) int {

	v, found := m[k]
	if !found {
		panic(fmt.Errorf("MapValueToInt8 failed! key=%s not found", k))
	}

	r, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		panic(err)
	}

	return int(r)
}

func MapValueToInt8(m map[string]string, k string) int8 {

	v, found := m[k]
	if !found {
		panic(fmt.Errorf("MapValueToInt8 failed! key=%s not found", k))
	}

	r, err := strconv.ParseInt(v, 10, 8)
	if err != nil {
		panic(err)
	}

	return int8(r)
}

func MapValueToInt16(m map[string]string, k string) int16 {
	v, found := m[k]
	if !found {
		panic(fmt.Errorf("MapValueToInt16 failed! key=%s not found", k))
	}

	r, err := strconv.ParseInt(v, 10, 16)
	if err != nil {
		panic(err)
	}
	return int16(r)
}

func MapValueToInt32(m map[string]string, k string) int32 {
	v, found := m[k]
	if !found {
		panic(fmt.Errorf("MapValueToInt32 failed! key=%s not found", k))
	}

	r, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		panic(err)
	}

	return int32(r)
}

func MapValueToInt64(m map[string]string, k string) int64 {
	v, found := m[k]
	if !found {
		panic(fmt.Errorf("MapValueToInt64 failed! key=%s not found", k))
	}

	r, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		panic(err)
	}

	return int64(r)
}

func MapValueToUint(m map[string]string, k string) uint {

	v, found := m[k]
	if !found {
		panic(fmt.Errorf("MapValueToInt8 failed! key=%s not found", k))
	}

	r, err := strconv.ParseUint(v, 10, 32)
	if err != nil {
		panic(err)
	}

	return uint(r)
}

func MapValueToUint8(m map[string]string, k string) uint8 {
	v, found := m[k]
	if !found {
		panic(fmt.Errorf("MapValueToUint8 failed! key=%s not found", k))
	}

	r, err := strconv.ParseUint(v, 10, 8)
	if err != nil {
		panic(err)
	}
	return uint8(r)
}

func MapValueToUint16(m map[string]string, k string) uint16 {
	v, found := m[k]
	if !found {
		panic(fmt.Errorf("MapValueToUint16 failed! key=%s not found", k))
	}

	r, err := strconv.ParseUint(v, 10, 16)
	if err != nil {
		panic(err)
	}
	return uint16(r)
}

func MapValueToUint32(m map[string]string, k string) uint32 {
	v, found := m[k]
	if !found {
		panic(fmt.Errorf("MapValueToUint32 failed! key=%s not found", k))
	}

	r, err := strconv.ParseUint(v, 10, 32)
	if err != nil {
		panic(err)
	}

	return uint32(r)
}

func MapValueToUint64(m map[string]string, k string) uint64 {
	v, found := m[k]
	if !found {
		panic(fmt.Errorf("MapValueToUint64 failed! key=%s not found", k))
	}

	r, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		panic(err)
	}

	return uint64(r)
}

func MapValueToFloat32(m map[string]string, k string) float32 {
	v, found := m[k]
	if !found {
		panic(fmt.Errorf("MapValueToFloat32 failed! key=%s not found", k))
	}

	r, err := strconv.ParseFloat(v, 32)
	if err != nil {
		panic(err)
	}
	return float32(r)
}

func MapValueToFloat64(m map[string]string, k string) float64 {
	v, found := m[k]
	if !found {
		panic(fmt.Errorf("MapValueToFloat64 failed! key=%s not found", k))
	}

	r, err := strconv.ParseFloat(v, 64)
	if err != nil {
		panic(err)
	}
	return float64(r)
}
