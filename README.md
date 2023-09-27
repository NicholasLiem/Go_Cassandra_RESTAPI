# Building simple demo REST API using Go and Cassandra

## **How to Run The Program**
1. Clone this repository
```sh
https://github.com/NicholasLiem/Go_Cassandra_RESTAPI.git
```
2. Build the containers
```sh
docker-compose up --build
```
4. Make a new .env file based on .env.example (you can just remove .example from the file's name)
5. If for some reason the database is not seeded, you can manually seed using
```sh
docker exec -it single_service_app bash
```
```sh
yarn migration:run
```

## **Endpoints**
| Endpoint             | Method   | Description                                        |
|----------------------|----------|----------------------------------------------------|
| /login               | POST     | Login verification                                 |
| /self                | GET      | Get session status                                 |
| /barang              | GET      | Get a list of barang registered based on query     |
| /barang/:id          | GET      | Get the detail of barang of the given id           |
| /barang              | POST     | Create a new barang                                |
| /barang/:id          | PUT      | Update the detail of a barang of the given id      |
| /barang/:id          | DELETE   | Delete barang of the given id                      |
| /perusahaan          | GET      | Get a list of perusahaan registered based on query |
| /perusahaan/:id      | GET      | Get the detail of perusahaan of the given id       |
| /perusahaan          | POST     | Create a perusahaan                                |
| /perusahaan/:id      | PUT      | Update the detail of perusahaan of the given id    |
| /perusahaan/:id      | DELETE   | Delete perusahaan of the given id                  |


## **Tech Stack**
Go-Gorilla, Cassandra

## **Extras**
- Amazon EC2 service is used for this backend service and am using Supabase service to deploy my PostgresSQL server. You can interact with the API through this public IPv4 address: N/A (use http instead of https)
