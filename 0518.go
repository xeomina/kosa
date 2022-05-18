package main // 패키지명이 main 인 경우, 컴파일러는 해당 패키지를 공유 라이브러리가 아닌 실행(executable) 프로그램으로 만든다.
import "fmt"

//  Go의 표준 라이브러리인 fmt 패키지, 노출된 식별자는 . 연산자를 통해 접근

func main() {
	/*
			// 섀도잉 변수 - 블록 내에 이름이 같은 변수가 있는 것을 의미
			x := 10
			if x > 5 {
				fmt.Println(x) // 10
				x = 20
				x := 5
				fmt.Println(x) // 5
			}
			fmt.Println(x) // 20

		x := 10
		if x > 5 {
			x, y := 5, 20
			fmt.Println(x, y) // 5 20
		}
		fmt.Println(x) // 10
	*/

	// if문
	/*
		n := rand.Intn(10)

		fmt.Println(n)
		if n == 0 {
			fmt.Println("That's too low")
		} else if n > 5 {
			fmt.Println("That's too big")
		} else {
			fmt.Println("That's a good number")
		}
		fmt.Println(n)	// n 출력 (error x)
	*/

	/*
		if n := rand.Intn(10); n == 0 {		// n : if/else 블록 범위내에서만 사용가능한 변수
			fmt.Println("That's too low")
		} else if n > 5 {
			fmt.Println("That's too big:", n)
		} else {
			fmt.Println("That's a good number:", n)
		}
		// fmt.Println(n) // eundefined: n error - if/else 문을 벗어났기 때문
	*/
	/*
	   	// for 문
	   	for i := 0; i < 10; i++ {
	   		fmt.Println(i)
	   	}

	   	// for-range 문
	   	evenVals := []int{2, 4, 6, 8, 10, 12}
	   	for i, v := range evenVals { // range가 두개의 값을 반환
	   		fmt.Println(i, v)
	   	}

	   	// 맵 순회
	   	m := map[string]int{
	   		"a": 1,
	   		"c": 3,
	   		"b": 2,
	   	}

	   	for i := 0; i < 3; i++ {
	   		fmt.Println("Loop", i)
	   		for k, v := range m {
	   			fmt.Println(k, v)
	   		}
	   	}

	   	// 문자열 순회 - 룬 순회 (바이트 x)
	   	samples := []string{"hello", "apple_π!"}
	   	for _, sample := range samples {
	   		for i, r := range sample {
	   			fmt.Printf("%T\n", r) // 타입 출력 : reflect.Type or Printf("%T")
	   			fmt.Println(i, r, string(r))
	   		}
	   		fmt.Println()
	   	}

	   	//for-range 값은 복사본 - 원본 수정 x
	   	evenVals2 := []int{2, 4, 6, 8, 10, 12}
	   	for _, v := range evenVals2 {
	   		v *= 2
	   	}
	   	fmt.Println(evenVals2) // [2 4 6 8 10 12]

	   	// for 문 레이블링
	   	samples2 := []string{"hello", "apple_π!"}
	   outer:
	   	for _, sample := range samples2 {
	   		for i, r := range sample {
	   			fmt.Println(i, r, string(r))
	   			if r == 'l' { // 문자열이 l이면 outer
	   				continue outer
	   			}
	   		}
	   		fmt.Println()
	   	}
	*/
	// switch 문

	words := []string{"a", "cow", "smile", "gopher",
		"octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}
	// loop 레이블 사용
loop:
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "is even")
		case i%3 == 0:
			fmt.Println(i, "is divisible by 3 but not 2")
		case i%7 == 0:
			fmt.Println("exit the loop!")
			break loop
		default:
			fmt.Println(i, "is boring")
		}
	}
	// 공백 switch 문 - 비교 되는 값 명시 x
	words2 := []string{"hi", "salutations", "hello"}
	for _, word := range words2 {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "is a short word!")
		case wordLen > 10:
			fmt.Println(word, "is a long word!")
		default:
			fmt.Println(word, "is exactly the right length.")
		}
	}

	// goto 문 - 잘 사용안함
}
