package main

import "fmt"

//3. 복합타입
func main() {
	//배열
	/*
		var x [3]int
		fmt.Println(x)
		var y = [3]int{10, 20, 30}
		fmt.Println(y)
		var z = [12]int{1, 5: 4, 6, 10: 100, 15}
		fmt.Println(z)

		var a = [...]int{1, 2, 3}
		var b = [...]int{1, 2, 3}

		fmt.Println(a == b) //true

		var c [2][3]int
		c[0] = 10	// 배열의 끝을 넘어서거나 음수의 인덱스 불가
		fmt.Println(c)

		var x [][]int
		var y = [][]int{{1, 2}, {4, 5, 6}}
		fmt.Println(x)
		fmt.Println(y)
	*/

	//슬라이스
	/*
		var x []int
		fmt.Println(x == nil) //true : nil = 값의 부재를 표현한 식별자
		fmt.Println(len(x))   //0

		x = append(x, 10)
		fmt.Println(x)

		x = append(x, 5, 6, 7)
		fmt.Println(x)

		y := []int{20, 30, 40}
		x = append(x, y...)
		fmt.Println(x)
	*/

	//수용력
	/*
		var x []int
		fmt.Println(x, len(x), cap(x))
		x = append(x, 10)
		fmt.Println(x, len(x), cap(x))
		x = append(x, 20)
		fmt.Println(x, len(x), cap(x))
		x = append(x, 30)
		fmt.Println(x, len(x), cap(x))
		x = append(x, 40)
		fmt.Println(x, len(x), cap(x))
		x = append(x, 50)
		fmt.Println(x, len(x), cap(x))
	*/

	//make - 수용력 미리
	/*
		var x []int
		//y := make([]string, 5)

		y := make([]int, 5)
		y[0] = 10
		y = append(y, 10)

		z := make([]int, 5, 10)
		z[0] = 10
		z = append(y, 10)

		fmt.Println(x, y, z)                // [] [10 0 0 0 0 10] [10 0 0 0 0 10 10]
		fmt.Println(len(x), len(y), len(z)) // 0 6 7
		fmt.Println(cap(x), cap(y), len(z)) // 0 10 7

		var a []int
		a = append(x, 5, 6, 7, 8)
		b := make([]int, 0, 10)
		b = append(b, 5, 6, 7, 8)

		fmt.Println(a, b)           // [5 6 7 8] [5 6 7 8]
		fmt.Println(len(a), len(b)) // 4 4
		fmt.Println(cap(a), cap(b)) // 4 10
	*/

	//슬라이스 선언 - x의 메모리 공유
	/*
		x := []int{1, 2, 3, 4}
		y := x[:2]
		z := x[1:]
		d := x[1:3]
		e := x[:]
		fmt.Println("x:", x)
		fmt.Println("y:", y)
		fmt.Println("z:", z)
		fmt.Println("d:", d)
		fmt.Println("e:", e)
		x[1] = 20
		y[0] = 10
		z[1] = 30
		fmt.Println("x:", x)
		fmt.Println("y:", y)
		fmt.Println("z:", z)
	*/

	/*
		x := []int{1, 2, 3, 4}
		y := x[:2]
		fmt.Println(cap(x), cap(y))

		y = append(y, 30) // y에만 추가 -> x에도 반영 : 같은 메모리 공유
		fmt.Println("x:", x)
		fmt.Println("y:", y)

		y = append(y, 30, 40, 50)
		fmt.Println("x:", x)
		fmt.Println("y:", y)

		y = append(y, 999)
		fmt.Println("x:", x)
		fmt.Println("y:", y)

		x := make([]int, 0, 5)
		x = append(x, 1, 2, 3, 4)
		// y := x[:2]	// 원본
		// z := x[2:]

		y := x[:2:2] // 수용력 = 2
		z := x[2:4:4]
		fmt.Println("x:", x)
		fmt.Println("y:", y)
		fmt.Println("z:", z)

		fmt.Println(cap(x), cap(y), cap(z))

		y = append(y, 30, 40, 50)
		x = append(x, 60) // 수용력 추가
		z = append(z, 70)
		fmt.Println("x:", x)
		fmt.Println("y:", y)
		fmt.Println("z:", z)
	*/

	//배열을 슬라이스로 변환
	/*
		x := [4]int{5, 6, 7, 8}
		y := x[:2]
		z := x[2:]
		x[0] = 10
		fmt.Println("x:", x)
		fmt.Println("y:", y)
		fmt.Println("z:", z)
	*/

	//copy
	/*
		x := []int{1, 2, 3, 4}
		y := make([]int, 4)
		num := copy(y, x)   // x: target source, num = 복사 원소 갯수
		fmt.Println(y, num) // [1 2 3 4] 4
	*/
	/*
		x := []int{1, 2, 3, 4}
		y := make([]int, 2)
		copy(y, x[:2])    // 세네번째만 복사
		fmt.Println(x, y) // [1 2 3 4] [1 2]

		num2 := copy(x[:3], x[1:])
		fmt.Println(x, num2)	// [2 3 4 4] 3
	*/
	/*
		x := []int{1, 2, 3, 4}
		d := [4]int{5, 6, 7, 8}
		y := make([]int, 2) // [0, 0]
		copy(y, d[:])
		fmt.Println(y) // [5 6]
		// copy(d[:], x)
		// fmt.Println(d) // [1 2 3 4]
		copy(d[1:], x[1:3])
		fmt.Println(d) // [5 2 3 8]
	*/

	//문자열과 룬 그리고 바이트
	/*
		x := "안녕하세요"
		fmt.Println(len(x)) // 15

		var s string = "Hello there"
		var b byte = s[6] // 정수

		fmt.Println(s, b)         // Hello world 116
		fmt.Println(s, string(b)) // Hello world t

		var s2 string = s[4:7]
		var s3 string = s[:5]
		var s4 string = s[6:]

		fmt.Println(s, s2, s3, s4) // Hello there o t Hello there
	*/
	/*
		var a rune = 'x'
		var s string = string(a)
		var b byte = 's'
		var s2 string = string(b)

		fmt.Println(a, s)  // 120 x
		fmt.Println(b, s2) // 115 s

		x := 65
		fmt.Println(x) // 65
	*/

	// 문자열에서 슬라이스로 변환
	/*
		var s string = "Hello, 🌞"
		var bs []byte = []byte(s)
		var rs []rune = []rune(s)
		var rs2 rune = rune(s[0])
		fmt.Println(bs)  // [72 101 108 108 111 44 32 240 159 140 158]
		fmt.Println(rs)  // [72 101 108 108 111 44 32 127774]
		fmt.Println(rs2) // 72
	*/
	/*
		x := "Hello 세상아"
		rs := []rune(x)
		bs := []byte(x)

		fmt.Println(x)              // Hello 세상아
		fmt.Println(len(x))         // 15
		fmt.Println(len([]rune(x))) // 9

		fmt.Println(rs)
		fmt.Println(bs)
		fmt.Println(string(rs)) // Hello 세상아
		fmt.Println(string(bs)) // Hello 세상아
	*/

	// 맵
	/*
		var nilMap map[string]int

		args := make(map[int][]string, 10) //make는 비어있는 맵(길이 0, 초기 지정 크기 이상으로 커질 수 있음)
		values := make([]int, 10)          //슬라이스는 제로값 채워줌

		fmt.Println(nilMap) // map[]
		fmt.Println(args)   // map[]
		fmt.Println(values) // [0 0 0 0 0 0 0 0 0 0]
	*/

	//맵 읽고 쓰기
	// := 사용 불가 - (새로운 변수 선언)
	/*
		totalWins := map[string]int{}
		totalWins["Orcas"] = 1
		totalWins["Lions"] = 2
		fmt.Println(totalWins["Orcas"])
		fmt.Println(totalWins["Kittens"])
		totalWins["Kittens"]++
		fmt.Println(totalWins["Kittens"])
		totalWins["Lions"] = 3
		fmt.Println(totalWins["Lions"])

		// 콤마 ok 관용구 및 맵 요소 삭제
		m := map[string]int{
			"hello": 5,
			"world": 0,
		}
		v, ok := m["hello"]
		fmt.Println(v, ok) // 5 true

		v2, ok := m["world"]
		fmt.Println(v2, ok) // 0 true

		v3, ok := m["good"]
		fmt.Println(v3, ok) // 0 false

		delete(m, "hello") // 반환값 없음
		fmt.Println(m)     // map[world:0]
	*/

	// 맵을 셋으로 이용
	/*
		intSet := map[int]bool{}
		vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
		for _, v := range vals {
			intSet[v] = true
		}
		fmt.Println(len(vals), len(intSet)) // 11 8 : 중복 허용 x

		fmt.Println(intSet[5])   // true : 5 key (o)
		fmt.Println(intSet[500]) // false : 500 key (x)

		fmt.Println(intSet) // unique 값만

		v, ok := intSet[5]
		fmt.Println(v, ok) // true true

		v2, ok := intSet[500]
		fmt.Println(v2, ok) // false false

		if intSet[100] { // false
			fmt.Println("100 is in the set")
		}
	*/
	// 구조체 - 객체지향
	/*
			type person struct {
				name string
				age  int
				pet  string
			}

			var fred person

			bob := person{}

			julia := person{
				"Julia", // 구조체 항목 값은 ,로 구분
				40,
				"cat", // 마지막 구조체에도 ,
			}

			beth := person{
				age:  30, // 항목 이름 명시하여 값 할당
				name: "Beth",
			}

			fmt.Println(fred)
			fmt.Println(bob)   // { 0 }
			fmt.Println(julia) // {Julia 40 cat}
			fmt.Println(beth)  // {Beth 30 } : 값이 지정되지 않은 변수는 제로 값

			//익명 구조체
			var person2 struct {
				name string
				age  int
				pet  string
			}

			person2.name = "bob"
			person2.age = 50
			person2.pet = "dog"

			fmt.Println(person2)

		type person3 struct {
			name string
			age  int
			pet  string
		}

		bob := person3{}

		fmt.Println(bob) // {0}
	*/

	// 구조체 비교와 변환
	type firstPerson struct {
		name string
		age  int
	}

	f := firstPerson{
		name: "Bob",
		age:  50,
	}

	var g struct {
		name string
		age  int
	}

	g = f
	fmt.Println(f == g) // true

}
