# URL Shortner with Golang

#### This project utilizes the Gin framework in Go to build a robust URL shortening service. It employs REST APIs for seamless creation and retrieval of shortened URLs. Data storage is managed with SQLite, leveraging the GORM library for efficient ORM-based data handling. The combination of Gin's performance, RESTful design, and GORM's ORM capabilities ensures a scalable and maintainable solution for URL management.

### Setup
- Clone the project using `git clone https://github.com/Aadhieaswar/go-url-shortner.git`
- In the project directory, install necessary golang modules for the project using `go mod tidy`
- Start the url shortening service using `go run main.go`

### Services
#### URL Shortening Service
- Endpoint: POST `/shorten`
- Input: Form data
  - `longUrl` : string --> URL that needs to be shortened
- Returns: Generated short URL in `JSON` format

#### Short URL Redirect
- Endpoint: GET `/:shortCode`
- Input: __Parameter__ `shortCode` --> shortened code for the URL generated
- Returns: Redirects you to the shortened URL if found in the db

#### Discovery
- Endpoint: GET `/codes/discovery`
- Returns: List of all the shortened URLs and their respective short codes
