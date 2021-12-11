export tag=v1.0.3

release:
	echo "building httpserver container"
	docker build -t vincent616/httpserver:${tag} .

push: release
	echo "pushing vincent616/httpserver"
	docker push vincent616/httpserver:${tag}

