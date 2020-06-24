package main
import (
	"fmt"
	"os"
	"path/filepath"
	s "strings"
)



// Построить список имен файлов, название которых начинается на заданную букву.
// Построить список файлов с заданным «расширением».
func searchFiles(symbol string, fileExemption string)([]string, []string)  {
	var bySymbol []string
	var byExemption []string
	filepath.Walk("/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if s.HasPrefix(info.Name(), symbol){
			bySymbol = append(bySymbol, info.Name())
		}
		if s.HasSuffix(info.Name(), fileExemption){
			byExemption = append(byExemption, info.Name())
		}

		return nil
	})
	return bySymbol, byExemption
}

func main()  {
	bySymbol,byExemption := searchFiles("a", ".png")
	fmt.Println("Files starts by a:", bySymbol)
	fmt.Println("files finished by png", byExemption)
}