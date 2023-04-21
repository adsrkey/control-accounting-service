# Control Accounting Service

### How to run service

```
cd .\deploy\
```

```
make up_build
```

```
make migrate_up
```

### How to stop service

```
cd .\deploy\
```

```
make down
```

### How to down migration

```
   make migrate_down
```

---
### Endpoints:
<br>

#### POST:
Description:

- Create new operator

```
http://localhost:8080/api/v1/operators
```
#### Request Headers:
```Content-Type application/json```

Body raw (json)
```
{
"first_name":"Daniel",
"last_name":"Holopov",
"middle_name":"Poncho",
"city":"Moscow",
"phone_number":"+7-(906)-906-00-00",
"email":"email@temp.com"
}
```
Response:
```
{
    "id": "14d75941-3c9d-43ff-8db4-2f8e6932f16e",
    "generated_password": "G0_u3G.,"
}
```
---
#### GET
Description:

- Get operator by uuid

```
http://localhost:8080/api/v1/operators/d96ab496-58ea-40f8-ab33-41383a809dc0
```
Response:
```
{
    "id": "d96ab496-58ea-40f8-ab33-41383a809dc0",
    "created_at": "2023-04-21T12:03:27.251738Z",
    "modified_at": "2023-04-21T12:03:27.251738Z",
    "first_name": "Arnold",
    "last_name": "Schwarzenegger",
    "middle_name": "Alois",
    "city": "Thal",
    "phone_number": "19049009390",
    "email": "emailf@email.com"
}
```
---

#### GET
Description:

- Get all operators with offset and limit

```
http://localhost:8080/api/v1/operators
``` 
## Query Params
- offset 
- limit

```
http://localhost:8080/api/v1/operators?offset=2&limit=2
```
Response:
```
{
    "count": 2,
    "limit": 2,
    "offset": 2,
    "operators": [
        {
            "id": "4cfad44f-2a4d-4462-912e-113d8ce428fd",
            "created_at": "2023-04-21T12:03:27.251738Z",
            "modified_at": "2023-04-21T12:03:27.251738Z",
            "first_name": "Michael",
            "last_name": "Vosnetsov",
            "middle_name": "Romanovich",
            "city": "Syktyvkar",
            "phone_number": "79019019090",
            "email": "emailb@email.com"
        },
        {
            "id": "319e4715-0c9e-4c99-9fb4-a79ba58bf133",
            "created_at": "2023-04-21T12:03:27.251738Z",
            "modified_at": "2023-04-21T12:03:27.251738Z",
            "first_name": "Sergey",
            "last_name": "Khlebov",
            "middle_name": "Fedorovich",
            "city": "Moscow",
            "phone_number": "89029029090",
            "email": "emailc@email.com"
        }
    ]
}  
```
---

#### PUT

Description:

- Update operator by uuid

```
http://localhost:8080/api/v1/operators
```
#### Request Headers:
```Content-Type application/json```

Body raw (json)
```
{
    "id": "ea4a3096-bb06-4a43-9fbc-6bdeb95e182a",
    "first_name": "Pavel", 
    "city": "St. Petersburg", 
    "phone_number": "79349009090",
    "email": "emaile1@gmail.com"
}
```

---
#### DELETE

Description:

- Delete operator by uuid

```
http://localhost:8080/api/v1/operators/ea4a3096-bb06-4a43-9fbc-6bdeb95e182a
```
 
```
"deleted"
```

---
#### POST

Description:

- Create new project

```
http://localhost:8080/api/v1/projects
```

#### Request Headers:
```Content-Type application/json```

Body raw (json)
```
{
"project_name": "project",
"project_type":1
}
```

Response:
```
"4e9db8cf-045e-45db-9a88-553e948afc4c"
```

---
#### GET

Description:

- Get project by uuid (with project operators)

```
http://localhost:8080/api/v1/projects/e0a1c218-1aa4-4261-bb03-dacd1d106c40
```

Response:
```
{
    "id": "e0a1c218-1aa4-4261-bb03-dacd1d106c40",
    "created_at": "2023-04-21T12:03:27.251738Z",
    "modified_at": "2023-04-21T12:03:27.251738Z",
    "project_name": "R2D2",
    "project_type": "auto",
    "operators": [
        {
            "id": "e6745a46-6eb6-4ca4-a769-549259c26e56",
            "created_at": "2023-04-21T12:03:27.251738Z",
            "modified_at": "2023-04-21T12:03:27.251738Z",
            "first_name": "Mark",
            "last_name": "Hamill",
            "middle_name": "Richard",
            "city": "Star Wars",
            "phone_number": "11023035591",
            "email": "emailk@email.com"
        }
    ]
}
```

---
#### GET

Description:

- Get all projects with offset and limit

```
http://localhost:8080/api/v1/projects
``` 
## Query Params
- offset
- limit

```
http://localhost:8080/api/v1/projects?limit=1&offset=1
```
Response:
```
{
    "count": 1,
    "limit": 1,
    "offset": 1,
    "project": [
        {
            "id": "aac75330-8a90-40e2-9440-066cca9a1f8a",
            "created_at": "2023-04-21T12:03:27.251738Z",
            "modified_at": "2023-04-21T12:03:27.251738Z",
            "project_name": "T1000",
            "project_type": "out",
            "operators": [
                {
                    "id": "d96ab496-58ea-40f8-ab33-41383a809dc0",
                    "created_at": "2023-04-21T12:03:27.251738Z",
                    "modified_at": "2023-04-21T12:03:27.251738Z",
                    "first_name": "Arnold",
                    "last_name": "Schwarzenegger",
                    "middle_name": "Alois",
                    "city": "Thal",
                    "phone_number": "19049009390",
                    "email": "emailf@email.com"
                }
            ]
        }
    ]
}
```
---
#### PUT

Description:

- Update project by uuid

```http://localhost:8080/api/v1/projects```

#### Request Headers:
```Content-Type application/json```

Body raw (json)
```
{
    "id":"e0a1c218-1aa4-4261-bb03-dacd1d106c40",
    "project_name":"Pool",
    "project_type":"2"
}
```
Response:
```
"ok"
```

---
#### DELETE

Description:

- Delete project by uuid

```
http://localhost:8080/api/v1/projects/e0a1c218-1aa4-4261-bb03-dacd1d106c40
```

Response:
```
"deleted"
```

---
#### DELETE

Description:

- Get projects operators by uuids

```
http://localhost:8080/api/v1/projects/operator
```

#### Request Headers:
```Content-Type application/json```

Body raw (json)

```
{
    "id":"29525e05-6507-4820-80c7-b0a56bc7f5e2",
    "operator_ids":[
        "167ed880-3a5c-4949-bad8-ad2634ab7ab9",
        "bef0c626-1af5-4f78-8821-a4877c73a2b0",
        "319e4715-0c9e-4c99-9fb4-a79ba58bf133"
    ]
}
```
Response:
```
"deleted"
```

---
#### POST

Description:

- Assign operators to project by project uuid

```
http://localhost:8080/api/v1/projects/operator
```

#### Request Headers:
```Content-Type application/json```

Body raw (json)

```
{
    "id":"29525e05-6507-4820-80c7-b0a56bc7f5e2",
    "operator_ids":[
        "167ed880-3a5c-4949-bad8-ad2634ab7ab9",
        "bef0c626-1af5-4f78-8821-a4877c73a2b0",
        "319e4715-0c9e-4c99-9fb4-a79ba58bf133"
    ]
}
```
Response:
```
"assigned"
```
