
# Data Whale

Handling data, data types, data validation,
data visualization, and data sources all in
one platform.

## Project vision and mission
Data Whale is a backend tool for Go developers that streamlines
tasks like handling REST APIs. It merges JSON request/response
definitions, validation, and code generation into one process.
Developers define JSON types and validation logic, and the tool
auto-generates the code. Data Whale enables dynamic API route
creation with validation, including nested structures. It supports
multi-database connections and data pipelines for PostgreSQL,
MySQL, Cassandra, Redis, and MongoDB. Reports are generated in
CSV format using Google APIs and stored on Google Drive with
shareable links. It also features CSV-to-DB conversion without
manual queries and allows filtering JSON arrays based on
conditions.




## Run Locally

Clone the project

```bash
  git clone https://github.com/ManojaD2004/gofr-dev-practice1.git
```

Go to the project directory

```bash
  cd gofr-dev-practice1
```

Install client dependencies

```bash
  cd front-end
  npm install
```
Run client app

```bash
  npm run dev
```

Start the GoFr server, you can also use **air** to run the server

```bash
  cd ../back-end
  go mod tidy
  go build -o ./tmp/main .
  ./tmp/main

```
