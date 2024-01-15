#   Pet-project

gRPC сервер для сложения 2х целочисленных значений, т.е. проект без какой-либо реального практического применения 
или пользы для бизнеса, а только для иллюстрации использования (aka `тренеровка на кошках`):
- [gRPC](https://grpc.io/) - RPC framework
- [Protocol Buffers](https://protobuf.dev/) - протокол сериализации структурированных данных
- [Wire](https://github.com/google/wire) — это инструмент генерации кода, который автоматизирует соединение компонентов с помощью внедрения зависимостей.

##  ToDo
- replace dependencies to interfaces
- tests

#### Утилиты
Для удобства рутинных сценариев (кодогенерация proto, wire) используется [Taskfile](https://taskfile.dev/) - это не
касается разработки, просто для удобства, ибо кросс-платформенное решение.

##  TaskFile

```shell
$ task
task: [default] task --list
task: Available tasks for this project:
* install-utils:       Installation utility packages
* lint:                Запуск линтера
* lint-fix:            Запуск линтера с авто-исправлением предупреждений
* upgrade:             Обновление пакетов приложения      (aliases: upg)
* wiregen:             wire di codegen                    (aliases: wg)

```