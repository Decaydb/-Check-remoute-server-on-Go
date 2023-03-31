package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	cmd1 := exec.Command("wmic", "computersystem", "get", "totalphysicalmemory") //Сама команда
	byteTmem, err := cmd1.Output()                                               //Вывод, по умолчанию в байтах
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	cmd2 := exec.Command("wmic", "OS", "get", "FreePhysicalMemory") //Сама команда
	byteFmem, err := cmd2.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	strTotalMem := string(byteTmem) //Преобразование в строку из байтов
	strFreeMem := string(byteFmem)
	replacer := strings.NewReplacer("\n", "", "TotalPhysicalMemory", "", " ", "", "FreePhysicalMemory", "", "\r", "") //Вырезка из строки указанных элементов и текста
	strTotalMem = replacer.Replace(strTotalMem)
	strFreeMem = replacer.Replace(strFreeMem)

	TotalMem, err := strconv.Atoi(strTotalMem) //Конвертация из строки в int
	if err != nil {
		fmt.Println("Ошибка преобразования строки в число")
	}
	FreeMem, err := strconv.Atoi(strFreeMem) //Конвертация из строки в int
	if err != nil {
		fmt.Println("Ошибка преобразования строки в число")
	} else {
		//togig := 1024 * 1024 * 1024
		fmt.Println("Всего оперативной памяти: ", TotalMem/(1024*1024*1024), "Свободно оперативной памяти: ", FreeMem/(1024*1024))

	}

}
