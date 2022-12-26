package main

import (
	"fmt"
	"strings"
)

// intCh:結果を送信する送信専用のチャネル
func doubleInt(src int, intCh chan<- int) {
	result := src * 2
	intCh <- result
}

func doubleString(src string, strCh chan<- string) {
	result := strings.Repeat(src, 2)
	strCh <- result
}
// 送受信ペアパターン
func main() {
	// make関数を使ってはじめてchに送受信できるチャネルが格納される
	ch1, ch2 := make(chan int), make(chan string)
	// お役御免になったら、チャネルを閉じる
	defer close(ch1)
	defer close(ch2)

	// goroutineの起動
	go doubleInt(1, ch1)
	go doubleString("hello", ch2)

	// どちらか受信できるほうから値を受信
	for i := 0; i < 2; i++ {
		//  結果をチャネルから受信
		// doubleIntが終わっていない場合は待機する
		select {
		case numResult := <-ch1:
			fmt.Println(numResult)
		case strResult := <-ch2:
			fmt.Println(strResult)
		}
	}
}
