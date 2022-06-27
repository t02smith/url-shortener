# URL Shortener

A url shortener written in Go with SQLite3 that provides a RESTful API to allow users to easily
interact with it.

Typically a shortened url will have an N character hash
appended to the hosting domain and this link will simply
redirect to the correct page.

> *e.g. **t02smith.com** may be allocated **link.t02smith.com/56735***

This project *should* currently being run on a **Raspberry Pi 3B** at [http://link.t02smith.com:8080](http://link.t02smith.com:8080)

![Go Badge](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=fff&style=for-the-badge)
![SQLite Badge](https://img.shields.io/badge/SQLite-003B57?logo=sqlite&logoColor=fff&style=for-the-badge)
![Docker Badge](https://img.shields.io/badge/Docker-2496ED?logo=docker&logoColor=fff&style=for-the-badge)
![Raspberry Pi Badge](https://img.shields.io/badge/Raspberry%20Pi-A22846?logo=raspberrypi&logoColor=fff&style=for-the-badge)
![HTML5 Badge](https://img.shields.io/badge/HTML5-E34F26?logo=html5&logoColor=fff&style=for-the-badge)
![CSS3 Badge](https://img.shields.io/badge/CSS3-1572B6?logo=css3&logoColor=fff&style=for-the-badge)

## How to Run

First clone the repository:

```bash
git clone https://github.com/t02smith/url-shortener.git
```

Then to run you can either:

1. Run locally:

    ```bash
    go run main.go
    ```

2. Run on a Docker container:

    ```bash
    docker-compose up --build
    ```

### Config

Constants within the application can be changed to suit
your needs. Currently included constants are:

- DATABASE_LOCATION = path to SQLite3 database
- DOMAIN = Domain the application is being hosted on
- HASH_SIZE = The length of the hash for the shortened URL
- API_PATH = The prefix to API call paths
- PORT = The port to listen on
       = NOTE: should match port in Dockerfile & docker-compose.yml

## API

### getURL

```javascript
path: /getURL
content: {
    url: string // the url you want to shorten
    
    // TODO
    request: string // request a mapping
}
```

## TODO

- shortened urls should have time limits *e.g. by day*
- request field in the get request to ask for a specific URL
- Webpages:
  - enter URL and receive shortened copy
  - error pages (redirect error)
- Unit testing
- Serve static folder
