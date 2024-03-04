# go-logs
An Async Logging utility, for non blocking logs


## Usage

### Handler

```go
logs.Handler.Start()

//....

logs.Handler.Join()

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