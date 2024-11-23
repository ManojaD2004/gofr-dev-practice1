package pipelines

import "context"

// PostgresSQL To MYSQL
func AddPipeline(prev func(c1 context.Context), f func(c1 context.Context)) func(c1 context.Context) {
	return func(c1 context.Context) {
		if prev != nil {
			prev(c1)
		}
		f(c1)
	}
}
