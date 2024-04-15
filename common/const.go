package common

const (
	// Table names
	TbNamePosts    = "posts"
	TbNameUsers    = "users"
	TbNameSessions = "user_sessions"
	TbNameImages   = "images"

	//Len of salts
	LenSalt         = 30
	LenRefreshToken = 30

	// Keys
	KeyRequester = "requester"

	// Contexts
	CtxWithPubSub = "pubsub"

	// Channels
	ChannelUserChangedAvatar = "UserChangedAvatar"
)
