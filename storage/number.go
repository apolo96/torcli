package storage

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

func GetNumber() float64 {
	data, err := os.ReadFile(viper.GetString("file_path"))
	if err != nil {
		fmt.Println(err)
		return 0
	}
	number, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 0
	}
	return number
}

func SaveNumber(number float64) error {
	file, err := os.OpenFile(viper.GetString("file_path"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("%f", number))
	if err != nil {
		return err
	}
	return nil
}
