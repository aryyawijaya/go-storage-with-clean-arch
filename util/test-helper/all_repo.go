package testhelper

import fileusecase "github.com/aryyawijaya/go-storage-with-clean-arch/modules/file/use-case"

type Repo interface {
	fileusecase.Repo
}
