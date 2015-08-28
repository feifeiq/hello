package main

import (
	"ext/cli"
	"ext/util/console"
	"ext/util/files"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	client  = new(http.Client)
	ch      = make(chan int)
	strUrl  = cli.Arg("url", "download url.").String()
	saveDir = cli.Arg("dir", "save dir.").String()
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Mkdirs(path string) (isExist bool) {
	fi, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Printf("create dir:%s\n")
		os.MkdirAll(path, 0666)
		return true
	}
	if err == nil {
		return fi.IsDir()
	}

	return false
}
func isDirExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}

	return false
}
func initFile(strUrl, dstDir string) (fd *os.File, err error) {
	Mkdirs(dstDir)
	fpath := path.Join(dstDir, path.Base(strUrl))
	if files.IsExist(fpath) {
		fmt.Printf("open:%s\n", fpath)
		fd, err = os.OpenFile(fpath, os.O_RDWR|os.O_APPEND, os.FileMode(0666))
	} else {
		fmt.Printf("create:%s\n", fpath)
		fd, err = os.Create(fpath)
	}
	return fd, err
}
func GetContentLength(link string) (int64, error) {
	resp, err := http.Get(link)
	checkErr(err)
	defer resp.Body.Close()
	return strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 32)
}

func Download(strUrl, dstDir string) (err error) {
	fd, err := initFile(strUrl, dstDir)
	checkErr(err)
	defer fd.Close()

	fi, err := fd.Stat()
	checkErr(err)
	contentLength, err := GetContentLength(strUrl)
	checkErr(err)

	req := new(http.Request)
	req.ProtoMajor = 1
	req.ProtoMinor = 1
	req.Header = http.Header{}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-", fi.Size(), contentLength))
	req.Method = "GET"
	req.URL, err = url.Parse(strUrl)
	checkErr(err)
	res, err := client.Do(req)
	checkErr(err)
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return fmt.Errorf("Error status %s %s", res.Status, strUrl)
	}
	pw := console.NewProgress(fd, fi.Size(), contentLength)

	defer pw.Close()
	_, err = io.Copy(pw, res.Body)
	return
}

/*获取当前文件执行的路径*/
func GetCurrPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	splitstring := strings.Split(path, "\\")
	size := len(splitstring)
	splitstring = strings.Split(path, splitstring[size-1])
	ret := strings.Replace(splitstring[0], "\\", "/", size-1)
	return ret
}
func init() {
	defer func() {
		if e := recover(); e != nil {
			cli.Usage()
		}
	}()
	cli.Version("0.1.0")
	cli.Parse()
	if *saveDir == "" {
		fpath, err := filepath.Abs(filepath.Dir(os.Args[0]))
		println(fpath)
		checkErr(err)
	}
}
func main() {
	//	*strUrl = "https://static.rust-lang.org/dist/rust-1.0.0-beta.4-x86_64-pc-windows-gnu.msi"

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			return
		}
	}()

	//	if strUrl == nil {
	//		return
	//	}
	//	go func() {
	//		err := Download(*strUrl, *saveDir)
	//		if err != nil {
	//			fmt.Errorf(err.Error())
	//		}
	//		ch <- 1
	//	}()
	//	<-ch

}
