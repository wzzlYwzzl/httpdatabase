package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/wzzlYwzzl/httpdatabase/resource/user"
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
	log.Println(body)
	err = json.Unmarshal(body, &ns)
	if err != nil {
		log.Println("json Unmarshal  error :", err)
		return ns, err
	}
	return ns, nil
}

func (c Client) GetNSAll(name string) ([]string, error) {
	ns := make([]string, 0, 10)
	url := "http://" + c.Host + "/api/v1/user/ns/all/" + name
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

func (c Client) GetAllInfo(name string) (*user.User, error) {
	user := new(user.User)
	url := "http://" + c.Host + "/api/v1/user/allinfo/" + name
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return user, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ReadAll in func GetAllInfo err: ", err)
		return user, err
	}

	err = json.Unmarshal(body, user)
	if err != nil {
		log.Println("json Unmarshal err: ", err)
		return user, err
	}

	return user, nil
}

func (c Client) DeleteUser(name string) (bool, error) {
	client := &http.Client{}
	url := "http://" + c.Host + "/api/v1/user/" + name
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Println(err)
		return false, err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return false, err
	}

	if resp.StatusCode == http.StatusOK {
		return true, nil
	}

	return false, nil
}
