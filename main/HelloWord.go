package main

import (
	"fmt"
)

type circle struct {
	radius float64
}

func nums() (int, int, string) {
	a, b, c := 1, 2, "1"
	return a, b, c
}

func add(base int) func(int) *int {

	fmt.Printf("%p\n", &base) //打印变量地址，可以看出来 内部函数时对外部传入参数的引用

	f := func(i int) *int {
		fmt.Printf("%p\n", &base)
		base += i
		return &base
	}

	return f
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func reDefind(nums *[3]int) {
	nums[1] = 1
}

func deleteNum(nums []int, i int) []int {
	copy(nums[i:], nums[i+1:])
	return nums[:len(nums)-1]
}
func main() {

	q := [...]int{1, 2, 3}
	reDefind(&q)
	//fmt.Println(q)
	//branchtest
	//fmt.Printf(fmt.Sprintf("%d", 123))
	//int string转换
	//strconv.FormatInt(123, 10)
	//strconv.Itoa(1)
	//y, err := strconv.ParseInt("123", 10, 64)
	// 类型断言
	/*p := new(int)
	var num interface{}
	num = p
	_, ok := num.(*int)
	fmt.Println(ok)*/
	/*input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
	}*/
	//fmt.Println(strings.Join(os.Args[0:], ""))
	/*arrList := [10]string{"1", "2", "3"}
	arrList1 := make([]int, 5, 10)
	fmt.Println(len(arrList1), cap(arrList))
	for i := range arrList1 {
		fmt.Println(arrList1[i])
	}*/
	//t1 := add(10)
	//fmt.Println(*t1(1), *t1(2))
	/*for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}*/
	/*fmt.Println("Hello Golang!")
	s1 := "point"
	strings := []string{"a", "b"}
	for i, s := range strings {
		fmt.Printf("%d %s\n", i, s)
	}

	var ptr = &s1
	fmt.Println(ptr, &s1, *ptr)

	testMap := map[string] int {}
	testMap["a"] = 1
	testMap["b"] = 2
	for k := range testMap {
		fmt.Println(k, testMap[k])
	}

	value, ok := testMap["a"]
	if ok {
		fmt.Println("存在", value)
	} else {
		fmt.Println("不存在")
	}

	fmt.Println(int32(testMap["a"]))*/

}
