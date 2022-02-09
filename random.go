package util

import (
	"math/rand"
	"reflect"
)

//轮盘概率接口
type IWheelItem interface {
	GetWeight() int
}

//RandomClosed get random num [min,max]
func RandomClosed(min int, max int) int {
	if min > max {
		min, max = max, min
	}
	return rand.Intn(max-min+1) + min
}

func RandClosed(min int32, max int32) int32 {
	if min > max {
		min, max = max, min
	}
	return rand.Int31n(max-min+1) + min
}

//RandomSemiClosed get random num [min,max)
func RandomSemiClosed(min int, max int) int {
	if min > max {
		min, max = max, min
	}
	if min == max {
		return min
	}
	return rand.Intn(max-min) + min
}

func RandSemiClosed(min int32, max int32) int32 {
	if min == max {
		return min
	}

	if min > max {
		min, max = max, min
	}
	return rand.Int31n(max-min) + min
}

//RandomFloat get random num [min,max]
func RandomFloatClosed(min float64, max float64) float64 {
	if min > max {
		min, max = max, min
	}
	//f [0.0,1.0]
	f := float64(rand.Int63()) / (1 << 63)
	return min + (max-min)*f
}

//RandCheck 随机测试, ratio / base 的概率返回true
func RandCheck(ratio int, base int) bool {
	if base <= 0 {
		return false
	}
	//[0,base)
	randNum := rand.Intn(base)
	return ratio > randNum
}

func RandOk(ratio int32, base int32) bool {
	if base <= 0 {
		return false
	}
	//[0,base)
	return ratio > rand.Int31n(base)
}

//GetRandomElement 使用轮盘算法随机从容器中获取一个元素 元素必须实现IWheelItem接口
func GetRandomElement(container interface{}) interface{} {
	if !AssertContainer(container) {
		return nil
	}
	cValue := reflect.ValueOf(container)
	cType := cValue.Type()
	if cValue.Len() == 0 {
		return nil
	}
	isIWheelItem := func(t reflect.Type) bool {
		var ok bool
		if t.Kind() == reflect.Ptr {
			_, ok = reflect.New(t).Elem().Interface().(IWheelItem)
		} else {
			_, ok = reflect.New(t).Interface().(IWheelItem)
		}
		return ok
	}
	ok := isIWheelItem(cType.Elem())
	if !ok {
		return nil
	}
	total := 0
	ForEach(container, func(k interface{}, v interface{}) {
		total += v.(IWheelItem).GetWeight()
	})
	randNum := RandomClosed(0, total-1)
	tmpNum := 0
	var ret interface{} = nil
	ForEach(container, func(k interface{}, v interface{}) {
		tmpNum += v.(IWheelItem).GetWeight()
		if tmpNum > randNum && ret == nil {
			ret = v
		}
	})
	return ret
}

//M个里取N个
func GetNFromM(container interface{}, num int) []interface{} {
	ret := []interface{}{}
	if !AssertContainer(container) {
		return ret
	}
	cValue := reflect.ValueOf(container)
	if cValue.Len() == 0 {
		return ret
	}
	index := 0
	//蓄水池
	ForEach(container, func(k interface{}, v interface{}) {
		if index < num {
			ret = append(ret, v)
		} else {
			if RandCheck(num, index+1) {
				swapIndex := RandomClosed(0, num-1)
				ret[swapIndex] = v
			}
		}
		index++
	})
	return ret
}

//Shuffle 洗牌
func Shuffle(array interface{}) {
	if !AssertArray(array) {
		return
	}
	cValue := reflect.ValueOf(array)
	len := cValue.Len()
	for i := 0; i < len-1; i++ {
		randNum := RandomClosed(i, len-1)
		if randNum != i {
			tmp := cValue.Index(i).Interface()
			cValue.Index(i).Set(reflect.ValueOf(cValue.Index(randNum).Interface()))
			cValue.Index(randNum).Set(reflect.ValueOf(tmp))
		}
	}
}
