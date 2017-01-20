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
    "bytes"
	"time"
)

// data structure authentication request
type signOn struct {
	Username string `json:"username"`
	Password string `json:"password"`
	RelayState string `json:"relayState"`
	options struct {
		MultiFactor bool `json:"multiOptionalFactorEnroll"`
		WarnPWExpire bool `json:"warnBeforePasswordExpired"`
		} `json:"options"`
	context struct {
		DeviceToken string `json:"deviceToken"`
		} `json:"context"`
}


// data structure for registration request
type NewUser struct {
	Profile struct {
		Firstname string  `json:"firstName"`
		Lastname  string  `json:"lastName"`
		Email     string  `json:"email"`
		Login     string  `json:"login"`
		Phone     string  `json:"mobilePhone"`
	} `json:"profile"`
	Credentials struct {
		Password struct {
			Value string  `json:"value"`
		} `json:"password"`
		Recovery_question struct {
			Question string  `json:"question"`
			Answer   string  `json:"answer"`
		} `json:"recovery_question"`
	} `json:"credentials"`
}

//data structure for login response and session data
type oktaSession struct {
	ExpiresAt    time.Time `json:"expiresAt"`
	Status       string `json:"status"`
	SessionToken string `json:"sessionToken"`
	Embedded     struct {
	     User struct {
		  ID              string `json:"id"`
		  PasswordChanged time.Time `json:"passwordChanged"`
		  Profile         struct {
			  Login     string `json:"login"`
			  FirstName string `json:"firstName"`
			  LastName  string `json:"lastName"`
			  Locale    string `json:"locale"`
			  TimeZone  string `json:"timeZone"`
		  } `json:"profile"`
	  } `json:"user"`
    } `json:"_embedded"`
}


// set your okta org URL and API key here
const (
    oktaOrg string = "https://vorton.okta.com"
    oktaKey string = "SSWS 00U_IlbOhvpKH7EA8KwuKYZs2dmnBy47dSaUQr-Zvw"
    authEndPoint string = "/api/v1/authn"
    userEndPoint string = "/api/v1/users?activate=false"
    loginRedirect string = "/home/vannortondemo_samplegolangapp_1/0oaarewi1qL8QyiKg0x7/alnarf8ei1vAq9y9n0x7"
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

	signOnData := signOn{}
	signOnData.Username = r.Form["username"][0]
	signOnData.Password = r.Form["password"][0]
	signOnData.RelayState = "/"
	signOnData.options.MultiFactor = false
	signOnData.options.WarnPWExpire = false
	signOnData.context.DeviceToken = "26q43Ak9Eh04p7H6Nnx0m69JqYOrfVBY"

	jsonReq, _ := json.Marshal(signOnData)

	fmt.Println(bytes.NewBuffer(jsonReq))

//        payload := strings.NewReader("{\n    \"username\": \""+r.Form["username"][0]+"\",\n    " +
//                "\"password\": \""+r.Form["password"][0]+"\",\n    " +
//                "\"relayState\": \""+"/\",\n    " +
//                "\"options\": {\n       " +
//                "\"multiOptionalFactorEnroll\": false,\n       " +
//                "\"warnBeforePasswordExpired\": false \n        },\n    \"context\": {\n          " +
//                "\"deviceToken\": \"26q43Ak9Eh04p7H6Nnx0m69JqYOrfVBY\" \n   }}")

 //       fmt.Println(payload)

        url := oktaOrg + authEndPoint

        req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))

        req.Header.Add("accept", "application/json")
        req.Header.Add("content-type", "application/json")
        req.Header.Add("authorization", oktaKey)
        req.Header.Add("cache-control", "no-cache")
        req.Header.Add("postman-token", "db3d87a4-41d9-eb33-91a6-6881bda94fac")

        res, _ := http.DefaultClient.Do(req)

        defer res.Body.Close()

        body, _ := ioutil.ReadAll(res.Body)

        var loginRes oktaSession

        err := json.Unmarshal(body, &loginRes)

        if res.Status == "200 OK" {

            fmt.Println("SessionToken is ...", loginRes.SessionToken)
            fmt.Println("Status is ...", loginRes.Status)
	    fmt.Println("Login ID is ...", loginRes.Embedded.User.Profile.Login)
	    fmt.Println("Users name is ...", loginRes.Embedded.User.Profile.FirstName, " ", loginRes.Embedded.User.Profile.LastName)

            loginURL := oktaOrg + "/login/sessionCookieRedirect?token=" + loginRes.SessionToken + "&redirectUrl=" + oktaOrg + loginRedirect

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

        newUserData := NewUser{}
	newUserData.Profile.Firstname = r.Form["element_1_1"][0]
	newUserData.Profile.Lastname = r.Form["element_1_2"][0]
	newUserData.Profile.Email = r.Form["element_3"][0]
	newUserData.Profile.Login = r.Form["element_2"][0]
        newUserData.Profile.Phone = r.Form["element_4_1"][0] + "-" + r.Form["element_4_2"][0] + "-" + r.Form["element_4_3"][0]
	newUserData.Credentials.Password.Value = "Password1"
	newUserData.Credentials.Recovery_question.Question = "What is your favorite language"
	newUserData.Credentials.Recovery_question.Answer = "golang"

        jsonReq, _ := json.Marshal(newUserData)

	fmt.Println(bytes.NewBuffer(jsonReq))

        url := oktaOrg + userEndPoint

        req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))

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