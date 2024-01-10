package p2p;

import (
    "encoding/json"
    "fmt"
    "strconv"
    "time"

    dhash "github.com/epsilondylan/blockchain/hash"

    "github.com/epsilondylan/blockchain/common"
    "github.com/epsilondylan/blockchain/models"
    pto "github.com/epsilondylan/blockchain/proto"
)

type P2P_Server struct {
    pto.UnimplementedP2PServer
}
