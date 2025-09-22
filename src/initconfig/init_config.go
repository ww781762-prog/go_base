package initconfig

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func UnMarshal(fData []byte, sData interface{}) (err error) {
	if len(fData) == 0 {
		err = errors.New("配置文件为空")
		return
	}
	if reflect.TypeOf(sData).Kind() != reflect.Ptr {
		err = errors.New("传入的不是指针")
	}
	if reflect.TypeOf(sData).Elem().Kind() != reflect.Struct {
		err = errors.New("传入的不是结构题指针")
	}

	cMap := make(map[string]map[string]interface{})
	//fmt.Println(string(fData))
	for lineOn, line := range strings.Split(string(fData), "\n") {
		lineData := strings.TrimSpace(line)
		selectName := ""
		if len(lineData) == 0 {
			continue
		}
		if strings.HasPrefix(lineData, "#") || strings.HasPrefix(line, ";") {
			continue
		}
		if strings.HasPrefix(lineData, "[") {
			//fmt.Println(lineData)
			selectName, err = PassSelect(lineData)
			if err != nil {
				err = fmt.Errorf("解析%d行出错: %s", lineOn, err)
				return
			}
			cMap[selectName] = make(map[string]interface{})
			fmt.Println(cMap)
		} else {
			iMapData := map[string]interface{}{}
			err = PassConfigItem(lineData, iMapData)
			if err != nil {
				err = fmt.Errorf("解析%d行出错: %s", lineOn, err)
				return
			}
			cMap[selectName] = iMapData
			fmt.Println(cMap)

		}

	}
	fmt.Println(cMap)
	return
}

func PassSelect(lineData string) (selectName string, err error) {
	if !strings.HasSuffix(lineData, "]") {
		err = errors.New("没有以 ] 结尾")
		return
	}
	selectName = strings.TrimSpace(lineData[1 : len(lineData)-1])
	if len(selectName) == 0 {
		err = errors.New("select 不能为空")
		return
	}

	return
}
func PassConfigItem(lineData string, m map[string]interface{}) (err error) {
	index := strings.Index(lineData, "=")
	//fmt.Println(len(lineData), index)
	if len(lineData[index+1:]) == 0 {
		err = errors.New("配置项不能为空")
	}
	//fmt.Println(lineData[:index], lineData[index+1:])
	m[lineData[:index]] = lineData[index+1:]
	return
}

func Marshal() {

}
