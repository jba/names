package names

import nm "google.golang.org/grpc/naming"

func Resolve() nm.Watcher {
	return nil
}

type Watcher = nm.Watcher

func Alias() Watcher {
	return nil
}
