## Golang API Skeleton

This is a go api quick start. We have a `.toml` file where we are going to add the configuration (for now we only have the port)

- It is based in golang chi, you just have to add your routes in `routing.go`.

```
{
  Pattern: "/",
  Handlers: routing.Handlers{
    {
      Method:  http.MethodGet,
      Handler: handler.HelloWorld(),
    },
    {
      Method:  http.MethodPost,
      Handler: handler.HelloWorldPostWhatever(),
    },
  },
}
```

In handler we have our handlers

With this you have the skeleton to develop a api in golang
