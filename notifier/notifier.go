package notifier

// Notifier is a notification interface
type Notifier interface {
	Notify(body string) (err error)
}
