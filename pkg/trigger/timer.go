package trigger

import "time"

// TimerTriggerRequest — тело запроса, если функцию вызывает таймер
type TimerTriggerRequest struct {
	Messages []struct {
		EventMetadata struct {
			EventId   string    `json:"event_id"`
			EventType string    `json:"event_type"`
			CreatedAt time.Time `json:"created_at"`
			CloudId   string    `json:"cloud_id"`
			FolderId  string    `json:"folder_id"`
		} `json:"event_metadata"`
		Details struct {
			TriggerId string `json:"trigger_id"`
		} `json:"details"`
	} `json:"messages"`
}
