## Adjust HTTP-MD5  Tool
This CLI tool makes http requests and prints the address of the request along with the
MD5 hash of the response.
- The tool is able to limit the number of parallel requests, to prevent exhausting local resources.
- The tool accepts a flag to indicate this limit, and its default is 10 if the flag is not provided.
### Usage

- Without `parallel` flag:

``` $ go run . http://www.adjust.com http://google.com```

Output: 
```
http://www.adjust.com d1b40e2a2ba488a054186e4ed0733f9752f66949
http://google.com 9d8ec921bdd275fb2a605176582e08758eb60641
```

- With `parallel` flag:

```
$ go run . -parallel 3 adjust.com google.com facebook.com yahoo.com yandex.com twitter.com
reddit.com/r/funny reddit.com/r/notfunny baroquemusiclibrary.com
```
Output:
```
http://google.com 8ff1c478ccca08cca025b028f68b352f
http://adjust.com 6b2560b9a5262571258cc173248b7492
http://yandex.com 4baab01ff9ff0f793bf423aeef539c9d
http://facebook.com ccae5ffa91c4936aef3efd5091a43f3e
http://twitter.com 857efe81a54c8b5c2241846ac4f08e66
http://reddit.com/r/funny ff3b2b7dcd9e716ca0adcbd208061c9a
http://reddit.com/r/notfunny ff3b2b7dcd9e716ca0adcbd208061c9a
http://yahoo.com e2d50a30b7bfbfda097d72e32578c6a6
http://baroquemusiclibrary.com 8e5138a0111364f08b10d37ed3371b11
```

### Tests

```
$ go test -cover
PASS
coverage: 47.3% of statements
```