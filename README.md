# go-resp

`go-resp` is a lightweight Go package that simplifies JSON responses in HTTP APIs. It provides standardized response formats for success, errors, validation failures, pagination, and custom structures, enabling consistent and maintainable API development.

## 📦 Installation

To install `go-resp`, use the following command:

```bash
go get github.com/aquaheyday/go-resp
```

## 🚀 Usage

```go
package main

import (
    "net/http"
    "github.com/aquaheyday/go-resp"
)

func handler(w http.ResponseWriter, r *http.Request) {
    data := map[string]string{"message": "Hello, World!"}
    resp.Success(w, data)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

## 🎯 Features

- **Success**: Send a standard success response.

  ```go
  resp.Success(w, data)
  ```

- **Error**: Send a standard error response with a message.

  ```go
  resp.Error(w, "An error occurred")
  ```

- **ValidationError**: Send a response indicating validation errors.

  ```go
  errors := map[string]string{
      "email": "Invalid email address",
  }
  resp.ValidationError(w, errors)
  ```

- **Paginated**: Send a paginated response with data and pagination info.

  ```go
  pagination := resp.Pagination{
      Page:    1,
      PerPage: 10,
      Total:   100,
  }
  resp.Paginated(w, data, pagination)
  ```

- **Custom**: Send a response with a custom status code and payload.

  ```go
  payload := map[string]interface{}{
      "custom": "value",
  }
  resp.Custom(w, http.StatusTeapot, payload)
  ```

- **StatusSuccess / StatusError**: Send responses with `status`, `message`, and optional `data`.

  ```go
  resp.StatusSuccess(w, "Operation successful", data)
  resp.StatusError(w, "Operation failed")
  ```

## 📘 Documentation

For detailed documentation and examples, visit the [GoDoc](https://pkg.go.dev/github.com/aquaheyday/go-resp) page.

## 🛡 License

This project is licensed under the [MIT License](LICENSE).

## 🙋‍♂️ Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.
