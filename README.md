# SETUP
Running avaliable only on *NIX-like systems

## Step 0
Intall Docker, Golang v1.20

## Step 1
Configure .env file
`
POSTGRES_DATABASE="postgres"
POSTGRES_USER="admin"
POSTGRES_PASSWORD="root"
SALT="YOURE SECRET"
`

## Step 2
`docker-compose up`

## Step 3
Intall dependencies
`go mod tidy`

## Step 4
Run the app
`go run cmd/calendarBack/.`
