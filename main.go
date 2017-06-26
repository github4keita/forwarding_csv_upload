package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	iconv "github.com/djimenez/iconv-go"
)

// エラーログをはいて、exit(1)
func fatal(err error) {
	if err != nil {
		log.Fatal("Error:", err)
	}
}

func main() {
	log.Println("test")
	// ログファイル名を日付で作成

	// ログファイルの作成と標準出力とログファイル２つに書き込むように設定

	// csvファイル指定
	file, err := os.Open("/Users/keitamatsuo/Desktop/D204420170426100000.csv")
	fatal(err)
	defer file.Close()

	// ファイルをsjisからutf-8へ変換
	// これがだめでsignal killed iconv-goをどうにかしないと
	converter, err := iconv.NewReader(file, "sjis", "utf-8")
	fatal(err)

	// csvリーダー作成
	reader := csv.NewReader(converter)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			// ここのエラーは止めない、ログ書き込み
			log.Println("Error:", err)
		} else {
			// 配列に詰める
			log.Printf("%#v", record)
		}
	}

	// 配列を回し、1行ずつAPIを叩く
}
