# Backend Developer Tests

These are my answers / implementations of the backend development tests over at [StackPath](https://github.com/stackpath/backend-developer-tests)

## Unit Testing

I updated the implementation at `unit-testing/pkg/fizzbuzz.go` to account for negative inputs
passed into each of the input arguments. In this case negatives are still treated as valid steps
to print `Fizz` or `Buzz` since the remainder would still be `0`.

I have added tests in the `unit-testing/pkg/fizzbuzz_test.go` file to have coverage over these
inputs and many more. Running `go test -v -cover ./...` in the `unit-testing` directory will produce
successful results.

Optional improvements could be to have more of a table based testing structure similar to a
parameterized test in Java.

## Web Services

The code in the `rest-service` directory creates a web service with the following responses:

- Respond with JSON output
- Respond to `GET /people` with a 200 OK response containing all people in the 
  system
- Respond to `GET /people/:id` with a 200 OK response containing the requested 
  person or a 404 Not Found response if the `:id` doesn't exist
- Respond to `GET /people?first_name=:first_name&last_name=:last_name` with a 
  200 OK response containing the people with that first and last name or an 
  empty array if no people were found
- Respond to `GET /people?phone_number=:phone_number` with a 200 OK response 
  containing the people with that phone number or an empty array if no people 
  were found

The REST service has been implemeted with the go std lib.

I have included a `Makefile` so I can run `make test` to run commands to perform manual
testing to get the desired responses above and some other edge cases. In order to have the 
tests run successfully the server should be started by running `go run main.go` from the 
`rest-service` directory.

An assumption I made was that if query params are passed outside of the ones listed above
the response will produce no results. In this case we are not exposing any extra unintended
behavior when a client tries to query by params not listed in the spec.

## Input Processing

The `input-processing` directory has an implementation that takes STDIN as a named pipe
and outputs every line that contains the word "error" to STDOUT.

The file has two implementations by running `go run main.go` for a quick and dirty approach
or a `go run main.go -e` for a more efficient implementation. The second approach works
better for large streams without `\n` characters and doesn't put everything into memory.
It performs very well for these situations but for a file with many new lines that is very
large the runtime can suffer.

The program also contains a pprof profiler for cpu and memory utilization that I used for
trying to optimize the runtime of the program. I included a `Makefile` with a few recipies
for building the application from a `Dockerfile` and creating an input file of 5GB. I used
the containerized version of the application to test out implementations with memory restrictions
to verify the performance of the application.

I also have some simple tests that I ran with the `test.sh` script and the `input_files` directory
to be able to make sure when I made changes to the source code that I was not breaking functionality.

