package main

import (
	"flag"
	"fmt"
	yd "github.com/ebproofix/ydfree/internal/control"
	"log"
	"os"
	"path"
	"sync"
	"time"
)

var (
	YANDEX_DIR = "./testing/yd_dir"
	BACKUP_DIR = "./testing/rc_dir"
	MOD        = "print"
	listBakaps []string
	timeNow    = time.Now()
)

const(
	MOD_PRINT = "print"
	MOD_DEL = "del"
)

func init() {
	flag.StringVar(&YANDEX_DIR, "y", YANDEX_DIR, "Путь до каталога \"Яндекс Диска\"")
	flag.StringVar(&BACKUP_DIR, "b", BACKUP_DIR, "Путь до каталога c бекапом")
	flag.StringVar(&MOD, "m", MOD, "print or del")
}

func main() {

	flag.Parse()

	yandex := getDirectoryName(YANDEX_DIR)
	backup := getDirectoryName(BACKUP_DIR)


	var wg sync.WaitGroup

	for _, v := range yandex {
		wg.Add(1)
		go func(){
		errStr := ""
		yandexFullpath:= path.Join(YANDEX_DIR, v)

		err := yd.CheckOnTimeOut(yandexFullpath)
		if err != nil {
				errStr += fmt.Sprintf("%s\n\t%s\n", v, err.Error())
		}
		err = yd.CheckOnNotDeleteWord(yandexFullpath)
		if err != nil {
				errStr += fmt.Sprintf("\t%s\n", err.Error())
		}
		err = yd.CheckOnbackup(yandexFullpath, backup)
		if err != nil {
			errStr += fmt.Sprintf("\t%s\n", err.Error())
		}

		if errStr != "" {
			if MOD == MOD_PRINT {
				log.Println(errStr)
			}
			wg.Done()
			return
		}else if MOD == MOD_PRINT {
			fmt.Println("%s - может быть удален", v)
		}
		if MOD == MOD_DEL {
			os.ReadDir(yandexFullpath)
		}
		wg.Done()
	}()
	}
	 wg.Wait()

}

func getDirectoryName(pathDir string) (dirs []string){
	entryInfo, err := os.ReadDir(pathDir)
	if err != nil {
		log.Fatalf("Не могу прочитать директорию: %s\n%s\n", pathDir, err.Error())
	}
	for _, v := range entryInfo {
		if v.IsDir() {
			dirs = append(dirs, v.Name())
		}
	}
	return
}
