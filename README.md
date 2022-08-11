# wishlist-golang
A wishlist app built with a Go backend, React frontend, PostgreSQL and Docker.

## About
This project was built with the intent of having a fully functional + completed web application built with a Go backend and React frontend. In the end, I just wanted to be able to practice using some new skills and this project reflects that. The web app only has the simple functionality of:
1. Listing all items in the wishlist on the page.
2. Adding items to the wishlist.
3. Deleting items from the wishlist.
4. #TODO: Updating items from the wishlist.
By implementing these simple operations, I was able to practice developing a RESTful API using Go.

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
```bash
psql -U <postgres username>
```

### React Frontend
To get the React frontend started, navigate to the `frontend/` directory and execute the following command:
```bash
npm start
```
Now, you should see the web app open up in your default browser!

