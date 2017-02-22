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
//	"github.com/astaxie/beego/session"
//	"github.com/astaxie/beego/session"
)

// data structure authentication request
type signOn struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RelayState string `json:"relayState"`
	Options    struct {
			   MultiFactor  bool `json:"multiOptionalFactorEnroll"`
			   WarnPWExpire bool `json:"warnBeforePasswordExpired"`
		   } `json:"options"`
	Context    struct {
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

type LoginReg struct {
	ErrorMsg string
}

// set your okta org URL and API key here
// these can be provided at build time or in a config
// setting these values during the build is the most secure way of doing it
const (
	oktaOrg string = "https://orton.oktapreview.com"
	oktaKey string = "SSWS 00M5lbj56GhXacnJ4_7SaJ8kOHuoMj05Ftjur96TDQ"
	authEndPoint string = "/api/v1/authn"
	userEndPoint string = "/api/v1/users?activate=true"
	loginRedirect string = "/home/oidc_client/0oa9jl4it7ORWEETT0h7/aln5z7uhkbM6y7bMy0g7"
	groupID string = "00g9jl6vo98RI9pGQ0h7"
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

// This will give us a 44 byte, base64 encoded output
var sessionKey, _ = GenerateRandomString(32)

// Initilize a new CookieStore for session data using a random
// generated 44 byte string
var store = sessions.NewCookieStore([]byte(sessionKey))

// appHomne is the aplications home page which uses homeData as an interface
// to the Home struct for presenting data to the page
func appHome(w http.ResponseWriter, r *http.Request) {

	//Parse url parameters passed, then parse the response packet for the POST body (request body)
	// attention: If you do not call ParseForm method, then r.Form with not contain the URL params
	r.ParseForm()

	// Get the user's session data if any
	// If the user is not logged in then a new session will be initialized with nothing in it
	session, err := store.Get(r, "session-name")

	// print out session data
	// this will be empty if the user has not logged in yet
	fmt.Printf("%v", session)

	// do nothing with the error info from get session
	_ = err

	// if the is no access token in the session by the name token then send the user to the loginAndRegister screen
	// to either login or register a new account
	if session.Values["token"] == nil {

		// print out the fact the user has no session
		fmt.Println("There is no valid end user session!\nRedirecting to login screen")

		// set the http verb to GET so the loginAndRegister form does not submit automatically
		r.Method = "GET"

		// send the user to the loginAndRegister page
		loginAndRegister(w, r)

	// if the user has a value for token in the current session
	} else {

		// print out all the session data debug only and can be removed
		fmt.Println("\nThe token from the session is...", session.Values["token"])
		fmt.Println("The name of the user is...", session.Values["name"])
		fmt.Println("The login ID of the user is...", session.Values["loginID"])

		// print out all the request parameters set back from okta
		// this should be SAMLResponse and RelayState
		for k, v := range r.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", strings.Join(v, ""))
		}

		// initialize the interface to the Hone struct for exporting data to the page
		homeData := Home{}

		// set the title of the page
		homeData.Title = "My Sample Go Application"

		// set the full name of the user from the session
		homeData.Name = session.Values["name"].(string)

		// set the user ID used to login from the session
		homeData.Username = session.Values["loginID"].(string)

		// parse the html temnplate
		t, _ := template.ParseFiles("html/home.html")

		// write the parsed template data and the homeData interface to the browser
		t.Execute(w, homeData)
	}
}

// login will parse form data from loginAndRegister when the user clicks "LOG IN"
func login(w http.ResponseWriter, r *http.Request) {

	// print out the http verb.  this should always be POST
	fmt.Println("method:", r.Method) //get request method

	// This will parse all the form data from the user sent in the request
	r.ParseForm()

	// create a new active session.
	// the same method is used for getting an active session in appHome
	// TODO somthing more descriptive than session-name
	session, err := store.Get(r, "session-name")
	_ = err

	//  set session to expire after one day
	// time here is represented in seconds so 2 days would be 86400 * 2 and so on
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400,
		HttpOnly: true,
	}

	// Initialize the interface to the signOn struct
	signOnData := signOn{}

	// Initialize the data used to build the json request bodyu
	// sing the above interface
	signOnData.Username = r.Form["username"][0]
	signOnData.Password = r.Form["password"][0]
	signOnData.RelayState = "/"
	signOnData.Options.MultiFactor = false
	signOnData.Options.WarnPWExpire = false
	signOnData.Context.DeviceToken = "26q43Ak9Eh04p7H6Nnx0m69JqYOrfVBY"

	// Build the json request body
	jsonReq, _ := json.Marshal(signOnData)

	// print out the json object for debugging
	// this can be removed later
	fmt.Println(bytes.NewBuffer(jsonReq))

	// construct the URL for okta authentication
	// oktaOrg is the URL for the okta org Ex. https://company.okta.com
	// authEndPoint is the rest of the URL needed to hit the AuthN API Ex. /api/v1/authn
	url := oktaOrg + authEndPoint

	// create the new http request object that will call the API with the
	// above URL and the marshaled json
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))

	// add headers to the request including the admin API key
	// the postman token still needs to be dealt with
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", oktaKey)
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "db3d87a4-41d9-eb33-91a6-6881bda94fac")

	// POST the request to the API
	// all data in the request is contained in the req object
	res, httpErr := http.DefaultClient.Do(req)

	// this will leave the http request open until we exit login function
	defer res.Body.Close()

	// this will convert the ioReader object returned form the req and convert it into a byte array
	// that can be unmarsheled into a json response which will represent the initial okta session
	body, _ := ioutil.ReadAll(res.Body)

	// declare a variable of type oktaSession to hole the okta session json
	var loginRes oktaSession

	// unmarshal the []byte body which represents the json returned in the response
	// loginRes is a interface to the oktaSession struct
	json.Unmarshal(body, &loginRes)


	// if the login response has a status of 200 OK then let's prep the user to access the app
	if res.Status == "200 OK" {

		// set some session values to ID the user
		// accessToken, full name, and user ID for now.
		session.Values["token"] =  loginRes.SessionToken
		session.Values["name"] = loginRes.Embedded.User.Profile.FirstName + " " + loginRes.Embedded.User.Profile.LastName
		session.Values["loginID"] = loginRes.Embedded.User.Profile.Login

		// Save it before we write to the response/return from the handler.
		session.Save(r, w)

		// build a string value that represents the URL the user will be redirected to
		// oktaOrg is https://company.okta.com
		// loginRes.Session is the session token returned from the login request
		// IMPORTANT: oktaOrg + loginRedirect is the application embed link from okta which points back to appHome
		loginURL := oktaOrg + "/login/sessionCookieRedirect?token=" + loginRes.SessionToken + "&redirectUrl=" + oktaOrg + loginRedirect

		// redirect the user's browser
		// this allows okta to set the okta session cookies so other applications
		// can be accessed without prompting for reauthentication
		http.Redirect(w, r, loginURL, 301)

	} else {

		// print out the response status from the login response when not "200 OK"
		fmt.Println("http error code is...", res.Status)
		session.AddFlash("Invalid username/password. For help use the \"Forgot Password Link\"")


		// dump any error information set be the httpRequest execution
		fmt.Println(httpErr)

		// set the http verb in the request to GET so the loginAndRegister form is not automatically resubmited
		// creating an endless loop
		r.Method = "GET"

		//  send the user back to the loginAndRegister page to try again
		loginAndRegister(w, r)
	}
}

