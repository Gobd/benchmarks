package benchmarks

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
	_ "unsafe"

	"github.com/pingcap/go-ycsb/pkg/generator"

	"github.com/maypok86/benchmarks/client"
)

//go:noescape
//go:linkname fastrand runtime.fastrand
func fastrand() uint32

const dataLength = 3_000_000

var (
	values []string
	datas  []data

	clients = []client.Client[string, string]{
		&client.Otter[string, string]{},
		&client.Theine[string, string]{},
		&client.Ristretto[string, string]{},
	}
)

func init() {
	values = make([]string, 0, dataLength)
	for i := 0; i < dataLength; i++ {
		v := newValue(strconv.Itoa(i))
		values = append(values, v)
	}

	datas = []data{
		newZipfData(),
	}
}

func newKey(k string) string {
	return "0i02-3rj203rn230rjx0m238ex10eu1-x n-9u" + k
}

func newValue(v string) string {
	return "ololololooooooookkeeeke_9njijuinugyih" + v
}

type benchCase struct {
	name           string
	readPercentage int
	setPercentage  int
}

var benchCases = []benchCase{
	{"reads=100%,writes=0%", 100, 0},
	{"reads=75%,writes=25%", 75, 25},
	{"reads=50%,writes=50%", 50, 45},
	{"reads=25%,writes=75%", 25, 75},
	{"reads=0%,writes=100%", 0, 100},
}

type data struct {
	name string
	keys []string
}

func newZipfData() data {
	// populate using realistic distribution
	z := generator.NewScrambledZipfian(0, dataLength/3, generator.ZipfianConstant)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	keys := make([]string, 0, dataLength)
	for i := 0; i < dataLength; i++ {
		k := newKey(strconv.Itoa(int(z.Next(r))))
		keys = append(keys, k)
	}

	return data{
		name: "Zipf",
		keys: keys,
	}
}

func runParallelBenchmark(b *testing.B, benchFunc func(pb *testing.PB)) {
	b.ResetTimer()
	b.ReportAllocs()
	b.RunParallel(benchFunc)
}

func runCacheBenchmark(
	b *testing.B,
	benchCase benchCase,
	keys []string,
	c client.Client[string, string],
) {
	c.Init(dataLength / 10)

	for i := 0; i < dataLength; i++ {
		c.Set(keys[i], values[i])
	}

	b.ReportAllocs()
	runParallelBenchmark(b, func(pb *testing.PB) {
		// convert percent to permille
		setThreshold := 10 * benchCase.readPercentage
		for pb.Next() {
			op := int(fastrand() % 1000)
			i := int(fastrand() % uint32(dataLength))
			if op >= setThreshold {
				c.Set(keys[i], values[i])
			} else {
				c.Get(keys[i])
			}
		}
	})
}

func BenchmarkCache(b *testing.B) {
	for _, data := range datas {
		for _, benchCase := range benchCases {
			for _, c := range clients {
				name := fmt.Sprintf("%s_%s_%s", data.name, c.Name(), benchCase.name)
				b.Run(name, func(b *testing.B) {
					runCacheBenchmark(b, benchCase, data.keys, c)
				})
			}
		}
	}
}