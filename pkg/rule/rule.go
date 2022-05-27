package rule

// Rule — правило отправки сообщения
type Rule struct {
	// schedule — расписание отправки в виде CRONTAB выражения
	schedule string
}