// This will take the form values from the page and create a new user in okta
func register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method

	// parse register form from request so we can read
	// the calues from r.Form
	r.ParseForm()

	// initialize interface to NewUser struct
	newUserData := NewUser{}

	// set values from the register form to the newUserData interface
	newUserData.Profile.Firstname = r.Form["firstname"][0]
	newUserData.Profile.Lastname = r.Form["lastname"][0]
	newUserData.Profile.Email = r.Form["email"][0]
	newUserData.Profile.Login = r.Form["username"][0]
	newUserData.Profile.Phone = r.Form["phone"][0]
	newUserData.Credentials.Password.Value = r.Form["password"][0]

	// TODO add security question to registration form as we are setting static values for now
	newUserData.Credentials.Recovery_question.Question = "What is your favorite language"
	newUserData.Credentials.Recovery_question.Answer = "golang"

	// create a string array from the groupID in okta that the new user will be put into
	// for applicatoin access
	stringArray := []string {groupID}

	// add the array for groupIDs to the newUserData interface
	// for now this is a single group for a single app
	newUserData.GroupIds = stringArray

	// marshal the NewUser to create a json object used in the add user API call
	jsonReq, _ := json.Marshal(newUserData)

	// print out the json that has been marshaled
	// converted to a []byte for printing
	fmt.Println(bytes.NewBuffer(jsonReq))

	// build the URL to the add user API end point in okta
	// see constants at the top
	url := oktaOrg + userEndPoint

	//  create a new http request object with the []byte from json.Marshal and the url
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))

	// add http header info for the API call including the admin API access key
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", oktaKey)
	req.Header.Add("cache-control", "no-cache")

	// TODO this still needs to be dealt with
	req.Header.Add("postman-token", "db3d87a4-41d9-eb33-91a6-6881bda94fac")

	// make the http call to the API putting the response in res
	res, _ := http.DefaultClient.Do(req)

	// close the connection when the register function exits
	defer res.Body.Close()

	// print out the body of the response after converting it to a []byte
	body, _ := ioutil.ReadAll(res.Body)

	// print out the complete response object for debbuging
	// can be removed
	fmt.Println(res)

	// print out the string representation of the response body
	// TODO this should be Unmarshaled into a json object for getting profile data
	fmt.Println(string(body))

	// log the new user in with the username and password supplied in the registration form
	// TODO we need more data validation here to ensure the registration was successful
	login(w, r)
}

// This will render the page for logging and registering a new user
func loginAndRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		session, err := store.Get(r, "session-name")
		loginScreen := LoginReg{}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// Get the previously flashes, if any.
		if flashes := session.Flashes(); len(flashes) > 0 {
			loginScreen.ErrorMsg = flashes[0].(string)
		}

		t, _ := template.ParseFiles("html/loginAndRegister.html")
		t.Execute(w, loginScreen)
	}
}

// process the logout request when the user clicks logout
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

// all the money is here
func main() {
	http.HandleFunc("/", appHome) // setting router rule
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	http.HandleFunc("/logout", logout)
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("./html"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img"))))
	http.HandleFunc("/loginAndRegister", loginAndRegister)
	beego.SetStaticPath("/html", "/html")
	err := http.ListenAndServeTLS(":9090", "server.crt", "server.key", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}