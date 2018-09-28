.PHONY: goApp

goApp:
	docker build -t goapp .
	sleep 5
	docker run -d --name goapp goapp

clean:
	docker stop goapp
	docker rm goapp

removeAll:
	docker system prune -a


