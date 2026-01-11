package postgres

type ToDomain[D any] interface {
	ToDomain() *D
}

func MapToDomain[M ToDomain[D], D any](models []M) []D {
	items := make([]D, len(models))
	for i := range models {
		items[i] = *models[i].ToDomain()
	}
	return items
}
