package main

import (
	"fmt"
	"strings"
)

func main() {
    fmt.Println(strings.Repeat("=", 50))
    fmt.Println("ПЕРСОНАЛЬНЫЙ КАЛЕНДАРЬ v0.1")
    fmt.Println(strings.Repeat("=", 50))
    
    fmt.Println("\nКоманды:")
    fmt.Println("1. Показать все события")
    fmt.Println("2. Добавить событие")
    fmt.Println("3. Удалить событие")
    fmt.Println("4. Выйти")
    
    fmt.Println("\nПрограмма запущена. Выберите команду...")
}