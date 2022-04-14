package main

import (
	"io"
	"net/http"
	"os"
	"path"
)

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}

/*
对resp.Body.Close延迟调用我们已经见过了，在此不做解释。
上例中，通过os.Create打开文件进行写入，在关闭文件时，我们没有对f.close采用defer机制，
因为这会产生一些微妙的错误。许多文件系统，尤其是NFS，写入文件时发生的错误会被延迟到文件关闭时反馈。
如果没有检查文件关闭时的反馈信息，可能会导致数据丢失，而我们还误以为写入操作成功。
如果io.Copy和f.close都失败了，我们倾向于将io.Copy的错误信息反馈给调用者，因为它先于f.close发生，
更有可能接近问题的本质。

练习5.18：不修改fetch的行为，重写fetch函数，要求使用defer机制关闭文件。
*/
func fetch1(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// 由于 defer 后定义先执行，如果直接使用 defer 必然会丢失关闭文件时的错，因为一下不会执行了
	// 所以要使用 defer 可以使用匿名函数修改返回值来做
	// 但是这么写不沙雕吗
	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()
	return local, n, err
}

func main() {

}
