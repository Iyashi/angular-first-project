# Angular First Project
This is my first angular project.

## Getting Started

Clone this repository and navigate to the directory:
```bash
git clone https://github.com/Iyashi/angular-first-project.git
cd angular-first-project
```

Copy `example.env` to `.env` and configure it for your needs:
```bash
cp example.env .env
```

Install `node_modules` for development environment:
```bash
cd frontend
npm i
cd ..
```

Run the stack
```bash
docker compose up --build
```

## !!!ATTENTION!!!
Currently, I have only implemented the deployment using `compose.proxy.yml`.

Right now, the `UserService` has the API Url hardcoded as `api.angular-first-project.local`.
Because of this, using direct port mappings (`compose.ports.yml`) is not possible.

You can either change the API Url in `UserService` to `http://localhost:8081` or use the `compose.proxy.yml` way.

If you changed the API Url, run the stack and access the application at [http://localhost:8080](http://localhost:8080).

In order to use `compose.proxy.yml`, you have to do the following:

- Add the following domains to `/etc/hosts`:
  ```hosts
  127.0.0.1  angular-first-project.local api.angular-first-project.local traefik.angular-first-project.local
  ```
- Set `COMPOSE_FILE=compose.yml:compose.proxy.yml` in `.env`

Then run the stack and access the application at [http://angular-first-project.local](http://angular-first-project.local).
