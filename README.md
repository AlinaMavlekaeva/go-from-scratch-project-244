### Hexlet tests and linter status:
[![Actions Status](https://github.com/AlinaMavlekaeva/go-from-scratch-project-244/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/AlinaMavlekaeva/go-from-scratch-project-244/actions)


## genDiff

### About

gendiff - CLI утилита, предназначенная для сравнения двух JSON или YAML файлов. Утилита поддерживает 3 формата вывода результатов сравнения.

**Доступные флаги:**
1. **--format (-f)** - флаг выбора форматирования, доступные варианты: "stylish", "plain", "json".
2. **--help (-h)** - выводит на экран описание утилиты. 

### Install

Для установки утилиты нужно выполнить следующую команду:
```
go install github.com/AlinaMavlekaeva/go-from-scratch-project-244@latest
```

Проверить работу утилиты можно запустив команду:
```
gendiff -h
``` 
В результате должна появиться информационная справка вида:
```
NAME:
   gendiff - Compares two configuration files and shows a difference.

USAGE:
   gendiff [global options]

GLOBAL OPTIONS:
   --format string, -f string  output  format: 'stylish', 'plain', 'json'. (default: "stylish")
   --help, -h                  show help
   ```

 ### Examples

Для корректной работы утилиты ей необходимо передать относительный путь до сравниваемых файлов и при необходимости желаемый формат вывода через флаг "--format". 

**Работа утилиты без передачи флагов. Вывод по умолчания в формате "stylish".**

Команда:   
 ```
 gendiff ./testdata/fixture/file1.json ./testdata/fixture/file2.json
 ```            
Вывод:      
```
{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}
```

**Вывод данных в формате "plain"**

Команда:
```
gendiff ./testdata/fixture/file1.json ./testdata/fixture/file2.json -f plain
```        
Вывод: 
```
Property 'follow' was removed
Property 'proxy' was removed
Property 'timeout' was updated. From 50 to 20
Property 'verbose' was added with value true
```

**Вывод данных в формате "json"**

Команда:
```
gendiff ./testdata/fixture/file1.json ./testdata/fixture/file2.json -f json
```        
Вывод: 
```
{
  "key": "",
  "status": "",
  "oldValue": null,
  "value": [
    {
      "key": "follow",
      "status": "removed",
      "oldValue": null,
      "value": false,
      "deep": 1
    },
    {
      "key": "host",
      "status": "same",
      "oldValue": null,
      "value": "hexlet.io",
      "deep": 1
    },
    {
      "key": "proxy",
      "status": "removed",
      "oldValue": null,
      "value": "123.234.53.22",
      "deep": 1
    },
    {
      "key": "timeout",
      "status": "updated",
      "oldValue": 50,
      "value": 20,
      "deep": 1
    },
    {
      "key": "verbose",
      "status": "added",
      "oldValue": null,
      "value": true,
      "deep": 1
    }
  ],
  "deep": 0
}
```
**Пример работы утилиты с YAML файлами**

Команда: 
```
./bin/gendiff ./testdata/fixture/filepath1.yml ./testdata/fixture/filepath2.yml
```      
Вывод: 
```
{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}
```
**Пример использвания :**
[https://asciinema.org/a/6ekEuy0b6GiSBQlL](https://asciinema.org/a/nCnn4kVT1MHV18Ae)



