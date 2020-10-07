package main
 
import (
	"encoding/json"
	"io/ioutil"
    "os"
)
 
 
type Data struct {
	SERVER_PORT, SERVER_LISTEN_BACKLOG, LOOP_LAPSE, LOOP_PERIOD,ID,SERVER_ADDRESS string
}
 
func main() {

    CLI_SERVER_ADDRESS:= os.Getenv("CLI_SERVER_ADDRESS")
    SERVER_PORT:= os.Getenv("SERVER_PORT")
    SERVER_LISTEN_BACKLOG:=os.Getenv("SERVER_LISTEN_BACKLOG")
    CLI_LOOP_LAPSE:=os.Getenv("CLI_LOOP_LAPSE")
    CLI_LOOP_PERIOD:=os.Getenv("CLI_LOOP_PERIOD")
    CLI_ID:=os.Getenv("CLI_ID")
    
	data := Data{
        SERVER_ADDRESS: CLI_SERVER_ADDRESS,
        SERVER_PORT: SERVER_PORT,
        SERVER_LISTEN_BACKLOG: SERVER_LISTEN_BACKLOG,
        LOOP_LAPSE: CLI_LOOP_LAPSE,
        LOOP_PERIOD: CLI_LOOP_PERIOD,
        ID: CLI_ID,
    }
 
	file, _ := json.MarshalIndent(data, "", " ")
 
	_ = ioutil.WriteFile("data1/test.json", file, 0644)
    

}