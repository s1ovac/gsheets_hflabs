# gsheets_hflabs
Program which parse html table and send that to Google Sheet


Для реализации данного проекта испоьзовалось Google Sheets API.
"google.golang.org/api/sheets/v4"

Для синхронизации "Базы знаний" был создан HTTP API 
Для обновления можно использовать событие `onchange` в JavaScript 
и отправлять запросы на мой сервер с методом <POST по пути "/api_update">
