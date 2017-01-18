package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
    "github.com/astaxie/beego"
    "io/ioutil"
    "encoding/json"
)

// data structure to hold the response json obeject from AuthN
type Message struct {
    ExpiresAt    string
    Status       string
    RelayState   string
    SessionToken string
}

// set your okta org URL and API key here
const (
    oktaOrg string = "https://vorton.okta.com"
    oktaKey string = "SSWS 00U_IlbOhvpKH7EA8KwuKYZs2dmnBy47dSaUQr-Zvw"
    authEndPoint string = "/api/v1/authn"
    userEndPoint string = "/api/v1/users?activate=false"
)

func appHome(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()  //Parse url parameters passed, then parse the response packet for the POST body (request body)
    // attention: If you do not call ParseForm method, the following data can not be obtained form
    //fmt.Println(r.Form) // print information on server side.
    //fmt.Println("path", r.URL.Path)
    //fmt.Println("scheme", r.URL.Scheme)
    //fmt.Println(r.Form["url_long"])

    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }

    fmt.Fprintf(w, "Hello Vann!") // write data to response
}

func login(w http.ResponseWriter, r *http.Request) {

    fmt.Println("method:", r.Method) //get request method

    if r.Method == "GET" {
        t, _ := template.ParseFiles("html/login.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        // logic part of log in

        payload := strings.NewReader("{\n    \"username\": \""+r.Form["username"][0]+"\",\n    " +
                "\"password\": \""+r.Form["password"][0]+"\",\n    " +
                "\"relayState\": \""+"/\",\n    " +
                "\"options\": {\n       " +
                "\"multiOptionalFactorEnroll\": false,\n       " +
                "\"warnBeforePasswordExpired\": false \n        },\n    \"context\": {\n          " +
                "\"deviceToken\": \"26q43Ak9Eh04p7H6Nnx0m69JqYOrfVBY\" \n   }}")

        fmt.Println(payload)

        url := oktaOrg + authEndPoint

        req, _ := http.NewRequest("POST", url, payload)

        req.Header.Add("accept", "application/json")
        req.Header.Add("content-type", "application/json")
        req.Header.Add("authorization", oktaKey)
        req.Header.Add("cache-control", "no-cache")
        req.Header.Add("postman-token", "db3d87a4-41d9-eb33-91a6-6881bda94fac")

        res, _ := http.DefaultClient.Do(req)

        defer res.Body.Close()

        body, _ := ioutil.ReadAll(res.Body)

        var loginRes Message

        err := json.Unmarshal(body, &loginRes)

        if res.Status == "200 OK" {

            fmt.Println("SessionToken is ...", loginRes.SessionToken)
            fmt.Println("Status is ...", loginRes.Status)

            loginURL := oktaOrg + "/login/sessionCookieRedirect?token=" + loginRes.SessionToken + "&redirectUrl=http://localhost:9090"

            http.Redirect(w, r, loginURL, 301)

        } else {
            fmt.Println("http error code is...", res.Status)
            fmt.Println(err)
            r.Method = "GET"
            login(w,r)
        }
    }
}

func register(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method

    // would like to change this to switch on VERB but GET and POST are fine for now.

    if r.Method == "GET" {
        t, _ := template.ParseFiles("html/register.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()

        // logic part of register

        payload := strings.NewReader("{\n  \"profile\": " +
                "{\n    \"firstName\": \""+r.Form["element_1_1"][0]+"\",\n    " +
                "\"lastName\": \""+r.Form["element_1_2"][0]+"\",\n    " +
                "\"email\": \""+r.Form["element_3"][0]+"\",\n    " +
                "\"login\": \""+r.Form["element_2"][0]+"\"\n  },\n  " +
                "\"credentials\": " +
                "{\n    \"password\" : { \"value\": \""+"Password1"+"\" },\n    " +
                "\"recovery_question\": {\n      " +
                "\"question\": \"What is your favorite language\",\n      " +
                "\"answer\": \"golang\"\n    }}}")

        fmt.Println(payload)

        url := oktaOrg + userEndPoint

        fmt.Println(url)

        req, _ := http.NewRequest("POST", url, payload)

        req.Header.Add("accept", "application/json")
        req.Header.Add("content-type", "application/json")
        req.Header.Add("authorization", oktaKey)
        req.Header.Add("cache-control", "no-cache")
        req.Header.Add("postman-token", "db3d87a4-41d9-eb33-91a6-6881bda94fac")

        res, _ := http.DefaultClient.Do(req)

        defer res.Body.Close()
        body, _ := ioutil.ReadAll(res.Body)

        fmt.Println(res)
        fmt.Println(string(body))

        appHome(w,r)

    }
}
func main() {
    http.HandleFunc("/", appHome) // setting router rule
    http.HandleFunc("/login", login)
    http.HandleFunc("/register", register)
    beego.SetStaticPath("/html", "/html")
    err := http.ListenAndServe(":9090", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}