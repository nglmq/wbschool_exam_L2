Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
nil
false

err это интерфейс. Интерфейс это структура из двух полей: itab и data. Интерфейс равен nil, если оба поля равны nil.
 
tab - указатель на структуру itab, которая хранит данные о типе и списке методов
data - указатель на фактическую переменную с конкретным типом

В данном случае data равен nil, но itab не равен nil, поэтому err != nil
```