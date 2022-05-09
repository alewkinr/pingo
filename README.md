# pingo
Бот для Jet Brains Space, который умеет отправлять сообщения в групповой канал. Работает поверх Yandex.Serverless функций. Запускается по триггеру timer

## Настройка
Для настройки приложения, установите значение переменных окружения

|Название | Описание                                                                                                                                                                                          | Значение                       |
|---------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------|
|ENVIRONMENT| Текущее окружение, <br/>в котором работает приложение                                                                                                                                             | `development`, `staging`, `production` |
|SPACE_HOST| Хост вашей инсталляции Jet Brains Space                                                                                                                                                           | https://space.example.com      |
|SPACE_TOKEN| API-токен для работы с API Space 	                                                                                                                                                                | "myToken" 	                     |
|SPACE_DEBUG_CHANNEL| Идентификатор групппового канала, <br/>который будет использоваться для отладки сообещений. В него отправляются сообщения, <br/>если значение переменной среды ENVIRONMENT установлено  `development`	 | `4anH4q38FKzj` 	               |

