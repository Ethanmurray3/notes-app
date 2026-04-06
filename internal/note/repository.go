package note

// think about adding context.Context eventually. maybe look up what this means

type Repository interface {
	Create(note Note) (Note, error)
	GetByID(int64) (Note, error)
	List() ([]Note, error)
	Update(note Note) error
	Delete(int64) error
}
