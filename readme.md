# Go code quality

## Задание 1. Использование стандартного форматирования кода
### Отформатируйте код вашего проекта с использованием команды ​​​​go fmt​​​​.

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
### Запустите ​​​​go vet​​​​ для вашего проекта и исправьте обнаруженные проблемы.

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

## Задание 3. Использование линтера
### Интегрируйте линтер (например, ​​​​golangci-lint​​​​) в свой проект и исправьте обнаруженные проблемы.

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

## Задание 4. Проведение аудита производительности с ​​​​go test​​​​
### Используйте инструмент ​​​​go test​​​​ с флагом ​​​​bench​​​​ для проведения аудита производительности некоторых функций в вашем проекте.

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
### Используйте инструмент ​​​​go test​​​​ с флагом ​​​​cover​​​​ для проверки покрытия кода вашими тестами.

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
### Запустите ​​​​errcheck​​​​ для вашего проекта и обработайте все обнаруженные необработанные ошибки.

## Задание 7. Работа с циклом обратной связи от статического анализатора кода
### Интегрируйте статический анализатор кода (например, ​​​​staticcheck​​​​) и решите обнаруженные проблемы.

## Задание 8. Анализ циклом обратной связи от пользователей
### Проведите анализ обратной связи от пользователей и решите обнаруженные проблемы или улучшите код на основе предложений.

## Задание 9. Использование ​​​​go mod tidy​​​​
### Очистите и обновите зависимости с помощью команды ​​​​go mod tidy​​​​.

## Задание 10. Использование Git Hooks для автоматической проверки качества кода
### Настройте Git Hooks (например, с использованием ​​​​pre-commit​​​​), чтобы автоматически выполнять проверку качества кода перед каждым коммитом.
