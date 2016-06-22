package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
	Host string
}

func (c Client) JudgeName(name string) (bool, error) {
	url := "http://" + c.Host + "/api/v1/user/" + name
	log.Println("url is", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return false, err
	}
	log.Println("response status is ", resp.StatusCode)
	if resp.StatusCode == http.StatusOK {
		return true, nil
	}

	return false, nil
}

func (c Client) CreateNS(name, namespace string) (bool, error) {
	url := "http://" + c.Host + "/api/v1/user/" + name + "/" + namespace
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}

	if resp.StatusCode == http.StatusCreated {
		return true, nil
	}

	return false, nil
}

func (c Client) GetNS(name string) ([]string, error) {
	ns := make([]string, 0, 10)
	url := "http://" + c.Host + "/api/v1/user/ns/" + name
	resp, err := http.Get(url)
	if err != nil {
		log.Println("http get error", err)
		return ns, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ReadAll error :", err)
		return ns, err
	}

	err = json.Unmarshal(body, &ns)
	if err != nil {
		log.Println("json Unmarshal  error :", err)
		return ns, err
	}
	return ns, nil
}
