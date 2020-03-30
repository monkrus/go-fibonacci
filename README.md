# go-fibonacci
HTTP Server testing in Go


1. Fibonacci function: In the sequence of  x,y,z numbers  number z returns the sum of x+y.
   Our webserver returns the fibonacci sequence (or error) for any given number.

# SETUP
2. Create an endpoint (HandleFunc, "/fib", server by handleFib function )

3. Start a server 8080 (ListenAndServe)

4. Write a "Response-Request" function. We are getting a `num` parameter from request.(FormValue function)

5. Need to convert into `int` since function takes `int` (FormValue)

6. Atoi function returns `int` or error (StatusBadRequest)

7. If convertion is good, returns a line

8. Converting again: io.WriteString (int-->line)

10. Define Fib function

11. Create cache and initialize it with init function.

# TESTING

1. Testing function Handle Fib (using HTTPPIE, curl etc.)

2. Not testing a server, just a HandleFib function

3. Http test has recorder

4. Write test cases. We use `bytes` since response comes in bytes.

5. Handler returns --serve http-->fake request

6. Add cycle

7. http->recorder

8. Fake request (GET, sprintf)

9. Call ServeHTTP
-assert
-body ->returns bytes
-code status