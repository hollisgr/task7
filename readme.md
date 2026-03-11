# Go code quality

## Задание 1. Использование стандартного форматирования кода
### Отформатировать код с использованием команды ​​​​go fmt​​​​.

- Исходный код

```go
package main
import "fmt"
func main() {
      fmt.Println(   "Hello"  )
    if true{
  fmt.Println("World")
    }
}
```

- Применяем `go fmt 1/main.go`

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello")
	if true {
		fmt.Println("World")
	}
}
```

## Задание 2. Проведение анализа кода с помощью ​​​​go vet​​​​
### Запустить ​​​​go vet​​​​ и исправить обнаруженные проблемы.

- Исходный код
```go
func main() {
	name := "Alex"
	fmt.Printf("Привет, %d!", name)
}
```
- Применяем `go vet 2/main.go`
```bash
2/main.go:7:2: fmt.Printf format %d has arg name of wrong type string
```
*Утилита ругается на некорректный спецификатор %d, для исправления ошибки укажем корректный спецификатор - %s*
- Исправленный код
```go
func main() {
	name := "Alex"
	fmt.Printf("Привет, %s!", name)
}
```

## Задание 3. Использование линтера
### Интегрировать линтер ​​​​golangci-lint​​​​ и исправить обнаруженные проблемы.

- Исходный код
```go
package main

import "fmt"

