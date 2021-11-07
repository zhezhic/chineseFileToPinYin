package main

import (
	"flag"
	"fmt"
	"github.com/Chain-Zhang/pinyin"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func main() {
	//os.Args是一个[]string
	//if len(os.Args) > 0 {
	//	for index, arg := range os.Args {
	//		fmt.Printf("args[%d]=%v\n", index, arg)
	//	}
	//}
	// 解析命令行参数
	var dir string
	flag.StringVar(&dir, "d", "", "文件夹路径")
	flag.Parse()
	if dir == "" {
		flag.Usage()
		return
	}
	exists := DirExists(dir)
	if !exists {
		fmt.Println("此路径不存在")
		return
	}
	//遍历文件夹内所文件
	fileInfos, _ := ioutil.ReadDir(dir)
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}
		fileName := fileInfo.Name()
		fileSuffix :=path.Ext(fileName)
		targetName :=strings.TrimSuffix(fileName,fileSuffix)
		finalName,_ :=pinyin.New(targetName).Split("-").Mode(pinyin.WithoutTone).Convert()
		filePath := dir+"/"+fileName
		newFilePath := dir+"/" + finalName+ fileSuffix
		_=os.Rename(filePath,newFilePath)
	}
	//_ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
	//	//if !info.IsDir() {
	//	//	fileInfo,_ :=os.Stat(path)
	//	//	fmt.Println(fileInfo.Name())
	//	//}
	//	if info.IsDir() {
	//		return nil
	//	}
	//	fmt.Println(info.Name())
	//	return nil
	//})
	//// 遍历文件路径，修改文件名
	//for _, path := range paths {
	//	var chinese string = "4人麻将[星空汉化](JP)[TAB](0.18Mb)"
	//	str, err := pinyin.New(chinese).Split("-").Mode(pinyin.WithoutTone).Convert()
	//	if err != nil {
	//		fmt.Println("修改名称错误")
	//	}
	//	err = os.Rename(path, str)
	//	if err != nil {
	//		return
	//	}
	//}
}

// 判断所给路径文件夹是否存在

func DirExists(path string) bool {
	filInfo, err := os.Stat(path) //os.Stat获取文件信息
	if !filInfo.IsDir() {
		fmt.Println("此路径不是一个文件夹")
		return false
	}
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
