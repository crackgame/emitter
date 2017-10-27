package main

import (
	"fmt"

	"github.com/crackgame/emitter"

	"github.com/chuckpreslar/emission"
)

func test1() {
	emit := emitter.NewEmitter()
	emit.On("test", func(args ...interface{}) {
		fmt.Println("on test1")

		emit.On("test2", func(args ...interface{}) {
			fmt.Println("on test2")
		})

		emit.Emit("test2", nil)
	})

	emit.Emit("test", nil)
}

func test2() {
	emit := emission.NewEmitter()
	emit.On("test", func(args ...interface{}) {
		fmt.Println("on test1")

		emit.On("test2", func(args ...interface{}) {
			fmt.Println("on test2")
		})

		emit.Emit("test2", nil)
	})

	emit.Emit("test", nil)
}

func test3() {
	emitter := emission.NewEmitter()

	hello := func(to string) {
		fmt.Printf("Hello %s!\n", to)
	}

	count := func(count int) {
		for i := 0; i < count; i++ {
			fmt.Println(i)
		}
	}

	emitter.On("hello", hello).
		On("count", count).
		Emit("hello", "world").
		Emit("count", 5)
}

func test4() {
	emitter := emitter.NewEmitter()

	hello := func(to string) {
		fmt.Printf("Hello %s!\n", to)
	}

	count := func(count int) {
		for i := 0; i < count; i++ {
			fmt.Println(i)
		}
	}

	emitter.On("hello", hello)
	emitter.On("count", count)
	emitter.Emit("hello", "world")
	emitter.Emit("count", 5)
}

func main() {
	test4()
}
