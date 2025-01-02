package main

import "github.com/osoka34/homework-bot/pkg/utils"

func main() {
	l, err := utils.InitJSONLogger()
    if err != nil {
        panic(err)
    }

    defer l.Sync()




}
