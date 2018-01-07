package one_dimensional

import (
	"github.com/golang/glog"
	"fmt"
	//"errors"
	"math/rand"
	"errors"
)

//  0,  1,  2,  3,  4,  5,  6,  7
//128,129,130,131,132,133,134,135

func LifeStart() {
	//state := []byte{ 128,1,  130,  3,  132,  133,  6,  135}
	//state := []byte{0, 129, 130, 3, 132, 5, 6, 135}
	state := []byte{0, 129, 130, 131, 132, 133, 134, 7}
	stateMapping, err := buildStateMapping(state, 3, 0)
	if err != nil {
		fmt.Errorf("%s", err)
		return
	}
	const length int = 100
	space := make([]byte, length)
	aliveNum := int(rand.Float32()*40 + 10)
	initSpace(&space, aliveNum)
	for i := 0; i < length; i++ {
		printSpace(&space)
		nextTime(&space, 1, stateMapping)
	}
	glog.Info(space[0])
}

func initSpace(space *[]byte, tNum int) {
	var rate float32
	rate = float32(tNum) / float32(len(*space))
	s := *space
	count := 1
	spaceLen := len(s)
	for count <= tNum {
		for i := 0; i < spaceLen; i += count {
			if rand.Float32() <= rate && !(s[i] == 1) && count <= tNum {
				s[i] = 1
				count++
			}
		}
	}
}

func printSpace(space *[]byte) {
	s := *space
	length := len(*space)
	for i := 0; i < length; i++ {
		if s[i] == 1 {
			fmt.Print("O")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func getBeforeNum(current, totalLength, before int) int {
	beforeNum := (current + totalLength - before ) % totalLength
	return beforeNum
}

func getNextNum(current, totalLength, next int) int {
	nextNum := (current + next) % totalLength
	return nextNum
}

func status(index []byte, stateMapping *[][]byte) byte {
	length := len(*stateMapping)
	for i := 0; i < length; i++ {
		state := (*stateMapping)[i]
		j := 1
		for ; j <= len(index); j++ {
			if state[j] != index[j] {
				break
			}
		}
		if j == len(index) {
			return state[0]
		}
	}
	return 0
}

func nextTime(space *[]byte, selfIndex int, stateMapping *[][]byte) {
	stateLength := len(*stateMapping)
	length := len(*space)
	s := *space
	tmp := make([]byte, length)
	for i := 0; i < length; i++ {
		index := make([]byte, stateLength)
		for j := 0; j < stateLength; j++ {
			index[getNextNum(selfIndex, stateLength, j)] = s[getNextNum(i, length, 0)]
		}
		tmp[i] = status(index, stateMapping)
	}
	for i := 0; i < length; i++ {
		s[i] = tmp[i]
	}
}

var place = []byte{1, 2, 4, 8, 16, 32, 64, 128}

func buildStateMapping(state []byte, stateLength uint, defaultValue byte) (*[][]byte, error) {
	if stateLength > 7 {
		return nil, errors.New("state Length > 7")
	}
	length := 1 << stateLength

	scope := make([]int, length*2)
	scopeLength := length*2
	for i,j := 0,0; i < scopeLength; i,j = i+2, j+1 {
		scope[i] = j
		scope[i+1] = j + 128
	}

	stateMapping := make([][]byte, length)
	stateSet := make(map[byte]interface{})
	for _, v := range state {
		stateSet[v] = v
	}
	// 0 ~ 255
	for i,v := range scope {
		fmt.Println(i,v,length)
		if stateSet[byte(i)] == nil {
			stateMapping[i] = *oneState(byte(v), stateLength, defaultValue)
		} else {
			stateMapping[i] = *oneState(byte(v), stateLength, bytePlaceNum(byte(v), 7))
		}
	}

	return &stateMapping, nil
}

func oneState(num byte, stateLength uint, value byte) *[]byte {
	stateMapping := make([]byte, stateLength+1)

	stateMapping[0] = value
	for i := uint(1); i <= stateLength; i++ {
		stateMapping[i] = bytePlaceNum(num, i)
	}
	return &stateMapping
}

func bytePlaceNum(b byte, n uint) byte {
	if (b & place[n]) == place[n] {
		return 1
	}
	return 0
}
