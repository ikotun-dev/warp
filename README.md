## Warp

> Also not recommended for production _yet_

This project was inspired by chimney [https://github.com/aosasona/chimney](https://github.com/aosasona/chimney).

It's also an attempt to deeply understand the inner workings of http/mimetypes/request lifecycle and more.

##### Deployed with warp

[larchive](https://larchive-fe.fly.dev) on [https://fly.io](https://fly.io)

### Why warp?

- Rewrites work by default. No special configuration is required.
- As tiny as possible as well
- Logs all incoming requests _( can be disabled )_
- Accurate Mimetypes for all files

### Installation

##### Using Docker

###### Runing the docker image

```bash
docker pull ikotun/warp:latest
docker run -p PORT:8080 ikotun/warp:latest
```

This should have warp running on the specified port.
You can confirm by checking your browser.

![warp](https://res.cloudinary.com/dbd7rcwwx/image/upload/v1714760406/Screenshot_2024-05-03_at_7.17.21_PM_jlks9r.png)

> Create a warp.yaml file in the root directory of your project.

> This is an example

```bash
port : PORT
root : "index.html"
fallbackDocument : "404.html"
routes :
  - "/login"
  - "/signup"

```

##### To use warp in your project

```bash

#Pull warp
FROM ikotun/warp:latest

#copy your built project
COPY ./dist /public

#copy yout warp config file
COPY warp.yaml ./warp.yaml

#Expose 80
EXPOSE 8080
```
