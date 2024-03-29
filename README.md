# pingo
Бот который умеет отправлять регулярные сообщения различные групповые каналы. Работает поверх Yandex.Serverless функций. Запускается по триггеру timer.

## Настройка
Для настройки приложения, установите значение переменных окружения

|Название | Описание                                                          | Значение                                       |
|---------|-------------------------------------------------------------------|------------------------------------------------|
|ENVIRONMENT| Текущее окружение, <br/>в котором работает приложение             | `development`, `staging`, `production`         |
|REMOTE_CONFIG_URL| URL до файла конфигурации шаблонов (если файл находится удаленно( | `https://example.s3.ru/configs/templates.yaml` |


## Поддержка альтернативных каналов

Для поддержки альтернативных каналов, нужно установить указанные ниже переменные окружения. 

❗️ Если переменные окружения для определенного канала не будут указаны, `pingo` их проигнорирует.

| Название            | Описание                                                                                                   | Значение         |
|---------------------|------------------------------------------------------------------------------------------------------------|------------------|
| SLACK_TOKEN         | API токен для отправки уведомлений в Slack                                                                 | "slack_token"    |
| TELEGRAM_TOKEN      | API токен бота в Telegram для отправки уведомлений в Telegram                                              | "telegram_token" |
| SMTP_HOST           | Настройки для отправки Email через SMTP-сервер, в частности SMTP-хост                                      | smtp.example.ru  |
| SMTP_PORT           | Настройки для отправки Email через SMTP-сервер, в частности SMTP-порт                                      | 983              |
| SMTP_USERNAME       | Настройки для отправки Email через SMTP-сервер, в частности `Имя пользователя` учетной записи для рассылок | `alewkin`        |
| SMTP_PASSWORD       | Настройки для отправки Email через SMTP-сервер, в частности `Пароль` от учетной записи для рассылок        | `qwerty123`        |
| SPACE_HOST          | Хост вашей инсталляции Jet Brains Space                                                                                                                                                           | https://space.example.com      |
| SPACE_TOKEN         | API-токен для работы с API Space 	                                                                                                                                                                | "myToken" 	                     |

## Шаблоны сообщений
Для поддержки мультиканальности, в приложении есть специальный формат описания шаблонов сообщений.

### Email
Отправка сообщений по Email осуществляется соглсано `mailto:` схеме. Поддерживаются параметры:
 * `from`
 * `subject`
 * [`unsubscribeLink`](https://support.google.com/mail/answer/81126)

Примеры:
 * mailto:"John Wayne"<john@example.org>?subject=test-subj&from="Notifier"<notify@example.org>
 * mailto:addr1@example.org,addr2@example.org?subject=test-subj&from=notify@example.org&unsubscribeLink=http://example.org/unsubscribe

### Slack
Отправка сообщений в Slack осуществляется согласно `slack:` схеме. Сообщение можно отправить в канал, указав его название/идентификатор или в личные сообщения пользователю.
Выбор канала осужествляется по аналогии с mailto схемой: `slack:someChannelName` или `slack:someChannelID` или `slack:someUserID`

Поддерживаемые параметры:
 * `title`
 * `titleLink`
 * `attachmentText`
 
Примеры:
 * slack:channel?title=title&attachmentText=test%20text&titleLink=https://example.org

### Telegram
Отправка сообщений в Telegram осуществляется согласно `telegram:` схеме. 
Сообщение можно отправить в канал (указав его название) или групповой чат (указав его идентификатор).
Выбор канала осужествляется по аналогии с mailto схемой: `telegram:channel` или `telegram:channel?parseMode=HTML`.

Поддерживаемые параметры:
* `parseMode` (парсинг текста сообщения, по умолчанию `Markdown`, можно указать `HTML`)
  
Примеры:
* telegram:someChannel?parseMode=HTML


### Space
Отправка сообщений в Space осуществляется согласно `space:` схеме. Сообщение можно отправить в канал (указав его идентификатор). 
Выбор канала осужествляется по аналогии с mailto схемой: `space:channelID`.

Примеры:
* space:123456