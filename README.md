🔒 P2P Чат с Шифрованием
Децентрализованный мессенджер без серверов


📌 Особенности
Прямое peer-to-peer соединение

Сквозное шифрование (ECDH + AES-GCM)

Сохранение истории (BoltDB)

Работает за NAT

🚀 Запуск
# Установка
git clone https://github.com/ваш-ник/p2p-chat.git

cd p2p-chat

go mod download


# Запуск первого узла

go run main.go -port 8080 -name Alice

# Запуск второго узла

go run main.go -port 8081 -name Bob -connect /ip4/127.0.0.1/tcp/8080/p2p/QmAliceNodeID

📂 Структура

/crypto   - Шифрование

/p2p      - Сетевое взаимодействие

/wal      - Хранение сообщений

main.go   - Точка входа
