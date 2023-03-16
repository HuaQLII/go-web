package generate

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var stRootDir string
var stSeparator string       //定义一个分隔符
var iJsonData map[string]any //定义一个变量来接收加载的Json文件

const stJsonFileName = "dir.json" //定义一个文件名称

func loadJson() { //小写包外不可见

	stSeparator := string(filepath.Separator)                         //分隔符实例化强转为string
	stWorkDir, _ := os.Getwd()                                        // Getwd返回一个有根的路径名称，对应于当前目录。
	stRootDir = stWorkDir[:strings.LastIndex(stWorkDir, stSeparator)] //string可以按照切片读取某一个片段的内容，从0-当前的这个目录最后一个分隔符对应的位置
	//fmt.Println(stSeparator)
	//fmt.Println(stWorkDir)
	//fmt.Println(stRootDir)
	//加载Json文件
	gnJsonBytes, _ := os.ReadFile(stWorkDir + stSeparator + stJsonFileName)
	err := json.Unmarshal(gnJsonBytes, &iJsonData)
	if err != nil {
		panic("Load Json Data Error" + err.Error())
	}

}

// 解析Json文件
// mapData 把Json看成map文件所以参数这里要接受某一条map数据,map类型，key是string，类型是any
// stParentDir  解析数据的父级路径传进来
func parseMap(mapData map[string]any, stParentDir string) {
	stSeparator := string(filepath.Separator) //分隔符实例化强转为string
	for k, v := range mapData {
		//做一个类型的断言
		switch v.(type) {
		case string:
			{
				path, _ := v.(string) //强转string，做一个类型断言
				if path == "" {
					continue
				}
				if stParentDir != "" {
					path = stParentDir + stSeparator + path
					if k == "text" {
						stParentDir = path
					}
				} else {
					stParentDir = path
				}
				createDir(path)
			}
		case []any:
			{
				parseArray(v.([]any), stParentDir)
			}
		}
	}

}

func parseArray(giJsonData []any, stParentDir string) {
	for _, v := range giJsonData {
		mapV, _ := v.(map[string]any)
		parseMap(mapV, stParentDir)
	}
}

func createDir(path string) {
	stSeparator := string(filepath.Separator) //分隔符实例化强转为string

	if path == "" {
		return
	}
	//fmt.Println(path)
	err := os.MkdirAll(stRootDir+stSeparator+path, fs.ModePerm)
	if err != nil {
		panic("Create DirError:" + err.Error())
	}
}

func TestGenerateDir(t *testing.T) {
	// T是传递给测试函数的一个类型，用于管理测试状态和支持格式化的测试日志。
	loadJson()
	parseMap(iJsonData, "")
}
