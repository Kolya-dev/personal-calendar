package storage

import (
    "fmt"
	"time"
    "personal-calendar/internal/models"
)

// MemoryStorage - хранилище событий в памяти
type MemoryStorage struct {
    events  map[int]*models.Event
    nextID  int
}

// NewMemoryStorage - создает новое хранилище
func NewMemoryStorage() *MemoryStorage {
    return &MemoryStorage{
        events: make(map[int]*models.Event),
        nextID: 1,
    }
}

// Add - добавляет новое событие
func (s *MemoryStorage) Add(title, description string, dateTime string) (*models.Event, error) {
    // Здесь потом будет парсинг даты, пока упростим
    event := &models.Event{
        ID:          s.nextID,
        Title:       title,
        Description: description,
        CreatedAt:   time.Now(),
    }
    
    s.events[s.nextID] = event
    s.nextID++
    
    return event, nil
}

// GetAll - возвращает все события
func (s *MemoryStorage) GetAll() []*models.Event {
    result := make([]*models.Event, 0, len(s.events))
    for _, event := range s.events {
        result = append(result, event)
    }
    return result
}

// GetByID - находит событие по ID
func (s *MemoryStorage) GetByID(id int) (*models.Event, error) {
    event, exists := s.events[id]
    if !exists {
        return nil, fmt.Errorf("событие с ID %d не найдено", id)
    }
    return event, nil
}

// Delete - удаляет событие по ID
func (s *MemoryStorage) Delete(id int) error {
    _, exists := s.events[id]
    if !exists {
        return fmt.Errorf("событие с ID %d не найдено", id)
    }
    delete(s.events, id)
    return nil
}