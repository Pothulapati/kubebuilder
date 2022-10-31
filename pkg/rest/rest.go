package rest

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	installerv1alpha1 "github.com/mrsimonemms/kubebuilder/api/v1alpha1"
)

type clientsOnboard struct {
	ClientsOnboard int64 `json:"clientsOnboard,omitempty"`
}

func BindClient(clientResource *installerv1alpha1.Config, IP string) bool {
	data := map[string]string{
		"clientId": clientResource.Spec.ClientId,
		"IP":       IP,
	}

	jsonValue, _ := json.Marshal(data)
	resp, err := http.Post("http://"+IP+":8080/addClient", "application/json", bytes.NewBuffer(jsonValue))

	if err == nil && resp.StatusCode == http.StatusOK {
		return true
	} else {
		return false
	}
}

func HasClients(clientResource *installerv1alpha1.Config, IP string) bool {
	resp, err := http.Get("http://" + IP + ":8080/hasClients")

	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	var co = clientsOnboard{}
	json.Unmarshal(body, &co)

	if err == nil && co.ClientsOnboard > 0 {
		return true
	} else {
		return false
	}
}

func GetClient(clientResource *installerv1alpha1.Config, IP string) bool {
	resp, err := http.Get("http://" + IP + ":8080/client/" + clientResource.Spec.ClientId)

	if err == nil && resp.StatusCode == http.StatusOK {
		return true
	} else {
		return false
	}
}
