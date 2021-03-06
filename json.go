package lymon

import (
	"encoding/json"
	"log"
)

// JSON simplify json.Marshal
func (c Context) JSON(source interface{}) {
	b, err := json.Marshal(source)
	if err != nil {
		log.Println(err)
	}
	c.W.Write(b)
}
