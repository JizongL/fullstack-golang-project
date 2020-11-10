package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)



func TestSecureHeaders(t *testing.T){
	
	
	rr := httptest.NewRecorder()
	

	// is for general path middleware test
	r, err := http.NewRequest(http.MethodGet, "/", nil) 
	if err != nil {
		t.Fatal(err) }

	
	if err!=nil{	
	t.Fatal(err) }
		//step 1 setup handlers sandwitch for target middleware
		
	//step 1 setup handlers sandwitch for target middleware
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK")) })
		
	secureHeaders(next).ServeHTTP(rr, r)
	

	//step 2 get result and start real tests
	rs := rr.Result()
	frameOptions := rs.Header.Get("X-Frame-Options")
	if frameOptions != "deny" {
	t.Errorf("want %q; got %q", "deny", frameOptions) }

	xssProtection := rs.Header.Get("X-XSS-Protection")
	if xssProtection != "1; mode=block" {
	t.Errorf("want %q; got %q", "1; mode=block", xssProtection) }

	if rs.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, rs.StatusCode) }
	// step 3 
	defer rs.Body.Close()
	body, err := ioutil.ReadAll(rs.Body) 
	if err != nil {
	t.Fatal(err) }
	if string(body) != "OK" {
	t.Errorf("want body to equal %q", "OK")
	}
}


func TestRequireAuthentication(t *testing.T){
	testUser:="test2"
	email:="test2@gmail.com"
	password:="2020Test##"
	
	app:= newTestApplication(t)
	ts:= newTestServer(t,app.routes())
	defer ts.Close()


	
	
	// authenticate(next)


	
	 app.users.Insert(testUser, email, password)
	//  id,err := app.users.Authenticate(email,password)
	//  if(err!=nil){
	// 	 fmt.Println(err)
	//  }
	//  fmt.Println(id,"test id")
	//  app.session.Put(rs, "authenticatedUserID", id)


	//  next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("OK")) })
	
	
	// authenticate(next)
	// requireAuthentication(next).ServeHTTP(rr, r)	
	
	//step 2 get result and start real tests
	

	// cacheControl := rs.Header.Get("Cache-Control")
	// if cacheControl != "no-store" {
	// t.Errorf("want %q; got %q", "no-store", cacheControl) }

	

	// defer rs.Body.Close()
	// body, err := ioutil.ReadAll(rs.Body)
	
	// if err != nil{
	// 	t.Fatal(err)
	// }
	// if string(body) != "OK"{
	// 	t.Errorf("want body to equal %q","OK")
	// }
}
