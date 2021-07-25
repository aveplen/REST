package handles

import (
	"github.com/aveplen/REST/internal/store"
	"github.com/sirupsen/logrus"
)

type IServer interface {
	GetLogger() *logrus.Logger
	GetStore() *store.Store
}
