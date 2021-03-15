package main

import "fmt"

/*
	테스트를 하기 위해서는 외부의 세계(side-effect) 로부터 도메인 코드를 분리시켜야 한다
	fmt.Println 이 side-effect 이고 우리가 전달하는 문자열은 우리의 도메인이다
*/
const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return spanishHelloPrefix + name + "!"
	}

	if language == "Korean" {
		return "안녕, " + name + "!"
	}

	return englishHelloPrefix + name + "!"
}

func main() {
	fmt.Println(Hello("World", ""))
}
