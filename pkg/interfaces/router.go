package interfaces

type CronRouter interface {
	Every(every int) EveryScheduler
	Cron(expression string) CronScheduler
}
