# chain-notes

## Project structure

Предварительно

chain-notes/
│
├── cmd/                     # CLI и исполняемые бинарники
│   ├── chain-cli/           # CLI-утилита
│   └── commit-service/      # Сервис обработки коммитов
│
├── internal/                # Переиспользуемая логика и библиотеки
│   ├── crypto/              # Подписи, хэширование, шифрование
│   ├── model/               # Структуры Commit/Blob/Tree
│   ├── storage/             # Работа с БД (Cockroach)
│   └── transport/           # NATS-обёртки, сериализация
│
├── api/                     # gRPC/protobuf/JSON schema
│   ├── events/              # JSON Schema для сообщений в NATS
│   └── proto/               # Protobuf (если потребуется)
│
├── deploy/                  # Инфраструктура и правила диплоя
│   ├── ansible/             # Ansible-роли и инвентори
│   ├── systemd/             # systemd unit-файлы
│   ├── docker/              # docker-compose или Dockerfile
│   ├── nats/                # конфиги брокера, subject ACL, токены
│   └── README.md            # инструкция по развертыванию
│
├── scripts/                 # Утилиты и вспомогательные скрипты
│   ├── gen-keys.sh          # генерация ключей
│   ├── load-test.sh         # нагрузочное тестирование
│   └── reset-db.sh          # очистка CockroachDB
│
├── .github/                 # CI/CD (GitHub Actions)
│   └── workflows/
│       ├── build.yml        # сборка и тест
│       └── deploy.yml       # CD (ansible, docker, systemd)
│
├── tests/                   # Интеграционные и e2e-тесты
│   └── commit-flow/         # тестовая цепочка коммитов
│
├── README.md
├── Makefile
├── LICENSE
└── go.mod
