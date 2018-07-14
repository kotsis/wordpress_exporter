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
  "regexp"
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
    fmt.Printf("Read :%v bytes\n", len(dat))

    //We must locate with regular expressions the MySQL connection credentials and the table prefix
    //define('DB_HOST', 'xxxxxxx');
    r, _ := regexp.Compile(`define\(['"]DB_HOST['"].*?,.*?['"](.*?)['"].*?\);`)
    res := r.FindStringSubmatch(string(dat[:len(dat)]))
    if(res == nil){
        fmt.Fprintf(os.Stderr, "Error could not find DB_HOST in wp-config.php ...\n")
        os.Exit(1)
    }
    db_host := res[1]
    
    //define('DB_NAME', 'xxxxxxx');
    r, _ = regexp.Compile(`define\(['"]DB_NAME['"].*?,.*?['"](.*?)['"].*?\);`)
    res = r.FindStringSubmatch(string(dat[:len(dat)]))
    if(res == nil){
        fmt.Fprintf(os.Stderr, "Error could not find DB_NAME in wp-config.php ...\n")
        os.Exit(1)
    }
    db_name := res[1]

    //define('DB_USER', 'xxxxxxx');
    r, _ = regexp.Compile(`define\(['"]DB_USER['"].*?,.*?['"](.*?)['"].*?\);`)
    res = r.FindStringSubmatch(string(dat[:len(dat)]))
    if(res == nil){
        fmt.Fprintf(os.Stderr, "Error could not find DB_USER in wp-config.php ...\n")
        os.Exit(1)
    }
    db_user := res[1]

    //define('DB_PASSWORD', 'xxxxxxx');
    r, _ = regexp.Compile(`define\(['"]DB_PASSWORD['"].*?,.*?['"](.*?)['"].*?\);`)
    res = r.FindStringSubmatch(string(dat[:len(dat)]))
    if(res == nil){
        fmt.Fprintf(os.Stderr, "Error could not find DB_PASSWORD in wp-config.php ...\n")
        os.Exit(1)
    }
    db_password := res[1]

    //$table_prefix  = 'wp_';
    r, _ = regexp.Compile(`$table_prefix.*?=.*?['"](.*?)['"];`)
    res = r.FindStringSubmatch(string(dat[:len(dat)]))
    if(res == nil){
        fmt.Fprintf(os.Stderr, "Error could not find $table_prefix in wp-config.php ...\n")
        os.Exit(1)
    }
    table_prefix := res[1]
  }

  //This section will start the HTTP server and expose
  //any metrics on the /metrics endpoint.
  http.Handle("/metrics", promhttp.Handler())
  log.Info("Beginning to serve on port :8888")
  log.Fatal(http.ListenAndServe(":8888", nil))
}

