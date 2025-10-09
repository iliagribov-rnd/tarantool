package task

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

func searchPathRepeat(
	phone *Phone, vis Visit, phoneToCall []rune,
	pos int, x, y int,
	moveCnt, moveSm int, ans chan [2]int,
	uniquePlaces [4][2]int, curPlaceIdx int, wg *sync.WaitGroup,
) {
	defer wg.Done()
	places := make(map[[2]int]bool)
	fullPlaces := true
	for _, place := range uniquePlaces {
		if place[0] == -1 && place[1] == -1 {
			fullPlaces = false
			break
		}
		places[place] = true
	}
	if len(places) <= 2 && fullPlaces {
		return
	}
	vis.set(x, y)
	if pos >= len(phoneToCall) {
		ans <- [2]int{moveCnt, moveSm * moveCnt}
		return
	}
	possibleMoves := []Move{
		Move{x: -2, y: -1},
		Move{x: -2, y: 1},
		Move{x: 2, y: -1},
		Move{x: 2, y: 1},
		Move{x: 1, y: -2},
		Move{x: 1, y: 2},
		Move{x: -1, y: -2},
		Move{x: -1, y: 2},
	}
	phoneNum := int(phoneToCall[pos] - rune('0'))
	for _, move := range possibleMoves {
		if phone.checkFiled(x+move.x, y+move.y) {
			if vis.get(x+move.x, y+move.y) {
				visCopy := vis.copy()
				uniquePlaces[curPlaceIdx] = [2]int{x + move.x, y + move.y}
				wg.Add(1)
				go searchPathRepeat(
					phone, visCopy, phoneToCall,
					pos, x+move.x, y+move.y,
					moveCnt+1, moveSm, ans,
					uniquePlaces, (curPlaceIdx+1)%4, wg,
				)
			}
			if phone.getNum(x+move.x, y+move.y) == phoneNum && !vis.get(x+move.x, y+move.y) {
				visCopy := vis.copy()
				uniquePlaces[curPlaceIdx] = [2]int{x + move.x, y + move.y}
				wg.Add(1)
				go searchPathRepeat(
					phone, visCopy, phoneToCall,
					pos+1, x+move.x, y+move.y,
					moveCnt+1, moveSm+phoneNum, ans,
					uniquePlaces, (curPlaceIdx+1)%4, wg,
				)
			}
		}
	}
}

func startSearchingRepeat(phone *Phone, phoneToCall []rune) (int, int) {
	wg := &sync.WaitGroup{}
	ans := make(chan [2]int)
	n, m := phone.getSize()
	phoneNum := int(phoneToCall[0] - rune('0'))
	vis := Visit{}
	vis.create(phone.n, phone.m)
	uniquePlaces := [4][2]int{
		[2]int{-1, -1},
		[2]int{-1, -1},
		[2]int{-1, -1},
		[2]int{-1, -1},
	}
	curPlaceIdx := 0
	for idx := range n {
		for jdx := range m {
			if phone.getNum(idx, jdx) == phoneNum {
				visCopy := vis.copy()
				wg.Add(1)
				uniquePlaces[curPlaceIdx] = [2]int{idx, jdx}
				go searchPathRepeat(
					phone, visCopy, phoneToCall,
					1, idx, jdx,
					0, phoneNum, ans,
					uniquePlaces, (curPlaceIdx+1)%4, wg,
				)
			}
		}
	}

	go func(wg *sync.WaitGroup) {
		wg.Wait()
		close(ans)
	}(wg)

	pathFound := false
	minCnt, minSm := math.MaxInt, math.MaxInt
	for out := range ans {
		if out[0] < minCnt {
			minCnt, minSm = out[0], out[1]
			pathFound = true
		}
	}
	if pathFound {
		return minCnt, minSm
	}
	return -1, -1
}

func TaskRepeat() (int, int) {
	// input phone size
	var line string
	var nums []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	line = scanner.Text()
	nums = strings.Split(line, " ")
	if len(nums) < 2 {
		log.Print("Bad matrix size!")
		os.Exit(1)
	}
	N, err := strconv.Atoi(nums[0])
	if err != nil {
		log.Print("Bad num rows parameter!")
		os.Exit(1)
	}
	M, err := strconv.Atoi(nums[1])
	if err != nil {
		log.Print("Bad num cols parameter!")
		os.Exit(1)
	}

	// create phone
	phone := Phone{}
	phone.create(N, M)

	// fill gaps in phone
	for idx := range N {
		scanner.Scan()
		line = scanner.Text()
		nums = strings.Split(line, " ")
		if len(nums) < M {
			fmt.Println("Missing matrix values!")
		}
		for jdx := range M {
			val, err := strconv.Atoi(nums[jdx])
			if err != nil {
				log.Print("Bad matrix value!")
				os.Exit(1)
			}
			phone.setNum(val, idx, jdx)
		}
	}

	// get phone to call
	scanner.Scan()
	line = scanner.Text()
	nums = strings.Split(line, " ")
	// slen, _ := strconv.Atoi(nums[0])
	phoneToCall := []rune(nums[1])

	// start
	minCnt, minSm := startSearchingRepeat(&phone, phoneToCall)

	return minCnt, minSm
}

func TaskRepeatManual(phone Phone, phoneToCall string) (int, int) {

	minCnt, minSm := startSearchingRepeat(&phone, []rune(phoneToCall))

	return minCnt, minSm
}
