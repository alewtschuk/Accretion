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
func (m *MemPool) Len() int {

	// done
	return len(m.Pending)
}

func (m *MemPool) Add(t common.ITx) {

	// fast fail
	if !t.Verify() {

		// logs
		fmt.Printf("[mempool] Someone is a dirty fraudster!\n")

		// done
		return
	}

	// verified news
	if !slices.Contains(m.Pending, t) {

		// initialized data
		// var data []byte = []byte{}

		// serialize json
		// data, _ = t.MarshalJSON()

		// new news
		// fmt.Printf("[mempool] \"%s\"\n", string(data))
	} else {

		// old news
		fmt.Printf("[mempool] Yeah, we know\n")
	}

	// block building
	{

		// initialized data
		var err error = nil
		var b common.IBlock = nil
		var bb common.IBlockBuilder = common.NewBlockBuilder()

		// add the transaction to the block
		bb.AddTransaction(t)

		// build the block
		b, err = bb.Build()
		if err != nil {
			fmt.Printf("Failed to build block!\n")
		}

		// print the block
		fmt.Printf("%+v\n", b)
	}
}

func (m *MemPool) List(lim int) []common.ITx {

	// done
	return m.Pending[0:lim]
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
