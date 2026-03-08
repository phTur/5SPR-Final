package actioninfo

import (
	"fmt"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			fmt.Println("Ошибка парсинга:", err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			fmt.Println("Ошибка обработки:", err)
			continue
		}
		fmt.Println(info)
	}

}
