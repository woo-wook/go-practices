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
