# deepfind-core
The goal of this project is to provide a user-friendly way to identify deep fakes to help slow the spread of disinformation.

## Design goals
We will rely heavily on machine learning to identify the deepfakes. Since ML is far from instantaneous, our UX must display a queue of the video requests and their progress. The design of the backend queue is TBD.

## Proposed Backend Architectures
### K8s
[microservices arch](./k8s_proposal.png "microservices arch")
#### Benefits
* Easily scalable
* Separation of concerns

### K8s (Authless)
[authless microservices arch](./k8s_proposal(authless).png "authless microservices arch")
#### Benefits
* Easily scalable
* Separation of concerns
* Less complicated than with auth

### Pseudo "Monolith"
[monolith arch](./monolith_proposal.png "monolith arch")
#### Benefits
* Significantly simpler

## Tech stack
TBD after deciding on a backend architecture.

## Set up
TBD after deciding on the tech stack.

## Wrangling the data
The compressed training set is ~100GB. We need to find a way to work with the data in the data service as painlessly as possible. Free S3 instances are limited at 5GB so we probably can't use that.

