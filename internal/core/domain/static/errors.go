package static

import "errors"

var (
	ErrDocumentNotFound                = errors.New("document not found")
	ErrNotContainTranslationStoreSRCId = errors.New("not contain translation store srcId")
)
