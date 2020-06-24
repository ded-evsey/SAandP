package main

import (
	"fmt"
	"io/ioutil"
	"log"
)
// Реализовать функцию, которая из указанного каталога строит ассоциативный массив, в котором ключами являются имена файлов, а значениями их размер.
func fileSize(dir string) map[string]int64 {
	var filesList = make(map[string]int64)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		filesList[file.Name()] = file.Size()
	}
	return filesList
}

func main()  {
	fmt.Println(fileSize("/home/user/Documents"))
}