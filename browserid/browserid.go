/*

BrowserID for GO HTTP servlets and for AppEngine Applications written in GO

(c) 2011, Elias Athanasopoulos  <elathan@ics.forth.gr>

This code is free to use in any of your projects.
*/

package browserid

import (
  "fmt"
  "http"
  "strings"
  "io/ioutil"
  "json"
  "appengine"
  "appengine/urlfetch"
)
  
func Verify(r *http.Request) string {
  
  if err := r.ParseForm(); err != nil {
    return "Anonymous"
  }
 
  token := r.FormValue("assertion")
  url := "https://browserid.org/verify"
  bodytype := "application/x-www-form-urlencoded" 
  body := strings.NewReader("assertion=" + token + "&audience=" + r.Host)

  var response_body []byte;
  res, err := http.Post(url, bodytype, body) 
  if err != nil { 
   fmt.Println("err=", err) 
   return "Anonymous" 
  } else {
   response_body, _ = ioutil.ReadAll(res.Body);
   res.Body.Close();
  }

  var f interface{}
  json.Unmarshal(response_body, &f)

  m := f.(map[string]interface{})
  return fmt.Sprintf("%s", m["email"])
}

func AppEngineVerify(r *http.Request) string {
  
  if err := r.ParseForm(); err != nil {
    return "Anonymous"
  }

  token := r.FormValue("assertion")
  url := "https://browserid.org/verify"
  bodytype := "application/x-www-form-urlencoded" 
  body := strings.NewReader("assertion=" + token + "&audience=" + r.Host)

  var response_body []byte;
  c := appengine.NewContext(r)
  client := urlfetch.Client(c)
  res, err := client.Post(url, bodytype, body)
  if err != nil { 
   fmt.Println("err=", err) 
   return "Anonymous" 
  } else {
   response_body, _ = ioutil.ReadAll(res.Body);
   res.Body.Close();
  }

  var f interface{}
  json.Unmarshal(response_body, &f)

  m := f.(map[string]interface{})
  return fmt.Sprintf("%s", m["email"])
}