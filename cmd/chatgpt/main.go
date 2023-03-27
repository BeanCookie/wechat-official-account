package main

import chatgpt "cn.lzzz.chatgpt"

func main() {
	err := chatgpt.Run()
	if err != nil {
		panic(err)
	}
}
