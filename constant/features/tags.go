package features

func Tags() (tags []string) {
	if CMFA {
		tags = append(tags, "cmfa")
	}
	if WithLowMemory {
		tags = append(tags, "with_low_memory")
	}
	if NoFakeTCP {
		tags = append(tags, "no_fake_tcp")
	}
	if WithGVisor {
		tags = append(tags, "with_gvisor")
	}
	if WithCGO {
		tags = append(tags, "with_cgo")
	}
	return
}
