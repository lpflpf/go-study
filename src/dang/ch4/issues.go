// issues
package main

import (
	"dang/ch4/github"
	"fmt"
	"log"
	"os"
	"time"
)

type myTime int

// 使用枚举定义几个时间段
const (
	OneMonth myTime = iota
	OneYear
	AfterOneYear
)

func main() {
	fmt.Println("Hello World!")
	result, err := github.SearchIssues(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

	output := make(map[myTime]([](*github.Issue)))

	for _, item := range result.Items {
		switch {
		case time.Now().AddDate(-1, 0, 0).After(item.CreateAt):
			output[AfterOneYear] = append(output[AfterOneYear], item)
		case time.Now().AddDate(0, -1, 0).After(item.CreateAt):
			output[OneYear] = append(output[OneYear], item)
		default:
			output[OneMonth] = append(output[OneMonth], item)
		}
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Printf("After One Year:\n")
	for _, item := range output[AfterOneYear] {
		fmt.Printf("#%-5d %s %9.9s %.55s\n",
			item.Number, item.CreateAt, item.User.Login, item.Title)
	}

	fmt.Printf("\nAfter One Month:\n")
	for _, item := range output[OneYear] {
		fmt.Printf("#%-5d %s %9.9s %.55s %s\n",
			item.Number, item.CreateAt, item.User.Login, item.Title)
	}

}
