package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "time"
    "personal-calendar/internal/models"
    "personal-calendar/internal/storage"
)

func main() {
    // Создаем хранилище
    store := storage.NewMemoryStorage()
    
    // Добавляем несколько тестовых событий для примера
    addTestEvents(store)
    
    // Создаем reader для ввода с клавиатуры
    reader := bufio.NewReader(os.Stdin)
    
    for {
        // Очищаем экран (просто вывод пустых строк)
        fmt.Print("\033[H\033[2J")
        
        // Показываем меню
        printMenu()
        
        // Читаем команду
        fmt.Print("\nВыберите команду (1-4): ")
        command, _ := reader.ReadString('\n')
        command = strings.TrimSpace(command)
        
        switch command {
        case "1":
            showEvents(store)
        case "2":
            addEvent(store, reader)
        case "3":
            deleteEvent(store, reader)
        case "4":
            fmt.Println("\nДо свидания!")
            return
        default:
            fmt.Println("\nНеизвестная команда!")
            fmt.Println("Нажмите Enter для продолжения...")
            reader.ReadString('\n')
        }
    }
}

func printMenu() {
    fmt.Println(strings.Repeat("-", 50))
    fmt.Println("ПЕРСОНАЛЬНЫЙ КАЛЕНДАРЬ v0.1")
    fmt.Println(strings.Repeat("-", 50))
    
    fmt.Println("\n📅 МЕНЮ:")
    fmt.Println("1. Показать все события")
    fmt.Println("2. Добавить событие")
    fmt.Println("3. Удалить событие")
    fmt.Println("4. Выйти")
}

func addTestEvents(store *storage.MemoryStorage) {
    // Добавляем пару тестовых событий
    tomorrow := time.Now().AddDate(0, 0, 1)
    nextWeek := time.Now().AddDate(0, 0, 7)
    
    store.Add("Встреча с командой", "Обсуждение прогресса", tomorrow)
    store.Add("Дедлайн проекта", "Сдать финальную версию", nextWeek)
}

func showEvents(store *storage.MemoryStorage) {
    events := store.GetAll()
    
    if len(events) == 0 {
        fmt.Println("\n📭 Нет событий в календаре")
    } else {
        fmt.Println("\n📋 ВСЕ СОБЫТИЯ:")
        fmt.Println(strings.Repeat("-", 40))
        for _, event := range events {
            fmt.Printf("[ID: %d] %s\n", event.ID, event.Title)
            fmt.Printf("   📝 %s\n", event.Description)
            fmt.Printf("   📅 %s\n", event.Date.Format("02.01.2006 15:04"))
            fmt.Println(strings.Repeat("-", 40))
        }
    }
    
    fmt.Println("\nНажмите Enter для продолжения...")
    bufio.NewReader(os.Stdin).ReadString('\n')
}

func addEvent(store *storage.MemoryStorage, reader *bufio.Reader) {
    fmt.Println("\n➕ ДОБАВЛЕНИЕ НОВОГО СОБЫТИЯ:")
    fmt.Println(strings.Repeat("-", 40))
    
    // Вводим название
    fmt.Print("Введите название события: ")
    title, _ := reader.ReadString('\n')
    title = strings.TrimSpace(title)
    
    if title == "" {
        fmt.Println("❌ Название не может быть пустым!")
        fmt.Println("\nНажмите Enter для продолжения...")
        reader.ReadString('\n')
        return
    }
    
    // Вводим описание
    fmt.Print("Введите описание события: ")
    description, _ := reader.ReadString('\n')
    description = strings.TrimSpace(description)
    
    // Вводим дату (упрощенно)
    fmt.Print("Введите дату (в формате ДД.ММ.ГГГГ ЧЧ:ММ, например 15.03.2026 18:00): ")
    dateStr, _ := reader.ReadString('\n')
    dateStr = strings.TrimSpace(dateStr)
    
    // Парсим дату
    date, err := time.Parse("02.01.2006 15:04", dateStr)
    if err != nil {
        fmt.Println("❌ Неправильный формат даты! Использую текущее время + 1 день")
        date = time.Now().AddDate(0, 0, 1)
    }
    
    // Добавляем событие
    event := &models.Event{
        ID:          store.GetNextID(),
        Title:       title,
        Description: description,
        Date:        date,
        CreatedAt:   time.Now(),
    }
    
    store.AddEvent(event)
    
    fmt.Println("\n✅ Событие успешно добавлено!")
    fmt.Printf("   ID: %d\n", event.ID)
    fmt.Printf("   Название: %s\n", event.Title)
    fmt.Printf("   Дата: %s\n", event.Date.Format("02.01.2006 15:04"))
    
    fmt.Println("\nНажмите Enter для продолжения...")
    reader.ReadString('\n')
}

func deleteEvent(store *storage.MemoryStorage, reader *bufio.Reader) {
    fmt.Println("\n🗑️  УДАЛЕНИЕ СОБЫТИЯ:")
    fmt.Println(strings.Repeat("-", 40))
    
    // Показываем существующие события
    events := store.GetAll()
    if len(events) == 0 {
        fmt.Println("📭 Нет событий для удаления")
        fmt.Println("\nНажмите Enter для продолжения...")
        reader.ReadString('\n')
        return
    }
    
    fmt.Println("Существующие события:")
    for _, event := range events {
        fmt.Printf("  [%d] %s (%s)\n", 
            event.ID, 
            event.Title, 
            event.Date.Format("02.01.2006"))
    }
    
    // Вводим ID для удаления
    fmt.Print("\nВведите ID события для удаления: ")
    idStr, _ := reader.ReadString('\n')
    idStr = strings.TrimSpace(idStr)
    
    var id int
    fmt.Sscanf(idStr, "%d", &id)
    
    // Удаляем событие
    err := store.Delete(id)
    if err != nil {
        fmt.Printf("❌ Ошибка: %s\n", err)
    } else {
        fmt.Println("✅ Событие успешно удалено!")
    }
    
    fmt.Println("\nНажмите Enter для продолжения...")
    reader.ReadString('\n')
}