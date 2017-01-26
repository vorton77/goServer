package main

// TODO decode saml in app
// TODO add session persistance

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
	"github.com/gorilla/sessions"
	"encoding/base64"
	"crypto/rand"
)

// data structure authentication request
type signOn struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RelayState string `json:"relayState"`
	options    struct {
			   MultiFactor  bool `json:"multiOptionalFactorEnroll"`
			   WarnPWExpire bool `json:"warnBeforePasswordExpired"`
		   } `json:"options"`
	context    struct {
			   DeviceToken string `json:"deviceToken"`
		   } `json:"context"`
}


// data structure for registration request
type NewUser struct {
	Profile     struct {
			    Firstname string  `json:"firstName"`
			    Lastname  string  `json:"lastName"`
			    Email     string  `json:"email"`
			    Login     string  `json:"login"`
			    Phone     string  `json:"mobilePhone"`
		    } `json:"profile"`
	Credentials struct {
			    Password          struct {
						      Value string  `json:"value"`
					      } `json:"password"`
			    Recovery_question struct {
						      Question string  `json:"question"`
						      Answer   string  `json:"answer"`
					      } `json:"recovery_question"`
		    } `json:"credentials"`
	GroupIds []string `json:"groupIds"`
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

type Home struct {
	Title string
	Name string
	Username string
}


// set your okta org URL and API key here
const (
	oktaOrg string = "https://vorton.okta.com"
	oktaKey string = "SSWS 00U_IlbOhvpKH7EA8KwuKYZs2dmnBy47dSaUQr-Zvw"
	authEndPoint string = "/api/v1/authn"
	userEndPoint string = "/api/v1/users?activate=true"
	loginRedirect string = "/home/vannortondemo_samplegolangapp_1/0oaarewi1qL8QyiKg0x7/alnarf8ei1vAq9y9n0x7"
	groupID string = "00gauck8lf1xT4Vcr0x7"
)

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.

func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

// Example: this will give us a 44 byte, base64 encoded output
var sessionKey, _ = GenerateRandomString(32)

var store = sessions.NewCookieStore([]byte(sessionKey))


func appHome(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	//Parse url parameters passed, then parse the response packet for the POST body (request body)
	// attention: If you do not call ParseForm method, the following data can not be obtained form

	//fmt.Println(r.Form) // print information on server side.
	//fmt.Println("path", r.URL.Path)
	//fmt.Println("scheme", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])

	session, err := store.Get(r, "session-name")
	fmt.Printf("%v", session)
	_ = err

	if session.Values["token"] == nil {
		fmt.Println("There is no valid end user session!\nRedirecting to login screen")
		r.Method = "GET"
		loginAndRegister(w, r)
	} else {

		fmt.Println("\nThe token from the session is...", session.Values["token"])
		fmt.Println("The name of the user is...", session.Values["name"])
		fmt.Println("The login ID of the user is...", session.Values["loginID"])

		for k, v := range r.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", strings.Join(v, ""))
		}
		homeData := Home{}
		homeData.Title = "My Sample Go Application"
		homeData.Name = session.Values["name"].(string)
		homeData.Username = session.Values["loginID"].(string)
		t, _ := template.ParseFiles("html/home.html")
		t.Execute(w, homeData)

//		httpOutput := "Hello " + session.Values["name"].(string) + ", you are loged in as " + session.Values["loginID"].(string)
//		fmt.Fprintf(w, httpOutput ) // write data to response
	}
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

			// get an active session or create a new one.
			session, err := store.Get(r, "session-name")
			_ = err

			//  set session to expire after one day
			session.Options = &sessions.Options{
				Path:     "/",
				MaxAge:   86400,
				HttpOnly: true,
			}

			// set some session values to ID the user
			session.Values["token"] =  loginRes.SessionToken /*set value*/
			session.Values["name"] = loginRes.Embedded.User.Profile.FirstName + " " + loginRes.Embedded.User.Profile.LastName
			session.Values["loginID"] = loginRes.Embedded.User.Profile.Login

			// Save it before we write to the response/return from the handler.
			session.Save(r, w)

			loginURL := oktaOrg + "/login/sessionCookieRedirect?token=" + loginRes.SessionToken + "&redirectUrl=" + oktaOrg + loginRedirect

			http.Redirect(w, r, loginURL, 301)

		} else {
			fmt.Println("http error code is...", res.Status)
			fmt.Println(err)
			r.Method = "GET"
			loginAndRegister(w, r)
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
		newUserData.Profile.Firstname = r.Form["firstname"][0]
		newUserData.Profile.Lastname = r.Form["lastname"][0]
		newUserData.Profile.Email = r.Form["email"][0]
		newUserData.Profile.Login = r.Form["username"][0]
		newUserData.Profile.Phone = r.Form["phone"][0]
		newUserData.Credentials.Password.Value = r.Form["password"][0]
		newUserData.Credentials.Recovery_question.Question = "What is your favorite language"
		newUserData.Credentials.Recovery_question.Answer = "golang"
		stringArray := []string {groupID}
		newUserData.GroupIds = stringArray

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

		login(w, r)

	}
}

func loginAndRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("html/loginAndRegister.html")
		t.Execute(w, nil)
	}
}

func logout(w http.ResponseWriter, r *http.Request){
	url := oktaOrg + "/login/signout"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	// get an active session or create a new one.
	session, err := store.Get(r, "session-name")
	_ = err

	session.Values["token"] = nil

	session.Options = &sessions.Options{
		MaxAge:   -1,
	}

	session.Save(r, w)

	r.Method = "GET"
	loginAndRegister(w, r)
}

func main() {
	http.HandleFunc("/", appHome) // setting router rule
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	http.HandleFunc("/logout", logout)
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("./html"))))
	http.HandleFunc("/loginAndRegister", loginAndRegister)
	beego.SetStaticPath("/html", "/html")
	err := http.ListenAndServeTLS(":9090", "server.crt", "server.key", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}