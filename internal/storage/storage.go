package storage

import (
	"encoding/json"
	"log/slog"
	"os"
	"sync"
	"time"
)

const storeFile = "data/history.json"

type RequestEntry struct {
	ProfileURL string    `json:"profile_url"`
	Status     string    `json:"status"` // "sent", "connected", "ignored"
	SentAt     time.Time `json:"sent_at"`
	NoteSent   bool      `json:"note_sent"`
}

type Storage struct {
	mu       sync.RWMutex
	Requests map[string]*RequestEntry `json:"requests"`
}

func New() *Storage {
	s := &Storage{
		Requests: make(map[string]*RequestEntry),
	}
	s.load()
	return s
}

func (s *Storage) load() {
	data, err := os.ReadFile(storeFile)
	if err != nil {
		if os.IsNotExist(err) {
			// Ensure dir exists
			_ = os.MkdirAll("data", 0755)
			return
		}
		slog.Error("Failed to load storage", "error", err)
		return
	}
	if err := json.Unmarshal(data, s); err != nil {
		slog.Error("Failed to parse storage", "error", err)
	}
}

func (s *Storage) save() {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		slog.Error("Failed to marshal storage", "error", err)
		return
	}
	if err := os.WriteFile(storeFile, data, 0644); err != nil {
		slog.Error("Failed to save storage", "error", err)
	}
}

// AddRequest records a sent connection request
func (s *Storage) AddRequest(url string, note bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Requests[url] = &RequestEntry{
		ProfileURL: url,
		Status:     "sent",
		SentAt:     time.Now(),
		NoteSent:   note,
	}
	s.save()
}

// HasPendingRequest checks if we already sent a request
func (s *Storage) HasPendingRequest(url string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	entry, exists := s.Requests[url]
	if !exists {
		return false
	}
	// Assume if sent, we shouldn't send again soon unless specific logic
	return entry.Status == "sent" || entry.Status == "connected"
}

// MarkAsConnected updates status to connected
func (s *Storage) MarkAsConnected(url string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if entry, exists := s.Requests[url]; exists {
		entry.Status = "connected"
		s.save()
	} else {
		// New connection not tracked previously? Track it.
		s.Requests[url] = &RequestEntry{
			ProfileURL: url,
			Status:     "connected",
			SentAt:     time.Now(),
		}
		s.save()
	}
}

// GetPendingFollowups returns list of connected users who haven't received a follow-up
// We reuse "NoteSent" to mean "Followup Sent" for "connected" users for simplicity in this POC,
// or we can add a new field. Let's add a new field.
func (s *Storage) GetPendingFollowups() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var urls []string
	for url, req := range s.Requests {
		if req.Status == "connected" && !req.NoteSent { // modifying semantics: NoteSent=FollowupSent for connected
			urls = append(urls, url)
		}
	}
	return urls
}

// MarkFollowupSent updates the flag
func (s *Storage) MarkFollowupSent(url string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if entry, exists := s.Requests[url]; exists {
		entry.NoteSent = true
		s.save()
	}
}

// CheckDailyLimit checks if we exceeded limit (simple count for today)
func (s *Storage) CountRequestsToday() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	count := 0
	now := time.Now()
	for _, req := range s.Requests {
		if req.Status == "sent" &&
			req.SentAt.Year() == now.Year() &&
			req.SentAt.Month() == now.Month() &&
			req.SentAt.Day() == now.Day() {
			count++
		}
	}
	return count
}
