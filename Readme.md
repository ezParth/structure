# 📦 Go-Structure

A simple Go package to create project structures for Windows and Linux/Mac.

---

## 🚀 Features

✅ Create a structured project directory for Go projects.  
✅ Works on Windows, Linux, and Mac.  
✅ No external dependencies – pure Go standard library.  
✅ Easy to integrate into existing Go projects.  

---

## 📥 Installation

To install the package, run:

```sh
go get github.com/your_username/go-structure@latest
```

---

## 🛠 Usage

### 1️⃣ Import the package

In your Go project, import the package:

```go
package main

import (
    "fmt"
    "github.com/your_username/go-structure/windowsCommand"
    "github.com/your_username/go-structure/bashCommand"
    "runtime"
)

func main() {
    fmt.Println("Creating project structure...")

    if runtime.GOOS == "windows" {
        windowsCommand.MakeFile()
    } else {
        bashCommand.MakeFile()
    }

    fmt.Println("Project setup complete!")
}
```

### 2️⃣ Run the program

```sh
go run main.go
```

This will create the required project structure based on the operating system.

---

## 📂 Project Structure Created

After execution, the following structure will be created:

```
your_project/
│── cmd/
│   ├── server/
│       ├── main.go
│── internal/
│   ├── models/
│   ├── services/
│   ├── database/
│── pkg/
│   ├── logger/
│── api/
│   ├── handlers/
│── config/
│── scripts/
│── tests/
│── go.mod
│── README.md
```

---

## 🔗 Contributing

1. Fork the repository  
2. Create a new branch (`git checkout -b feature-branch`)  
3. Commit your changes (`git commit -m "Added new feature"`)  
4. Push to the branch (`git push origin feature-branch`)  
5. Open a **Pull Request**  

---

## 📜 License

This project is open-source and available under the [MIT License](LICENSE).

---

## 💡 Support

⭐ Star the repo if you found it useful!  
🐛 Found a bug? Open an [issue](https://github.com/your_username/go-structure/issues).  

🚀 **Happy Coding!** 🎉

