package initconfig

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func UnMarshalFromFile(fileName string, result interface{}) (err error) {
	fData, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("读取文件失败,", err)
	}
	err = UnMarshal(fData, result)
	return
}
func UnMarshal(fData []byte, sData interface{}) (err error) {
	if len(fData) == 0 {
		err = errors.New("配置文件为空")
		return
	}
	if reflect.TypeOf(sData).Kind() != reflect.Ptr {
		err = errors.New("传入的不是指针")
		return
	}
	if reflect.TypeOf(sData).Elem().Kind() != reflect.Struct {
		err = errors.New("传入的不是结构题指针")
		return
	}

	for i := 0; i < reflect.TypeOf(sData).Elem().NumField(); i++ {
		if reflect.TypeOf(sData).Elem().Field(i).Type.Kind() != reflect.Struct {
			err = fmt.Errorf("tag %s 对应的类型不是结构体", reflect.TypeOf(sData).Elem().Field(i).Tag.Get("ini"))
			return
		}
	}

	//cMap := make(map[string]map[string]interface{})
	selectName := ""
	//rType := reflect.TypeOf(sData).Elem()
	//rVal := reflect.ValueOf(sData).Elem()
	//fmt.Println(string(fData))
	for lineOn, line := range strings.Split(string(fData), "\n") {
		lineData := strings.TrimSpace(line)
		if len(lineData) == 0 {
			continue
		}
		if strings.HasPrefix(lineData, "#") || strings.HasPrefix(line, ";") {
			continue
		}
		if strings.HasPrefix(lineData, "[") {
			//fmt.Println(lineData)
			selectName, err = PassSelect(lineData, sData)
			if err != nil {
				err = fmt.Errorf("解析%d行出错: %s", lineOn, err)
				return
			}
			//fmt.Println(selectName)
			if selectName == "" {
				continue
			}
		} else {
			key, value, errT := PassConfigItem(lineData, sData, selectName)
			_ = key
			_ = value
			if errT != nil {
				err = fmt.Errorf("解析%d行出错: %s", lineOn, err)
				return
			}

			//for i := 0; i < rType.NumField(); i++ {
			//	if selectName == rType.Field(i).Tag.Get("ini") {
			//		for j := 0; j < rType.Field(i).Type.NumField(); j++ {
			//			if key == rType.Field(i).Type.Field(j).Tag.Get("ini") {
			//				//fmt.Printf("%s 的类型为 %s\n", key, rType.Field(i).Type.Field(j).Type.Kind())
			//				switch rType.Field(i).Type.Field(j).Type.Kind() {
			//				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			//					vInt, err1 := strconv.ParseInt(value, 10, 64)
			//					if err1 != nil {
			//						err = fmt.Errorf("%s的类型不能识别%s", key, rType.Field(i).Type.Field(j).Type.Kind())
			//					}
			//					rVal.Field(i).Field(j).SetInt(vInt)
			//				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			//					vUint, err1 := strconv.ParseUint(value, 10, 64)
			//					if err1 != nil {
			//						err = fmt.Errorf("%s的类型不能识别%s", key, rType.Field(i).Type.Field(j).Type.Kind())
			//					}
			//					rVal.Field(i).Field(j).SetUint(vUint)
			//
			//				case reflect.Float32, reflect.Float64:
			//					vFloat, err1 := strconv.ParseFloat(value, 64)
			//					if err1 != nil {
			//						err = fmt.Errorf("%s的类型不能识别%s", key, rType.Field(i).Type.Field(j).Type.Kind())
			//					}
			//					rVal.Field(i).Field(j).SetFloat(vFloat)
			//				case reflect.String:
			//					rVal.Field(i).Field(j).SetString(value)
			//				default:
			//					err = fmt.Errorf("%s的类型不能识别%s", key, rType.Field(i).Type.Field(j).Type.Kind())
			//				}
			//			}
			//		}
			//	}
			//}
		}

	}
	return
}

