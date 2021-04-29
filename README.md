Command line tool to help parse URIs so that I can use them in shell scripts.

API is designed to be similar to `printf`.

```
go install github.com/aizatto/urlf
```

How to use:

```sh
urlf "h P" http://example.com/
# example.com 80
```

Usage:
```sh
urlf <url>
  Prints outs the components of a url

urlf <format> <url...>
  Prints outs the url according to the given format

  Format:
    - s: Scheme
    - u: User
    - p: Password
    - h: Host
    - P: Port
    - p: Path
    - q: Query
    - f: Fragment

  Supports multiple urls

  Example: 
    urlf "h P" http://example.com:443 https://aizatto.com/
```

How to use in shell scripts:
```sh
telnet $(urlf "h P" http://example.com:100/)
telnet $(urlf "h P" $ENV_APP)
```