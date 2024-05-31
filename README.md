# URL Shortener in Go

## Overview

This project is a simple URL shortener service written in Go. It provides an API to shorten URLs and redirect shortened URLs to their original destinations.

## Features

- **Shorten URLs:** A POST endpoint to generate a shortened URL for a given long URL.
- **Redirect URLs:** A GET endpoint to redirect a shortened URL to its original URL.

## Project Structure

- `main.go`: The main file containing the HTTP server setup and the handler functions.
- `URL` struct: Defines the structure for storing original and shortened URLs.
- `urls` map: An in-memory store for mapping shortened URLs to original URLs.
- `urlMux` mutex: Ensures thread-safe access to the `urls` map.
- `letters` slice: Contains characters for generating random shortened URLs.

## Installation

1. Ensure you have Go installed on your machine. You can download it from [here](https://golang.org/dl/).

2. Clone this repository:

   ```bash
   git clone https://github.com/Niall1985/GolangURLshortener.git
   cd GolangURLshortener
   ```

3. Install the mux package:
    ```bash
   go get -u github.com/gorilla/mux
   ```

4. Run the application:

   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`.

## Usage

### Shorten a URL

You can shorten a URL by sending a POST request to `/shorten` with a JSON body containing the URL to be shortened.

#### Using PowerShell

```powershell
$longUrl = "YOUR_LONG_URL"
$body = @{ url = $longUrl } | ConvertTo-Json  
$response = Invoke-RestMethod -Uri http://localhost:8080/shorten -Method Post -Body $body -ContentType "application/json"
Write-Host "Original URL: " $response.original
Write-Host "Shortened URL: http://localhost:8080/$($response.short)"
```

#### Using Curl

```bash
curl -X POST http://localhost:8080/shorten -H "Content-Type: application/json" -d '{"url":"YOUR_LONG_URL"}'
```

### Redirect to Original URL

Access the shortened URL in your browser or use a GET request to be redirected to the original URL. For example, if the shortened URL is `http://localhost:8080/abc123`, navigating to this URL will redirect you to the original URL.

## API Endpoints

- **POST /shorten**

  Request Body:
  
  ```json
  {
    "url": "YOUR_LONG_URL"
  }
  ```

  Response:

  ```json
  {
    "original": "YOUR_LONG_URL",
    "short": "abc123"
  }
  ```

- **GET /{shortURL}**

  Redirects to the original URL associated with `shortURL`.

## Example

1. Shorten a URL:

   ```bash
   curl -X POST http://localhost:8080/shorten -H "Content-Type: application/json" -d '{"url":"https://www.example.com"}'
   ```

   Response:

   ```json
   {
     "original": "https://www.example.com",
     "short": "abc123"
   }
   ```

2. Redirect using the shortened URL:

   Open `http://localhost:8080/abc123` in your browser. You will be redirected to `https://www.example.com`.

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

---