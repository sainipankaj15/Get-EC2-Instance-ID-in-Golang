package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"math/rand"
)

func getInstanceID() string {
	resp, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")
	if err != nil {
	    rand.Seed(time.Now().UnixNano())
	    instanceID := fmt.Sprintf("%d", rand.Int())
	    return instanceID
	}

	defer resp.Body.Close()

	instanceID, err := ioutil.ReadAll(resp.Body)
	if err != nil {
            rand.Seed(time.Now().UnixNano())
	    instanceID := fmt.Sprintf("%d", rand.Int())
            return instanceID
	}

	return string(instanceID)
}

func main() {
	instanceID := getInstanceID()
	fmt.Println("EC2 Instance ID:", instanceID)
}
