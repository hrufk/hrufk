package mytests
import (
 "fmt"
 _ "github.com/mattn/go-sqlite3"
 "database/sql"
)
 
func main() {
		db, err := sql.Open("sqlite3", "db.db")
		if err!=nil {
			fmt.Println(err)
		}
		pr:= `
				PRAGMA encoding = "UTF-8";
				`
		_, err = db.Exec(pr);
				if err!=nil {
			fmt.Println(err)
		}
		fmt.Println("1111111111");
		fmt.Scanln()
}