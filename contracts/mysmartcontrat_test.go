package api

import (
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	api "github.com/sajanjswl/ethereum-notes/api"
	"github.com/stretchr/testify/suite"
)

type HelloworldTestSuite struct {
	suite.Suite
	auth            *bind.TransactOpts
	address         common.Address
	gAlloc          core.GenesisAlloc
	sim             *backends.SimulatedBackend
	MySmartContract *api.Api
}

func TestRunHelloworldSuite(t *testing.T) {
	suite.Run(t, new(HelloworldTestSuite))
}

func (s *HelloworldTestSuite) SetupTest() {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Println("error generating key")
	}
	// dd := big.NewInt(67)
	s.auth, err = bind.NewKeyedTransactorWithChainID(key, big.NewInt(997))

	if err != nil {
		log.Println("error generating authentications transaction options")
	}

	s.address = s.auth.From
	s.gAlloc = map[common.Address]core.GenesisAccount{
		s.address: {Balance: big.NewInt(100000000000)},
	}

	s.auth.GasFeeCap = big.NewInt(1000000000000000)

	s.sim = backends.NewSimulatedBackend(s.gAlloc, uint64(8750000000000000000))

	_, _, hw, e := api.DeployApi(s.auth, s.sim)
	s.MySmartContract = hw
	s.Nil(e)
	s.sim.Commit()
}

func (s *HelloworldTestSuite) TestSay() {
	str, err := s.MySmartContract.Hello(nil)
	if err != nil {
		log.Println("something")
	}
	_ = str
	s.Equal("Hello World", str)
	s.Nil(6)
}
