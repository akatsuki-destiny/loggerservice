loggerservice:
	docker build -t loggerservice .
	docker run -p 3003:3003 loggerservice