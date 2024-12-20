# Sea


## 📦 Описание

Sea - менеджер архитектур для ваших проектов, расчитан на Go, но возможен и для других ящыков

## 🚀 Начало работы

Инструкции по установке и настройке Sea для локальной разработки.

### 🎈 Предварительные требования


- [Go](https://go.dev/) 

### 📦 Установка

1. Клонируйте репозиторий: ```git clone https://gitlab.com/osamikoyo/sea.git```
2. Выполните команду:  ```go install cmd/sea/sea.go```

### 💧 Использование 
- Создайте папку в домашней дериктории с именем - ```sea create``` 
- Переместите туда .toml файл с описанием шаблона архитектуры(см. ниже) ```sea install {filename.toml}```
- Теперь вы можете использовать этот шаблон в директории, в которой хотите создать эту архитектуру ```sea search {name без .toml} {название проекта}```
## Шаблоны 

    directories = ["dir1", "dir2", "dir3"]
    files = ["dir1/file1.txt", "file2.txt"]
    comands = ["echo hello", "command2"]

    deps = ["gorm.io/gorm", "github.com/go-chi/chi/v5"]

    [[contents]]
    file = "dir1/file1.txt"
    data = "hello"

    [[contents]]
    file = "file2.txt"
    data = "fd"
### directories
- Имена директорий, которые вы хотите видеть в проекте

### files
- Имена файлов, которые вы хотите видеть в проекте, указывайте полный путь от корня проекта 

### commands
- Команды для терминала, которые вы хотите использовать при парсинге шаблона в архитектуру ```!!!ВАЖНО!!! команда go mod init вшита в проект, указывать её ненужно```

### deps
- Зависимости проекта, подключаемые с помощью ```go get```

### contents
- Текст для каких либо файлов, в поле ```file``` указывается полное имя файла от корня проекта, а в поле ```data``` может быть как одностроковый ```"hello"```, так и многостроковый```"""hello"""``` ввод
- Символ ```$``` на уровне парсинга заменяется на имя вашего проекта, которое пользователь ввод как аргумент при команде ```sea search templname name```

