# gocronlib
A library that supports common functions for the gocron service.
https://github.com/jsirianni/gocron

## Structs
### Config
Used when goron reads the config.yml file.
```
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
```

### Cron
Used to assign query string parameters
```
type Cron struct {
      Cronname    string   // Name of the cronjob
      Account     string   // Account the job belongs to
      Email       string   // Address to send alerts to
      Ipaddress   string   // Source IP address
      Frequency   int      // How often a job should check in    
      Lastruntime int      // Unix timestamp                     
      Alerted     bool     // set to true if an alert has already been thrown
      Site        bool     // Set true if service is a site (Example: Network gateway)
}
```

## Functions
Most functions are used by the gocron front and backend services. Each function takes an argument
called `verbose`, which is a boolean. When true, all error messages and log messages are printed to
standard out.

### GetConfig()
Returns a config struct

### DatabaseString()
Returns a postgres connection string

### QueryDatabase()
Returns the result of a query and whether or not the query was successful

### StringToInt()
Converts a String to an Integer. Returns -1 if the conversion fails.

### CronLog()
Writes a message to syslog

### CheckError()
Sends a message of type `err` to CronLog
