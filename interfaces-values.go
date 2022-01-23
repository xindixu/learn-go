package main

import "fmt"

type Person interface {
	greet()
}

type Chinese struct {
	name     string
	location string
}

type American struct {
	name     string
	location string
}

type Japanese struct {
	name     string
	location string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
// value as receiver
func (t Chinese) greet() {
	fmt.Printf("你好,%s。我在%s\n。", t.name, t.location)

}

func (t American) greet() {
	fmt.Printf("Hello, %s. I'm in %s\n.", t.name, t.location)
}

// pointer as receiver
func (t *Japanese) greet() {
	if t == nil {
		fmt.Printf("こんにちわ\n")
		return
	}
	fmt.Printf("こんにちわ、%s。 %sにいます。\n", t.name, t.location)
}

// NOTE: Interface values is like a tuple of a value and a concrete type:
// (value, type)
func describe(i Person) {
	// value, concrete type
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	// Chineses, American, *Japanese implement Person
	var ch Person = Chinese{"笛子", "纽约"}
	var us Person = American{"Xindi", "NY"}
	var jp Person = &Japanese{"奈々子", "ニューヨーク"}
	var noValueJp Person = &Japanese{}

	// non-nil interface: no value, yes concrete type
	var nilJp *Japanese
	var noValueCh Chinese
	// nil interface: no value, no concrete type
	var nilPerson Person

	describe(ch)        // ({笛子 纽约}, main.Chinese)
	describe(us)        // ({Xindi NY}, main.American)
	describe(jp)        // (&{奈々子 ニューヨーク}, *main.Japanese)
	describe(noValueJp) // (&{ }, *main.Japanese)
	describe(nilJp)     // (<nil>, *main.Japanese)
	describe(noValueCh) // ({ }, main.Chinese)
	describe(nilPerson) // (<nil>, <nil>)

	ch.greet()        // 你好,笛子。我在纽约。
	us.greet()        // Hello, Xindi. I'm in NY.
	jp.greet()        // こんにちわ、奈々子。 ニューヨークにいます。
	noValueJp.greet() // こんにちわ、。 にいます。
	nilJp.greet()     // こんにちわ
	noValueCh.greet() // 你好,。我在
	nilPerson.greet() // Error

}
