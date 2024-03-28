# Bitespeed - Identity Reconciliation 

## Overview

This would provide a way to identify and keep track of a customer's identity across multiple purchases.

It would link different orders made with different contact information to the same person.

## Prerequisites

Please ensure that you have the following installed:\
[golang](https://go.dev/doc/install)\
[mysql](https://www.postgresql.org/download/)

## Installation

```bash
git clone https://github.com/bhaveshparsad/Identity_Reconciliation_BiteSpeed
.git
cd Identity_Reconciliation_BiteSpeed
go mod tidy
```

## Setup

1. open the `pkg/config/app.go` file and replace the following values with your own

```go
const (
	DBUserName = "your_username"
	DBPassword = "your_password"
	DBName     = "your_db_name"
	DBHost     = "localhost"
	DBPort     = "5432" // Default PostgreSQL port
)
```

2. Create a database with the name you provided in the previous step

3. Run the following command to start the server

```
go run main.go
```

4. Open Postman and navigate to `http://localhost:8080/`

5. You can now use the API

## Usage

There is only one API endpoint for this project.\
It is a POST request to `http://localhost:8080/identify`

## References

https://bitespeed.notion.site/Bitespeed-Backend-Task-Identity-Reconciliation-53392ab01fe149fab989422300423199
