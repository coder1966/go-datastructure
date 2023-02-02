package a003recursion

import (
	"fmt"
	"godatastructure/a002stackarray/stackarray"
	"godatastructure/a004queue/queue"
	"io/ioutil"
	"path/filepath"
)

func DoWalkFolder() {
	fmt.Println("// 递归方式")
	path := "."
	files := []string{}
	files, _ = GetAll(path, files)
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}

	fmt.Println("// 压栈代替递归")
	stack := stackarray.NewStack(1000)
	stack.Push(path)
	files = []string{}

	for !stack.IsEmpty() {
		data := stack.Pop() // 取出栈内一个数据
		if data == nil {
			break
		}

		read, err := ioutil.ReadDir(data.(string)) // 读取文件件
		if err != nil {
			fmt.Println("ioutil.ReadDir 读文件件 error: ", err)
		}

		// 循环读到的文件|文件夹
		for _, fi := range read {

			if fi.IsDir() {
				// 文件夹
				// fullDir := data.(string) + "/" + fi.Name() // 构造新的路径
				fullDir := filepath.Join(data.(string), fi.Name()) // 构造新的路径
				files = append(files, fullDir)                     // 无论文件 or 子目录，都加上
				stack.Push(fullDir)                                // 递归
			} else {
				// fullDir := data.(string) + "/" + fi.Name() // 构造新的路径
				fullDir := filepath.Join(data.(string), fi.Name()) // 构造新的路径
				files = append(files, fullDir)                     // 无论文件 or 子目录，都加上
			}
		}

	}

	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}

	fmt.Println("// 队列代替递归")
	myq := queue.NewQueue()
	myq.EnQueue(path)
	files = []string{}

	for !myq.IsEmpty() {
		data := myq.DeQueue() // 取出栈内一个数据
		if data == nil {
			break
		}

		read, err := ioutil.ReadDir(data.(string)) // 读取文件件
		if err != nil {
			fmt.Println("ioutil.ReadDir 读文件件 error: ", err)
		}

		// 循环读到的文件|文件夹
		for _, fi := range read {

			if fi.IsDir() {
				// 文件夹
				// fullDir := data.(string) + "/" + fi.Name() // 构造新的路径
				fullDir := filepath.Join(data.(string), fi.Name()) // 构造新的路径
				files = append(files, fullDir)                     // 无论文件 or 子目录，都加上
				myq.EnQueue(fullDir)                               // 递归
			} else {
				// fullDir := data.(string) + "/" + fi.Name() // 构造新的路径
				fullDir := filepath.Join(data.(string), fi.Name()) // 构造新的路径
				files = append(files, fullDir)                     // 无论文件 or 子目录，都加上
			}
		}

	}

	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}

}

func GetAll(path string, files []string) ([]string, error) {
	read, err := ioutil.ReadDir(path) // 读取文件件
	if err != nil {
		fmt.Println("ioutil.ReadDir 读文件件 error: ", err)
	}

	// 循环读到的文件|文件夹
	for _, fi := range read {

		if fi.IsDir() {
			// 文件夹
			// fullDir := path + "/" + fi.Name() // 构造新的路径
			fullDir := filepath.Join(path, fi.Name()) // 构造新的路径
			files = append(files, fullDir)            // 无论文件 or 子目录，都加上
			files, _ = GetAll(fullDir, files)         // 递归
		} else {
			// fullDir := path + "/" + fi.Name() // 构造新的路径
			fullDir := filepath.Join(path, fi.Name()) // 构造新的路径
			files = append(files, fullDir)            // 无论文件 or 子目录，都加上
		}
	}
	return files, nil
}
