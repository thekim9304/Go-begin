# Go tutorial
# [A Tour of Go](https://tour.golang.org/welcome/1)
## 기본 자료형
```go
bool

string

int int8(byte) int16 int32(rune) int64
uint uint8 uint16 uint32 uint64 uintptr

float32 float64

complex64 complex128
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
    - var(
          c bool = true
          python bool = false
          java string = "no!"
      )

### 짧은 변수 선언
- 함수 내에서 := 라는 짧은 변수 선언은 암시적 type으로 var 선언처럼 사용될 수 있음
- 함수 밖에서는 모든 선언이 키워드(var, func, 기타 등등)로 시작하므로 := 구문 사용 불가
- Ex
    - k := 3
    - c, python, java := true, false, "no!"


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
- 이름이 주어진 반환 값
    ```go
    func add(x, y int)(result int){
        result = x + y
        return
    }
    ```

## etc
- factored import 문 : import를 그룹 짓는
- naked return : 인자가 없는 return 문
    - 이름이 주어진 반환 값을 반환
    ```go
    func swap(str1, str2 string)(res1, res2 string){
        res1 = str2
        res2 = str1
        return
    }
    ```
- 함수 이름이 대문자로 시작하면 Exportable(Public)
- 함수 이름이 소문자로 시작하면 Not exportable(Private)

[한 눈에 끝내는 고랭 기초](https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88)
