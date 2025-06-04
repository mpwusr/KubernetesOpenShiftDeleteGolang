package main

import (
    "crypto/tls"
    "crypto/x509"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

func main() {
    apiServer := "https://api.openshift.example.com:6443"
    namespace := "default"
    resource := "deployments"
    name := "example-deployment"
    token := os.Getenv("BEARER_TOKEN")
    caPath := "/path/to/ca.crt"

    // Load CA
    caCert, _ := ioutil.ReadFile(caPath)
    caPool := x509.NewCertPool()
    caPool.AppendCertsFromPEM(caCert)

    client := &http.Client{
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{RootCAs: caPool},
        },
    }

    url := fmt.Sprintf("%s/apis/apps/v1/namespaces/%s/%s/%s", apiServer, namespace, resource, name)
    req, _ := http.NewRequest("DELETE", url, nil)
    req.Header.Set("Authorization", "Bearer "+token)
    req.Header.Set("Accept", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)

    fmt.Printf("Status: %s\nBody: %s\n", resp.Status, string(body))
}
