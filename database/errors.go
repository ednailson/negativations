package database

import "github.com/pkg/errors"

var ErrReadByDocument = errors.New("failed to read by document")
var ErrReadResult = errors.New("failed to decode data from database")
var ErrSaveDocument = errors.New("failed to save document")
var ErrInitCollection = errors.New("failed to init collection")
var ErrConnecting = errors.New("failed to connect to database")
var ErrInitDatabase = errors.New("failed to init database")
