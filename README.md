# api_for_random_nerds

A warm welcome to you from my little project and me.
This project is an api with a very specific task, providing random numbers and calculating the standard deviation from them
You need to enter a API_KEY_RANDOM, which is the key to the api from the RANDOM.ORG website, you need to create an account for this.
I prepared the ENV file to make it easier for you, you also need to add the same key to docker-compose.yaml and dockerfile if you want to run the Docker container.

Once the stack is up, you must send a get request:
```
localhost:8080/random/mean?requests=20&length=4
```
to Get functional response:
```
[
	{
		"stddev": 1.4142135623730951,
		"data": [
			2,
			4,
			3,
			5,
			6
		]
	},
	{
		"stddev": 1.4142135623730951,
		"data": [
			2,
			4,
			3,
			5,
			6
		]
	}
]
```
