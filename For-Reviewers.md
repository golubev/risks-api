# For Reviewers

Hi there 👋

I was coding this project without any hands-on experience with Go. I'm currently using Python but I'm comfortable with switching tech stacks.

For this project I focused on satisfying all functional requirements while trying to structure the code.

I also decided to add Swagger UI (OpenAPI) for the next reasons:
- it provides a fast high level overview of the API;
- it gives a playground out of the box so there's no need to use either curl or Postman to try the API;
- OpenAPI is the best practice and a vital part of the developers' tools, so the effort to integrate it will never be vain.

There are still some issues that might be done better:
- Validating the `State` enumeration could be implemented via a custom validator.
- It would be great to add a pagination for the `GET /v1/risks` endpoint.
- The project is running in debug mode. To run in production it should be also configured using the `GIN_MODE` environment variable.
- Docker build performance is far from optimal and should be optimized by splitting the Dockerfile into separate layers: install dependencies first and then copy all the application code.

You might also want to have a look at my open-source project. It is a bit similar but more holistic and complete, and it is written in Python - https://github.com/golubev/fastapi-advanced-rest-template

Thank you for your time and consideration.
