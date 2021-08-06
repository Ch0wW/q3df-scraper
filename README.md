# q3df-scraper

`Q3DF-scraper` is a small tool written in Golang that retrives and downloads the latest files from https://ws.q3df.org. It was heavily inspired by the scraper script used by Github user Q3defrag (https://github.com/q3defrag/defrag-server).

# Usage
```q3df-scraper [-max-downloads] [-output-directory /path/to/folder] [-refresh 3600]``` 

Parameters:
```txt
-max-downloads <number> : allows for a number of n downloads before the program exits.
-output-directory </path/to/folder> : downloads all files to the specific directory. If folder doesn't exist, it will try to create it.
-refresh <seconds> : Sets the required time between two checks (in seconds).
```

# Pre-Requisites for compilation
- Golang 1.11 or newer (previous versions weren't tested)
- Package `gofeed` from user mmcdole (`go get github.com/mmcdole/gofeed`)

# Licence
This program is licenced under GPLv3.