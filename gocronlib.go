package gocronlib
import (
      "os"

      "io/ioutil"
      "fmt"
      "os/exec"
      "gopkg.in/yaml.v2"
      "database/sql"; _ "github.com/lib/pq";
)


type Config struct {
      Dbfqdn       string
      Dbport       string
      Dbuser       string
      Dbpass       string
      Dbdatabase   string
      Smtpserver   string
      Smtpport     string
      Smtpaddress  string
      Smtppassword string
      Interval     int
}


type Cron struct {
      Cronname    string   // Name of the cronjob
      Account     string   // Account the job belongs to
      Email       string   // Address to send alerts to
      Ipaddress   string   // Source IP address
      Frequency   string   // How often a job should check in
      Lastruntime string   // Unix timestamp
      Alerted     bool     // set to true if an alert has already been thrown
}


const Version string = "1.0.1"
const sslmode string  = "disable"   // Disable or enable ssl
const syslog string   = "logger"    // Command to write to syslog
const confPath string = "/etc/gocron/config.yml"


// Read in the config file
func GetConfig(verbose bool) Config {
      var config Config
      yamlFile, err := ioutil.ReadFile(confPath)
      if err != nil {
           CheckError(err, verbose)
           os.Exit(1)
      }

      // Set the global config var
      err = yaml.Unmarshal(yamlFile, &config)
      if err != nil {
            CheckError(err, verbose)
            os.Exit(1)
      }

      return config
}


// Return a Postgres connection string
func DatabaseString(verbose bool) string {
      var config Config = GetConfig(verbose)
      var connectionString string = "postgres://" +
            config.Dbuser + ":" +
            config.Dbpass + "@" +
            config.Dbfqdn +
            "/gocron" +
            "?sslmode=" + sslmode

      return connectionString
}


// Function returns true if query is successful
func InsertDatabase(query string, verbose bool) bool {
      db, err := sql.Open("postgres", DatabaseString(verbose))
      if err != nil {
            CheckError(err, verbose)
      }
      defer db.Close()

      _, err = db.Query(query)
      if err != nil {
            CheckError(err, verbose)
            return false
      } else {
            return true
      }
}


// Function returns the result of a SELECT query
func SelectDatabase(query string, verbose bool) (*sql.Rows, bool) {
      var status bool // false if query is not successful

      db, err := sql.Open("postgres", DatabaseString(verbose))
      if err != nil {
            CheckError(err, verbose)
      }
      defer db.Close()

      rows, err := db.Query(query)
      if err != nil {
            CheckError(err, verbose)
            status = false
      } else {
            status = true
      }

      // Return the query result and the status
      return rows, status
}


// Function writes messages to syslog and (optionally) to standard out
func CronLog(message string, verbose bool) {
      err := exec.Command(syslog, message).Run()
      if err != nil {
            fmt.Println("Failed to write to syslog")
            fmt.Println(message)
      }
      if verbose == true {
            fmt.Println(message)
      }
}


// Function passes error messages to the cronLog() function
func CheckError(err error, verbose bool) {
      if err != nil {
            CronLog("Error: \n" + err.Error(), verbose)
      }
}