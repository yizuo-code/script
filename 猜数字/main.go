package main

import (
	"fmt"
	"math/rand"
	"time"
)

func guessNum() {
	/*
		需求：
		猜数字 生成随机数字0-100
		从控制台数据 与生成数字比较
		大 提示太大了
		小 提示太小了
		等于 成功, 程序结束，提示是否重新进行游戏
		最多猜测五次，未猜对，说太笨了，程序结束，提示是否重新进行游戏
	*/
	for {
		dateNow := time.Now().Unix()
		rand.Seed(dateNow)
		gameNum := rand.Intn(100)
		fmt.Printf("游戏答案为：%d \n", gameNum)

		var randNumber int
		for nums := 1; nums <= 5; nums++ {
			fmt.Println("请输入一个数字: ")
			fmt.Scan(&randNumber)

			switch {
			case randNumber > gameNum:
				fmt.Println("太大了")
			case randNumber < gameNum:
				fmt.Println("太小了")
			default:
				fmt.Println("回答正确，游戏结束")
				nums = 0
			}

			if nums == 5 {
				fmt.Println("猜错5次，游戏结束,你太笨了")
			}
			if nums == 0 {
				break
			}
		}

		var gamecheck string
		fmt.Println("是否重新进行游戏(yes/no): ")
		fmt.Scan(&gamecheck)
		if gamecheck == "yes" {
			continue
		} else {
			break
		}

	}

}

func main() {
	guessNum()
}
