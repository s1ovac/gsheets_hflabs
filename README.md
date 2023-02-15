# gsheets_hflabs
Program which parse html table and send that to Google Sheet


Для реализации данного проекта испоьзовалось Google Sheets API.
"google.golang.org/api/sheets/v4"

Для синхронизации "Базы знаний" был создан HTTP API 
Для обновления можно использовать событие `onchange` в JavaScript 
и отправлять запросы на мой сервер с методом `POST по пути "/api_update"`

Реализация данного проекта заняла порядка `6` рабочих часов.

Ссылка на Google Sheet: <a>https://docs.google.com/spreadsheets/d/1c9rQsoZddsCylo8LYvT8Qj28XmofNOGXGdLIbgA7ANw/edit#gid=0</a>
