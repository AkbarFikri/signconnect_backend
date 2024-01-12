# SignConnect Backend Apps

## 📒 Index

- [SignConnect Backend Apps](#signconnect-backend-apps)
  - [📒 Index](#-index)
  - [🔧 Development](#-development)
    - [📓 Tech Stack](#-tech-stack)
    - [📷 API \& Database Specification](#-api--database-specification)
    - [📁 File Structure](#-file-structure)
  - [🌟 Credit](#-credit)
  - [🔒License](#license)

## 🔧 Development

Here is a description of our apps development

### 📓 Tech Stack

List all the Tech Stack we use to build the system in this this project.

| No  | Tech                  | Details                                                           |
| --- | --------------------- | ----------------------------------------------------------------- |
| 1   | Go                    | To build a fast and efficient Backend App                         |
| 2   | Google Cloud Platform | To provide all application needs related to server infrastructure |

### 📷 API & Database Specification

You can just look at the `Docs` folder to check Our `API-Specification` and `Database-Specificasion`
Here is the link

1. [API-Specification](https://github.com/AkbarFikri/signconnect_backend/blob/main/Docs/API-Specification.md)
2. [API-Specification](https://github.com/AkbarFikri/signconnect_backend/blob/main/Docs/Database-Specification.md)

### 📁 File Structure

Here is our file Structure

```
├───Config/
│   ├───config.go
│   └───db.go
├───Controller/
│   ├───auth_Controller.go
│   ├───dictionary_Controller.go
│   ├───games_Controller.go
│   ├───home_Controller.go
│   ├───leaderboard_Controller.go
│   ├───lembaga_Controller.go
│   └───user_Controller.go
├───Database/
│   ├───database.go
│   └───migration.go
├───Docs/
│   ├───API-Specification.md
│   └───Database-Specification.md
├───Models/
│   ├───dictionary_Models.go
│   ├───leaderboard_Models.go
│   ├───lembaga_Models.go
│   ├───leveling_Models.go
│   ├───role_Models.go
│   ├───soal_Models.go
│   └───user_Models.go
├───Public/
├───routers/
│   ├───Middleware/
│   ├───index.go
│   └───router.go
├───.env
├───.env.example
├───.gitignore
├───DockerFile
├───go.mod
├───go.sum
├───main.go
└───README.md
```

## 🌟 Credit

The Member of our team

1. Akbar Fikri Abdillah
2. Firza Aurellia Iskandar
3. Muhammad Rafly Ash Shiddiqi
4. Arif Athaya Harahap

## 🔒License

© Inti Bumi - Hackfest by Google Indonesia 2024
