>[!NOTE]
># **Описание**
>Это **сервер-калькулятор,** написанный на языке **GO**, принимающий POST-запросы формата JSON.
>
>## **Возможности**
>Калькулятор поддеживает **числа формата float**, а также следующие операторы:
>
> | Оператор | Символ | Особенности |
> | -------- | ------ | -------- |
> | Сложение | + | ... |
> | Вычитание | - | ... |
> | Умножение | * | ... |
> | Деление | / | На ноль делить нельзя |
> | Возведение | ** | Чтобы извлечь корень, возведите в степень 0.5 |
> | Скобки | ( ) | ... |
> ### **Что может калькулятор?**
> Вот математические выражения, которые способен решить калькулятор:
>> + (70/7) * 10 /((3+2) * (3+7))
>> + ((7+1) / (2+2) * 4) / 8 * (32 - ((4+12)*2))
>> + ((1+2)*(5 * (7+3) - 70 / (3+4) * (1+2)) - (8-1)) + (10 * (5-1 * (2+3)))
___
>[!TIP]
># **Как работает калькулятор?**
>Давайте разбёрем код на примере **`2+2*(10-8)`**.
>
>\
>С помощью функции [ParserExpression](Calculator_Golang/pkg/parser.go) извлекаем из математического выражения все числа и операторы, при чём для операторов создан специальный класс:
>```go
>type operator struct {
>	name     string // Тип оператора
>	Index    int    // Его индекс в примере
>	priority int    // Приоритет его выполнения
>}
>```
> функция вернёт: **`[2, 2, 10, 8], [{+ 0 0}, {* 1 1}, {- 2 5}], nil`** список чисел и операторов соответственно. Список операторов будет отсортирован по приоритету с помощью функции [SortOperators](Calculator_Golang/pkg/operator.go): **`[{- 2 5}, {* 1 1}, {+ 0 0}]`**.
>
>\
> Дальше функция [Calc](Calculator_Golang/calc/calc.go) будет пошагово делать операции. Возьмём первый оператор: **`[{- 2 5}`**, с помощью его индекса мы можем получить оперируемые числа:
>```go
>n1, n2 := nums[opIndex], nums[op.Index+1]
>```
>выполним операцию с помощью функции [MakeOperation](Calculator_Golang/pkg/operator.go) и заменим оперируемые числа на результат операции, а использованный оператор уберём из списка. Список чисел и операторов станет выглядеть так: **`[2, 2, 2], [{* 1 1}, {+ 0 0}]`**. Так будет продолжаться до тех пор пока все операции не будут выполнены. Оставшееся число и будет результатом математического выражения.
>
>Так будет выглядеть работа всего калькулятора с самого начала:\
>**`[2 2 10 8] [{+ 0 0} {* 1 1} {- 2 5}]`**\
>**`[2 2 2] [{* 1 1} {+ 0 0}]`**\
>**`[2 4] [{+ 0 0}]`**\
>**`[6] []`**\
>Результат это **`6`**.

>[!IMPORTANT]
># **Инструкция**
>## **Как запустить сервер?**
> + ```shell
>   go run ./cmd/main/main.go
>   ```
> + ```shell
>   go run абсолютный путь до main.go
>   ```
>
>## **Как отправлять POST запросы?**
>### **Требования к запросам:**
>+ Сервер принимает запросы по адресу:
>   ```shell
>   http://localhost:8080/api/v1/calculate
>   ```
>+ Поддерживает JSON следующего вида:
>   ```json
>   {
>       "expression": "EXAMPLE"    
>   }
>   ```
>### **Пример запроса:**
>```shell
>curl -X POST -H "Content-type:application/json" --data "{\"expression\":\"2+2*2\"}" http://localhost:8080/api/v1/calculate
>```

>[!IMPORTANT]
># **Примеры использования:**
> Взависимости от запроса, сервер может вернуть следующий статус:
>
> | Статус | JSON запроса | Ответ сервера | Причина |
> | ---- | --------- | ------ | ------ |
> | 200 | {"expression": "2+2*2"} | {"result": 6} | ... |
> | 422 | {"expression": "2+2*"} | {"error": "Expression is not valid"} | Неккоректный пример
> | 500 | {"aboba": "aboba"} | {"error": "Internal server error"} | Неккоректный JSON или неизвестная ошибка

