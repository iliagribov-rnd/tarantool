# Knight Dialer

## Задача

Набор номера на клавиатуре N×M ходом шахматного коня. Требуется найти минимальное количество переходов и вычислить коэффициент KC = (сумма нажатых цифр) × K.

## Структура

- `structures.go` — структуры данных
- `tarantool_norepeat.go` — решение без повторений клавиш
- `tarantool_repeat.go` — решение с возможностью повторных нажатий
- `tarantool_test.go` — тесты

## Запуск

```bash
go run tarantool_norepeat.go structures.go < input.txt
go run tarantool_repeat.go structures.go < input.txt
```

## Тесты

```bash
go test
```

## Входные данные

```
N M
[клавиатура N×M из цифр]
Slen
[строка для набора]
```

## Выход

```
-1  # если набор невозможен
K C  # минимальное кол-во переходов и KC
```
