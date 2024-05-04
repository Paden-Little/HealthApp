up-build:
	sudo docker compose up --build --remove-orphans -d
up:
	sudo docker compose up -d
down:
	sudo docker compose down