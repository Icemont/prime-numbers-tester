## Prime Numbers Tester

Simple implementation of the Prime Numbers Tester

## Author
 [Ray Icemont](https://github.com/Icemont/)


## Configuration
To customize some application settings, you can copy the `.env.example` file to `.env`, make changes to it, and place it in the same folder as the compiled executable.


### Example query with cURL to check for prime numbers

	$ curl -X POST http://localhost:8888/ -F 'numbers=2' -F 'numbers=5'
