# deepfind-core
The goal of this project is to provide a user-friendly way to identify deep fakes to help slow the spread of disinformation.

## Design goals
We will rely heavily on machine learning to identify the deepfakes. Since ML is far from instantaneous, our UX must display a queue of the video requests and their progress. The design of the backend queue is TBD.

## Tech stack
golang server side
postgresql
mongodb
docker

## Set up
1. `docker-compose up -d`
1. `./bin/migrate.sh` to migrage atlas to latest version.


## API
```
curl -X 'GET' \
  'http://0.0.0.0:8080/check?youtubeUrl=<some_youtube_url>'
```
