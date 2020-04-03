package allocator

import (
	"dynamic-game/config"
	"io/ioutil"
	"net/http"
	"strings"
)

func getHttpClient() *http.Client {
	c := &http.Client{}
	return c
}

func getURL() string {
	url := "http://" + config.Config.Allocator_Addr
	return url
}

func RequestAllocate() (statusCode int, serverID, version string, err error) {
	url := getURL()
	client := getHttpClient()
	urlHandle := url + "/" + HANDLE_ALLOCATE
	req, _ := http.NewRequest("GET", urlHandle, strings.NewReader("name= cbj"))
	req.Header.Set(HEADER_FLEET, config.Config.FleetName)
	req.Header.Set(HEADER_NAMESPACE, config.Config.NameSpace)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	statusCode = resp.StatusCode
	serverID = resp.Header.Get("serverID")
	version = resp.Header.Get("version")
	return
}

func RequestDelete(key string, value string) (statusCode int, body string, err error) {
	url := getURL()
	client := getHttpClient()
	urlHandle := url + "/" + HANDLE_DELETE
	req, _ := http.NewRequest("DELETE", urlHandle, strings.NewReader("name= cbj"))
	req.Header.Set(key, value)
	req.Header.Set(HEADER_FLEET, config.Config.FleetName)
	req.Header.Set(HEADER_NAMESPACE, config.Config.NameSpace)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	statusCode = resp.StatusCode
	read, err := ioutil.ReadAll(resp.Body)
	body = string(read)
	defer resp.Body.Close()
	return
}

func RequestImageVersion() (statusCode int, version string, err error) {
	url := getURL()
	client := getHttpClient()
	urlHandle := url + "/" + HANDLE_VERSION
	req, _ := http.NewRequest("GET", urlHandle, strings.NewReader("name= cbj"))
	req.Header.Set(HEADER_FLEET, config.Config.FleetName)
	req.Header.Set(HEADER_NAMESPACE, config.Config.NameSpace)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	statusCode = resp.StatusCode
	defer resp.Body.Close()
	version = resp.Header.Get("version")
	return
}

func RequestGameVersion(name string) (statusCode int, version string, err error) {
	url := getURL()
	client := getHttpClient()
	urlHandle := url + "/" + HANDLE_GS_VERSION
	req, _ := http.NewRequest("GET", urlHandle, strings.NewReader("name= cbj"))
	req.Header.Set(HEADER_NAMESPACE, config.Config.NameSpace)
	req.Header.Set("Name", name)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	statusCode = resp.StatusCode
	defer resp.Body.Close()
	version = resp.Header.Get("version")
	return
}
