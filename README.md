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
```
## Использование

```bash
./gitAnalys [путь_к_репозиторию] [флаги]
```

### Примеры

```bash
./gitAnalys .                             # текущий репозиторий
./gitAnalys ./myrepo --ext=.go           # только .go файлы
./gitAnalys ~/project --sort=commits     # сортировка по коммитам
```

---

## Флаги

| Флаг      | Описание                                 |
|-----------|-------------------------------------------|
| `--ext`   | Фильтровать файлы по расширению (например `.go`) |
| `--sort`  | Сортировка: `lines` (по умолчанию), `commits`, `name` |

---

## Пример вывода

```
Author       Commits  Lines Added  Lines Deleted  Files Changed  
---------------------------------------------------------------
Alice        14        1040         200            12
Bob          10        800          150            8
```

---
