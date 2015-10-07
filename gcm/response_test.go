package gcm

import (
	"encoding/json"
	"testing"
)

func TestResponse(t *testing.T) {
	data := `{ "multicast_id": 216,
  			   "success": 3,
  			   "failure": 3,
  			   "canonical_ids": 1,
  			   "results": [
    			 { "message_id": "1:0408" },
    			 { "error": "Unavailable" },
    			 { "error": "InvalidRegistration" },
    			 { "message_id": "1:1516" },
    			 { "message_id": "1:2342", "registration_id": "32" },
    			 { "error": "NotRegistered"}
  			   ]
			 }`

	response := Response{}
	json.Unmarshal([]byte(data), &response)

	if !response.Results[1].ShouldPostpone() {
		t.Errorf("Expect true but found %t", response.Results[1].ShouldPostpone())
	}

	if !response.Results[4].ShouldUpdate() {
		t.Errorf("Expect true but found %t", response.Results[4].ShouldUpdate())
	}

	if !response.Results[5].ShouldRemove() {
		t.Errorf("Expect true but found %t", response.Results[5].ShouldRemove())
	}
}
