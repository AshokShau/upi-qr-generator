# UPI QR Code Generator API

A lightweight API for generating UPI QR codes and rendering a secure payment page. This project is built using Go and
the Gorilla Mux router.


<p align="center">
  <img src="https://raw.githubusercontent.com/AshokShau/upi-qr-generator/master/.github/pay.png" 
       alt="Pay Example" 
       width="300" 
       height="auto">
</p>
<p align="center"><em>Example of payment page</em></p>

## Features

- Generate QR codes for UPI payments.
- Serve a customizable payment page with UPI details.
- Easy-to-use RESTful endpoints.
- Static assets support for additional resources.

## Requirements

- Go - https://golang.org
- Modules:
    - `github.com/skip2/go-qrcode`

## Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/AshokShau/upi-qr-generator.git
   cd upi-qr-generator
   ```

2. Install dependencies:
    ```bash
    go mod download
    ```

3. Run the application:
   ```bash
   go run main.go
   ```

4. The server will start at `http://localhost:3000`.

## Project Structure

```
upi-qr-generator/
├── api/            // API endpoint handlers for vercel
│   ├── index.go    // Handles /
│   ├── pay.go      // Handles /pay
│   ├── qr.go       // Handles /qr
├── str/            // String constants and HTML templates
│   ├── files.go    // HTML templates for payment page
│   ├── handlers.go // Handlers for API endpoints
├── main.go         // Main entry point of application
├── README.md       // Project README file
└── LICENSE         // Project license file
```

## API Endpoints

### 1. `GET /`

Redirects to the homepage (`static/index.html`) that provides usage information.

### 2. `GET /qr`

Generates a UPI QR code based on the provided query parameters:

- **Parameters**:
    - `vpa` (string, required): UPI ID (e.g., `example@upi`).
    - `name` (string, optional): Payee name (e.g., `John Doe`).
    - `amount` (float, required): Payment amount (e.g., `100.50`).

- **Example**:
  ```
  /qr?vpa=example@upi&name=John+Doe&amount=100.50
  ```

- **Response**:
  A PNG image of the QR code.

### 3. `GET /pay`

Renders a payment page with UPI details:

- **Parameters**:
    - `vpa` (string, required): UPI ID (e.g., `example@upi`).
    - `name` (string, optional): Payee name (e.g., `John Doe`).
    - `amount` (float, required): Payment amount (e.g., `100.50`).

- **Example**:
  ```
  /pay?vpa=example@upi&name=John+Doe&amount=100.50
  ```

- **Response**:
  A rendered HTML page with the QR code and payment details.

## How It Works

1. The `/qr` endpoint generates a UPI QR code using the `qrcode` library.
2. The `/pay` endpoint serves a payment page with details dynamically rendered from the `template.html` file.
3. Static assets such as styles and templates are served from the `static` folder.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Author

[**Ashok Shau**](https://github.com/AshokShau)

Feel free to explore, use, or contribute to this project. If you encounter any issues or have suggestions, open a GitHub
issue or create a pull request!
