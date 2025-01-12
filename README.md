# JOS
JSON Object Store

[![Build](https://img.shields.io/github/actions/workflow/status/JakeRoggenbuck/JOS/build.yml?branch=main&style=for-the-badge)](https://github.com/JakeRoggenbuck/JOS/actions)
[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://github.com/JakeRoggenbuck?tab=repositories&q=&type=&language=go&sort=stargazers)


## Usage
```
export ADMIN_PASSWORD=<password>
```

### Upload a file
```
curl -u "Admin:$ADMIN_PASSWORD" -F "myFile=@path_to_your_file" http://localhost:8080/api/v1/upload-file
```

```json
{
  "hash": "<hash>",
  "message": "File uploaded and saved successfully!"
}
```

### Get a file
```
curl -u "Admin:$ADMIN_PASSWORD" http://localhost:8080/api/v1/get-file?hash=<hash>
```

### Upload JSON
```
curl -u "Admin:$ADMIN_PASSWORD" -X POST -H "Content-Type: application/json" -d '{"key": "value"}' \
     http://localhost:8080/api/v1/upload-json
```

```json
{
  "hash": "<hash>",
  "message": "JSON uploaded and saved successfully!"
}
```

### Get JSON
```
curl -u "Admin:$ADMIN_PASSWORD" http://localhost:8080/api/v1/get-json?hash=<hash>
```
