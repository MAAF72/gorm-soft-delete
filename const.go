package gorm_soft_delete

type ContextKey struct {
	name string
}

var (
	GORM_CTX_DELETED_BY = &ContextKey{"GORM_CTX_DELETED_BY"}
)
