package models

import (
	"time"
	"fmt"
)

// Event - структура, описывающая событие в календаре
type Event struct {
    ID          int       // Уникальный номер события
    Title       string    // Название
    Description string    // Описание
    Date        time.Time // Дата и время
    CreatedAt   time.Time // Когда создано
}

// NewEvent - функция-конструктор для создания нового события
func NewEvent(id int, title, description string, date time.Time) *Event {
    return &Event{
        ID:          id,
        Title:       title,
        Description: description,
        Date:        date,
        CreatedAt:   time.Now(),
    }
}

// String - метод для красивого вывода события
func (e *Event) String() string {
    return fmt.Sprintf("[%d] %s - %s\n    Описание: %s\n    Дата: %s",
        e.ID, 
        e.Title, 
        e.Date.Format("02.01.2006 15:04"),
        e.Description,
        e.CreatedAt.Format("02.01.2006 15:04"))
}