
## API Endpoints

### Create User Account

**Sample Request**

- `POST /User`
  
- `Content-Type:` application/json
  
- `JSON body:` 


```json
{
"username": "kwamz",
"email":    "k@mail.com",
"age":      "4",
}
```

|     **Key**    | **Type** | **Remarks** |
|:--------------:|:--------:|:-----------:|
| `username` | `string` | required  |
| `email` | `string` | required  |
| `age` | `string` | required  |


**Sample Response**

```json
{
"user": {
    "username": "kwamz",
    "email":    "k@mail.com",
    "age":      "4",
},
"message": "user succesfully created"
}
```

### Update user

**Sample Request**

- `PUT /User?username=<username>"`

- `Content-type:` application/json

- `Sample request JSON body:`

```json
{
"email": "j@mail",
"age":   "44",
}
```

|     **Key**    | **Type** | **Remarks** |
|:--------------:|:--------:|:-----------:|
| `email` | `string` | not required  |
| `age` | `string` | not required  |


**Sample Response**

```json
{
"user": {
    "email":    "k@mail.com",
    "age":      "4",
},
"message": "user succesfully updated."
}
```

### Filter user

**Sample Request**

- `GET User/?username=kb`

- `Content-type:` application/json

- `JSON body:`

```json
{
"by": "username"
}

```

|     **Key**    | **Type** |
|:--------------:|:--------:|
| `by` | `string` | required  |

**Sample Response**

```json
{
"user": {
    "email":    "k@mail.com",
    "age":      "4",
    "username": "kb"
},
"message": "user succesfully filtered."
}
```

### Create Game

**Sample Request**

- `POST /Game`
  
- `Content-Type:` application/json
  
- `JSON body:` 


```json
{
"age_rating":  "12+",
"url":         "ksadad@mail.com",
"description": "best game ever",
"publisher":   "EA Sports",
"title":       "Assassins Creed",
}
```

|     **Key**    | **Type** | **Remarks** |
|:--------------:|:--------:|:-----------:|
| `age_rating` | `string` | required  |
| `url` | `string` | required  |
| `description` | `string` | required  |
| `publisher` | `string` | required  |
| `title` | `string` | required  |


**Sample Response**

```json
{
"game": {
    "age_rating":  "12+",
    "url":         "ksadad@mail.com",
    "description": "best game ever",
    "publisher":   "EA Sports",
    "title":       "Assassins Creed",
},
"message": "game succesfully created"
}
```

### Update game

**Sample Request**

- `PUT /Game?title=<title>"`

- `Content-type:` application/json

- `Sample request JSON body:`

```json
{
"publisher":   "EA Sports",
}
```

|     **Key**    | **Type** | **Remarks** |
|:--------------:|:--------:|:-----------:|
| `publisher` | `string` | not required  |


**Sample Response**

```json
{
"game": {
"publisher": "EA Sports",
},
"message": "game succesfully updated"
}
```

### Filter user

**Sample Request**

- `GET Game/?title=assasinscreed`

- `Content-type:` application/json

- `JSON body:`

```json
{
"by": "title"
}

```

|     **Key**    | **Type** |
|:--------------:|:--------:|
| `by` | `string` | required  |

**Sample Response**

```json
{
"game": {
"age_rating":  "12+",
"url":         "ksadad@mail.com",
"description": "best game ever",
"publisher":   "EA Sports",
"title":       "assasinscreed",
},
"message": "game succesfully filtered."
}
```



***

## Error HTTP Status Codes

 **Error Codes**                              | **Meaning**
----------------------------------------------|------------------------------------------------------------------------------
 **[400 or StatusBadRequest](#400)**          | Malformed parameters or other bad requests to the server.
 **[500 or StatusInternalServerError](#500)** | Processing errors, probably relating to backing services.


**Sample 500 Response**

```json
{
"message": "could not read game record from database"
},
```

**Sample 400 request**
```json
{
"message":"could not parse request. check usage"
}
```
Above error occurs mostly from bad requests in request body

```json
{
"message":"no record found for username given"
}
```

Above error occurs mostly when a record is not found during filters and updates.

