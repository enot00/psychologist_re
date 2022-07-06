package validators

import (
	"encoding/json"
	"github.com/test_server/internal/infra/http/resources"
	"log"
)

func ValidationWorkingHours(bodyRead []byte) ([]resources.WorkingHoursDTO, error) {
	var wh []resources.WorkingHoursDTO
	err := json.Unmarshal(bodyRead, &wh)

	if err != nil {
		log.Printf("ValidationWorkingHoursUpdate: %s", err)
		return wh, err
	}
	return wh, nil
}
