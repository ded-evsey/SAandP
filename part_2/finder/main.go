package finder

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Finder struct {  // структура для хранения выбранной директории и хэша
	Homedir string
	Hash    map[string]string
}

func (f Finder) GetFile(w http.ResponseWriter, r *http.Request) bool{
	filename := r.FormValue("q")  // получаем данные с формы
	filePath, ok := f.Hash[filename]  // ищем
	if ok {
		w.Header().Set("Content-Disposition", "attachment; filename="+ filename)
		w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
		http.ServeFile(w,r,filePath)
		return true
	}
	return false

}


func (f Finder) MakeHash() Finder{
	if f.Homedir == "" {
		f.Homedir = "~/Documents/"
	}
	if f.Hash == nil {
		f.Hash = make(map[string]string)
	}
	filepath.Walk(f.Homedir, func(wPath string, info os.FileInfo, err error) error {
		// Обход директории без вывода
		if err != nil {
			return err
		}
		if strings.Contains(wPath, ".idea") {
			return nil
		}
		if strings.Contains(wPath, ".git") {
			return nil
		}
		if strings.Contains(wPath, "lib") {
			return nil
		}
		if strings.Contains(wPath, "venv") {
			return nil
		}
		if strings.Contains(wPath, "__pycache__") {
			return nil
		}
		if info.IsDir() {
			return nil
		}
		f.Hash[info.Name()] = wPath
		return nil
	})
	return f
}
