package main

import (
	"os"
	"io"
	"fmt"
	"log"
	"strconv"
        "net/url"
        "net/http"
        "io/ioutil"

	"github.com/tidwall/gjson"
)

func HTTPGET(URL string, auth string) string{
        reqURL, _ := url.Parse(URL)
        reqHeader := map[string][]string{
                "user-agent":{"mozilla/5.0 (x11; ubuntu; linux x86_64; rv:88.0) gecko/20100101 firefox/88.0"},
                "Accept":{"application/json, text/plain, */*"},
                "Accept-Language": {"en-US,en;q=0.5"},
                "Accept-Encoding": {"gzip, deflate"},
                "Authorization":{auth},
        }
        req := &http.Request{
                Method: "GET",
                URL: reqURL,
                Header: reqHeader,
        }
        res, err := http.DefaultClient.Do(req)
        if err != nil{
                log.Println("Error:",err);
                res,err= http.DefaultClient.Do(req)
        }
        if err != nil{
                log.Println("Error:",err)
        }
        data, _ := ioutil.ReadAll(res.Body)
        res.Body.Close()
        return string(data)
}

func DownloadFile(filepath, url string) error{
        resp, err := http.Get(url)
        if err != nil {return err}
        defer resp.Body.Close()
        out, err := os.Create(filepath)
        if err != nil {return err}
        defer out.Close()
        //log.Println(filepath,"✔️ 下载完成.")
        _, err = io.Copy(out,resp.Body)
        return err
}

func ReadFile(filename string)(string,error){
        content , err := ioutil.ReadFile(filename)
        if err!= nil {log.Fatal(err)}
        return string(content),err
}

func MkdirIfNotExist(path,name string){
        _, err := os.Stat(path+name)
        if os.IsNotExist(err){
                errDir:= os.MkdirAll(path+name,0755)
                if errDir != nil{
                        log.Println(err)
                }
        }
}

func DownloadSingleCourse(path, base, auth string, id int) {
        courseNameLink := fmt.Sprintf("https://%s/v1/courses/%d?nonce=00ea6d21-ae15-4317-a222-416e8d3a5ea5", base, id)
        messageLink := fmt.Sprintf("https://%s/v1/courses/%d/messages?nonce=00ea6d21-ae15-4317-a222-416e8d3a5ea5", base, id)
        courseJson := HTTPGET(courseNameLink,auth)
        courseName := gjson.Get(courseJson, `title`).String()
        courseNameJson := courseName+".json"
        fmt.Println("课程名称:",courseName)
        messages := HTTPGET(messageLink,auth)

	prefix := path+courseName+"/"
	MkdirIfNotExist(prefix, courseName)
        os.WriteFile(prefix+courseNameJson, []byte(messages), 0644)

        filedatas := messages
        fmt.Println(gjson.Get(filedatas, "#").String(),"个文件待下载.")
	filedata := gjson.Parse(filedatas).Array()
        for i:=0; i<len(filedata); i++{
		filename := ""
                s := strconv.Itoa(i)
		switch gjson.Get(filedata[i].String(), "category").String() {
		case "PLAIN_AUDIO":
			filename = fmt.Sprintf(s+".mp3")
	                DownloadFile(prefix+filename, gjson.Get(filedata[i].String(),"attachment.url").String())
		case "PLAIN_IMAGE":
			filename = fmt.Sprintf(s+".png")
	                DownloadFile(prefix+filename, gjson.Get(filedata[i].String(),"attachment.url").String())
		case "PLAIN_TEXT":
			filename = fmt.Sprintf(s+".txt")
			os.WriteFile(prefix+filename, []byte(gjson.Get(filedata[i].String(),"text").String()), 0644)
		}
		log.Println(filename, "✔️ 下载完成.")
        }
}

func DownloadAll(path, base, auth string) {
	courseListLink := fmt.Sprintf("https://%s/v1/courses/?nonce=faae98fd-e329-4039-ac6f-39054c32c664&q=&offset=0&limit=5000&sort=oldest-first", base)
	courseList := HTTPGET(courseListLink, auth)
	idList := gjson.Get(courseList,"courses.#.id").Array()
	for i:=0; i<len(idList); i++ {
		DownloadSingleCourse(path, base, auth, int(idList[i].Int()))
	}
	log.Println("✔️ 下载完成.")
}

func DownloadRange(path, base, auth string) {
	
}
