package stratum

// Share represents a share submitted by a worker
type Share struct {
	JobID        uint64
	WorkerHashID int64
	IP           uint32
	UserID       int32
	Share        uint64
	Timestamp    uint32
	BlockBits    uint32
	Result       int
}
