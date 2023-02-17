# go-practices

## Go 특징

### 정적 타입

> 자료형을 컴파일할 때 결정하면 정적 타입이고, 실행할 때 컴파일 하면 동적 타입이다. 이 중 Go는 **정적 타입**이라는 특징을 가지고 있다.

### 강 타입

> 약 타입은 타입의 값을 바꿀 수 있다. 약 타입 언어는 자료형이 달라도 컴파일 또는 실행 시점에 정해진 규칙에 따라 암시적으로 형 변환을 수행한다.
> 강 타입은 값 자체가 타입이며 타입을 바꿀 수 없다. Go는 **강 타입** 언어로 암시적 형 변환이 안된다.

### 컴파일 언어

> Go는 컴파일 언어로 컴파일을 수행하면 기계어로 된 바이너리 파일을 만들어낸다. 컴파일 언어이면서 네이티브 바이너리 형식이기 때문에 완전한 **실행 파일**을 만들어낸다.

### 가비지 컬렉션

> Go는 가비지 컬렉션을 제공한다. 메모리를 관리해주는 가비지 컬렉터가 **실행 파일 안에 내장**되어 있다.

### 병행성

> Go는 언어 차원에서 병행성을 제공한다. 병행성은 동시 처리의 논리적인 개념이며 여러개가 시간을 쪼개어 순차적으로 실행되어도 병행성을 만족하고, 하나의 스레드가 여러개 실행해도 만족한다.
> 병렬성은 물리적인 개념으로 작업을 여러 코어에 나눠서 동시에 처리하는 상태를 의미한다. Go는 **go** 키워드를 통해 함수 여러개를 동시에 실행할 수 있다. 이렇게 실행 된 함수를 **고루틴**이라고 한다.
> 스레드는 운영체제의 커널에서 제공하는 리소스기 떄문에 생성할수록 부담이 커지는데, Go는 적정량의 스레드를 생성해서 고루틴을 처리한다. 
> 또한 채널을 이용하여 고루틴끼리 통신이 가능하다. 

### 모듈화 및 패키지

> 소프트웨어 규모가 점점 크고 복잡해지면서 코드의 재사용성이 중요해졌다. Go 또한 언어 자체에서 모듈화를 제공하여 인터넷에 있는 소스코드를 바로 사용할 수 있다.

### 컴파일 속도

> C, C++은 처리해야 하는 헤더 파일이 많아 컴파일 속도가 매우 느리지만, Go는 그와 달리 헤더 파일이 없다. 소스코드를 패키지화하여 변경된 부분만 컴파일하므로 속도가 매우 빠르다. 또한, 문법적으로 복잡한 요소를 줄여 컴파일 속도에 유리하게 설계했다.

### 활용 범위

> 메모리 관리가 다소 느슨해도 되고, 규모가 크고 복잡하며 유지보수가 빈번한 곳에서 편리하게 사용할 수 있다.

## Go 시작하기

### [Hello, World!](src/start/hello_world.go)

> go 컴파일 방법은 **go build "소스파일"** 이다. 컴파일을 하고 나면 실행 파일이 생성되고, 생성 된 실행파일을 실행하면 된다.
> go는 모든 부분이 패키지로 구성되어 있고, 작성한 프로그램은 항상 **main** 패키지 부터 실행된다. import는 패키지를 사용하기 위한 키워드고, 다른 언어와 마찬가지로 main 함수부터 시작한다.

## 기본 문법 알아보기

> Go는 다른 언어와 달리 작성 스타일을 일부 강제한다. 대표적으로 중괄호는 반드시 같은 줄에 중괄호가 있어야 한다. 이렇게 하므로써 가독성이 좋아지는 장점이 있다.
> 그리고 컴파일러와 별개로 문법 스타일을 맞춰주는 **gofmt** 명령도 제공한다. 

### 세미콜론

> Go에서는 대부분 세미콜론을 생략한다. 한줄에 여러 구문을 사용할 경우에만 세미콜론으로 구분하면 된다.

### 주석

