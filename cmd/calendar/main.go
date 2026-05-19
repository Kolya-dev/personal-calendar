package main

import (
	"bufio"
	"fmt"
	"os"
	"personal-calendar/internal/models"
	"personal-calendar/internal/storage"
	"strings"
	"time"
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
		fmt.Print("\nВыберите команду (1-6): ")
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
			editEvent(store, reader)
		case "5":
			filterEventsByDate(store, reader)
		case "6":
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
	fmt.Println("4. Редактировать событие")
	fmt.Println("5. Показать события на дату")
	fmt.Println("6. Выйти")
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
		fmt.Printf("✅ Установлена дата: %s\n", date.Format("02.01.2006 15:04"))
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

func editEvent(store *storage.MemoryStorage, reader *bufio.Reader) {
	fmt.Println("\n✏️  РЕДАКТИРОВАНИЕ СОБЫТИЯ:")
	fmt.Println(strings.Repeat("-", 40))

	// Показываем существующие события
	events := store.GetAll()
	if len(events) == 0 {
		fmt.Println("📭 Нет событий для редактирования")
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

	// Вводим ID для редактирования
	fmt.Print("\nВведите ID события для редактирования: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)

	var id int
	fmt.Sscanf(idStr, "%d", &id)

	// Находим событие
	event, err := store.GetByID(id)
	if err != nil {
		fmt.Printf("❌ Ошибка: %s\n", err)
		fmt.Println("\nНажмите Enter для продолжения...")
		reader.ReadString('\n')
		return
	}

	fmt.Println("\n📝 ТЕКУЩИЕ ДАННЫЕ:")
	fmt.Printf("   Название: %s\n", event.Title)
	fmt.Printf("   Описание: %s\n", event.Description)
	fmt.Printf("   Дата: %s\n", event.Date.Format("02.01.2006 15:04"))

	// Вводим новые данные
	fmt.Println("\n(Если оставить пустым — значение не изменится)")

	fmt.Print("Новое название: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)
	if title != "" {
		event.Title = title
	}

	fmt.Print("Новое описание: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)
	if description != "" {
		event.Description = description
	}

	fmt.Print("Новая дата (ДД.ММ.ГГГГ ЧЧ:ММ): ")
	dateStr, _ := reader.ReadString('\n')
	dateStr = strings.TrimSpace(dateStr)
	if dateStr != "" {
		date, err := time.Parse("02.01.2006 15:04", dateStr)
		if err != nil {
			fmt.Println("❌ Неправильный формат даты! Дата не изменена")
		} else {
			event.Date = date
		}
	}

	fmt.Println("\n✅ Событие успешно обновлено!")
	fmt.Println("\nНажмите Enter для продолжения...")
	reader.ReadString('\n')
}

func filterEventsByDate(store *storage.MemoryStorage, reader *bufio.Reader) {
	fmt.Println("\n📅 ПОКАЗ СОБЫТИЙ НА ДАТУ:")
	fmt.Println(strings.Repeat("-", 40))

	fmt.Print("Введите дату (ДД.ММ.ГГГГ): ")
	dateStr, _ := reader.ReadString('\n')
	dateStr = strings.TrimSpace(dateStr)

	// Парсим только дату (без времени)
	targetDate, err := time.Parse("02.01.2006", dateStr)
	if err != nil {
		fmt.Println("❌ Неправильный формат даты! Используйте ДД.ММ.ГГГГ")
		fmt.Println("\nНажмите Enter для продолжения...")
		reader.ReadString('\n')
		return
	}

	events := store.GetAll()
	found := false

	fmt.Println("\n📋 СОБЫТИЯ НА", targetDate.Format("02.01.2006"))
	fmt.Println(strings.Repeat("-", 40))

	for _, event := range events {
		// Сравниваем только год, месяц, день
		if event.Date.Year() == targetDate.Year() &&
			event.Date.Month() == targetDate.Month() &&
			event.Date.Day() == targetDate.Day() {
			fmt.Printf("[ID: %d] %s\n", event.ID, event.Title)
			fmt.Printf("   📝 %s\n", event.Description)
			fmt.Printf("   ⏰ %s\n", event.Date.Format("15:04"))
			fmt.Println(strings.Repeat("-", 40))
			found = true
		}
	}

	if !found {
		fmt.Println("📭 Нет событий на указанную дату")
	}

	fmt.Println("\nНажмите Enter для продолжения...")
	reader.ReadString('\n')
}
