package utilrandom

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/aryyawijaya/go-storage-with-clean-arch/entity"
)

var r *rand.Rand

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Generates random integer [min, max]
func RandomInt(min, max int64) int64 {
	return min + r.Int63n(max-min+1)
}

// Generates random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[r.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Generates random enum from slice of T
func RandomEnum[T comparable](enum []T) T {
	l := len(enum)
	return enum[r.Intn(l)]
}

// Generates random time UTC [1970, 2070)
func RandomTime() time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()

	sec := RandomInt(min, max)
	return time.Unix(sec, 0).UTC()
}

func RandomPath() string {
	return fmt.Sprintf("%s/%s/", RandomString(5), RandomString(5))
}

func RandomFileExt() string {
	return fmt.Sprintf(".%s", RandomString(3))
}

func RandomByteSlices() []byte {
	return []byte(RandomString(5))
}

func RandomFile() *entity.File {
	return &entity.File{
		ID:        RandomInt(1, 100),
		Name:      RandomString(5),
		Access:    RandomEnum[entity.Access](entity.AllAccessValues()),
		Path:      RandomPath(),
		Ext:       RandomFileExt(),
		CreatedAt: RandomTime(),
		UpdatedAt: RandomTime(),
	}
}
