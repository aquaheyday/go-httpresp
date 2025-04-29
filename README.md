# goresp

`goresp` is a lightweight Go package that simplifies JSON responses in HTTP APIs. It provides standardized response formats for success, errors, validation failures, pagination, and custom structures, enabling consistent and maintainable API development.

## 📦 Installation

To install `goresp`, use the following command:

```bash
go get github.com/aquaheyday/goresp
```

## 🚀 Usage

```go
package main

import (
    "net/http"
    "github.com/aquaheyday/goresp"
)

func handler(w http.ResponseWriter, r *http.Request) {
    data := map[string]string{"message": "Hello, World!"}
    goresp.Success(w, data)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

## 🎯 Features

- **Success**: Send a standard success response.

  ```go
  goresp.Success(w, data)
  ```

- **Error**: Send a standard error response with a message.

  ```go
  goresp.Error(w, "An error occurred")
  ```

- **ValidationError**: Send a response indicating validation errors.

  ```go
  errors := map[string]string{
      "email": "Invalid email address",
  }
  goresp.ValidationError(w, errors)
  ```

- **Paginated**: Send a paginated response with data and pagination info.

  ```go
  pagination := goresp.Pagination{
      Page:    1,
      PerPage: 10,
      Total:   100,
  }
  goresp.Paginated(w, data, pagination)
  ```

- **Custom**: Send a response with a custom status code and payload.

  ```go
  payload := map[string]interface{}{
      "custom": "value",
  }
  goresp.Custom(w, http.StatusTeapot, payload)
  ```

- **StatusSuccess / StatusError**: Send responses with `status`, `message`, and optional `data`.

  ```go
  goresp.StatusSuccess(w, "Operation successful", data)
  goresp.StatusError(w, "Operation failed")
  ```

## 📘 Documentation

For detailed documentation and examples, visit the [GoDoc](https://pkg.go.dev/github.com/aquaheyday/goresp) page.

## 🛡 License

This project is licensed under the [MIT License](LICENSE).

## 🙋‍♂️ Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.
