package keyctl

import "github.com/lox/keyring/v2"

type Item = keyring.Item
type Metadata = keyring.Metadata

const KeyCtlBackend = keyring.KeyCtlBackend

var (
	ErrKeyNotFound          = keyring.ErrKeyNotFound
	ErrMetadataNotSupported = keyring.ErrMetadataNotSupported
	ErrUnavailable          = keyring.ErrUnavailable
)
