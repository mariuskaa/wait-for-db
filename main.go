package main
import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"os/exec"
	"os"
	"flag"
)

func main() {
	const (
		defaultTimeout = 5
		timeoutDescription = "seconds between retries"
	)

	var count int
	flag.IntVar(&count, "timeout", defaultTimeout, timeoutDescription)
	flag.IntVar(&count, "t", defaultTimeout, timeoutDescription)
	flag.Parse()


	if len(flag.Args()) < 2 {
		fmt.Printf("Usage: %s [options] <connection> <command>\n", os.Args[0])
		os.Exit(3)
	}

	var (
		connection = flag.Args()[0]
		command = flag.Args()[1]
	)

	fmt.Println("waiting for database to come online");

	db, err := sql.Open("mysql", connection)
	if err != nil {
		log.Fatal(err)
	}

	for {
		err = db.Ping()
		if err == nil {
			fmt.Println("service is ready to go!")
			cmd := exec.Command(command)
			output, err := cmd.CombinedOutput()
			if err != nil {
				log.Fatal(fmt.Sprint(err) + ":\n    " + string(output));
			}else{
				log.Printf("out: %s\n", output)
			}
			break;
		} else {
			fmt.Println("attempting to restart...")
			time.Sleep(time.Duration(count) * time.Second)
		}
	}
}
