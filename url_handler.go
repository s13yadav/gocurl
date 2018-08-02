package main

import (
        "encoding/base64"
        "fmt"
        "io/ioutil"
        "net/http"
        "os/exec"
        "strings"
        "time"
)

var req *http.Request
var resp *http.Response
var reqError error

func ExecUrlAddress(url string, RequestType string, Data string)(int){
        // create http client
        secureclient, _  := CreateSecureClient()

        // create non secure client in case certificates are not present, this is required for backward compatibility/upgrade

        body:=strings.NewReader(Data)

        Logger.Println("create new request")
        if Data == "" {
                req, _ = http.NewRequest(RequestType, url, nil)
        } else {
                req, _ = http.NewRequest(RequestType, url, body)
        }

        Logger.Printf("New request is - %s\n", *req)

        Logger.Printf("Set Authentication header")
        auth_encoded:=GetAuthHeader()
        req.Header.Add("Authorization", "Basic " + auth_encoded)

        // send request with retry and retry delay
        var i int
        for i = 1; i <= *numRetry; i++ {
                resp, reqError = secureclient.Do(req)
                if reqError != nil {
                        Logger.Printf("Failed to fetch the url %s", reqError.Error())
                } else {
                        break
                }
                if i >= *numRetry {
                        Logger.Println("Max retry done... unable to fetch the data ")
                        return 1
                }
                time.Sleep(time.Duration(*retryDelay) * time.Second)
        }

        defer resp.Body.Close()

        respBody, _ = ioutil.ReadAll(resp.Body)

        // read the response body and close
        Logger.Printf("Response code - %d\n", resp.StatusCode)
        Logger.Printf("%s\n", string(respBody))

        if *headerOnly == true {
                fmt.Printf("%s", resp.Header)
        } else {
                fmt.Printf("%s\n", string(respBody))
        }

        if resp.StatusCode == 200 {
                return 0
        } else {
                return 1
        }
}

func GetAuthHeader() string {
        ckey, err := ReadConf("/etc/ims/cmkeys")
        if err != nil {
                Logger.Printf("\n Failed to read /etc/ims/cmkeys ", err)
        }

        Pass := ckey["KEY1"]
        Time_stamp, _ := exec.Command("date +%s").Output()
        Username := "token#$#KEY1"
        Passcode := fmt.Sprintf("%s#$#%s", Pass, Time_stamp)

        auth := Username + ":" + Passcode

        return base64.StdEncoding.EncodeToString([]byte(auth))
}
