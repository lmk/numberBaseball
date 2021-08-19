package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Result struct {
	strikes int
	balls   int
}

type BaseballNumbers struct {
	nums [3]int
}

func (n *BaseballNumbers) Make() {
	rand.Seed(int64(time.Now().UnixMicro()))

	made := 0

	for i := range n.nums {

		n.nums[i] = rand.Intn(9) + 1

		for j := 0; j < made; j++ {

			if n.nums[i] == n.nums[j] {
				n.nums[i] = rand.Intn(9) + 1
				j = -1
				continue
			}
		}
		made++
	}
}

func (n *BaseballNumbers) KeyInput() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("\nEnter 3 numbers separated by spaces: ")
	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, " \r\n")

	nums := strings.Split(text, " ")

	for i := range n.nums {
		var error error
		n.nums[i], error = strconv.Atoi(nums[i])
		if error != nil {
			fmt.Println(error)
			continue
		}
	}
}

func (left *BaseballNumbers) Compare(right BaseballNumbers) Result {
	var result Result

	for i := range left.nums {
		for j := range right.nums {
			if left.nums[i] == right.nums[j] {
				if i == j {
					result.strikes++
				} else {
					result.balls++
				}
				break
			}
		}
	}

	return result
}

func main() {
	// 생성
	var my BaseballNumbers
	my.Make()

	// 시도 회수는 10번
	for i := 0; i < 10; i++ {

		var enemy BaseballNumbers

		// 사용자 입력
		enemy.KeyInput()

		// 비교
		result := my.Compare(enemy)

		if result.strikes == 3 {
			fmt.Println("Bingo")
			return
		} else {
			fmt.Printf("%d S %d B\n", result.strikes, result.balls)
		}

	}

	fmt.Println("Out!")

}
