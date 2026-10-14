package main
//! Marshal() -> from GOLANG to JSON
//! Unmarshal() -> from JSON to GOLANG
import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Todo struct { 
	UserId 	  int `json:"userId"`
	Id 		  int `json:"id"`
	Title 	  string `json:"title"`
	Completed bool `json:"completed"`
}


func main() {

	url := "https://jsonplaceholder.typicode.com/todos/1"
	resp, err := http.Get(url)
	if err != nil {
		//! xatoni ekranga chiqaradi va dasturni toxtadi fmt faqat xatoni ekranga chiqaradi va davom etadi
		log.Fatal(err)
	}

	defer resp.Body.Close()
	//! http.StatusOk -> xar doim 200 qaytaradi
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		todoItem := Todo{}
		todoErr := json.Unmarshal(bodyBytes, &todoItem)
		if todoErr != nil {
			log.Fatal(todoErr)
		}

		// json.NewDecoder(resp.Body).Decode(&todoItem)
		//! json.NewDecoder(resp.Body).Decode(&todoItem)
		//! 31-40 ni qatordagi code qilgan ishni bu xam qiladi 
		//! json.NewCoder -> Json tiliga change qiladi Decode json tilini golang tiliga change qiladi yani todoItem ga yozib qoyadi
		fmt.Printf("Data from API %+v\n",todoItem)

		//! -> from GOLANG to JSON
		todo, err := json.Marshal(&todoItem)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(todo)) 

	}	

}