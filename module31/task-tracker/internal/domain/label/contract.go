package label

import "context"

// LabelCreator представляет контракт для создания метки задачи.
type LabelCreator interface {
	Create(ctx context.Context, label *Label) (LabelID, error)
}
