up-build:
	sudo docker compose up --build --remove-orphans -d
dcu:
	sudo docker compose up --build --remove-orphans
up:
	sudo docker compose up -d
down:
	sudo docker compose down