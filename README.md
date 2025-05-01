# Goclock

Это терминальные часы, написанные на Go, отображающие текущее время в псевдографическом (ASCII-графика) стиле.

## 📷 Пример вывода

```
┌─┐ ┌─┐ ┌┐  ┌─┐ ┌─┐ ┌─┐
│ │ │ │  │  ┌─┘ ├─┤ ┌─┐
└─┘ └─┘  ┘  └─┘ └─┘ └─┘
```

## 🚀 Запуск

1. Убедитесь, что у вас установлен [Go](https://golang.org/dl/).
2. Склонируйте репозиторий:

   ```bash
   git clone https://github.com/yourusername/pseudographic-clock.git
   cd pseudographic-clock
   ```

3. Запустите приложение:

   ```bash
   go run clock.go
   ```
4. Или скомпилируйте исполняемый файл:

   ```bash
   go build -o goclock clock.go
   ```
## 🛑 Завершение

Нажмите `Ctrl + C` для остановки программы. При выходе курсор будет возвращён в нормальное состояние.

## ⚙️ Особенности

- Отображает время в формате `HH:MM:SS`
- Каждая цифра и двоеточие стилизованы с использованием UTF-8 псевдографики
- Обновление экрана каждую секунду
- ANSI-escape последовательности для стилизации и управления экраном
- Обработка `SIGINT` для корректного завершения

## 💻 Совместимость

- ✅ Linux

## Автор:
- **Grannik**

## Контакты:
- **Сайт**: [Granni](https://grannik.neocities.org/)

## Репозитории:
- **asciinema**:    [goClock](https://asciinema.org/a/717640)
- **Codeberg**:     [goClock](https://codeberg.org/Grannik/Goclock)
- **GitHub**:       [goClock]()
- **GitLab**:       [goClock]()
- **SourceForge**:  [goClock]()
- **NotABug**:      [goClock]()
- **Gitea (Demo)**: [goClock]()
- **Gogs (Demo)**:  [goClock]()

## Лицензия
Этот проект распространяется под лицензией **MIT**.
См. файл [`LICENSE`](LICENSE) для подробностей.
