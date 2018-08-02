/*
 * $Copyright:
 * Copyright (C) Nokia 2018
 * All rights reserved
 * Author : Sanjay Yadav
 * $
 *
 */

// Here scurl is designed to work as curl wrapper that loads certificate internally instead of reading from command line
// currently only below options used in the example is supported, it can be enhanced with full use cases in future
// Ex -  scurl -X POST --connect-timeout 120 --retry 5 --retry-delay 5 -H "Authorization: Basic $auth_enc" -d "{"key1":"Value1"}" https://10.255.14.111:5572/api/v1/executecmd

package main

import (
        "flag"
        "io"
        "log"
        "os"
)

var (
        CACert = "/opt/cacerts/cacert.crt"
        ClientCert = "/opt/certs/client.crt"
        ClientPrivateKey = "/opt/certs/clientprivkey.pem"
        CACertDirectories = []string{
                "/opt/cacerts",
        }
        PostData = ""
        logfile = "/tspinst/scurl.log"
        fo io.Writer
        Logger *log.Logger
        respBody []byte
        insecureFlag *bool
        headerOnly *bool
        maxTime *int
        connectTimeout *int
        retryDelay *int
        numRetry *int
)
func InitLogger(){

        if _, err := os.Stat(logfile); err == nil {
                fo, _ = os.OpenFile(logfile, os.O_APPEND|os.O_WRONLY, 0600)
        } else {
                fo, _ = os.Create(logfile)
        }
        Logger = log.New(fo, "", log.Lshortfile|log.Ldate|log.Ltime)

}

func main(){

    InitLogger()

    connectTimeout = flag.Int("connect-timeout", 120, "request timeout in seconds")
    maxTime = flag.Int("max-time", 120, "request timeout in seconds")
    data := flag.String("d", "", "data to be post on remote server")
    numRetry = flag.Int("retry", 3, "num of retry in case of failure")
    retryDelay = flag.Int("retry-delay", 10, "wait before next retry in seconds")
    reqMethod := flag.String("X", "POST", "request method i.e. GET, POST, PUT, DELETE")
    insecureFlag = flag.Bool("k", false, "This  option explicitly allows curl to perform insecure SSL connections and transfers. All SSL connections are attempted to be made secure by using the CA certificate bundle")
    headerOnly = flag.Bool("I", false, "Fetch the HTTP-header only!. When used on an FTP or FILE file, scurl displays the file size and last modification time only.")
    url := flag.String("u", "", "remote server url address")

    flag.Parse()

    if len(os.Args) < 2 {
        printHelp()
        os.Exit(1)
    }

    if *url == "" {
        printHelp()
        os.Exit(1)
    }
    if *data != "" {
        PostData=*data
    }

    Logger.Println("Starting secure curl --- ")

    Logger.Printf("connectTimeout  - %d\n", *connectTimeout)
    Logger.Printf("data  - %s\n", *data)
    Logger.Printf("numRetry  - %d\n", *numRetry)
    Logger.Printf("retryDelay  - %d\n", *retryDelay)
    Logger.Printf("reqMethod  - %s\n", *reqMethod)
    Logger.Printf("insecureFlag  - %v\n", *insecureFlag)
    Logger.Printf("url  - %s\n", *url)

    returnValue := ExecUrlAddress(*url, *reqMethod, PostData)
    os.Exit(returnValue)
}
