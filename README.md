# SCRUMS Ticket system

SCRUMS is a simple system that allows to create tickets, and assign them to a certain users with a simple notification.

## Preflight Check

- Docker
- Docker Compose
- make

Also for the dev you'll need:

- Go
- Node


## Running locally

Run `$ make build` script to install all the deps and run dockerized app.
The app will be available here: `localhost:8080`

## Development

### Frontend

Frontend lives in `/frontend` folder. It is available on app root, and serves from `/frontend/dist` folder.

### Backend

App is a set of microservices written in Go.
The backend lives in the `/backend` folder. Each subfolder is a separate go module, initialized with go mod.
APIs are available on `localhost:8080/api/v1/<api_name>`

#### Microservices

Subfolder   | API Location      | Description
------------|-------------------|------------
`!template` | N/A               | Microservice's template 
`users`     | `/api/v1/users/`  | Microservice that manages users
`auth`      | `/api/v1/auth/`   | Auth microservice. Please look at the [Auth Documentation](backend/auth/README.md) for the details.

#### Adding a new microservice

1. Copy `/backend/!template` folder, changing !template to your service's name
1. Implement the logic
1. Make changes in Dockerfile
1. Add a new record in .env file with a port that is different from others
1. Copypaste templates, making changes in the following files: 
    1. `docker-compose.yaml` - to add your new service into the virtual network
    1. `nginx/nginx.conf` - to add your API proxy pass, so your service will be visible

### Testing

We have a simple library called `testutils` that helps to maintain the most of the reusable testing methods.

Please refer to the [Testutils Documentation](backend/testing/README.md) for testing conventions, code organizations, code snippets, etc.
