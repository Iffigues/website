package linuxCommand

func addOptTab(r []string, e ...string) (b []string) {
	for _, val := range e {
		r = append(r, val);
	}
	return r
}

