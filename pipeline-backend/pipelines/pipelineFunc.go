package pipelines

// PostgresSQL To MYSQL
func AddPipeline(f func(), prev func()) func() {
	return func() {
		if prev != nil {
			prev()
		}
		f()
	}
}
