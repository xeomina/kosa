// 5. 함수
// 함수 선언과 호출

// 이름이 지정된 파라미터(named)와 선택적 파라미터 대응
// - Go는 키워드 파라미터(=named)를 지원 x
package main

import (
	"fmt"
)

// func div(numerator int, denominator int) int {
// 	if denominator == 0 {
// 		return 0
// 	}
// 	return numerator / denominator
// }

// func main() {
// 	result := div(5, 2)
// 	fmt.Println(result)
// }

/*
type MyfuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func Myfunc(opts MyfuncOpts) error {
	if opts.Age >= 20 {
		fmt.Println(opts.LastName, "은(는) 성인입니다.")
	}
	return nil
}
*/

// func main() {
// 	Myfunc(MyfuncOpts{
// 		LastName: "Patel",
// 		Age:      50,
// 	})
// 	Myfunc(MyfuncOpts{
// 		FirstName: "Joe",
// 		LastName:  "Smith",
// 	})
// }

// 가변 입력 파라미터와 슬라이스
/*
func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}
*/

// func main() {
// 	fmt.Println(addTo(3))
// 	fmt.Println(addTo(3, 2))
// 	fmt.Println(addTo(3, 2, 4, 6, 8))
// 	a := []int{4, 3}
// 	fmt.Println(addTo(3, a...))
// 	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))
// }

// 다중 반환값 - 함수 정의 시 반환값의 타입을 ,로 구분하고 ()로 묶어줌
/*
func divAndRemainder(n int, d int) (int, int, error) {
	if d == 0 {
		return 0, 0, errors.New("0으로 나눌 수 없습니다.")
	}
	return n / d, n % d, nil
}

func main() {
	result, remainder, err := divAndRemainder(5, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1) // exit(0): 오류 없이 종료, exit(1): 오류 종료
	}
	fmt.Println(result, remainder)
}
*/

// 이름이 지정된 반환값 - 해당 함수의 로컬 변수로 간주
/*
func divAndRemainder(n, d int) (result int, remainder int, err error) {
	if d == 0 {
		err = errors.New("0으로 나눌 수 없습니다.") // := 은 새변수 지정 - 섀도잉 문제
		return result, remainder, err
	}
	result, remainder = n/d, n%d
	return result, remainder, err
}

func main() {
	n, d, e := divAndRemainder(5, 0)
	fmt.Println(n, d, e)
}
*/

// 코너 케이스 - 1) 섀도잉 문제
/*
func divAndRemainder(n, d int) (result int, remainder int, err error) {
	if d == 0 {
		err := errors.New("0으로 나눌 수 없습니다.") // := 은 새변수 지정 - 섀도잉 문제(error)
	}
	return result, remainder, err
}

func main() {
	n, d, e := divAndRemainder(5, 0)
	fmt.Println(n, d, e)
}
*/

// 코너 케이스 - 2) 해당 변수들을 반환할 필요가 x
/*
func divAndRemainder(n, d int) (result int, remainder int, err error) {
	// assign some values
	result, remainder = 20, 30 // result, remainder에 값 할당
	if d == 0 {
		return 0, 0, errors.New("0으로 나눌 수 없습니다.")
	}
	return n / d, n % d, nil // result, remainder = 20, 30 반환 x , 최종적으로 n, d가 반환
	//return	// 20, 30 받으려면 return만.. - missing return error
}

func main() {
	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)
}
*/

//빈반환
/*
func divAndRemainder(n, d int) (result int, remainder int, err error) {
	// assign some values
	result, remainder = 20, 30 // result, remainder에 값 할당
	if d == 0 {
		return 0, 0, errors.New("0으로 나눌 수 없습니다.")
	}
	//return n / d, n % d, nil // result, remainder = 20, 30 반환 x , 최종적으로 n, d가 반환
	return // 빈반환 - 이름이 지정된 변수에 마지막으로 할당된 값 (20, 30) 반환
}

func main() {
	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)
}
*/

//함수는 값 - 맵에 저장 가능
/*
func add(i int, j int) int { return i + j }

func sub(i int, j int) int { return i - j }

func mul(i int, j int) int { return i * j }

func div(i int, j int) int { return i / j }

var opMap = map[string]func(int, int) int{
	"+": add, // 패키지 블록
	"-": sub,
	"*": mul,
	"/": div,
}

func main() {
	expressions := [][]string{
		[]string{"2", "+", "3"}, // 슬라이스
		[]string{"2", "-", "3"},
		[]string{"2", "*", "3"},
		[]string{"2", "/", "3"},
		[]string{"2", "%", "3"},
		[]string{"two", "+", "three"},
		[]string{"5"},
	}
	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := opFunc(p1, p2)
		fmt.Println(result)
	}
}
*/
// 익명 함수 - defer문과 고루틴
/*
func main() {
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println("printing", j, "from inside of an anonymous function")
		}(i)
	}
}
*/

// 위의 익명 함수와 동일 기능
/*
func show(j int) {
	fmt.Println("printing", j, "from inside of an anonymous function")
}
func main() {
	for i := 0; i < 5; i++ {
		// func(j int) {
		// 	fmt.Println("printing", j, "from inside of an anonymous function")
		// }(i)
		show(i)
	}
}
*/
//클로저 - 함수 내부에 선언된 함수(함수의 범위 제한)
//파라미터로 함수를 전달
/*
func main() {
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobbert", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Println(people)

	// sort by last name
	sort.Slice(people, func(i int, j int) bool { // people 참조(캡처)
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)

	// sort by age
	sort.Slice(people, func(i int, j int) bool { // people 슬라이스가 sort.Slice에 의해 변경
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)
}
*/

// 함수에서 함수(클로저) 반환
func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor // base라는 외부변수에 접근 가능
	}
}

func main() {
	twoBase := makeMult(2)
	threeBase := makeMult(3)
	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}
}
