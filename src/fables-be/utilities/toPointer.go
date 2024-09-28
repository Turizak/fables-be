package utilities

func ToPointer(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
