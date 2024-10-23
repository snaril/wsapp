package mysqlstore

import (
	"fmt"
	"github.com/snaril/wsapp/store"
	"github.com/snaril/wsapp/types"
	"github.com/snaril/wsapp/util/keys"
	"sync"
	"time"
)

var (
	_ store.IdentityStore        = (*Store)(nil)
	_ store.SessionStore         = (*Store)(nil)
	_ store.PreKeyStore          = (*Store)(nil)
	_ store.SenderKeyStore       = (*Store)(nil)
	_ store.AppStateSyncKeyStore = (*Store)(nil)
	_ store.AppStateStore        = (*Store)(nil)
	_ store.ContactStore         = (*Store)(nil)
)

type Store struct {
	*Container
	JID string

	preKeyLock sync.Mutex

	contactCache     map[types.JID]*types.ContactInfo
	contactCacheLock sync.Mutex
}

func NewSQLStore(c *Container, jid types.JID) *Store {
	return &Store{
		Container:    c,
		JID:          jid.String(),
		contactCache: make(map[types.JID]*types.ContactInfo),
	}
}

// IdentityStore

func (s *Store) PutIdentity(address string, key [32]byte) error {

	fmt.Println("--------------------PutIdentity")
	fmt.Println("--------------------PutIdentity")
	fmt.Println("--------------------PutIdentity")
	fmt.Println("--------------------PutIdentity")
	return nil
}

func (s *Store) DeleteAllIdentities(phone string) error {
	return nil
}

func (s *Store) DeleteIdentity(address string) error {
	return nil
}

func (s *Store) IsTrustedIdentity(address string, key [32]byte) (bool, error) {
	return true, nil
}

// SessionStore

func (s *Store) GetSession(address string) (session []byte, err error) {

	return
}

func (s *Store) HasSession(address string) (has bool, err error) {
	return
}

func (s *Store) PutSession(address string, session []byte) error {
	return nil
}

func (s *Store) DeleteAllSessions(phone string) error {
	return nil
}

func (s *Store) DeleteSession(address string) error {
	return nil
}

// PreKeyStore

func (s *Store) GenOnePreKey() (*keys.PreKey, error) {
	return nil, nil
}

func (s *Store) GetOrGenPreKeys(count uint32) ([]*keys.PreKey, error) {

	return nil, nil
}

func (s *Store) GetPreKey(id uint32) (*keys.PreKey, error) {
	return nil, nil
}

func (s *Store) RemovePreKey(id uint32) error {
	return nil
}

func (s *Store) MarkPreKeysAsUploaded(upToID uint32) error {
	return nil
}

func (s *Store) UploadedPreKeyCount() (count int, err error) {
	return 0, nil
}

// SenderKeyStore

func (s *Store) PutSenderKey(group, user string, session []byte) error {
	return nil
}

func (s *Store) GetSenderKey(group, user string) (key []byte, err error) {

	return nil, nil
}

// AppStateSyncKeyStore

func (s *Store) PutAppStateSyncKey(id []byte, key store.AppStateSyncKey) error {
	return nil
}

func (s *Store) GetAppStateSyncKey(id []byte) (*store.AppStateSyncKey, error) {

	return nil, nil
}

func (s *Store) GetLatestAppStateSyncKeyID() ([]byte, error) {

	return nil, nil
}

// AppStateStore

func (s *Store) PutAppStateVersion(name string, version uint64, hash [128]byte) error {
	return nil
}

func (s *Store) GetAppStateVersion(name string) (version uint64, hash [128]byte, err error) {

	return
}

func (s *Store) DeleteAppStateVersion(name string) error {
	return nil
}

func (s *Store) PutAppStateMutationMACs(name string, version uint64, mutations []store.AppStateMutationMAC) error {

	return nil
}

func (s *Store) DeleteAppStateMutationMACs(name string, indexMACs [][]byte) (err error) {

	return
}

func (s *Store) GetAppStateMutationMAC(name string, indexMAC []byte) (valueMAC []byte, err error) {

	return nil, nil
}

// ContactStore

func (s *Store) PutPushName(user types.JID, pushName string) (bool, string, error) {

	return false, "", nil
}

func (s *Store) PutBusinessName(user types.JID, businessName string) (bool, string, error) {

	return false, "", nil
}

func (s *Store) PutContactName(user types.JID, firstName, fullName string) error {
	return nil
}

func (s *Store) PutAllContactNames(contacts []store.ContactEntry) error {

	return nil
}

func (s *Store) GetContact(user types.JID) (types.ContactInfo, error) {

	return types.ContactInfo{}, nil
}

func (s *Store) GetAllContacts() (map[types.JID]types.ContactInfo, error) {

	return nil, nil
}

// ChatSettingsStore

func (s *Store) PutMutedUntil(chat types.JID, mutedUntil time.Time) error {

	return nil
}

func (s *Store) PutPinned(chat types.JID, pinned bool) error {
	return nil
}

func (s *Store) PutArchived(chat types.JID, archived bool) error {
	return nil
}

func (s *Store) GetChatSettings(chat types.JID) (settings types.LocalChatSettings, err error) {

	return
}

// MessageSecretStore

func (s *Store) PutMessageSecrets(inserts []store.MessageSecretInsert) (err error) {

	return
}

func (s *Store) PutMessageSecret(chat, sender types.JID, id types.MessageID, secret []byte) (err error) {
	return
}

func (s *Store) GetMessageSecret(chat, sender types.JID, id types.MessageID) (secret []byte, err error) {

	return
}

// PrivacyTokenStore

func (s *Store) PutPrivacyTokens(tokens ...store.PrivacyToken) error {

	return nil
}

func (s *Store) GetPrivacyToken(user types.JID) (*store.PrivacyToken, error) {
	return nil, nil
}
