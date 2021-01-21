# Go tutorial
# [A Tour of Go](https://tour.golang.org/welcome/1)
## 기본 자료형
- .
    ```go
    bool
    
    string
    
    int int8(byte) int16 int32(rune) int64
    uint uint8 uint16 uint32 uint64 uintptr
    
    float32 float64
    
    complex64 complex128
    ```

## 상수
- const 키워드와 함께 선언
- 상수는 := 를 통해 선언될 수 없음
- Ex
    ```go
    const Pi = 3.14
    const world = "hell world!"
    const check = false
    ```
- 숫자 상수는 명시적 캐스팅 등으로 타입이 주어지기 전까지 타입을 가지지 않음
- +alpha
    ```go
    const(
      Pi = 3.14
      world = "hello world!"
      check = false
      )
    ```
- ++alpha
    - 상수값을 0부터 순차적으로 부여하기 위해 iota라는 identifier를 사용할 수 있음
    ```go
    const(
      one = iota // 0
      two        // 1
      three      // 2
      four       // 3
      five       // 4
    )
    ```

## 변수 선언
- 초기화되지 않은 변수는 자동으로 zero value 가 주어짐
    - 숫자 type에는 0
    - boolean type에는 false
    - string에는 "" (빈 문자열)
### **var [변수명1, 변수명2, 변수명3, ...] [타입]**
- var 문은 변수에 대한 목록을 선언
- 문장 읽듯이 선언한다는 컨셉이라는데 음?
- 초기화 x
    - Ex1. 변수 a를 int 타입으로
        - var a int
    - Ex2. 변수 arr_a를 크기 5의 int 타입으로
        - var a [5]int
- 초기화 o
    - var a int = 1
    - var b, c int = 1, 2
    - var str_a string = "Hello"
    - var arr_a = [5]int{1, 2, 3, 4, 5}
    - var arr_b = [...]int{6, 7, 8}
- +alpha
    ```go
    var(
      c bool = true
      python bool = false
      java string = "no!"
    )
    ```

### 짧은 변수 선언
- 함수 내에서 := 라는 짧은 변수 선언은 암시적 type으로 var 선언처럼 사용될 수 있음
- 함수 밖에서는 모든 선언이 키워드(var, func, 기타 등등)로 시작하므로 := 구문 사용 불가
- Ex
    - k := 3
    - c, python, java := true, false, "no!"

### Type 변환
- Go에서는 다른 type의 요소들 간의 할당에는 명시적인 변환을 필요로함
- Ex1
    ```go
    var i int = 42
    var f float64 = float64(i)
    var u uint = uint(f)
    ``` 
- Ex2
    ```go
    i := 42
    f := float64(i)
    u := uint(f)
    ``` 
    

## 함수 선언
- 변수(매개변수) 이름 뒤에 type
    ```go
    func add(x int, y int) int{
        return x + y
    }
    ```
- 두 개 이상의 이름이 주어진 함수 매개변수가 같은 type 일 때 마지막 변수에만 type 표현 
    ```go
    func add(x, y int) int{
        return x + y
    }
    ```
- 함수는 여러개의 결과 반환 가능
    ```go
    func swap(x, y string)(string, string){
      return y, x
    }
    ```
- 이름이 주어진 반환 값 (naked return : 인자가 없는 return 문)
    ```go
    func add(x, y int)(result int){
        result = x + y
        return
    }
    ```

## Go 키워드
- 함수 밖에서는 모든 선언이 키워드(var, func, 기타 등등)로 시작
- Go 키워드들은 변수명, 상수명, 함수명등의 Identifier로 사용 불가
    ```go
    break        default      func         interface    select
    case         defer        go           map          struct
    chan         else         goto         package      switch
    const        fallthrough  if           range        type
    continue     for          import       return       var
    ```

## For
- Go의 반복 구조는 for 하나
- c와 비슷한 구조인데 for의 세 가지 구성 요소를 감싸는 ()만 없음 ({} 필수)
- for의 세 가지 구성 이름
    - 초기화 구문 : 첫 번째 iteration 전에 수행
    - 조건 표현 : 매 모든 iteration 이전에 판별
    - 사후 구문 : 매 iteration 마지막에 수행
- 초기화 구문은 짧은 변수 선언으로
    ```go
    sum := 0
    for i := 0; i < 10; i++{
    sum += i
    }
    ```
- while 처럼 쓰고싶으면 초기화 구문 사후 구분 생략
    ```go
    sum := 1
    for sum < 100{
        sum += sum
    } 
    ```
- 무한루프
    ```go
    for{
        // work
    } 
    ```

## If
- For 과 비슷하게 ()만 없고 {} 필수
- 짧은 구문의 If
    - 조건문 전에 수행될 짧은 구문으로 시작될 수 있음
    - 이 짧은 구문에서 선언된 변수들은 오직 if문(else 블럭 포함) 내에서만 사용 가능
    ```go
    if v := math.Pow(x, n); v < num_a{
        ret = true
    }
    ```
- else 블럭
    ```go
    if v := math.Pow(x, n); v < lim{
        return v
    } else {
        fmt.Println("%g >= %g\n", v, lim)
    }
    ```

