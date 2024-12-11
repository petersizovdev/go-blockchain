# Blockchain System in Go

<details>
<summary>RU</summary>

1. [Описание проекта](#описание-проекта)
2. [Основные компоненты](#основные-компоненты)
3. [Установка и запуск](#установка-и-запуск)
4. [Использование](#использование)
5. [Пример работы](#пример-работы)

## Описание проекта

Этот проект демонстрирует базовую реализацию блокчейна, который состоит из блоков, содержащих данные, хеши предыдущих блоков и nonce для Proof of Work. Блокчейн хранится в базе данных BadgerDB, что обеспечивает эффективное хранение и доступ к данным.
Проект также включает в себя интерфейс командной строки (CLI), который позволяет добавлять новые блоки в цепочку и просматривать существующие блоки.

## Основные компоненты

### 1. **Block (Блок)**

Каждый блок содержит:
- `Hash`: Хеш блока.
- `Data`: Данные, хранящиеся в блоке.
- `PrevHash`: Хеш предыдущего блока.
- `Nonce`: Значение, используемое для Proof of Work.

### 2. **BlockChain (Цепочка блоков)**

Цепочка блоков состоит из последовательности блоков, связанных хешами. Блокчейн хранится в базе данных BadgerDB.

### 3. **ProofOfWork (Proof of Work)**

Proof of Work — это механизм, который обеспечивает безопасность блокчейна. Он требует, чтобы для создания нового блока потребовалось выполнение вычислительной работы. В данном проекте используется алгоритм SHA-256 для генерации хешей.

### 4. **CLI (Интерфейс командной строки)**

Интерфейс командной строки позволяет пользователю взаимодействовать с блокчейном. Пользователь может добавлять новые блоки и просматривать существующие блоки в цепочке.

## Установка и запуск

1. **Клонируйте репозиторий:**

   ```bash
   git clone https://github.com/ваш-пользователь/go-blockchain.git
   cd go-blockchain
   ```

2. **Установите зависимости:**

   Убедитесь, что у вас установлен Go (рекомендуется версия 1.16 или выше). Зависимости проекта управляются с помощью Go Modules.

   ```bash
   go mod tidy
   ```

3. **Запустите проект:**

   ```bash
   go run main.go
   ```

## Использование

### Добавление нового блока

Чтобы добавить новый блок в цепочку, используйте команду `add`:

```bash
go run main.go add -block "Ваши данные"
```

### Просмотр цепочки блоков

Чтобы просмотреть все блоки в цепочке, используйте команду `print`:

```bash
go run main.go print
```

## Пример работы

```bash
$ go run main.go add -block "Первый блок"
Block added

$ go run main.go add -block "Второй блок"
Block added

$ go run main.go print
PrevHash:
BlockData: Genesis
BlockHash: 0000000000000000000000000000000000000000000000000000000000000000
PoW: true

PrevHash: 0000000000000000000000000000000000000000000000000000000000000000
BlockData: Первый блок
BlockHash: 0000000000000000000000000000000000000000000000000000000000000001
PoW: true

PrevHash: 0000000000000000000000000000000000000000000000000000000000000001
BlockData: Второй блок
BlockHash: 0000000000000000000000000000000000000000000000000000000000000002
PoW: true
```

---

</details>

<details>
<summary>ENG</summary>

1. [Project Description](#project-description)
2. [Key Components](#key-components)
3. [Installation and Setup](#installation-and-setup)
4. [Usage](#usage)
5. [Example](#example)

## Project Description

This project demonstrates a basic implementation of a blockchain, consisting of blocks containing data, previous block hashes, and a nonce for Proof of Work. The blockchain is stored in a BadgerDB database, ensuring efficient storage and access to data.
The project also includes a command-line interface (CLI) that allows users to add new blocks to the chain and view existing blocks.

## Key Components

### 1. **Block**

Each block contains:
- `Hash`: The hash of the block.
- `Data`: The data stored in the block.
- `PrevHash`: The hash of the previous block.
- `Nonce`: The value used for Proof of Work.

### 2. **BlockChain**

The blockchain consists of a sequence of blocks linked by hashes. The blockchain is stored in a BadgerDB database.

### 3. **ProofOfWork**

Proof of Work is a mechanism that ensures the security of the blockchain. It requires computational work to be performed to create a new block. In this project, the SHA-256 algorithm is used to generate hashes.

### 4. **CLI (Command-Line Interface)**

The command-line interface allows users to interact with the blockchain. Users can add new blocks and view existing blocks in the chain.

## Installation and Setup

1. **Clone the repository:**

   ```bash
   git clone https://github.com/your-username/go-blockchain.git
   cd go-blockchain
   ```

2. **Install dependencies:**

   Ensure you have Go installed (recommended version 1.16 or higher). Dependencies are managed using Go Modules.

   ```bash
   go mod tidy
   ```

3. **Run the project:**

   ```bash
   go run main.go
   ```

## Usage

### Adding a New Block

To add a new block to the chain, use the `add` command:

```bash
go run main.go add -block "Your data"
```

### Viewing the Blockchain

To view all blocks in the chain, use the `print` command:

```bash
go run main.go print
```

## Example

```bash
$ go run main.go add -block "First block"
Block added

$ go run main.go add -block "Second block"
Block added

$ go run main.go print
PrevHash:
BlockData: Genesis
BlockHash: 0000000000000000000000000000000000000000000000000000000000000000
PoW: true

PrevHash: 0000000000000000000000000000000000000000000000000000000000000000
BlockData: First block
BlockHash: 0000000000000000000000000000000000000000000000000000000000000001
PoW: true

PrevHash: 0000000000000000000000000000000000000000000000000000000000000001
BlockData: Second block
BlockHash: 0000000000000000000000000000000000000000000000000000000000000002
PoW: true
```

</details>
