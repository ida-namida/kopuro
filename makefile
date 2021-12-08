# TODO: Handle if Heroku CLI is not installed
# TODO: Check if `heroku whoami` already returns an identity
heroku-login:
	heroku login -i
	heroku container:login

docker-build-tag-push:
	docker build --build-arg BASE_FILE_PATH=$(BASE_FILE_PATH) -f deploy/docker/dockerfile -t kopuro .
	docker tag kopuro:latest registry.heroku.com/$(HEROKU_APP_NAME)/web
	docker push registry.heroku.com/$(HEROKU_APP_NAME)/web

heroku-release:
	heroku container:release -a $(HEROKU_APP_NAME) web

heroku-docker-all: heroku-login docker-build-tag-push heroku-release