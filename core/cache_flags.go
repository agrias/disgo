package core

// CacheFlags are used to enable/disable certain internal caches
type CacheFlags int

// values for CacheFlags
//goland:noinspection GoUnusedConst
const (
	CacheFlagsNone  CacheFlags = 0
	CacheFlagGuilds CacheFlags = 1 << iota
	CacheFlagDMChannels
	CacheFlagCategories
	CacheFlagTextChannels
	CacheFlagVoiceChannels
	CacheFlagStoreChannels
	CacheFlagRoles
	CacheFlagRoleTags
	CacheFlagEmojis
	CacheFlagVoiceStates
	CacheFlagStageInstances

	CacheFlagsChannels = CacheFlagDMChannels |
		CacheFlagCategories |
		CacheFlagTextChannels |
		CacheFlagVoiceChannels |
		CacheFlagStoreChannels

	CacheFlagsDefault = CacheFlagsChannels |
		CacheFlagRoles |
		CacheFlagEmojis

	CacheFlagsFullRoles = CacheFlagRoles |
		CacheFlagRoleTags

	CacheFlagsAll = CacheFlagsChannels |
		CacheFlagsFullRoles |
		CacheFlagEmojis |
		CacheFlagVoiceStates |
		CacheFlagStageInstances
)

// Add allows you to add multiple bits together, producing a new bit
func (c CacheFlags) Add(bits ...CacheFlags) CacheFlags {
	total := CacheFlags(0)
	for _, bit := range bits {
		total |= bit
	}
	c |= total
	return c
}

// Remove allows you to subtract multiple bits from the first, producing a new bit
func (c CacheFlags) Remove(bits ...CacheFlags) CacheFlags {
	total := CacheFlags(0)
	for _, bit := range bits {
		total |= bit
	}
	c &^= total
	return c
}

// HasAll will ensure that the bit includes all the bits entered
func (c CacheFlags) HasAll(bits ...CacheFlags) bool {
	for _, bit := range bits {
		if !c.Has(bit) {
			return false
		}
	}
	return true
}

// Has will check whether the Bit contains another bit
func (c CacheFlags) Has(bit CacheFlags) bool {
	return (c & bit) == bit
}

// MissingAny will check whether the bit is missing any one of the bits
func (c CacheFlags) MissingAny(bits ...CacheFlags) bool {
	for _, bit := range bits {
		if !c.Has(bit) {
			return true
		}
	}
	return false
}

// Missing will do the inverse of Bit.Has
func (c CacheFlags) Missing(bit CacheFlags) bool {
	return !c.Has(bit)
}
