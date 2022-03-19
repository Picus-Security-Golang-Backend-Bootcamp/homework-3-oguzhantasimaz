## Steps

### Clone the project
`git clone https://github.com/Picus-Security-Golang-Backend-Bootcamp/homework-2-oguzhantasimaz.git`

### To run the project
Change directory into project folder

### Change MySQLDB connection information from code
Change MYSQLDB connection information which is in main.go, line 34

### Start start.sh script
Write `./start.sh` code snippet to the terminal/command-line.

#### If sh script doesn`t work
Change directory to
`cd /cmd/homework-3-oguzhantasimaz`

##### Search for a book
`go run . search <Book Id || Book Title>`
##### List all books
`go run . list`
##### Buy a book
`go run . buy <id> <count>`
##### Delete a book
`go run . delete <id>`
