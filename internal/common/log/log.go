package log

import (
	"fmt"
	"os"
)

func LogError(err error) {
	strError := fmt.Sprintf("%+v", err)

	f, err := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Failed to open error log file:", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString(strError + "\n")
	if err != nil {
		fmt.Println("Failed to write to error log file:", err)
	}
}
