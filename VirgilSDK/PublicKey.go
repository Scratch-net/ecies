package VirgilSDK

import (
	"net/http"
	"encoding/json"
)



type PublicKeyReponse struct {
	CreatedAt string `json:"created_at"`
	ID        string `json:"id"`
	PublicKey string `json:"public_key"`
	Code string `json:"code"`
}


func (h *ServiceHub) GetPublicKey (id string) (key string){

	client := &http.Client{	}
	req, _ := http.NewRequest("GET", "https://keys.virgilsecurity.com/v3/public-key/"+ id, nil)
	req.Header.Add("X-VIRGIL-ACCESS-TOKEN", h.Token)
	response, _ := client.Do(req)
	defer response.Body.Close()
	obj := PublicKeyReponse{}
	json.NewDecoder(response.Body).Decode(&obj)
	return obj.PublicKey;
}
