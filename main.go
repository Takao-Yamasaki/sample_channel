package main

import (
	"fmt"
)

// intCh:結果を送信する送信専用のチャネル
func doubleInt(src int, intCh chan<- int) {
	result := src * 2
	intCh <- result
}

func main() {
	// make関数を使ってはじめてchに送受信できるチャネルが格納される
	ch := make(chan int)
	// お役御免になったら、チャネルを閉じる
	defer close(ch)

	// goroutineの起動
	go doubleInt(1, ch)
	//  結果をチャネルから受信
	// doubleIntが終わっていない場合は待機する
	result := <-ch

	fmt.Println(result)
}