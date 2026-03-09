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

// Add - добавляет новое событие (упрощенная версия)
func (s *MemoryStorage) Add(title, description string, date time.Time) (*models.Event, error) {
    event := &models.Event{
        ID:          s.nextID,
        Title:       title,
        Description: description,
        Date:        date,
        CreatedAt:   time.Now(),
    }
    
    s.events[s.nextID] = event
    s.nextID++
    
    return event, nil
}

// AddEvent - добавляет готовое событие
func (s *MemoryStorage) AddEvent(event *models.Event) {
    s.events[event.ID] = event
    if event.ID >= s.nextID {
        s.nextID = event.ID + 1
    }
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

// GetNextID - возвращает следующий доступный ID
func (s *MemoryStorage) GetNextID() int {
    return s.nextID
}