## Switch
- Go의 switch는 뒤이어 오는 모든 case를 실행하지 않고 오직 첫 번째로 선택된 케이스만들 실행함
- 자동으로 break가 제공됨
- Go의 스위치 케이스는 상수일 필요가 없고 정수일 필요도 없
    ```go
    switch os := runtime.GOOS; os{
    case "drawin":
        fmt.Println("OS X.")
    case "linux":
        fmt.Println("Linux.")
    default:
        fmt.Println("%s.\n", os)
    }
    ```
  
- 조건이 없는 Switch
    - 긴 if-else 체인을 작성할 때 좋은 방식
    ```go
    t := time.now()
    switch {
    case t.Hour() < 12:
        fmt.Println("Good morning!")
    case t.Hour() < 17:
        fmt.Println("Good afternoon.")
    default:
        fmt.Println("Good evening.")   
    }
    ```

## Defer
- defer문은 자신을 둘러싼 함수가 종료할 때까지 어떠한 함수의 실행을 연기
- 연기된 호출의 인자는 즉시 평가됨
- 함수가 수행되면서 바뀐 변화는 반영되지 않는듯
    ```go
    func main(){
        a, b := 1, 2
  
        defer fmt.Println(a + b) // 3
        
        fmt.Println("Hello")
        a, b = 3, 4 
    }   
    ```
- 연기된 함수 호출들은 스택에 쌓임 (후입선출)

## Pointers
- Go는 포인터를 지원함
- 포인터 산술은 지원하지 않음 + 메모리 주소 직접 대입 지원하지 않음
- 포인터의 zero value은 nil(Go에서는 NULL을 nil로 사용)
- 역참조(*) 사용해 주소의 값 
  ```go
  var numPtr1 *int // 빈 포인터
  var numPtr2 *int = new(int) // 메모리 할당
  
  num := 3
  numPtr1 = &num
  
  fmt.Println(*numPtr1)
  ```

## Structs
- 필드의 집합체 (1)
- 구조체의 필드는 .(dot)으로 접근할 수 있음 (2)
- 구조체 포인터를 통해서 구조체 필드를 접근할 수 있음 (3)
  ```go
  type Vertex struct { // (1)
	X int
    Y int
  }
  
  func main() { 
    v := Vertex{1, 2}
    v.X = 4 // (2)
  
    p := &v // (3)
    p.X = 1e9
  }
  ```

~~- Struct Literals~~
  - ~~왜 필요한지 모르겠으니 쓸 일 있으면 다시 보자~~

## Arrays
- 선언 형식
  ```go
  var num1 [5]int
  var num2 = [5]int{1, 4, 6, 8, 9, 10}
  num3 := [5]int{2, 3, 5, 7, 11, 13}
  ```
## Slices
- 동적 배열이라고 생각하자
- 연속된 메모리 공간에 동일한 타입의 데이터를 순차적으로 저장할 때 사용
- 길이가 고정적인 배열과 다른게 길이를 유동적으로 다룰 수 있음
- 배열은 길이의 변경이 필요할 때마다 새로운 길이를 가진 배열을 다시 할당하는 비효율적인 작업을 해야함
  - slice는 길이의 변경에 대비해 미리 특정 용량을 가진 배열을 할당해두고, 정해진 길이 만큼만 사용할 수 있도록 하여, 길이의 수정 만으로 배열을 재할당할 필요 없이 유동적으로 다룰 수 있도록 하는 것

#### Slice length and capacity
  - 슬리이스는 _length(길이)_ 와 _capacity(용량)_ 둘 다 보유함
  - 길이 : 슬라이스에 포함된 요소의 개수, len(s)
  - 용량 : 기본 배열의 요소 개수, cap(s)

#### Slice 생성
- 초기화하지 않은 slice는 len과 cap이 0인 nil slice가 됨
  - make와 리터럴로 초기화된 slice는 len과 cap이 0이어도 nil 아님
- make(type, len, cap) (용량 지정)
  ```go
  slice1 := make([]int, 5, 10)
  ```
- slice 리터럴 (용량 생략)
  ```go
  slice2 := []int{1, 2, 3, 4, 5}  
  ```

#### Slicing
- Go에서는 배열의 slice의 특정 영역을 slice 형태로 추출할 수 있도록 slicing 기능 제공
- slice는 어떤 데이터도 저장할 수 없음 단지 기본 배열(or slice)의 한 영역을 나타낼 뿐
- 그래서, slice의 요소를 변경하면 기본 배열(or slice)의 해당 요소가 수정됨
  ```go
  num := [5]int{1, 2, 3, 4, 5}
  num_slice := num[2:]
  num_slice[1] = 10
  fmt.Println(num, num_slice)
  // num : [1, 2, 3, 10, 5]
  // num_slice : [3, 10, 5] 
  ```

#### Append Slice
- slice에 값 추가
- **용량 값에 따라 다르게 동작**
  - slice의 용량이 새로운 요소들을 추가하기 충분하면 메모리를 공유하고 길이가 다른 slice 생성
  - 용량이 충분하지 않으면 상황에 따라 최적화 된 용량을 가진 slice를 생성
    - 이 slice는 기존 slice와 메모리를 공유하지 않음

## etc
- factored import 문 : import를 그룹 짓는

- 함수 이름이 대문자로 시작하면 Exportable(Public)
- 함수 이름이 소문자로 시작하면 Not exportable(Private)

[한 눈에 끝내는 고랭 기초](https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88)
