package snapshot

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strconv"

	"github.com/Taraxa-project/taraxa-snapshotter/internal/dir"
	"github.com/Taraxa-project/taraxa-snapshotter/internal/upload/ipfs"
	log "github.com/sirupsen/logrus"
)

type Snapshot struct {
	// Number is the snapshot number
	Number int `json:"number"`
	// DBDir is the path to the db directory
	DBDir *dir.DirEntry `json:"db_dir"`
	// StateDBDir is the path to the state_db directory
	StateDBDir *dir.DirEntry `json:"state_db_dir"`
}

// NewSnapshotFromDir creates a new instance of the Snapshot struct
func NewSnapshotFromDir(d *dir.DirEntry) *Snapshot {
	extract := regexp.MustCompile(`(\d+)`)
	matches := extract.FindStringSubmatch(d.Name)
	if len(matches) < 2 {
		log.Fatal("Failed to extract snapshot number from directory name")
	}

	i, err := strconv.Atoi(matches[1])
	if err != nil {
		log.Fatal("Failed to extract snapshot number from directory name")
	}

	stateDBDir := dir.NewDirEntry(path.Dir(d.Path) + "/state_db" + fmt.Sprint(i))

	return &Snapshot{
		Number:     i,
		DBDir:      d,
		StateDBDir: stateDBDir,
	}
}

// Upload uploads the snapshot to IPFS
func (s *Snapshot) Upload(ipfsClient *ipfs.IPFSClient) error {
	files := s.DBDir.Files()
	files = append(files, s.StateDBDir.Files()...)

	for _, f := range files {
		log.Info("Publishing file: ", f.Path)
		f, err := os.Open(f.Path)
		if err != nil {
			log.Fatal(err)
		}
		cid, err := ipfsClient.Add(f)
		if err != nil {
			log.WithError(err).Fatal("Failed to upload file to IPFS")
		}

		log.Info("CID: ", cid)
	}

	return nil
}
