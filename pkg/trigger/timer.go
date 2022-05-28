package trigger

import "time"

// TimerRequest — тело запроса для main.Handler
// если вызов приходит из Yandex.Cloud Functions в виде триггера Timer
type TimerRequest struct {
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
