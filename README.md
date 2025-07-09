# gitAnalyse

`gitAnalyse` — это инструмент командной строки, написанный на Go, который анализирует git-репозиторий и показывает вклад каждого автора:

- Количество коммитов
- Количество добавленных и удалённых строк
- Количество изменённых файлов

## Установка

```bash
git clone https://github.com/gorillassi/gitAnalys.git
cd gitAnalys
go build -o gitAnalys
