package capacity_manager

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

// CapacityManager manages the prover capacity concurrent-safely.
type CapacityManager struct {
	capacity              map[uint64]bool
	tempCapacity          []time.Time
	tempCapacityExpiresAt time.Duration
	maxCapacity           uint64
	mutex                 sync.RWMutex
	rpc 				  *rpc.Client
}

// New creates a new CapacityManager instance.
func New(capacity uint64, tempCapacityExpiresAt time.Duration, rpc *rpc.Client) *CapacityManager {
	return &CapacityManager{
		capacity:              make(map[uint64]bool),
		maxCapacity:           capacity,
		tempCapacity:          make([]time.Time, 0),
		tempCapacityExpiresAt: tempCapacityExpiresAt,
		rpc: rpc,
	}
}

// ReadCapacity reads the current capacity.
func (m *CapacityManager) ReadCapacity() uint64 {
	//m.mutex.RLock()
	//defer m.mutex.RUnlock()

	hasResource, err := m.rpc.ApusTask.HasResource(&bind.CallOpts{})

	log.Info("Reading capacity",
		"hasResource", hasResource,
		"err", err,
	)

	return hasResource.Uint64()
}

// ReleaseOneCapacity releases one capacity.
func (m *CapacityManager) ReleaseOneCapacity(blockID uint64) (uint64, bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return m.ReadCapacity(), true
}

// TakeOneCapacity takes one capacity.
func (m *CapacityManager) TakeOneCapacity(blockID uint64) (uint64, bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	task, _, err := m.rpc.ApusTask.GetTask(&bind.CallOpts{}, 0, big.NewInt(int64(blockID)))
	if err != nil {
		log.Error("Apus Market: onBlockProposed", "index", "get_task", "err", err)
		return 0, false
	}
	if task.UniqID.Uint64() == 0 || task.UniqID.Uint64() == blockID {
		return m.ReadCapacity(), true
	}
	log.Info("take one capacity", "blockID", blockID)
	capacity := m.ReadCapacity()
	return capacity, capacity > 0
}

func (m *CapacityManager) TakeOneTempCapacity() (uint64, bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	// clear expired tempCapacities

	m.clearExpiredTempCapacities()

	hasResource, err := m.rpc.ApusTask.HasResource(&bind.CallOpts{})
	if err != nil {
		log.Error("Take one Temp capacity",
			"error", err,
		)
		return 0, false
	}
	log.Info("Take one Temp capacity",
		"cur_capacity", hasResource.Uint64(),
		"tmp_capacity", len(m.tempCapacity),
		"can take tmp capacity", hasResource.Int64() > int64(len(m.tempCapacity)),
	)
	if hasResource.Int64() <= int64(len(m.tempCapacity)) || hasResource.Uint64() == 0 {
		return 0, false
	}

	m.tempCapacity = append(m.tempCapacity, time.Now().UTC())


	return hasResource.Uint64() - uint64(len(m.tempCapacity)), hasResource.Int64() + 1 >= int64(len(m.tempCapacity))
}

func (m *CapacityManager) clearExpiredTempCapacities() {

	tmp := make([]time.Time, 0)
	for _, c := range m.tempCapacity {
		if time.Now().UTC().Sub(c) > m.tempCapacityExpiresAt {
			//m.tempCapacity = append(m.tempCapacity[:i], m.tempCapacity[i+1:]...)
			log.Info("Cleared one temp capacity",
				"maxCapacity", m.maxCapacity,
				"currentCapacityAfterClearing", m.maxCapacity-uint64(len(m.capacity)),
				"currentUsageAfterClearing", len(m.capacity),
				"tempCapacity", m.maxCapacity-uint64(len((m.tempCapacity))),
				"tempCapacityUsage", len(m.tempCapacity),
			)
		} else {
			tmp = append(tmp, c)
		}
	}
	m.tempCapacity = tmp
}
