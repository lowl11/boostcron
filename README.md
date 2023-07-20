# boostcron

Get started
```go
cron := boostcron.New(boostcron.Config{
    ErrorHandler: func(err error) error {
        log.Debug("inner error:", err)
        return nil
    },
})

cron.Every(2).Seconds().Do(func() error {
    fmt.Println("every seconds action")
    return nil
})

cron.Every(5).Seconds().Do(func() error {
    return errors.New("some job error")
})

cron.Cron("0 40 16 ? * * *").Do(func() error {
    fmt.Println("cron expression job")

    return nil
})

cron.Run()
```