func main() {
	if true {
		Greet("test")
	} else {
	}
}
```
- Запускаем `golangci-lint run 3/main.go`
```bash
3/main.go:12:9: SA9003: empty branch (staticcheck)
	} else {
	       ^
1 issues:
* staticcheck: 1
```
*Получили ошибку линтера: пустая ветка else, для исправления удалим ее*

- Исправленный код
```go
func main() {
	if true {
		Greet("test")
	}
}
```

## Задание 4. Проведение аудита производительности с ​​​​go test​​​​
### Использовать инструмент ​​​​go test​​​​ с флагом ​​​​bench​​​​ для проведения аудита производительности некоторых функций.

- Исходный код
```go
func BenchmarkDefaultConcat(b *testing.B) {
	for b.Loop() {
		var s string
		for j := 0; j < 1000; j++ {
			s += "x"
		}
	}
}

func BenchmarkConcatWithStringsBuilder(b *testing.B) {
	for b.Loop() {
		var sb strings.Builder
		for j := 0; j < 1000; j++ {
			sb.WriteString("x")
		}
		_ = sb.String()
	}
}
```

- Запускаем тесты `go test -bench=. -benchmem ./4`
```bash
goos: linux
goarch: amd64
pkg: task7/4
cpu: AMD Ryzen 7 5700U with Radeon Graphics         
BenchmarkDefaultConcat-16               	    8360	    143438 ns/op	  530281 B/op	     999 allocs/op
BenchmarkConcatWithStringsBuilder-16    	  190896	      5789 ns/op	    3320 B/op	       9 allocs/op
PASS
ok  	task7/4	2.311s
```

*BenchmarkDefaultConcat-16 — Это название теста. Число 16 показывает, сколько логических ядер процессора использовал Go; \
190896 (Итерации) — Столько раз выполнился цикл внутри b.Loop(); \
5789 ns/op (Наносекунды) — Среднее время выполнения одного прохода цикла; \
3320 B/op (Байты) — Сколько оперативной памяти (в куче) выделилось за одну операцию. \
9 allocs/op (Аллокации) — Сколько раз программа просила у системы новый кусок памяти.*

## Задание 5. Проверка покрытия кода тестами
### Использовать инструмент ​​​​go test​​​​ с флагом ​​​​cover​​​​ для проверки покрытия кода.

- Исходный код
```go
func CheckPassword(pwd string) (bool, error) {
	passLen := len(pwd)

	if pwd == "" {
		return false, errors.New("password is empty")
	}

	if passLen < 5 {
		return false, errors.New("password length is less than 5")
	}

	if passLen > 20 {
		return false, errors.New("password length is more than 20")
	}

	forbidden := `;&|\"'`
	if strings.ContainsAny(pwd, forbidden) {
		return false, errors.New("password contains forbidden chars")
	}

	return true, nil
}
```
- Запускаем тесты `go test -coverprofile=cover.out ./5`
```bash
ok  	task7/5	0.003s	coverage: 100.0% of statements
```

## Задание 6. Использование ​​​​errcheck​​​​ для проверки необработанных ошибок
### Запустить ​​​​errcheck​​​​ и обработать все обнаруженные необработанные ошибки.

- Исходный код
```go
func main() {
	str := "hello, 123"
	n := 0
	fmt.Sscanf(str, "hello, %d", &n)
}
```
- Запускаем `errcheck ./6`
```bash
6/main.go:8:12:	fmt.Sscanf(str, "hello, %d", &n)
```
*Утилита ругается что не проверяем на ошибку, для исправления необходимо отработать ошибку*
- Исправленный код
```go
func main() {
	str := "hello, 123"
	n := 0
	_, err := fmt.Sscanf(str, "hello, %d", &n)
	if err != nil {
		return
	}
}
```

## Задание 7. Работа с циклом обратной связи от статического анализатора кода
### Интегрировать статический анализатор кода (например, ​​​​staticcheck​​​​) и решить обнаруженные проблемы.

- Исходный код

```go
func CheckStatus(s string) {
	if s == s {
		fmt.Println("omg")
		return
	}

	if s == "new" {
		fmt.Println("new")
		return
	}

	if s == "old" {
		fmt.Println("old")
		return
	}
}

func main() {
	CheckStatus("new")
}
```

- Используем `staticcheck ./7`
```bash
7/main.go:6:5: identical expressions on the left and right side of the '==' operator (SA4000)

```
*Утилита подсказывает нам, что мы проверяем одинаковые выражения, для исправления удалим эту проверку*

```go
func CheckStatus(s string) {
	if s == "new" {
		fmt.Println("new")
		return
	}

	if s == "old" {
		fmt.Println("old")
		return
	}
}

func main() {
	CheckStatus("new")
}
```

## Задание 8. Анализ циклом обратной связи от пользователей
### Провести анализ обратной связи от пользователей и решите обнаруженные проблемы или улучшите код на основе предложений.

- Исходный код
```go
func Join(ids []int) string {
	res := ""
	for _, id := range ids {
		res += strconv.Itoa(id) + ","
	}
	return res
}
```
*Представим что получили фидбек от пользователей: "В конце всегда запятая, надо исправить"*
- Исправленный код
```go
func Join(ids []int) string {
	strIDs := make([]string, len(ids))
	for i, id := range ids {
		strIDs[i] = strconv.Itoa(id)
	}
	return strings.Join(strIDs, ",")
}
```

## Задание 9. Использование ​​​​go mod tidy​​​​
### Очистить и обновить зависимости с помощью команды ​​​​go mod tidy​​​​.

*Для демонстрации работы проинициализируем проект `go mod init part9` \
Затем установим для него зависимости, например, cleanenv: `go get github.com/ilyakaznacheev/cleanenv`*

- Исходный код проекта
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}

```

- Содержимое файла `go.mod`
```
module part9

go 1.24.6

require (
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/ilyakaznacheev/cleanenv v1.5.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)
```

- Применяем команду `go mod tidy`, после чего смотрим содержимое файла `go.mod`
```
module part9

go 1.24.6
```

*В проекте не был импортирован пакет cleanenv, поэтому все зависимости этого пакета были удалены*

## Задание 10. Использование Git Hooks для автоматической проверки качества кода
### Настроить Git Hooks (например, с использованием ​​​​pre-commit​​​​), чтобы автоматически выполнять проверку качества кода перед каждым коммитом.

- Напишем скрипт, который будет проверять форматирование и проводить анализ (vet), создаем файл `.git/hooks/pre-commit`
```bash
#!/bin/sh

echo "Running pre-commit quality checks..."

# 1. Formatting
echo "Formatting code with 'go fmt'..."
go fmt ./...

# 2. Static analysis
echo "Running 'go vet'..."
if ! go vet ./...; then
    echo "Error: 'go vet' found issues. Fix them before committing."
    exit 1
fi

echo "All checks passed. Proceeding with commit..."
exit 0
```

- Далее даем права на выполнение: `chmod +x .git/hooks/pre-commit`

- Добавляем файл с кодом

```go
package main
import "fmt"

func main() {
      fmt.Println( "Hook test" )
    
    name := "Hollis"
    fmt.Printf("User: %d\n", name) 
}
```

- Пытаемся закоммитить измения и получаем ошибки
```bash
Running pre-commit quality checks...
Formatting code with 'go fmt'...
10/main.go
Running 'go vet'...
# task7/10
# [task7/10]
10/main.go:9:2: fmt.Printf format %d has arg name of wrong type string
Error: 'go vet' found issues. Fix them before committing.
```

- Исправляем ошибки и коммитим
```go
func main() {
	fmt.Println("Hook test")

	name := "Hollis"
	fmt.Printf("User: %s\n", name)
}
```

- В терминале видим
```bash
Running pre-commit quality checks...
Formatting code with 'go fmt'...
Running 'go vet'...
All checks passed. Proceeding with commit...
[main 8156bf3] final
 2 files changed, 69 insertions(+)
 create mode 100644 10/main.go
```
*Предварительный тест пройден и мы можем пушить изменения*