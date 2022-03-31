# Finance Collector (Go)

## Overview

### Purpose
Provide a way to simply keep track of stock data via API. This may be used by a software or my DIY stock ticker built from an ESP8266 and LED matrix display. It may be extended in the future to provide text and/or email messages on certain criteria or at day close. The information API that I am pulling from was written by me in python/flask and returns a ticker information in the format below - "ticker json format".
<br>
<br>

### Features
- [ ] Individual users 
- [ ] Add stock tickers and amounts to a database. 
- [ ] Return current ticker data plus current worth as JSON
- [ ] Keep a cached version (10 minutes) of each ticker
- [ ] If ticker data is newer than 10 minutes, do not update via python scraper
- [ ] Simple GUI to update database for user
- [ ] Adding a ticker with 0 amount is allowed (for showing on DIY stock ticker)
<br>
<br>

## Technicals

### Total json format

```
{
    "worth":            number,
    "dayPercentChange": string,
    "dayDollarChange":  number,
    "tickers" :         []
}
```

### ticker json format
```
{
    "dayChange": string,
    "name":      string,
    "open":      number,
    "prevClose": number,
    "price":     number,
    "symbol":    string,
    "time":      number,
    "todayHigh": number,
    "todayLow":  number
}
```
### Env
rename `env.example` to `.env` and change the appropriate information.

example 
```
TOKEN=yourToken
LISTENPORT=8080
REMOTEURL=http://1.2.3.4
REMOTEPORT=5000
```

<br> 

### DB schema
* coming soon
#
## Milestone Overview
- [ ] Sprint 1 - make a verison that works with the old nodejs version. NOTE: the format of json is not as above to make old DIY stock ticker work
- [ ] Sprint 2 - create a new branch to start work on the above json formats, basic IP throttling?
    - [ ] cleanup
- [ ] Sprint 3 - setup mysql database, provide basic read functionality
    - [ ] cleanup
- [ ] Sprint 4 - full functionality based on users from database
    - [ ] cleanup

More coming in the future