> 한줄 주석과 두줄 주석이 있다. 한줄 주석은 **(//)**, 두줄 주석은 **(/\* \*/)** 형태로 표기한다.

### 중괄호

> 중괄호는 구문의 맨 뒤에서 시작한다. **반드시 같은 줄에서 시작**해야 하며, 새 줄부터 시작하면 오류가 발생한다.

### 들여쓰기

> 들여쓰기에 탭을 사용한다. 스페이스로 들여쓰기를 수행해도 gofmt 명령을 사용하면 모두 탭으로 변경된다.

## 변수 사용하기

> 변수는 **var** 키워드를 사용하는 방식과 자료형을 생략하는 방식 두가지가 있다. 자료형이 변수명 뒤에 오고 변수명은 문자와 숫자로 이루어져 있다. 단 숫자로 시작할 수 없다.
> 자료형을 생략하면 변수의 자료형은 대입하는 값의 자료형으로 결정된다.

### [짧은 선언 사용하기](src/variables/short/short_declare.go)

> := 를 사용하면 var, 자료형 키워드를 사용하지 않고 간단하게 변수를 선언하고 초기화할 수 있다. 변수의 자료형은 대입하는 값의 자료형으로 결정된다.

### [변수 여러 개 선언하기](src/variables/many/many_declare.go)

> 변수를 여러 개 선언하고 초기화할 때는 변수와 값을 콤마로 구분하여 나열한다. 변수를 선언한 순서대로 대입되며, 반드시 선언한 변수의 개수와 대입할 값의 갯수가 같아야 한다.
> 짧은 선언도 사용 가능하다. 변수가 선언된 후에도 여러개에 값을 대입할 수 있다. Go는 이 기능을 **병렬 할당(Parallel assignment)** 이라고 한다.  
> (Tip. 사용하지 않고 선언만 한 변수가 있으면 컴파일 에러가 발생하는데. _ (밑줄 문자)에 대입하면 컴파일 에러가 발생하지 않는다.)

## 숫자 사용하기

> Go는 정수, 실수, 복소수를 지원한다.

### [정수](src/numbers/integer/integer.go)

> 정수를 저장할 땐 8진수는 숫자 앞에 0을 붙이고, 16진수는 숫자 앞에 0x를 붙인다.

### [실수](src/numbers/floating/floating.go)

> 실수는 소수점을 사용하거나 지수 표기법으로 표기할 수 있으며, 변수에 저장 될 때는 부동소수점 방식을 사용한다. 지수 표기법으로 e를 사용하고, +, -로 위치를 지정한다.
> 컴퓨터는 2진수를 사용하므로 실수를 정확히 표현할 수 없다. 근사값으로 표현하기 때문이다. 따라서 실수는 등호로 직접 비교하면 안된다. 연산값과 비교할 값의 차이를 구한 뒤 머신 앱실론보다 작거나 같은지를 비교해야 한다.

### [복소수](src/numbers/complex/complex.go)

> 복소수는 실수부 + 허수부 형태이며 허수부에는 마지막에 i를 붙인다. 실수부와 허수부는 고정소수점 및 부동소수점 형태로 표현한다. 복소수를 변수에 저장할 때 둘중 하나는 생략이 가능하다.

### [바이트](src/numbers/byte/byte.go)

> 바이트는 보통 16진수, 문자 값으로 저장한다. 실무에서는 바이너리에서 데이터를 읽거나, 암복호화 시 주로 사용한다. 문자열은 저장할 수 없다.

### [룬](src/numbers/rune/rune.go)

> 룬은 유니코드 문자 코드를 저장할 때 사용한다. 

### [숫자 연산자](src/numbers/operator/operator.go)

> 숫자 연산에는 덧셈(+), 뺄셈(-), 곱셉(*), 나눗셈(/), 나머지(%), 시프트(<< >>), 비트반전(^) 연산자를 사용할 수 있다.
> 숫자끼리 연산을 할 때 자료형이 다르면 컴파일 에러가 발생한다(강 타입 특성). 따라서 명시적으로 자료형을 변환해야 한다.
> 실수를 정수로 변환하면 소수점 아래를 버리고, 크기가 작은 자료형으로 변환하면 넘치는 비트를 모두 버린다.

### [오버플로우와 언더플로우](src/numbers/overflow/overflow.go)

> 각 자료형에서 저장할 수 있는 최대 크기를 넘어서는 것을 오버플로우, 최소 크기보다 작아지면 언더플로우가 발생한다. Go는 오버플로우가 발생하면 컴파일 오류가 발생한다. 
> 오버플로우 상황 자체를 허용하지 않는다. 언더플로우가 발생해도 오버플로우 컴파일 에러로 표시한다.

### [변수의 크기 구하기](src/numbers/sizeof/sizeof.go)

> 변수의 크기를 구하려면 unsafe 패키지의 SizeOf 함수를 이용하자. 바이트 단위로 크기를 반환한다.

## 문자열 사용하기

> Go 언어는 문자열을 저장할수 있는 변수를 제공한다. 문자열은 따옴표로 묶어야 하며, 유니코드 문자로도 지정이 가능하다.
> 여러줄로 된 문자열을 저장할 때는 백틱으로 묶어준다.

### [문자열 길이 구하기](src/strings/length/length.go)

> 문자열 변수에 저장된 문자열의 길이를 구할 때는 **len** 함수를 사용한다. 한글은 UTF-8로 저장 시 한 글자당 3byte의 길이를 가지고 있음에 유의하라.
> 한글의 실제 문자열의 길이를 구하려면, unicode/utf8 패키지의 **RuneCountInString** 함수를 사용하라.

### [문자열 연산하기](src/strings/operator/operator.go)

> 문자열을 비교할 때는 == 연산자를 사용한다. 여러 문자열을 붙일 때는 + 연산자를 사용한다. 배열과 동일하게 [] 대괄호로 각 문자를 가져올 수 있다.
> 한글은 rune 배열로 만든 후 가져와야 한다.

### [문자열 수정하기](src/strings/operator/operator.go)

> Go는 문자열 저장 후 내용 수정이 불가하다. 다른 문자열은 대입할 수 있지만, 내용을 수정하려면 컴파일 오류가 발생한다. (불변)

## [불(bool) 사용하기](src/bool/bool.go)

> Go는 참과 거짓을 표현하는 1byte bool 자료형을 제공한다. 보통 &&, ||, ! 논리 연산자화 함께 사용한다. 비교 연산자 또한 bool 값으로 나온다.

## [상수 사용하기](src/const/const.go)

> Go는 **const** 키워드로 상수를 만들 수 있다. 반드시 선언과 동시에 초기화가 되어야 하며 선언 뒤에는 값을 변경할 수 없다.
> 변수와 같이 자료형을 생략하면 대입하는 값의 자료형으로 결정된다.

## [열거형 사용하기](src/enum/enum.go)

> Go는 **const** 키워드로 열거형을 선언한다. 상수에 값을 일일히 대입하지 않고 순서대로 사용하려면 **iota**를 사용하면 된다.
> Go의 Zero Value issue로 0부터 사용하는 것은 좋지 않다. 항상 1부터 시작하자. 연산자와 특정 값을 조합하면 상수 값을 순서대로 생성하지 않고 특정한 순서로 생성할 수 있다.

## 패키지 사용하기

> 각종 기능과 라이브러리를 패키지로 만들어서 제공한다. 소스코드에서 패키지를 사용하려면 **import** 키워드를 사용한다.

### [여러 패키지 사용하기](src/package/many/many.go)

> 여러 패키지를 사용할 때에는 import 키워드를 여러개 사용하거나 import 키워드 후 괄호로 묶어서 사용이 가능하다.

### [전역 패키지 사용하기](src/package/global/global.go)

> 패키지를 가져 올 때 패키지 이름 앞에 .(점)을 사용하면 전역 패키지로 사용할 수 있다. 단 함수, 변수, 상수 이름이 중복될 수 있기에 주의해서 사용해야 한다.

### [패키지 별칭 사용하기](src/package/alias/alias.go)

> 패키지를 가져올때 별칭을 지정할 수 있다. 기본 패키지와 내가 만든 패키지가 중복되거나 할때 활용할 수 있다. 
> 또한 패키지를 가져온 뒤 밑줄 문자를 지정하면 컴파일 오류가 발생하지 않는다.

## [if 조건문](src/condition/condition.go)

> if 조건문은 특정 조건을 설정하여 프로그램의 흐름을 바꿀 때 사용한다. Go에서는 다른 언어와 마찬가지로 if 조건문을 제공한다.
> 조건문을 작성할 때 괄호를 사용하지 않고 조건문을 시작하는 줄에서 중괄호가 시작된다. else if, else 사용이 가능하다.
> 또한, 조건문 안에서 함수 사용이 가능하다. 조건문 안에서 함수를 실행한 뒤 조건을 판단하는 방식으로 가능하다. 해당 조건문 안에서 선언된 변수의 스코프는 if 안에서만 사용이 가능하다.

## 반복문

> Go는 타 언어와 달리 반복문은 for만 존재한다.

### [for 반복문](src/loop/for/for.go)

> C, C++ 등의 반복문과 동일하다. 단 반복문을 작성할 때 괄호를 사용하지 않으며, 시작하는 줄에서 중괄호가 시작된다.
> **초기값, 조건식, 변화식** 형태로 이루어져 있다. 조건식만 사용하므로써 while 처럼 작동하게 만들 수 있다.

### [break 사용하기](src/loop/break/break.go)

> for 키워드에 아무것도 없으면 무한 루프가 된다. 이때 반복문을 중단하기 위해 **break** 키워드를 사용할 수 있다.
> break 키워드에 **레이블**을 지정할 수 있다. 중첩 된 반복문에서 레이블을 지정하면 편리하다. 레이블과 for 키워드 사이에 다른 코드가 있으면 안된다.

### [continue 사용하기](src/loop/continue/continue.go)

> 반복문에서 특정 부분 이하는 실행하지 않고 다음 반복으로 넘어가려면 **continue**를 사용한다. 레이블도 사용할 수 있다.

### [반복문에서 변수 여러개 사용하기](src/loop/many/many_argument.go)

> 반복문의 변화식에서 여러 변수를 처리하려면 병렬 할당을 사용해야 한다. 단 주의할점은 변화식 부분에서 병렬 할당을 사용하지 않고 변화식을 그대로 나열하면 오류가 발생한다.

## [goto 사용하기](src/goto/goto.go)

> goto는 정해진 레이블로 곧장 이동한다. 일반적으로는 권장되지 않는다. 

## [switch 분기문 사용하기](src/switch/switch.go)

> 다양한 조건을 if, else if 조건문으로 나열하는 것 보다 switch를 사용하면 좀 더 간단하게 표현할 수 있다. 특이한 점은 Go 언어의 switch 문은 break를 생략한다.
> break를 사용은 가능하나, 문장 실행을 중단하는 용도로 쓰인다. **fallthrough**를 사용하여 다음 문장을 실행할 수 있다.

## [FizzBuzz 구현하기](src/fizzbuzz/fizzbuzz.go)

> 지금까지 배운 내용으로 FizzBuzz를 구현해보기

## [99병의 맥주 구현하기](src/bottles/bottles.go)

> 지금까지 배운 내용으로 99병의 맥주 구현해보기

## [배열](src/array/array.go)

> Go 언어의 배열은 길이가 고정되어 있고, 배열의 인덱스는 0부터 시작한다. var 키워드로 배열을 생성할 수 있으며. **var 배열명 [길이]자료형**의 형태를 가진다.
> 자료형 뒤에 중괄호로 값을 나열하면 선언과 동시에 초기화도 가능하다. 배열의 길이를 **...** 으로 설정하게 되면, 초기화 크기에 따라 자동으로 크기가 설정된다.

### [배열 순회하기](src/array/traverse/traverse.go)

> 배열 안의 요소를 차례대로 출력하기 위해서는 len 함수를 통해 배열의 길이를 가져오거나, **range** 키워드를 사용할 수 있다.

### [배열 복사하기](src/array/copy/copy.go)

> Go 배열 변수는 배열 전체를 뜻하며 첫번째 요소의 포인터가 아니다.(C 특징) 따라서, 배열을 다른 변수에 대입하면 전체가 복사된다.

## [슬라이스](src/slice/slice.go)

> 슬라이스는 배열과 같지만, 길이가 고정되어 있지 않고 동적으로 증가한다. 배열과는 달리 **레퍼런스 타입**이다. 배열과 다르게 길이를 지정하지 않고 **make** 함수를 통해 공간을 할당해서 값을 넣어야 한다.
> 배열과 마찬가지로 값을 생성하면서 초기화 하려면 중괄호를 사용하면 된다. 또한 배열은 길이와 용량을 설정할 수 있는데, 길이는 접근이 가능한 공간, 용량은 미리 메모리를 할당해둔다.
> 용량은 **cap** 함수로 구할 수 있다.

### [슬라이스에 값 추가하기](src/slice/append/append.go)

> **append** 함수를 이용해 슬라이스 맨 뒤에 값을 추가할 수 있다. **...** 을 사용하여 가변인자를 선언하는데 사용도 가능하지만, 반대로 배열을 가변 인자로 만들기도 한다.

### [레퍼런스 타입](src/slice/reference/reference_type.go)

> 슬라이스는 레퍼런스 타입이다. 내장 배열에 대한 **포인터** 이므로, 슬라이스끼리 대입하면 같은 객체를 참조한다.

### [슬라이스 복사하기](src/slice/copy/copy.go)

> 슬라이스를 복사할때는 **copy** 함수를 사용한다. 복사 될 슬라이스는 반드시 **make** 함수로 공간을 할당해야 한다. 그렇지 않은 슬라이스에 복사가 불가하다.
> 만약 공간보다 슬라이스가 작은 경우 일정 부분이 유실될 수 있다.

### [슬라이스와 용량](src/slice/volume/volume.go)

> 슬라이스는 길이와 용량이 구분되어 있는데. **append**를 통해 값을 추가하면 동적으로 크기가 증가한다. 이 때 동적 배열을 구현하기 위해 용량과 길이를 구분한다. (자체 알고리즘을 통해 슬라이스의 용량을 늘린다.)

### [부분 슬라이스 만들기](src/slice/part/part.go)

> 슬라이스는 기존 슬라이스에서 일정 위치를 지정하여 부분 슬라이스를 만들 수 있다. 시작 인덱스 부터 끝 인덱스까지 일부만 참조한다. 길이가 5인 슬라이스를 모두 참조하려면 **[0:5]**를 사용해야 한다.
> 이렇게 만든 슬라이스도 **레퍼런스 타입**이기에 반드시 주의해야 한다. 배열에도 사용이 가능하다. 다만, 배열도 부분 슬라이스는 레퍼런스 타입이기에 주의해야 한다. 시작 인덱스와 끝 인덱스는 모두 생략 가능하다.
> 마지막으로, 용량도 같이 한번에 지정할 수 있다.

## [맵](src/map/map.go)

> Go는 기본 자료형으로 맵을 지원한다. 선언은 **var 변수명 map[키_자료형]값_자료형** 형태로 선언이 가능하다. 슬라이스와 동일하게 make 함수를 사용하여 공간을 할당해야 한다. 맵 선언과 동시에 사용하면 자료형이 생략 가능하다.
> 중괄호를 사용하여 선언과 동시에 초기화가 가능하다.

### [맵에 데이터 저장하고 조회하기](src/map/get/get.go)

> 맵에 데이터를 넣으려면 대괄호 안에 키를 지정하고 값을 대입한다. 마찬가지로 값을 조회할 때도 대괄호에 키를 지정하면 된다.
> 맵에 데이터가 있는지 검사는 값을 조회한 뒤 2번째 bool 리턴값을 활용하여 검사할 수 있다.

### [맵 순회하기](src/map/traverse/traverse.go)

> 배열, 슬라이스와 마찬가지로 **range** 키워드를 사용한다. 

### [맵에서 데이터 제거하기](src/map/delete/delete.go)

> 맵에서 값을 제거하려면 **delete** 함수를 사용한다. 

### [맵 안에 맵 만들기](src/map/inmap/map_in_map.go)

> 맵 값 안에는 일반 자료형 뿐 아니라 맵 자체도 들어갈 수 있다. 

## [함수](src/function/function.go)

> 함수는 **func 함수명() {}** 형태로 선언한다. 함수 정의를 시작한 줄에서 중괄호가 시작된다. C와는 다르게 함수를 먼저 정의하지 않아도 된다.

### [매개 변수와 리턴값 사용](src/function/arguments/arguments.go)

> 매개 변수와 리턴값은 **func 함수명(매개변수명 자료형..) 리턴_자료형 {}** 형태로 선언한다. 
> Go 에서는 리턴 값에 이름을 지정할 수 있다. 리턴 값 변수를 사용할 때는 return 뒤에 리턴할 변수를 지정하지 아니한다. 

### [리턴 값 여러 개 사용하기](src/function/many/many_return_value.go)

> Go는 함수에서 값을 여러 개 리턴할 수 있다. **func 함수명(매개변수명 자료형..) (리턴_자료형, 리턴_자료형) {}** 형태로 선언한다.

### [가변인자 사용하기](src/function/variadic/variadic.go)

> 함수의 매개 변수가 개수가 정해져 있지 않고 유동적으로 변하는 형태를 가변인자라고 한다. 가변인자는 **...** 으로 만들 수 있다.

### [재귀호출 사용하기](src/function/recursive/recursive.go)

> 자기 자신을 다시 호출하는 재귀 함수를 만들 수 있다.

### [함수를 변수에 저장하기](src/function/variable/function_save_to_variable.go)

> 함수를 변수에 저장할 수 있다. 함수를 정의한 뒤 선언한 변수에 대입하기만 하면 된다. 변수 뿐 아니라 슬라이스, 맵에도 함수를 저장할 수 있다.

### [익명 함수 사용하기](src/function/anonymous/anonymous.go)

> Go는 이름이 없는 함수를 정의한 후 바로 호출할 수 있다. 익명 함수는 코드 양을 줄일 수 있다.

## [클로저](src/closure/closure.go)

> Go는 클로저를 지원한다. 함수 안에서 함수 선언이 가능하며, 밖 함수에 선언 된 변수에도 접근이 가능하다.
> 기본적으로 지역 변수는 함수 실행이 끝나면 소멸되지만, 클로저를 사용하게 되면 지역 변수가 소멸되지 않고 포획되었다가, 나중에 함수를 호출할 때 마다 계속 사용할 수 있다.

## [지연 호출 사용하기](src/defer/defer.go)

> 특정 함수를 현재 함수 끝나기 직전에 실행하는 기능이다. try - finally 구문과 비슷하지만, 문법이 좀 더 간단하다.
> **defer 함수명** 형태로 사용하며, 실행되는 순서는 스택과 동일하다. 즉 맨 나중에 지연 호출한 함수가 먼저 실행된다.

## [패닉과 복구 사용하기](src/panic/panic.go)

> 프로그램이 잘못되어 에러가 발생된 뒤 종료되는 상황을 패닉이라고 한다. 이런 오류는 런타임 뿐 아니라 **panic** 함수를 사용하여 사용자가 직접 오류를 발생시킬 수 있다. 
> 문법적인 오류는 아니지만, 로직에 따라 에러로 처리하고 싶을 때 사용한다. 이런 패닉을 **recover** 함수로 패닉이 발생 했을 때 예외처리가 가능하다.
> recover는 반드시 지연호출 해야함에 유의하라.

## [포인터](src/pointer/pointer.go)

> C, C++ 처럼 메모리 주소 표현 포인터를 지원한다. C와 달리 *를 자료형 앞에 붙인다. 포인터 변수는 초기화하면 nil로 설정된다.
> 빈 포인터 변수는 바로 사용이 불가하고 **new** 함수로 메모리를 할당해야 한다. 해당 함수는 지정한 자료형의 크기에 맞는 메모리 공간을 할당한다.
> 변수를 사용할 때 *를 붙이면 역참조가 된다. 일반 변수에 참조(&)를 사용하면, 포인터 변수에 대입할 수 있다. Go는 메모리 주소를 직접 대입하거나, 포인터 연산을 허용하지 않는다.

### [함수에서 포인터형 매개변수 사용하기](src/pointer/function/function.go)

> 함수의 매개변수에 값을 넘겨줄 때 포인터형 변수를 사용할 수 있다. 일반 자료형을 사용하면 값이 복사되므로 함수 바깥에 영향을 줄 수 없지만, 포인터형 매개변수는 영향을 줄 수 있다.

## [구조체 사용하기](src/struct/struct.go)

> 여러 변수를 담을 수 있는 구조체를 지원한다. **type 구조체명 struct {}** 형태로 사용한다. 중괄호 안에는 필드를 정의한다. 자료형이 같은 필드는 한줄로 나열이 가능하다.
> 구조체는 일반 변수를 선언하는것과 같이 사용이 가능하다. 포인터에 메모리 공간을 할당할 수 있으며, 생성할 때 값을 초기화 할 수 있다.

### [구조체 생성자 패턴 사용하기](src/struct/constructor/constructor.go)

> new 함수로 구조체의 메모리를 할당하는 동시에 값을 초기화 할 수 없다. 하지만 이러한 패턴을 사용하여 다른 언어의 생성자를 흉내낼 수 있다.

### [함수에서 구조체 매개변수 사용하기](src/struct/function/function.go)

> 보통 함수의 매개변수에 구조체 포인터를 받는다. 일반 인스턴스로 넘기게 되면 값이 모두 복사되므로 메모리 상 이점이 없다. 

### [구조체에 메서드 연결하기](src/struct/method/method.go)

> Go 언어에는 클래스가 없는 대신 구조체에 메서드를 연결할 수 있다. **func (리시버명 구조체타입) 함수명() 리턴값_자료형 {}** 형태로 생성한다.
> 메서드 안에서 리시버 변수로 인스턴스의 값에 접근할 수 있다. 그리고 구조체에 점을 사용하여 연결된 메서드를 호출한다. 리시버 변수는 구조체의 인스턴스 포인터가 들어온다.
> 포인터와 일반 구조체 모두로 받을 수 있지만, 원본 데이터를 수정할 수 있는 특성도 각각의 형태에 맞게 적용된다.

### [구조체 임베딩 사용하기](src/struct/embedding/embedding.go)

> Go 언어에는 클래스가 없기에 상속 또한 없다. 하지만 임베딩을 사용하면 상속과 같은 효과를 낼 수 있다. 구조체 안에 필드명 없이 타입만 선언하면 구조체가 타입을 포함하는 관계가 된다. 
> 부모 구조체의 메서드 이름과 중복된다면 상속 과정의 맨 아래 메서드가 호출된다.

## [인터페이스 사용하기](src/interface/interface.go)

> 인터페이스는 메서드의 집합이다. 메서드 자체를 구현하지는 않는다. 인터페이스를 선언하는 방법은 변수와 같다. 다른 자료형과 동일하게 모두 사용 가능하다.
> 선언하면서 동시에 초기화 하려면, 인터페이스 생성자에 매개 변수로 타입을 넣어주면 된다.

### [덕 타이핑](src/interface/ducktyping/ducktyping.go)

> 덕 타이핑은 실제 타입은 상관하지 않고 구현된 메서드로만 타입을 판단하는 방식이다. 어떠한 객체든 상관 없이 꽥, 깃털 메서드만 가졌다면 모두 같은 인터페이스 처리할 수 있다.
> 타입 없이 특정 인터페이스를 구현하는지 검사하려면, **interface{}(인스턴스).(인터페이스)** 형태로 사용할 수 있다. 첫번째 리턴 값은 검사했던 인스턴스, 두번째 리턴은 인터페이스를 구현하고 있는지 여부다.

### [빈 인터페이스](src/interface/empty/empty.go)

> 인터페이스에 아무 메서드도 정의되어 있지 않으면 모든 타입을 저장할 수 있다. 함수의 매개 변수를 arg interface{} 처럼 지정하여 빈 인터페이스를 받게 할 수 있다.
> 이렇게 모든 타입을 받게 되면, arg.(type)을 통해 저장 된 타입을 가져올 수 있다. 빈 인터페이스 변수는 그대로 사용할 수 없고, arg.(int) 등 타입을 원하는 형태로 변경해야 한다.
> 일반 자료형 뿐 아니라 구조체 인스턴스 및 포인터도 빈 인터페이스로 넘길 수 있다.

## [고루틴 사용하기](src/goroutine/goroutine.go)

> 고루틴은 함수를 동시에 실행시키는 기능이다. 다른 언어의 스레드 생성보다 간단하며 리소스를 적게 사용한다. 함수를 호출할 때 앞에 **go** 키워드를 붙이면 고루틴으로 실행된다.
> 고루틴으로 실행하면 다음 함수가 바로 실행된다. 고루틴을 종료하려면 함수에서 return으로 빠져나오거나, runtime.Goexit 함수를 사용하면 된다. 리턴값은 사용할 수 없다.

### [멀티코어 활용하기](src/goroutine/multicore/multicore.go)

> Go는 기본적으로 코어를 한개만 사용하도록 설정되어 있다. 모든 CPU의 코어를 사용하게 하려면 별도로 설정해야 한다. **runtime.NumCPU** 함수로 현재 시스템의 CPU 코어 갯수를 구한 뒤 runtime.GOMAXPROCS 함수에 설정해주자.
> 이렇게 하면 모든 코어에서 고루틴을 실행할 수 있다.

### [클로저를 고루틴으로 실행하기](src/goroutine/closure/closure.go)

> 함수 안에서 클로저를 정의한 뒤 고루틴으로 실행할 수 있다. 하지만 클로저를 고루틴으로 실행할 때 반복문 안에서의 변수 사용에 주의해야 한다. 일반 클로저는 반복문 안에서 순서대로 실행되지만,
> 고루틴으로 실행한 클로저는 반복문이 끝난 뒤에 고루틴이 실행된다. 따라서, 클로저를 고루틴으로 실행할 때 반복문에 의해 바뀌는 변수는 반드시 매개변수로 넘겨준다.

## [채널 사용하기](src/channel/channel.go)

> 채널은 고루틴끼리 데이터를 주고 받고 실행 흐름을 제어한다. 모든 타입을 채널로 사용할 수 있다. 채널은 **값이 아닌 레퍼런스 타입**이다.
> 채널은 **make(chan 자료형)** 형태로 만들 수 있다. 채널은 반드시 make 함수로 공간을 할당해야 하고, 이렇게 생성하면 **동기 채널**이 생성된다. 
> 채널을 매개변수로 받는 함수는 반드시 go 키워드를 사용하여 고루틴으로 실행해야 한다. 채널 변수에는 = 로 값을 대입하지 않고 <- 연산자를 사용한다. 
> 같은 연산자를 통해 채널에서 값을 뽑아올 수 있다. **<- 채널명** 형태로 사용하면, 채널에서 값이 들어올 때 까지 대기한다. 그리고 채널에 값이 들어오면 대기를 끝내고 다음 코드를 실행한다.

### [동기 채널](src/channel/sync/sync.go)

> 동기 채널은 값을 보내면 다른 곳에서 값을 꺼낼 떄 까지 대기한다. 따라서, 다음 구문이 실행되지 않는다. 보내는 쪽에서는 값을 받을 때 까지 대기하고, 받는 쪽에서는 채널에 값이 들어올 때 까지 대기하는 특성을 가진다.

### [채널 버퍼링](src/channel/buffering/buffering.go)

> 채널에 버퍼를 한개 이상 설정하면 비동기 채널이 생성된다. 비동기 채널은 보내는 쪽에서 버퍼가 가득차면 실행을 멈추고 받는 쪽에서는 버퍼에 값이 없으면 대기한다.

### [range와 close 사용하기](src/channel/range/range.go)

> 채널에 range 키워드를 사용하면 채널이 닫힐 때 가지 반복하면서 값을 꺼낸다. 이미 닫힌 채널에 값을 보내면 패닉이 발생함에 유의하라. 채널이 열려있고 값이 들어오지 않으면, range loop는 계속 대기한다.
> 만약 채널에 값이 들어오면 그때부터 실행된다. 정리하면, 채널에 값이 몇개나 들어올 지 모를 때 사용한다. 

### [전용 채널 사용하기](src/channel/receiveonly/receiveonly.go)

> 보내기 전용 채널과 받기 전용 채널이 있는데, 이러한 채널은 값의 흐름이 한방향으로 고정 된 채널이다. 보내기 및 받기 전용 채널은 채널 앞뒤로 <- 연산자를 붙여서 만든다. 채널 키워드를 기준으로 앞에 붙으면 받기, 뒤에 붙으면 보내기 전용 채널이다.

### [셀렉트 사용하기](src/channel/select/select.go)

> 셀렉트는 여러 채널을 손쉽게 사용할 수 있도록 제공하는 분기문이다. switch와 비슷하지만, 검사할 변수를 지정하지 않으며 각 채널에 값이 들어오면 해당 케이스가 실행된다.
> (close로 함수를 닫았을 때도 실행된다.) default 케이스도 지정 가능하며, 케이스에 지정된 채널에 데이터가 들어오지 않으면 즉시 실행된다. 단 적절한 처리를 하지 않으면 모든 코어를 점유하므로 주의하라.

## 동기화 객체 사용하기

> Go에서는 채널 이외에도 고루틴의 실행 흐름을 제어하는 동기화 객체를 제공한다. Mutex, RWMutex, Cond, Once, Pool, WaitGroup, Atomic 등이 있다.

### [뮤텍스 사용하기](src/sync/mutex/mutex.go)

> 뮤텍스는 여러 고루틴이 공유하는 데이터를 보호할 때 사용한다. 뮤텍스를 할당 한 후 고루틴에서 lock, unlock을 사용하여 처리할 수 있다. 뮤텍스는 반드시 짝을 맞춰야하며, 맞지 않으면 데드락이 발생할 수 있음에 주의하라.

### [읽기, 쓰기 뮤텍스 사용하기](src/sync/rwmutex/rwmutex.go)

> 읽기, 쓰기 뮤텍스는 읽기와 쓰기 동작을 나누어 잠금을 걸 수 있다. 읽기 락은 서로를 막지 않지만, 값이 바뀌면 안되므로 쓰기는 막는다, 쓰기 락은 둘 다 허용되지 않으므로 뮤텍스와 같이 동작한다.
> 중요한 쓰기 작업을 할 때 다른 곳에서 이전 데이터를 읽지 못하도록 방지하거나, 데이터가 바뀌는 상황을 방지할 때 사용한다. 특히, 읽기 동작이 많을 때 유리하다.

### [조건 변수 사용하기](src/sync/cond/cond.go)

> 조건 변수는 대기하고 있는 객체를 깨울 때 사용한다. 고루틴 안에서 wait 함수를 사용하면 대기하고, 그 고루틴을 깨울 때는 signal 함수를 사용한다. 모든 고루틴을 깨우는 broadcast 함수도 존재한다.

### [함수를 한번만 실행하기](src/sync/once/once.go)

> once 를 사용하면 함수를 한번만 실행할 수 있다. Once를 할당 한 이후 Do 함수로 사용한다. 한번만 실행되는 특성이 있으므로 복잡한 반복문 안에서 각종 초기화를 할때 유리하다.
 
### [풀 사용하기](src/sync/pool/pool.go)

> 풀은 객체를 사용한 후 보관해 두었다가 다시 사용하게 해주는 기능이다. 객체를 반복해서 할당하면 메모리 사용량이 증가하고, 메모리를 해제해야하는 GC에도 부담이 간다.
> 풀은 일종의 캐시다. 여러 고루틴에서 동시에 사용이 가능하다.

### [대기 그룹 사용하기](src/sync/waitgroup/waitgroup.go)

> 대기 그룹은 고루틴이 모두 끝날 때까지 기다릴때 사용한다. Add를 통해 설정한 값과 Done 함수의 호출 횟수가 동일하지 않으면 패닉이 발생한다.

### [원자적 연산 사용하기](src/sync/atomic/atomic.go)

> 원자적 연산은 더이상 쪼갤 수 없는 연산을 의미한다. 따라서 여러 스레드, CPU, 코어에서 수정할 때 서로 영향받지 않고 안전하게 연산할 수 있다. (보통 CPU의 기능을 사용하여 구현된다.)

## [리플렉션 사용하기](src/reflection/reflection.go)

> 리플렉션은 런타임 시점에 인터페이스나 구조체 등 타입 정보를 얻어내거나 결정하는 기능이다. 

### [구조체 태그 가져오기](src/reflection/tag/tag.go)

> 구조체에는 백틱을 통해 태그를 설정할 수 있는데, 리플렉션으로 구조체의 태그도 가져올 수 있다.

### [포인터와 인터페이스의 값 가져오기](src/reflection/pointer/pointer.go)

> 포인터는 일반 변수와는 다르게 값을 가져오려면 reflect.valueOf 함수로 값 정보를 얻어온 뒤 다시 Elem 함수로 값 정보를 가져와야 한다. 
> 그리고 변수에 타입의 맞는 함수를 통해 값을 가져온다.

## [동적으로 함수 생성하기](src/dynamicfunction/dynamicfunction.go)

> 리플렉션으로 동적으로 함수를 생성할 수 있다. 이렇게 만들 수 있는 함수는 반드시 reflect.Value를 매개변수로 받으며, 반환형도 동일해야 한다.
> 이렇게 사용하면 함수는 여러 종류지만 실제로는 하나의 구현으로 처리할 수 있다.
 
## [인터넷 소스 저장소의 패키지 사용하기](src/onlinepackage/onlinepackage.go)

> **import** 키워드는 로컬에 있는 패키지 뿐만 아니라 인터넷에 올라와있는 소스 저장소에 있는 패키지도 사용이 가능하다.
> go get 명령어를 통해 패키지를 받아오고, 그 받아온 패키지를 임포트하여 사용할 수 있다. 

## [패키지 만들기](src/makepackage/makepackage.go)

> 패키지 안에서 첫 글자를 영문 소문자로 지정하면 패키지 안에서만 사용 가능하고, 대문자로 지정하면 패키지 외부에서 사용이 가능하다.
> 패키지를 컴파일하여 라이브러리로 만들려면 **go install** 명령어를 사용하면 된다.

## [문서화 하기](src/documentation/documentation.go)

> Go는 패키지를 만들면서 문서화가 가능하다. package 키워드와 함수 정의부 위에 // 주석 형태로 작성하면,
> **godoc** 명령어를 통해 패키지 설명을 볼 수 있다. godoc -http:6060 형태로 실행하면, 웹 서버에서 문서를 볼 수 있다.

## [출력 함수 사용하기](src/outputfunction/outputfunction.go)

> Go fmt 패키지에서는 표준 출력 함수로 다음 함수를 제공한다. (Print, Println, Printf)
> Println은 표준 출력에 값을 한 줄로 출력한다. 변수를 여러개 지정해도 모두 한줄로 출력된다. Print 함수는 동일하지만 새 줄로 넘어가지 않고 그 자리에 출력한다.
> Printf는 출력할 형식을 지정하고 출력할 값을 형식 지정자를 통해 지정할 수 있다.

## [입력 함수 사용하기](src/inputfunction/inputfunction.go)

> fmt 패키지에는 Scan, Scanln, Scanf 입력 함수를 지원한다. Scan은 입력 받은 문자, 숫자를 저장하고, 레퍼런스 형태로 넣는다.
> Scanln은 한 줄에서 공백으로만 값을 구분하며 엔터키를 입력하여 새줄로 넘어가면 종료된다. Scanf는 형식을 지정하여 입력 받을 형식을 설정한다.

## [문자열 입출력 함수 사용하기](src/stringfunction/stringfunction.go)

> 표준 출력과 표준 입력 뿐 아니라 변수를 문자열로 만들거나 문자열에서 변수로 가져올 수 있다. fmt 패키지에는 Sprint, Sscan 과 같은 문자열 입출력 함수가 존재한다.

## [파일 입출력 함수 사용하기](src/filefunction/filefunction.go)

> 값을 파일로 저장하거나 파일에서 변수로 값을 가져올 수 있다. **os 패키지**에는 Create, Open, Close 등의 함수를 제공한다. 이후 fmt 패키지에서 제공하는 Fprint, Fscan 등을 사용하여 파일을 생성하거나 열 수 있다.

## 유니코드와 UTF-8 함수 사용하기

> 유니코드 패키지에는 한글을 처리할 수 있는 여러 방안을 제공한다.

### [유니코드 함수 사용하기](src/unicode/unicode/unicode.go)

> 유니코드는 전 세계 모든 문자를 표현할 수 있다. 문자를 종류별로 처리할 수 있도록 범위가 지정되어 있다. unicode 패키지에는 여러 함수를 제공한다.
> 그 외에도 범위 테이블을 제공한다. 한글인지, 한자인지, 일본어인지 등 간단하게 **Is**함수를 통해 구분할 수 있다. 유니코드 패키지에는 이외에도 특성을 가지고 비교할 수 있는 함수들을 지원한다.

### [UTF-8 함수 사용하기](src/unicode/utf8/utf8.go)

> UTF-8은 유니코드를 인코딩하는 방식 중 하나다. Go에서는 주로 사용한다. 영문자를 저장할때는 문자 하나가 1바이트를 차지한다, 하지만 한글, 한자, 일본어 같은 문자는 1바이트로 저장이 불가하기에 보통 2바이트로 저장한다.
> UTF-8은 가변 길이 문자 인코딩 방식이라 문자를 저장할 때 1~4바이트를 사용하며, 한글은 3바이트를 사용한다. 이러한 이유로 바이트를 슬라이스하는데 주의하라, 한글은 3바이트가 온전하지 못하면 표시되지 않을 수 있다.

## 문자열 처리하기

> Go는 기본 라이브러리에서 다양한 문자열 처리 함수를 제공한다. 이 함수들을 조합해서 사용하는 것 만으로도 코드를 간단하게 유지할 수 있다.

### [문자열 검색하기](src/manipulatestring/search/search.go)

> 문자열을 처리할 때 주로 사용하는 기능은 문자열 검색이다. strings 패키지에서는 많은 수의 문자열 검색 함수를 제공한다.

### [문자열 조작하기](src/manipulatestring/manipulate/manipulate.go)

> 문자열을 연결하고, 쪼개고, 반복하고 바꾸는 등 문자열을 조작하는 함수도 지원한다.

## [문자열 변환 함수 사용하기](src/transformation/transformation.go)

> 정수, 실수, 불 등을 문자열로 변환하는 함수도 지원한다. **strconv** 패키지에서 제공하는 문자열 변환 함수들이 있다.

## [정규표현식 사용하기](src/regexp/regexp.go)

> 정규표현식은 일정한 규칙을 가진 문자열을 표현하는 방법이다. 많은 문자열 속에서 특정한 규칙을 가진 문자열을 추출하거나 정해진 규칙에 맞는지를 판별한다.
> Go에서는 **regexp** 패키지에서 지원한다.

### [이메일 주소 검사하기](src/regexp/email/email.go)

> 이메일 주소가 정상적인지 검사하는 때가 많다. 이때 이메일의 규칙은 **계정@도메인.최상위도메인** 형식이다.

### [문자열 검색하기](src/regexp/string/string.go)

> regexp 패키지에는 정규표현식 문자열 검색 함수 또한 지원한다.

### [문자열 조작하기](src/regexp/replace/replace.go)

> regexp 패키지에는 정규표현식을 통한 문자열 조작 함수를 지원한다.

## 파일 처리하기

> Go는 다양한 방법으로 파일을 읽고 쓸 수 있다.

### [파일 쓰기](src/file/write/write.go)

> os 패키지에는 파일 함수와 파일 쓰기 함수를 제공한다. 파일을 열었으면 반드시 닫는 코드를 작성해야 한다. **defer**를 통해 닫는 코드를 미리미리 선언하자.

### [파일 읽기](src/file/read/read.go)

> os 패키지에는 파일 열기, 파일 읽기 등 함수를 제공한다. **os.Open** 으로 파일을 열었을 경우에는 파일을 읽기만 가능하다.

### [파일 읽기 쓰기](src/file/readwrite/readwrite.go)

> **os.OpenFile** 함수를 사용하면 읽기/쓰기 모드로 파일을 열고 읽고 쓸 수 있다.

### [ioutil 패키지 사용하기](src/file/ioutil/ioutil.go)

> **ioutil** 패키지를 사용하면 좀 더 간단하게 파일을 읽고 쓸 수 있다. WriteFile과 ReadFile 함수를 제공한다.
> 파일명, 데이터, 파일모드 등만 지정하면 간단하게 사용할 수 있다. 
> 현재는 ioutil은 deprecate 되었고, os.ReadFile, os.WriteFile로 대체되었다.

## 입출력 인터페이스 사용하기

> Go는 io.Reader, io.Writer 인터페이스를 활용하여 다양한 방법으로 입출력을 지원한다. 화면에 출력하는 부분부터 파일, 문자열까지 같은 인터페이스로 처리할 수 있다.
> 어떤 구조체든, 매개 변수로 바이트 슬라이스를 받고, 정수와 에러 값을 리턴하는 Read 함수를 가지고 있으면 io.Reader 인터페이스를 따른다고 볼 수 있다. Write도 마찬가지다.

### [파일 처리하기](src/io/file/file.go)

> bufio는 Buffered I/O를 뜻하며 io.Reader, io.Writer 인터페이스를 받는다. bufio.NewWriter 함수에 file 인스턴스를 넣으면 io.Write 인스턴스를 따르는 쓰기 인스턴스를 리턴한다.
> bufio를 사용하므로, 파일에 바로 저장하지 않고 임시 공간에 쌓아둔다. 버퍼의 내용을 완전히 파일에 저장하기 위해서는 flush 메서드를 사용한다.

### [문자열을 파일로 저장하기](src/io/stringtofile/stringtofile.go)

> strings.NewReader를 통해 String을 io.Reader 형태로 변경할 수 있다. 이를 사용하여 문자열을 파일에 보다 쉽게 저장할 수 있다.

### [문자열을 화면에 출력하기](src/io/stringtoconsole/stringtoconsole.go)

> io.Reader를 그대로 화면에 출력할 수 있다. os.Stdout도 io.Writer 인터페이스를 따르기 때문에 io.Reader를 복사해주면 그대로 출력된다.

### [기본 입출력 함수 사용하기](src/io/stdio/stdio.go)

> fmt.Fscanf는 파일 인스턴스 뿐 아니라 io.Reader를 따르는 모든 인스턴스를 사용할 수 있다. 따라서 문자열로 만든 io.Reader도 사용할 수 있다. 
> 반대로 fmt.Fprintf 함수는 io.Writer를 따르는 모든 인스턴스를 사용할 수 있다. 이러한 인터페이스 사용법에 익숙해지자.

### [읽기, 쓰기 인터페이스를 함께 사용하기](src/io/readwriter/readwriter.go)

> 지금까지는 각각 사용했는데. 이번에는 **io.ReadWriter** 인터페이스를 통해 읽기/쓰기를 처리해보자. 해당 인터페이스는 Reader, Writer를 임베딩 하고있다.

## [JSON 문서 사용하기](src/json/json.go)

> JSON 형식은 인터넷 기반으로 발전하며 널리 쓰이고 있다. Go 언어에서는 encoding/json 패키지에서 json을 제공한다. 
> 문서를 읽기 위해서는 **Unmarshal** 함수를 사용하면 된다. JSON 문서의 데이터를 저장할 공간이 필요한데, 간단하게 맵으로 사용할 수 있다.
> 반대로 데이터를 JSON 형태로 변경하기 위해서는 **Marshal, MarshelIndent** 함수를 사용하여 변환할 수 있다. (indent는 사람이 보기 좋게 만들어 준다.)

### [구조체 활용하기](src/json/struct/struct.go)

> 구조체를 사용한 복잡한 형태의 JSON 문서 또한 지원이 가능하다. 구조체 또한 JSON 형태로 변환이 가능하다. `json:"태그"` 형식으로 키 이름을 지정하면, 
> JSON 문서를 읽거나 만들 때 해당 명칭으로 사용이 가능하다.

### [JSON 파일 사용하기](src/json/file/file.go)

> JSON 파일을 만들고 읽는 것 또한 가능하다. os.WriteFile과 os.ReadFile을 적절히 사용하여 처리가 가능하다.

## [압축 사용하기](src/compression/compression.go)

> Go는 다양한 압축 알고리즘을 패키지로 제공한다. **compress** 패키지로 제공한다. 압축 알고리즘은 모두 io.Reader, io.Writer 인터페이스를 따르므로, 같은 방법으로 압축 및 해제가 가능하다.
> bzip, zlib, flate, lzw 등을 지원한다.

## 암호화 사용하기

> Go는 다양한 암호화 알고리즘을 패키지로 제공한다. 해시 알고리즘, 대칭키 알고리즘, 공개키 알고리즘 등 대표적인 알고리즘 방식을 알아보자.

### [해시 알고리즘 사용하기](src/crypto/hash/hash.go)

> 대표적인 해시 알고리즘인 **sha512** 알고리즘을 사용해보자. 간단하게 함수를 사용하여 변경할 수 있다.

### [AES 대칭 키 알고리즘 사용하기](src/crypto/symmetrickey/symmetrickey.go)

> 대표적인 대칭 키 알고리즘인 **AES** 알고리즘을 사용해보자. AES는 블록 암호화 알고리즘으로 키와 암호화할 데이터의 크기가 일정해야 한다. 길이가 긴 데이터는 16바이트씩 잘라서 암호화 하면 되는데,
> 이럴 경우 보안에 취약하다. 이러한 방식을 ECB라고 한다. 안전한 암호화 운용 방식 중 하나인 CBC에 대해 알아보자.

### [RSA 공개 키 알고리즘 사용하기](src/crypto/publickey/publickey.go)

> 대칭 키는 암호 키가 유출되면 암호화 된 데이터를 모두 풀 수 있다. 특히 네트워크로 암호키를 주고 받으면 노출 될 위험이 커진다. 이러한 경우에는 공개키 알고리즘을 많이 사용한다.
> 공개 키 알고리즘은 암호화 할 때 공개키를 사용하고, 복호화 할 때 개인키를 사용한다. 공개키를 외부에 공개하여 전달한 뒤 암호화를 하는 방식이다. RSA 알고리즘은 메시지를 서명하고 인증할 수 있다.

## [정렬 활용하기](src/sort/sort.go)

> 데이터를 처리하다 보면 정렬을 할일이 많다. Go에서는 다양한 데이터를 정렬하기 위해 **sort** 패키지를 통해 제공한다.

### [정렬 인터페이스를 이용하여 구조체 정렬하기](src/sort/interface/interface.go)

> 정렬 인터페이스를 구현하여 구조체가 담긴 슬라이스를 정렬할 수 있다. 또한, 같은 객체를 임베디드 하여 또 다른 형태의 정렬도 구현이 가능하다. 임베디드 하여 필요한 메서드만 다시 구현하는 방식으로
> 상속을 통해 별도의 정렬을 구현한다.

### [정렬 키로 정렬하기](src/sort/key/key.go)

> 인터페이스를 교체하는 방법 말고 정렬 키 함수를 사용하는 방법도 있다. 함수 타입을 만들어 간단하게 정렬해야하는 키 메서드를 교체하여 사용할 수 있다. 
> Go는 하나의 데이터 타입을 여러 인터페이스로 바꿔가면서 OOP를 구현한다.

## 컨테이너 사용하기

> 다양한 데이터를 다룰 때 자료구조를 필수적으로 사용한다. 자료구조를 일일이 구현하기엔 시간과 노력이 필요한데, Go 에서는 기본 자료구조 패키지 **container**를 제공한다.

### [연결 리스트 사용하기](src/container/list/list.go)

> Go 언어에서는 이중 연결 리스트를 지원한다. 따라서 앞, 뒤 양쪽 방향으로 순회할 수 있다.

### [힙 사용하기](src/container/heap/heap.go)

> **container/heap** 패키지에서 힙을 제공한다. 힙은 다른 자료구조와 다르게 heap.Interface를 구현해야 한다. 힙은 완전 이진 트리를 사용한 구조며, 최대 힙과 최소 힙이 있다.

### [링 사용하기](src/container/ring/ring.go)

> **container/ring** 패키지에서 링을 제공한다. 링은 원형으로 연결 된 이중 연결 리스트다. 따라서 처음과 끝이 없고 nil을 가리키는 노드가 없다. 

## TCP 프로토콜 사용하기

> Go는 기본 패키지에서 다양한 네트워크 프로토콜을 제공한다, 그 중 TCP 프로토콜은 실시간 처리가 중요한데에 많이 사용된다. 흔히 우리가 사용하는 HTTP, FTP 등도 Layer 4 에서는 TCP 기반으로 구현되어 있다.

### [서버 작성하기](src/tcp/server/server.go)

> 클라이언트에서 보낸 메시지를 다시 응답하는 TCP 서버를 만들어 보자. **net** 패키지에서 TCP 함수를 제공한다. 클라이언트와 연결되면 Accept 함수를 통해 커넥션이 반환된다. 해당 커넥션을 통해 클라이언트와 송수신이 가능하다.
> 보통 데이터를 주고 받을 때에는 패킷의 최대 크기를 약속한다. (여기서는 4096)

### [클라이언트 작성하기](src/tcp/client/client.go)

> net.Dial 함수를 통해 연결 방식과 연결 서버를 지정한다. 서버에 연결되면 net.Conn이 반환되고, 데이터를 보내서나 받을 수 있다. 지연 호출을 사용하여 연결을 반드시 닫아주자. 

## RCP 프로토콜 사용하기

> RCP는 **원격 프로시저 호출**을 의미하며, 원격에서 함수를 실행하는 기술이다. 서버에 있는 함수를 로컬에서 실행할 수 있기 때문에 매우 편리하다. 기본 패키지로 지원하므로 간단하게 구현이 가능하다.

### [서버 작성하기](src/rpc/server/server.go)

> 원격에서 덧셈 함수를 호출하여 두 수를 더해보자. **net/rpc** 패키지에서 관련 함수를 제공한다. RPC 서버에 함수를 등록하려면 함수만으로는 안되고, 구조체나 일반 자료형과 같은 타입에 메서드 형태로 구성되어 있어야 한다.
> 빈 자료형이나 빈 구조체로 정의해도 된다. RPC로 호출되는 함수는 매개변수 두 개와 error 형 리턴 값 형식이다. 매개변수를 위한 구조체와, 리턴값을 위한 구조체를 정의해서 매개변수에 넣어줘야 한다.

### [클라이언트 작성하기](src/rpc/client/client.go)

> 서버와 동일한 매개변수와 리턴값 구조체를 정의해야 한다. 이 경우 공유하는 패키지로 만들면 편하게 사용할 수 있다. **client.Call**을 사용하면 동기 방식으로 작동하고. **client.Go**를 사용하면 비동기 방식으로 작동한다.
> Done 채널 변수를 통해 값이 들어올 때 까지 기다리면 된다.

## [HTTP 서버 사용하기](src/http/http.go)

> Go는 기본 라이브러리만 가지고 HTTP 서버를 간단하게 만들 수 있다. **net/http** 패키지에서 제공한다. **http.HandleFunc**를 통해 간단하게 경로와 처리할 함수를 설정할 수 있다.

## [명령줄 옵션 사용하기](src/cmdargs/cmdargs.go)

> 프로그램을 만들다 보면 실행할 떄 옵션을 설정하고 싶을 때가 있다. 명령줄에서 간단하게 설정한 옵션을 사용할 수 있다. Go의 메인 함수에는 매개변수가 없다. 명령줄에서 설정한 옵션을 가져오려면,
> **os.Args** 슬라이스를 사용한다. os.Args의 첫번째 요소는 항상 실행파일의 경로고, 두 번째 요소부터 옵션이다. 이를 일일이 분석하여 사용하기엔 매우 어렵다. 편리하게 사용할 수 있도록 **flag** 패키지를 제공한다.

## [에러 처리하기](src/error/error.go)

> fmt.Errorf, log.Fatal 함수 등을 사용해 에러를 처리할 수 있다. 보통 Go 언어의 함수에서는 첫 번째 리턴을 결과로 사용하고, 두 번째 리턴을 에러값으로 사용한다. 
> 에러 타입을 리턴할때는 **fmt.Errorf** 함수를 사용한다. 여기에 형식을 지정하여 에러 메시지를 만든다. 이렇게 발생한 메시지를 **log.Fatal** 함수로 출력한다. 이 때 프로그램이 exit code 1로 종료된다.
> 프로그램을 정상 종료시키지 않고 런타임 오류를 발생시키려면 **log.Panic** 함수를 사용한다. 시간과 에러 문자열을 출력한 뒤 패닉을 발생시킨다.
