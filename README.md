# goServer
This my first go server application using okta APIs and Bootstrap

To try it out you need modify the constants for your okta org, your okta API key, application embed URL
and the groupID that gives access to the app.  

Add an okta group and note the ID

Add a saml template app that points to https://localhost:9090

Add the okta saml template app to the okta group

You'll need the groupID and the application embed link from your okta org

Next you'll need to generate a server key and put it in the same directory where goServer will run

openssl genrsa -out server.key 2048

Next you'll need to use the sew server key to server certificate.  The certificate needs to be in the same directory where goServer will run.

openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650

Give it a shot!
