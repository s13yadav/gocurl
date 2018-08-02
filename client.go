package main

import (
        "crypto/tls"
        "net/http"
        "io/ioutil"
        "crypto/x509"
        "os"
        "os/exec"
        "time"
)

func CreateSecureClient()(*http.Client, error){

        // load all the tls configuration - Refer tls.Config for more details
        tlsConfig, _ := CreateTlsConfig()

        // create a secure transport  -  Refer http.Transport for more details
        var secure_transport = &http.Transport{
                TLSClientConfig : tlsConfig,
                TLSHandshakeTimeout: time.Duration(*connectTimeout) * time.Second,
        }

        // create secure client  -  Refer http.Client for more details
        var secureclient = &http.Client{
                Transport: secure_transport,
                Timeout: time.Duration(*maxTime) * time.Second,
        }

        return secureclient, nil
}

func CreateTlsConfig() (*tls.Config, error) {
        Logger.Println("loading client certificates...")

        var PublicKeyFile = ClientCert
        var PrivateKeyFile = ClientPrivateKey
        var config tls.Config

        config.Certificates = make([]tls.Certificate, 1)

        publicKey_data, err := ioutil.ReadFile(PublicKeyFile)
        decrypted_Privkey, err := Decryptprivatekey(PrivateKeyFile)

        clientCert, err := tls.X509KeyPair(publicKey_data, decrypted_Privkey)
        if err != nil {
                return nil, err
        }

        config.Certificates[0] = clientCert

        certPool := x509.NewCertPool()

        // read cacerts directory to support multiple signing authority
        var firstErr error
        for _, directory := range CACertDirectories {
                fis, err := ioutil.ReadDir(directory)
                if err != nil {
                        if firstErr == nil && !os.IsNotExist(err) {
                                firstErr = err
                        }
                        continue
                }
                rootsAdded := false
                for _, fi := range fis {
                        data, err := ioutil.ReadFile(directory + "/" + fi.Name())
                        if err == nil && certPool.AppendCertsFromPEM(data) {
                                rootsAdded = true
                        }
                }
                if rootsAdded {
                        config.RootCAs = certPool
                }
        }
        if *insecureFlag == true {
                config.InsecureSkipVerify = true
        } else {
                config.InsecureSkipVerify = false
        }
        return &config, nil
}

func Decryptprivatekey(PrivateKeyFile string) ([]byte, error) {
        output, err := exec.Command("/opt/keyutils/bin/keyutil", "-d", "-P", PrivateKeyFile).Output()
        if err != nil {
                Logger.Println("Error in decrypting the privatekey")
                return nil, err
        }
        return []byte(output), nil
}
