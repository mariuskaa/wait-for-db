# wait-for-db

Go application used as an aid to control startup order in Docker-Compose.
*wait-for-db* will attempt to establish a connection to the specified database and run the command as soon as the db is online.

You may use this as an runner in a ```docker-compose.yml``` to ensure the specified database is available before attempting to start your service.
## Install: 
```bash
go get github.com/mariuskaa/wait-for-db
```
## Usage:
```bash
wait-for-db <connection> <command>
```
Example usage:
```bash
wait-for-db "root:password@tcp(127.0.0.1:3306)/dbname" start-my-application
```
