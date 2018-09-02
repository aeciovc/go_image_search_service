package rabbitmq

import(
	"encoding/json"
	"log"
)

func buildJSON(struc interface{}) ([]byte, error) {
	resp, err := json.Marshal(struc)

	if err != nil {
		log.Printf("[rabbitmq] Couldn't marshall object. %s", err.Error())
		return nil, err
	}

	return resp, err
}