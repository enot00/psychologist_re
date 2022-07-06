# Test Routes

## Ping

`http://localhost:8081/ping`

## In order to run the app execute the following:

1. `sudo docker run --rm --name pgsqldocker -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=userpsy -e POSTGRES_DB=psychologists -d -p 5432:5432 -v $HOME/docker/volumes/postgres:/var/lib/postgresql/data postgres`


## Endpoints


### 1.Psychologist:

`GetAll psychologists: GET http://localhost:8081/v1/psychologists`

`Create psychologists: POST http://localhost:8081/v1/psychologists`

`Update psychologists: PUT http://localhost:8081/v1/psychologists/update`

`Get one psychologists: GET http://localhost:8081/v1/psychologists/{psyId}`

`Delete psychologists: DELETE http://localhost:8081/v1/psychologists/delete/{psyId}`


### 2.Meeting:

`Create by psychologist: POST http://localhost:8081/v1/meeting/create_by_psychologist`

`Create by client: POST http://localhost:8081/v1/meeting/create_by_client`

`Paginate: GET http://localhost:8081/v1/meeting/page/{pagenum}`

`Paginate by psychologist: GET http://localhost:8081/v1/meeting/psychologist/{psychologist_id}/page/{pagenum}`

`Paginate by client: GET http://localhost:8081/v1/meeting/client/{client_id}/page/{pagenum}`

`Get one: GET http://localhost:8081/v1/meeting/{meeting_id}`

`Update by psychologist: PUT http://localhost:8081/v1/meeting/{meeting_id}/update_by_psychologist`

`Update by client: PUT http://localhost:8081/v1/meeting/{meeting_id}/update_by_client`

`Delete by psychologist: DELETE http://localhost:8081/v1/meeting/{meeting_id}/delete_by_psychologist`

`Delete by client: DELETE http://localhost:8081/v1/meeting/{meeting_id}/delete_by_client`

### POST request example(JSON)

### Authorization(Postman) for POST methods(post, put, delete)

* Type: Api Key
* Key: user_id
* Value: {user_id} - psychologist_id for psychologist, client_id for client
* Add to: Header


#### Create/Update by psychologist
#
    {
        "client_id": 44,
        "meeting_date": "2022-06-16T00:00:00Z",
        "start_time": 9.00,
        "end_time": 10,
        "status": "not completed" // default value
    } 
#

#### Create/Update by client
#
    {
        "psychologist_id": 1,
        "meeting_date": "2022-06-15T00:00:00Z",
        "start_time": 10.30,
        "end_time": 11.30,
        "status": "successfully completed"
    } 
#

### 3.Client:

`Create: POST http://localhost:8081/v1/client`

`Paginate: GET http://localhost:8081/v1/client/page/{pagenum}`

`Get one: GET http://localhost:8081/v1/client/{id}`

`Update: PUT http://localhost:8081/v1/client/{id}`

`Delete: DELETE http://localhost:8081/v1/client/{id}`

### POST request example(JSON)

#### POST
#

    {
        "user_name": "test_user",
        "phone_number": "+380672156691",
        "email": "test_user@mail.net",
        "avatar": "http://www.storage.net/ava.jpg",
        "password": "My_password123"
    }
#

#### PUT
#
    {
        "user_name": "test_user",
        "phone_number": "+380672156691",
        "email": "test_user@mail.net",
        "avatar": "http://www.storage.net/ava.jpg",
        "old_password": "My_password123",
        "password": "My_new_password123"
    }
#    

