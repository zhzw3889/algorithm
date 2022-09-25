package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//生成count个[start,end)结束的不重复的随机数
func generateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}
	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start
		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}

func OpenFile(filename string) (*os.File, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("文件不存在")
		return os.Create(filename) //创建文件
	}
	fmt.Println("文件存在")
	return os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
}

//生成若干个不重复的随机数
func main() {
	// 生成随机数slice
	// now := time.Now()
	// nums := generateRandomNumber(1, 100001, 100000)
	// since := time.Since(now)
	// fmt.Println(nums)
	// fmt.Println(since)
	// 将随机写入文件

	// var buf bytes.Buffer
	// for i, v := range nums {
	// 	if i > 0 {
	// 		buf.WriteString(" ")
	// 	}
	// 	fmt.Fprintf(&buf, "%d", v)
	// }
	// fmt.Println(buf.String())

	// f1, err1 := OpenFile("./random.txt")
	// if err1 != nil {
	// 	log.Fatal(err1.Error())
	// }
	// defer f1.Close()
	// n, err1 := io.WriteString(f1, buf.String()) //写入文件(字符串)
	// if err1 != nil {
	// 	log.Fatal(err1.Error())
	// }
	// fmt.Printf("写入 %d 个字节\n", n)

	// 读取文件时间消耗
	now := time.Now()
	data, err := ioutil.ReadFile("./random.txt")
	if err != nil {
		fmt.Println("read file err:", err.Error())
		return
	}

	// 打印文件内容
	res := strings.Split(string(data), " ")
	fmt.Println(res)
	since := time.Since(now)
	fmt.Println(since)
}
