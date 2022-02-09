package util

import (
	"airbattle/internal/logic/define"
	"strconv"
)

var (
	colorMap map[int32]string
	pointMap map[int32]string
)

func init() {
	colorMap = map[int32]string{
		define.COLOR_DIAMOND: "方块",
		define.COLOR_CLUB:    "草花",
		define.COLOR_HEART:   "红心",
		define.COLOR_SPADE:   "黑桃",
	}
	pointMap = map[int32]string{
		define.POINT_ACE:   "A",
		define.POINT_TWO:   "2",
		define.POINT_THREE: "3",
		define.POINT_FOUR:  "4",
		define.POINT_FIVE:  "5",
		define.POINT_SIX:   "6",
		define.POINT_SEVEN: "7",
		define.POINT_EIGHT: "8",
		define.POINT_NINE:  "9",
		define.POINT_TEN:   "10",
		define.POINT_J:     "J",
		define.POINT_Q:     "Q",
		define.POINT_K:     "K",
	}
}

func MakeCard(color int32, point int32) int32 {
	return (color << 4) | point
}

func GetColor(card int32) int32 {
	return card >> 4
}

func GetPoint(card int32) int32 {
	return card % 16
}

func GetPokerName(card int32) string {
	color := GetColor(card)
	point := GetPoint(card)
	colorName, colorExist := colorMap[color]
	pointName, pointExist := pointMap[point]
	if colorExist && pointExist {
		return colorName + pointName
	} else {
		return "未知" + strconv.Itoa(int(card))
	}
}
