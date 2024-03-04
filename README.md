# go-logs
An Async Logging utility, for non blocking logs

> [About Publishing Go-Modules](https://go.dev/blog/publishing-go-modules)

## Usage

### Handler

```go
logs.Handler.Start() // At the start of the Program, starts the Log-Handler Go-Routine and creates the Log-Channels

//....

logs.Handler.Join() // At the End of the Program, Waits for all submitted Print Jobs to be finished an closes channels

```

### Logs

```go
logs.Debug(<MYMSG(string)>, <KEY(string)>, <VALUE(any)>...)
logs.Info(<MYMSG(string)>, <KEY(string)>, <VALUE(any)>...)
logs.Warn(<MYMSG(string)>, <KEY(string)>, <VALUE(any)>...)
logs.Error(<MYMSG(string)>, <KEY(string)>, <VALUE(any)>...)
logs.Fatal(<MYMSG(string)>, <KEY(string)>, <VALUE(any)>...)

logs.SendLog(<CHANNEL(MessageChannel)>, <MSG(string)>, <DATA([]any)>) // --> the underlaying function for all above
```