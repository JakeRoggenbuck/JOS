# JOS
JSON Object Store

[![Build](https://img.shields.io/github/actions/workflow/status/JakeRoggenbuck/JOS/build.yml?branch=main&style=for-the-badge)](https://github.com/JakeRoggenbuck/JOS/actions)
[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://github.com/JakeRoggenbuck?tab=repositories&q=&type=&language=go&sort=stargazers)


## Usage
```
export ADMIN_PASSWORD=<password>
```

```
curl -u "Admin:$ADMIN_PASSWORD" -X GET http://localhost:8080/api/v1/
```

```
curl -u "Admin:$ADMIN_PASSWORD" -X POST http://localhost:8080/api/v1/upload-file \
     -F "myFile=@example.txt"
```