func PassSelect(lineData string, sDta interface{}) (selectName string, err error) {
	if !strings.HasSuffix(lineData, "]") {
		err = errors.New("没有以 ] 结尾")
		return
	}
	fSelectName := strings.TrimSpace(lineData[1 : len(lineData)-1])
	if len(fSelectName) == 0 {
		err = errors.New("select 不能为空")
		return
	}
	rType := reflect.TypeOf(sDta).Elem()
	for i := 0; i < rType.NumField(); i++ {
		if fSelectName == rType.Field(i).Tag.Get("ini") {
			//fmt.Println(rType.Field(i).Name)
			selectName = rType.Field(i).Name
		}
	}
	return
}
func PassConfigItem(lineData string, sDta interface{}, selectName string) (key string, value string, err error) {
	index := strings.Index(lineData, "=")
	if len(lineData[index+1:]) == 0 {
		err = errors.New("配置项不能为空")
	}
	key = strings.TrimSpace(lineData[:index])
	value = strings.TrimSpace(lineData[index+1:])
	rVal := reflect.ValueOf(sDta).Elem().FieldByName(selectName)
	for i := 0; i < rVal.Type().NumField(); i++ {
		sKey := rVal.Type().Field(i).Tag.Get("ini")
		if sKey == key {
			switch rVal.Type().Field(i).Type.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				vInt, err1 := strconv.ParseInt(value, 10, 64)
				if err1 != nil {
					err = fmt.Errorf("%s的类型不匹配%s", key, rVal.Type().Field(i).Type.Kind())
				}
				rVal.Field(i).SetInt(vInt)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				vUint, err1 := strconv.ParseUint(value, 10, 64)
				if err1 != nil {
					err = fmt.Errorf("%s的类型不匹配%s", key, rVal.Type().Field(i).Type.Kind())
				}
				rVal.Field(i).SetUint(vUint)

			case reflect.Float32, reflect.Float64:
				vFloat, err1 := strconv.ParseFloat(value, 64)
				if err1 != nil {
					err = fmt.Errorf("%s的类型不匹配%s", key, rVal.Type().Field(i).Type.Kind())
				}
				rVal.Field(i).SetFloat(vFloat)
			case reflect.String:
				rVal.Field(i).SetString(value)
			default:
				err = fmt.Errorf("%s的类型不能识别%s", key, rVal.Type().Field(i).Type.Kind())
			}
		}
	}
	return
}

func MarshalToFile(fileName string, conf interface{}) (err error) {
	var buf []byte
	buf, err = Marshal(conf)
	if err != nil {
		return
	}
	err = os.WriteFile(fileName, buf, os.ModePerm)
	return
}
func Marshal(conf interface{}) (buf []byte, err error) {

	if reflect.TypeOf(conf).Kind() != reflect.Struct {
		err = errors.New("需要传入结构体")
		return
	}
	for i := 0; i < reflect.TypeOf(conf).NumField(); i++ {
		if reflect.TypeOf(conf).Field(i).Type.Kind() != reflect.Struct {
			err = errors.New("需要传入结构体的里面的字段只能是结构体")
			return
		}
	}
	var rest []string = []string{}
	for i := 0; i < reflect.TypeOf(conf).NumField(); i++ {
		//fmt.Print(reflect.TypeOf(conf).Field(i))
		if reflect.TypeOf(conf).Field(i).Tag.Get("ini") == "" {
			continue
		}
		rest = append(rest, fmt.Sprintf("\n[%s]\n", reflect.TypeOf(conf).Field(i).Tag.Get("ini")))
		for j := 0; j < reflect.TypeOf(conf).Field(i).Type.NumField(); j++ {
			if reflect.TypeOf(conf).Field(i).Type.Field(j).Tag.Get("ini") == "" {
				continue
			}
			rest = append(rest, fmt.Sprintf("%s = %v\n", reflect.TypeOf(conf).Field(i).Type.Field(j).Tag.Get("ini"), reflect.ValueOf(conf).Field(i).Field(j).Interface()))
		}
	}
	for line := range rest {
		buf = append(buf, []byte(rest[line])...)
	}
	return
}
