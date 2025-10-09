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

func searchPathNoRepeat(
	phone *Phone, vis Visit, phoneToCall []rune,
	pos, x, y int,
	moveCnt, moveSm int, ans chan [2]int,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
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
		if phone.checkFiled(x+move.x, y+move.y) &&
			phone.getNum(x+move.x, y+move.y) == phoneNum &&
			!vis.get(x+move.x, y+move.y) {
			visCopy := vis.copy()
			wg.Add(1)
			go searchPathNoRepeat(
				phone, visCopy, phoneToCall,
				pos+1, x+move.x, y+move.y,
				moveCnt+1, moveSm+phoneNum, ans,
				wg,
			)
		}
	}
}

func startSearchingNoRepeat(phone *Phone, phoneToCall []rune) (int, int) {
	wg := &sync.WaitGroup{}
	ans := make(chan [2]int)
	n, m := phone.getSize()
	phoneNum := int(phoneToCall[0] - rune('0'))
	vis := Visit{}
	vis.create(phone.n, phone.m)
	for idx := range n {
		for jdx := range m {
			if phone.getNum(idx, jdx) == phoneNum {
				visCopy := vis.copy()
				wg.Add(1)
				go searchPathNoRepeat(
					phone, visCopy, phoneToCall,
					1, idx, jdx,
					0, phoneNum, ans,
					wg,
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

func TaskNoRepeat() (int, int) {
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
	minCnt, minSm := startSearchingNoRepeat(&phone, phoneToCall)

	return minCnt, minSm
}

func TaskNoRepeatManual(phone Phone, phoneToCall string) (int, int) {

	minCnt, minSm := startSearchingNoRepeat(&phone, []rune(phoneToCall))

	return minCnt, minSm
}
