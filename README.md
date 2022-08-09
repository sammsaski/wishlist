# wishlist-golang
A wishlist app built with a Go backend, React frontend, PostgreSQL and Docker.

##  How to Deploy
Deployment is split into two stages, both of which will have their steps documented here. The two stages are:
1. PostgreSQL + Go Backend API
2. React Frontend
We'll begin by detailing how to get the PostgreSQL + Go Backend API deployed and running.
### PostgreSQL + Go Backend API
The deployment of both PostgreSQL and our Go backend API have been joined together in a single `docker-compose.yml` in `backend/`. In order to get them up and running, we recommend the following steps:
1. cd backend/
2. docker-compose build --no-cache
3. docker-compose up -d
Now, we can check that the two containers deployed properly by `exec`ing into each container. In the DB container, we can check that Postgres is running by executing:
```
psql -U <postgres username>
```
In the second container, we recommend executing some API calls either with `curl` or Postman to confirm that the backend is active.
