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

func main() {
	test1()
}
