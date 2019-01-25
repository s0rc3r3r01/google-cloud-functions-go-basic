package functions

import ( 
	"log"
	"net/http"
	"github.com/s0rc3r3r01/google-cloud-functions-go-basic/common"
)

func OutputMessage(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	message := query.Get("message")
	if message == ""{
			message = "No message"
	}
	jw := common.NewMessageWriter(message)
	jsonString, err := jw.JSONString()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonString))
}
