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

## etc
- factored import 문 : import를 그룹 짓는

- 함수 이름이 대문자로 시작하면 Exportable(Public)
- 함수 이름이 소문자로 시작하면 Not exportable(Private)

[한 눈에 끝내는 고랭 기초](https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88)
