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
