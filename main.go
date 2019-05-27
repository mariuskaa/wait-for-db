package main
import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"os/exec"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <connection> <command>\n", os.Args[0])
		os.Exit(3)
	}

	fmt.Println("waiting for service to come online");
	db, err := sql.Open("mysql", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	for {
		err = db.Ping()
		if err != nil {
			fmt.Println("attempting to restart...")
			time.Sleep(10 * time.Second)
		} else {
			fmt.Println("service is ready to go!")
			cmd := exec.Command(os.Args[2])
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println(fmt.Sprint(err) + ": " + string(output));
				log.Fatal("ERROR")
			}else{
				log.Printf("%s\n", output)
			}
			break;
		}
	}
}
