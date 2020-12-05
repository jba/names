package names

import "google.golang.org/grpc/naming"

func Resolve() naming.Watcher {
	return nil
}

type Watcher = naming.Watcher

func Alias() Watcher {
	return nil
}
