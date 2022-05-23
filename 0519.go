package main

import (
	"fmt"
	"time"
)

/*
// defer
func main() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}

// defer는 go에서 여러 클로저를 지연 - 후입선출(LIFO)
// 마지막 defer로 등록된 것이 가장 먼저 실행
// defer 클로저 내의 코드는 return 문이 실행된 후에 실행
func example() {
	defer func() int {
		return 2 // 해당 값을 읽을 방법이 없다
	}()
}

// 값에 의한 호출을 사용하는 go
type person struct {
	age  int
	name string
}

// 구조체를 이용 - 값 수정 시도
func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "Goodbye"
	p.name = "Bob"
}

func main() {
	p := person{}
	i := 2
	s := "Hello"
	modifyFails(i, s, p)
	fmt.Println(i, s, p) // 2 Hello {0 } : 값 변경 x - 변수 복사본이기 때문
}

// 맵 이용 - 값 수정
func modMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}

func main() {
	m := map[int]string{
		1: "first",
		2: "second",
	}
	modMap(m)
	fmt.Println(m)

	s := []int{1, 2, 3}
	modSlice(s)
	fmt.Println(s) // map[2:hello 3:goodbye]	[2 4 6]
	// : 맵은 변수 수정 가능 but 슬라이스는 수정은 되나 길이를 늘리는 것은 불가
	// 왜? 맵과 슬라이스는 포인터로 구현되었기 때문
}

*/

// 6. 포인터 - 값이 저장된 메모리의 위치 값 변수
// 6.1
/*
func main() {
	x := 10
	pointerToX := &x          // & : 주소 연산자
	x2 := *pointerToX         // * : 간접 연산자(역참조)
	fmt.Println(pointerToX)   // x의 메모리 주소 출력 - 0xc0000ba000
	fmt.Println(&pointerToX)  // pointerToX의 메모리 주소 출력 - 0xc0000b6018
	fmt.Println(*pointerToX)  // pointerToX가 가리키는 값 반환 - 10
	fmt.Println(*&pointerToX) // 0xc000018030
	z := 5 + x2
	fmt.Println(z) // 15
}
*/

/*
func main() {
	var x *int            // 포인터의 제로 값은 nil
	fmt.Println(x == nil) // true
	fmt.Printf("%T\n", x) // *int
	//fmt.Println(*x)       // 패닉

	var y = new(int)      // new() : 주어진 타입(int)의 제로값을 가리키는 포인터 변수 생성
	fmt.Println(y == nil) // false
	fmt.Printf("%T\n", y) // *int
	fmt.Println(y, *y)    // 0xc0000140d0 0
}
*/

/*
func main() {
	type person struct {
		FirstName  string
		MiddleName *string
		LastName   string
	}

	nName := "Perry"	// 1) 미리 변수에 상수 값

	p := person{
		FirstName: "Pat",
		//MiddleName: "Perry", // 해당 라인은 컴파일 되지 않음
		//MiddleName: &"Perry", // 기본 타입의 리터럴이나 상수는 주소가 없으므로 주소연산자를 사용할 수 없음
		//MiddleName: &nName,	// 1) 미리 변수에 상수 값
		MiddleName: stringp("Perry"),	// 2) 헬퍼 함수를 이용해 상수를
		LastName:   "Peterson",
	}
	fmt.Println(p)

	func stringp(s string) *string {
		return &s	// 변수의 주소 파라미터는 변수 주소를 가지고 주소 반환
	}
}
*/

// 포인터는 변경 가능한 파라미터를 가리킨다
/*
func failedUpdate(g *int) {
	x := 10
	g = &x
}


func main() {
	var f *int      // f = nil
	failedUpdate(f) // f 값인 nil 복사 > g에 할당 = nil > 이후 새로운 x 선언 > g가 x 주소 할당 받아도 f는 nil
	fmt.Println(f)  // <nil>
}
*/
/*
func failedUpdate(px *int) {
	x2 := 20
	px = &x2
}

func update(px *int) {
	*px = 20
}

func main() {
	x := 10
	failedUpdate(&x) // x의 주소 복사하여 px에 할당 > px가 x2를 가리키도록
	fmt.Println(x)   // 10 - 아직 x 값 변화 없음
	update(&x)       // x의 주소 복사하여 px에 다시 할당 > 역참조로 20이 되도록 변경
	fmt.Println(x)   // 20
}
*/

/*
// 포인터는 최후의 수단
type Foo struct {
	Field1 string
	Field2 int
}

// 포인터 권장 x
func MakeFoo(f *Foo) error {
	f.Field1 = "val"
	f.Field2 = 20
	return nil
}

// 구조체 권장
func MakeFoo2() (Foo, error) {
	f := Foo{
		Field1: "val2",
		Field2: 200,
	}
	return f, nil
}

// 둘다 실행은 되지만 포인터는 데이터 흐름 이해를 어렵게 하고 GC에게 추가 작업 부하를 건다(힙)
func main() {
	myFoo := Foo{}
	fmt.Println(myFoo) // { 0}

	MakeFoo(&myFoo)
	fmt.Println(myFoo) // {val 20}

	myFoo, err := MakeFoo2()
	if err != nil {
		fmt.Println("에러")
	}
	fmt.Println(myFoo) // {val2 200}
}
*/
/*
// 버퍼 슬라이스
file, err := os.Open(fileName)
if err != nil {
	return err
}
defer file.Close()

data := make([]byte, 100)

for {
	count, err := file.Read(data)

	process(data[:count])
}
*/

// 7. 타입, 메서드, 인터페이스
/*
// 타입을 위한 메서드 - 반드시 패키지 블록 레벨에서 정의
type Person struct { // 구조체를 이용한 사용자 정의 타입
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

func main() {
	fred := Person{
		FirstName: "Fred",
		LastName:  "Fredson",
		Age:       52,
	}

	output := fred.String()
	fmt.Println(output) // Fred Fredson, age 52
}
*/

// 포인터 리시버와 값 리시버
type Counter struct { // 구조체
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() { // 포인터 리시버
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func main() {
	var c Counter           // 초기값 없이 변수만 선언 : 0
	fmt.Println(c.String()) // total: 0, last updated: 0001-01-01 00:00:00 +0000 UTC
	c.Increment()           // 포인터 > 현재시간
	fmt.Println(c.String()) // total: 1, last updated: 2022-05-19 17:44:26.2516101 +0900 KST m=+0.010208501
}
