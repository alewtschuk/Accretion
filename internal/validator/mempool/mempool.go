// package declataion
package mempool

// imports
import (
	"accretion/internal/common"
	"fmt"
	"os"
	"slices"
	"strconv"
	"time"
)

// structure definitions
type MemPool struct {
	Pending []common.ITx
}

// function definitions
func NewMemPool() *MemPool {

	var result *MemPool = &MemPool{
		Pending: make([]common.ITx, 0),
	}

	// start a periodic job to clear the mempool
	go func(m *MemPool) {

		// initialized data
		var err error = nil
		var evictionString string = os.Getenv("MEMPOOL_EVICTION")
		var evictionTime int = 0

		// ascii to int
		evictionTime, err = strconv.Atoi(evictionString)
		if err != nil {

			// sane default
			evictionTime = 5
		}

		// logs
		fmt.Printf("[mempool] Employing the janitor\n")

		// forever
		for {

			// evict the block
			m.Clean()

			// periodically
			time.Sleep(time.Duration(evictionTime) * time.Second)
		}
	}(result)

	// done
	return result
}

// method definitions
func (m *MemPool) Add(t common.ITx) {

	// fast fail
	if !t.Verify() {

		// logs
		fmt.Printf("[mempool] Someone is a dirty fraudster!\n")

		// done
		return
	}

	// initialized data
	var data []byte = []byte{}

	// serialize json
	data, _ = t.MarshalJSON()

	// verified news
	if !slices.Contains(m.Pending, t) {

		// new news
		fmt.Printf("[mempool] \"%s\"\n", string(data))
	} else {

		// old news
		fmt.Printf("[mempool] Yeah, we know\n")
	}
}

func (m *MemPool) List() []common.ITx {

	// logs
	fmt.Printf("[mempool] Spilling the tea\n")

	// done
	return m.Pending
}

func (m *MemPool) Clean() {

	// fast fail
	if len(m.Pending) == 0 {
		return
	}

	// logs
	fmt.Printf("[mempool] Forgetting\n")
	m.Pending = make([]common.ITx, 0)
}
