package main

import (
  "net/http"

  log "github.com/Sirupsen/logrus"
  "github.com/prometheus/client_golang/prometheus/promhttp"

  "flag"
  "fmt"
  "os"
  "io/ioutil"
  "strings"
)

func main() {

  wpConfPtr := flag.String("wpconfig", "", "Path for wp-config.php file of the WordPress site you wish to monitor")

  flag.Parse()

  if *wpConfPtr == "" {
    //no path supplied error
    fmt.Fprintf(os.Stderr, "flag -wpconfig=/path/to/wp-config/ required!\n")
    os.Exit(1)
  } else{
    var wpconfig_file strings.Builder
    wpconfig_file.WriteString(*wpConfPtr)

    if strings.HasSuffix(*wpConfPtr, "/") {
        wpconfig_file.WriteString("wp-config.php")
    }else{
        wpconfig_file.WriteString("/wp-config.php")
    }

    //try to read wp-config.php file from path
    dat, err := ioutil.ReadFile(wpconfig_file.String())
    //check(err)
    if(err != nil){
        panic(err)
    }
    //fmt.Print(string(dat))

    //We must locate with regular expressions the MySQL connection credentials and the table prefix
    //to-do
    
  }

  //This section will start the HTTP server and expose
  //any metrics on the /metrics endpoint.
  http.Handle("/metrics", promhttp.Handler())
  log.Info("Beginning to serve on port :8888")
  log.Fatal(http.ListenAndServe(":8888", nil))
}

