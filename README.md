# Go Journey — дневник обучения Go

Учебный репозиторий по пути Python-разработчик → Junior Go Backend Developer.

## Прогресс

| День | Тема | Что сделано |
|---|---|---|
| [Week 1, Day 1](week01/day01-functions-errors) | Основы: функции, zero values, type inference | `divide()` с обработкой ошибки через `error`, проверка zero values (int/string/bool/float64), сравнение `var x int`/`var y`/`z :=` |
| [Week 1, Day 2](week01/day02-flowcontrol) | Flowcontrol: for, if, switch, defer | FizzBuzz через `switch`, короткая форма `if`, механика `defer` (LIFO), 3 варианта `for` (классический/while/бесконечный), проверка на простое число (`IsPrime`) |
| [Week 1, Day 3](week01/day03-slices-strings) | Слайсы, массивы, строки, range, мапы (базово) | `append` и его влияние на len/cap (в т.ч. случай, когда append выделяет новый массив и рвёт связь с исходным слайсом), `ReverseString` через `[]rune` (корректная работа с кириллицей), `CountWords` через мапу, матрица 3×3 через `[][]int`, эксперимент с `b := a` — разбор внутреннего устройства слайса (указатель+len+cap) и почему `b[0]=100` меняет оба слайса, а `append` иногда — нет |
| [Week 1, Day 4](week01/day04-pointers) | Указатели: `&`, `*`, pointer vs value receiver | `Increment` через `*int`, сравнение `modifyValue`/`modifyPointer`, метод `(c *Counter) Increment()` на структуре, `doubleAll` на слайсе, разбор receiver-синтаксиса |