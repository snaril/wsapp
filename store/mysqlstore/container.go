package mysqlstore

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/snaril/wsapp/v2/store"
	"github.com/snaril/wsapp/v2/types"
	"github.com/snaril/wsapp/v2/util/keys"
	waLog "github.com/snaril/wsapp/v2/util/log"
	"go.mau.fi/util/random"
	"math/rand"
)

// Container is a wrapper for a SQL database that can contain multiple whatsmeow sessions.
type Container struct {
	db      gdb.DB
	dialect string
	log     waLog.Logger

	DatabaseErrorHandler func(device *store.Device, action string, attemptIndex int, err error) (retry bool)
}

var _ store.DeviceContainer = (*Container)(nil)

func New(db gdb.DB, log waLog.Logger) (*Container, error) {
	container := &Container{
		db:  db,
		log: log,
	}

	return container, nil
}

func (c *Container) NewDevice() *store.Device {
	device := &store.Device{
		Log:       c.log,
		Container: c,

		DatabaseErrorHandler: c.DatabaseErrorHandler,

		NoiseKey:       keys.NewKeyPair(),
		IdentityKey:    keys.NewKeyPair(),
		RegistrationID: rand.Uint32(),
		AdvSecretKey:   random.Bytes(32),
	}
	device.SignedPreKey = device.IdentityKey.CreateSignedPreKey(1)
	return device
}

// GetAllDevices finds all the devices in the database.
func (c *Container) GetAllDevices() ([]*store.Device, error) {
	return nil, nil
}

// GetFirstDevice is a convenience method for getting the first device in the store. If there are
// no devices, then a new device will be created. You should only use this if you don't want to
// have multiple sessions simultaneously.
func (c *Container) GetFirstDevice() (*store.Device, error) {
	return nil, nil
}

// GetDevice finds the device with the specified JID in the database.
//
// If the device is not found, nil is returned instead.
//
// Note that the parameter usually must be an AD-JID.
func (c *Container) GetDevice(jid types.JID) (*store.Device, error) {

	return nil, nil
}

func (c *Container) Close() error {
	return nil
}

func (c *Container) PutDevice(device *store.Device) error {

	return nil
}

func (c *Container) DeleteDevice(store *store.Device) error {

	return nil
}
