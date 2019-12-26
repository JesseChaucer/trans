/**
 * 向文件中写数据
*/
package util

import (
	"fmt"
	"os"
	"bufio"
)

func WriteFile(str string) {
	filePath := "/Users/admin/code/go_workspace/src/inAction/translation/asset/future_delivery.json"
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
			fmt.Println(err)
			return
	}
	defer file.Close()

	// 写入时，使用带缓冲的 *Writer
	writer := bufio.NewWriter(file)
	writer.WriteString(str + "\r\n")  // \r\n表示回车换行

	// 因为writer是带缓冲的，因此在调用WriteString方法时，其实内容是先写到缓冲区的，
	// 需要调用Flush方法，将缓冲区的内容写入到文件中
	writer.Flush()
}