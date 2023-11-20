package im

type AddBlackListBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED
	BlackListInfos []AddBlackListBodyBlackListInfosItem `json:"BlackListInfos"`

	// REQUIRED; 用户 ID
	UserID int64 `json:"UserId"`

	// 信箱，用做逻辑隔离。默认值为 0
	InboxType *int32 `json:"InboxType,omitempty"`
}

type AddBlackListBodyBlackListInfosItem struct {

	// REQUIRED; 加入黑名单的用户 ID
	BlackListUserID int64 `json:"BlackListUserId"`

	// 黑名单用户扩展字段
	BlackListUserExt map[string]string `json:"BlackListUserExt,omitempty"`
}

type AddBlackListRes struct {

	// REQUIRED
	ResponseMetadata AddBlackListResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result AddBlackListResResult `json:"Result"`
}

type AddBlackListResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                `json:"Version"`
	Error   *AddBlackListResResponseMetadataError `json:"Error,omitempty"`
}

type AddBlackListResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type AddBlackListResResult struct {

	// REQUIRED
	FailedInfos []AddBlackListResResultFailedInfosItem `json:"FailedInfos"`
}

type AddBlackListResResultFailedInfosItem struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`

	// REQUIRED; 添加黑名单失败的用户的 ID
	UserID int64 `json:"UserId"`
}

type AddFriendBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 好友 UserId，一次最多添加 10 个好友
	FriendUserIDs []int64 `json:"FriendUserIds"`

	// REQUIRED; 用户 UserId
	UserID int64 `json:"UserId"`

	// 好友扩展字段
	Ext map[string]string `json:"Ext,omitempty"`

	// 信箱，用做逻辑隔离 默认值为 0
	InboxType *int32 `json:"InboxType,omitempty"`
}

type AddFriendRes struct {

	// REQUIRED
	ResponseMetadata AddFriendResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result AddFriendResResult `json:"Result"`
}

type AddFriendResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                             `json:"Version"`
	Error   *AddFriendResResponseMetadataError `json:"Error,omitempty"`
}

type AddFriendResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type AddFriendResResult struct {

	// 添加失败的好友信息
	FailedInfos []AddFriendResResultFailedInfosItem `json:"FailedInfos"`
}

type AddFriendResResultFailedInfosItem struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`

	// REQUIRED; 好友 UserId
	UserID int64 `json:"UserId"`
}

type BatchAddBlockParticipantsBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; key 为群成员 ID，value 为禁言或者拉黑时长， 单位为秒
	ParticipantBlockInfos map[string]int64 `json:"ParticipantBlockInfos"`

	// 操作行为 0：禁言。用户无法在直播群中发言。 1：拉黑。用户无法加入直播群。 默认值为 0
	BlockAction *int32 `json:"BlockAction,omitempty"`
}

type BatchAddBlockParticipantsRes struct {

	// REQUIRED
	ResponseMetadata BatchAddBlockParticipantsResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result BatchAddBlockParticipantsResResult `json:"Result"`
}

type BatchAddBlockParticipantsResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                             `json:"Version"`
	Error   *BatchAddBlockParticipantsResResponseMetadataError `json:"Error,omitempty"`
}

type BatchAddBlockParticipantsResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type BatchAddBlockParticipantsResResult struct {

	// REQUIRED; 禁言或者拉黑失败的会话成员 UserId
	FailedParticipantUserIDs []int64 `json:"FailedParticipantUserIds"`
}

type BatchAddConversationParticipantBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 执行加群操作人的 UserId
	Operator int64 `json:"Operator"`

	// REQUIRED; 群成员信息
	ParticipantInfos []BatchAddConversationParticipantBodyParticipantInfosItem `json:"ParticipantInfos"`

	// 是否开启屏障。如设置屏障，新加入用户无法看到历史会话消息。 false：不开启。 true：开启。 默认值为false。
	Barrier *bool `json:"Barrier,omitempty"`
}

type BatchAddConversationParticipantBodyParticipantInfosItem struct {

	// REQUIRED; 添加成员所属 UserId，UserId 必须大于 0。
	ParticipantUserID int64 `json:"ParticipantUserId"`

	// 成员扩展字段
	Ext map[string]string `json:"Ext,omitempty"`

	// 成员等级
	Level *int32 `json:"Level,omitempty"`

	// 成员昵称
	NickName *string `json:"NickName,omitempty"`

	// 成员已读位置点
	ReadIndex *int64 `json:"ReadIndex,omitempty"`

	// 成员身份，可取值为：0，1，2。 0: 普通成员。 1: 群主。添加群主时，需确保会话中的 ownerUid 与群主的 UserId 相同。 2：群管理员。 默认值为0，值不合法时自动调整为默认值。
	Role *int64 `json:"Role,omitempty"`
}

type BatchAddConversationParticipantRes struct {

	// REQUIRED
	ResponseMetadata BatchAddConversationParticipantResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result BatchAddConversationParticipantResResult `json:"Result"`
}

type BatchAddConversationParticipantResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                                   `json:"Version"`
	Error   *BatchAddConversationParticipantResResponseMetadataError `json:"Error,omitempty"`
}

type BatchAddConversationParticipantResResponseMetadataError struct {

	// REQUIRED; 错误码
	Code string `json:"Code"`

	// REQUIRED; 错误���息
	Message string `json:"Message"`
}

type BatchAddConversationParticipantResResult struct {

	// REQUIRED; 成功时为空，失败时返回失败用户 ID
	FailedUserIDs []int64 `json:"FailedUserIds"`
}

type BatchAddManagerBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 要添加为管理员的群成员 userId
	ManagerUserIDs []int64 `json:"ManagerUserIds"`

	// REQUIRED; 操作人 UserId
	Operator int64 `json:"Operator"`
}

type BatchAddManagerRes struct {

	// REQUIRED
	ResponseMetadata BatchAddManagerResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result BatchAddManagerResResult `json:"Result"`
}

type BatchAddManagerResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                   `json:"Version"`
	Error   *BatchAddManagerResResponseMetadataError `json:"Error,omitempty"`
}

type BatchAddManagerResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type BatchAddManagerResResult struct {

	// REQUIRED; 添加失败的 UserId
	FailedManagerUserIDs []int64 `json:"FailedManagerUserIds"`
}

type BatchAddWhitelistParticipantBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 操作人 UserId
	Operator int64 `json:"Operator"`

	// REQUIRED; 要添加白名单成员 UserId
	ParticipantUserIDs []int64 `json:"ParticipantUserIds"`
}

type BatchAddWhitelistParticipantRes struct {

	// REQUIRED
	ResponseMetadata BatchAddWhitelistParticipantResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result BatchAddWhitelistParticipantResResult `json:"Result"`
}

type BatchAddWhitelistParticipantResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                                `json:"Version"`
	Error   *BatchAddWhitelistParticipantResResponseMetadataError `json:"Error,omitempty"`
}

type BatchAddWhitelistParticipantResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type BatchAddWhitelistParticipantResResult struct {

	// REQUIRED; 添加失败的 UserId
	FailedUserIDs []int64 `json:"FailedUserIds"`
}

type BatchDeleteBlockParticipantsBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 取消禁言或者取消拉黑的会员成员 UserId 列表
	ParticipantUserIDs []int64 `json:"ParticipantUserIds"`

	// 操作行为。 0：取消禁言。 1：取消拉黑。 默认值为 0。
	BlockAction *int32 `json:"BlockAction,omitempty"`
}

type BatchDeleteBlockParticipantsRes struct {

	// REQUIRED
	ResponseMetadata BatchDeleteBlockParticipantsResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result BatchDeleteBlockParticipantsResResult `json:"Result"`
}

type BatchDeleteBlockParticipantsResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                                `json:"Version"`
	Error   *BatchDeleteBlockParticipantsResResponseMetadataError `json:"Error,omitempty"`
}

type BatchDeleteBlockParticipantsResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type BatchDeleteBlockParticipantsResResult struct {

	// REQUIRED; 取消禁言或者取消拉黑失败的会话成员 UserId
	FailedParticipantUserIDs []int64 `json:"FailedParticipantUserIds"`
}

type BatchDeleteConversationParticipantBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 执行删除群成员操作人的 UserId
	Operator int64 `json:"Operator"`

	// REQUIRED; 需要删除群成员的 UserId
	ParticipantUserIDs []int64 `json:"ParticipantUserIds"`
}

type BatchDeleteConversationParticipantRes struct {

	// REQUIRED
	ResponseMetadata BatchDeleteConversationParticipantResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result BatchDeleteConversationParticipantResResult `json:"Result"`
}

type BatchDeleteConversationParticipantResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                                      `json:"Version"`
	Error   *BatchDeleteConversationParticipantResResponseMetadataError `json:"Error,omitempty"`
}

type BatchDeleteConversationParticipantResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type BatchDeleteConversationParticipantResResult struct {

	// REQUIRED; 成功时为空，失败时返回失败用户 ID
	FailedUserIDs []int64 `json:"FailedUserIds"`
}

type BatchGetBlockParticipantsBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 分批起始位置
	Cursor int64 `json:"Cursor"`

	// REQUIRED; 每批查询条数，最大值为 20。
	Limit int64 `json:"Limit"`

	// 操作行为 0：获取禁言列表 1：获取拉黑列表 默认为 0
	BlockAction *int32 `json:"BlockAction,omitempty"`
}

type BatchGetBlockParticipantsRes struct {

	// REQUIRED
	ResponseMetadata BatchGetBlockParticipantsResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result BatchGetBlockParticipantsResResult `json:"Result"`
}

type BatchGetBlockParticipantsResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                             `json:"Version"`
	Error   *BatchGetBlockParticipantsResResponseMetadataError `json:"Error,omitempty"`
}

type BatchGetBlockParticipantsResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type BatchGetBlockParticipantsResResult struct {

	// REQUIRED; 是否还有更多数据
	HasMore bool `json:"HasMore"`

	// REQUIRED
	NextCursor int64 `json:"NextCursor"`

	// REQUIRED
	Participants []BatchGetBlockParticipantsResResultParticipantsItem `json:"Participants"`
}

type BatchGetBlockParticipantsResResultParticipantsItem struct {

	// REQUIRED; 直播群群成员头像
	AvatarURL string `json:"AvatarUrl"`

	// REQUIRED; 禁言/拉黑到何时对应时间戳，单位为秒
	BlockTime int64 `json:"BlockTime"`

	// REQUIRED; 禁言/拉黑设置对应的时间戳，单位为秒
	CreateTime int64 `json:"CreateTime"`

	// REQUIRED; 直播群群成员扩展字段
	Ext interface{} `json:"Ext"`

	// REQUIRED; 直播群群成员昵称
	NickName string `json:"NickName"`

	// REQUIRED; 群成员 ID
	ParticipantUserID int64 `json:"ParticipantUserId"`
}

type BatchGetConversationParticipantBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 需要查询会话成员的 UserId
	ParticipantUserIDs []int64 `json:"ParticipantUserIds"`
}

type BatchGetConversationParticipantRes struct {

	// REQUIRED
	ResponseMetadata BatchGetConversationParticipantResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result BatchGetConversationParticipantResResult `json:"Result"`
}

type BatchGetConversationParticipantResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                                   `json:"Version"`
	Error   *BatchGetConversationParticipantResResponseMetadataError `json:"Error,omitempty"`
}

type BatchGetConversationParticipantResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type BatchGetConversationParticipantResResult struct {

	// REQUIRED
	Participants []BatchGetConversationParticipantResResultParticipantsItem `json:"Participants"`
}

type BatchGetConversationParticipantResResultParticipantsItem struct {

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 操作人对应的 UserId
	Operator int64 `json:"Operator"`

	// REQUIRED; 群成员 ID
	ParticipantUserID int64 `json:"ParticipantUserId"`

	// REQUIRED; 成员身份。 0: 普通成员。 1: 群主。 2：群管理员。
	Role int64 `json:"Role"`

	// 禁言时间戳，单位为秒。0表示不禁言
	BlockTime *int64 `json:"BlockTime,omitempty"`

	// 成员进群时间对应时间戳，单位为秒
	CreateTime *int64 `json:"CreateTime,omitempty"`

	// 成员的扩展字段
	Ext map[string]string `json:"Ext,omitempty"`

	// 成员等级
	Level *int32 `json:"Level,omitempty"`

	// 成员昵称
	NickName *string `json:"NickName,omitempty"`

	// 成员状态。 0：正常 1：退出
	Status *int32 `json:"Status,omitempty"`
}

type BatchGetConversationsBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID []int64 `json:"ConversationShortId"`

	// 是否忽略获取会话成员数。 true： 忽略。 false：不忽略。 默认值为 true。
	SkipMemberCount *bool `json:"SkipMemberCount,omitempty"`
}

type BatchGetConversationsRes struct {

	// REQUIRED
	ResponseMetadata BatchGetConversationsResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result BatchGetConversationsResResult `json:"Result"`
}

type BatchGetConversationsResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                         `json:"Version"`
	Error   *BatchGetConversationsResResponseMetadataError `json:"Error,omitempty"`
}

type BatchGetConversationsResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type BatchGetConversationsResResult struct {

	// REQUIRED
	ConversationCoreInfos []BatchGetConversationsResResultConversationCoreInfosItem `json:"ConversationCoreInfos"`
}

type BatchGetConversationsResResultConversationCoreInfosItem struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 会话类型。 1：单聊 2:群聊 100:直播群
	ConversationType int32 `json:"ConversationType"`

	// REQUIRED; 群聊创建时间戳，单位为秒
	CreateTime int64 `json:"CreateTime"`

	// REQUIRED; 创群人 UserId
	CreatorUserID int64 `json:"CreatorUserId"`

	// REQUIRED; 信箱,用于逻辑隔离
	InboxType int32 `json:"InboxType"`

	// REQUIRED; 修改时间戳，单位为秒
	ModifyTime int64 `json:"ModifyTime"`

	// REQUIRED; 群主 UserId
	OwnerUserID int64 `json:"OwnerUserId"`

	// 群头像 url
	AvatarURL *string `json:"AvatarUrl,omitempty"`

	// 群描述
	Description *string `json:"Description,omitempty"`

	// 会话的扩展字段。
	Ext map[string]string `json:"Ext,omitempty"`

	// 会话成员数
	MemberCount *int64 `json:"MemberCount,omitempty"`

	// 群名
	Name *string `json:"Name,omitempty"`

	// 群公告
	Notice *string `json:"Notice,omitempty"`

	// 直播群在线人数。
	OnlineCount *int64 `json:"OnlineCount,omitempty"`

	// 单聊会话另一个UserId
	OtherUserID *int64 `json:"OtherUserId,omitempty"`

	// 会话状态。 0：正常 1：已解散
	Status *int32 `json:"Status,omitempty"`
}

type BatchGetWhitelistParticipantBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 查询起始位置
	Cursor int64 `json:"Cursor"`

	// REQUIRED; 每批查询条数。最大值为 20。
	Limit int64 `json:"Limit"`
}

type BatchGetWhitelistParticipantRes struct {

	// REQUIRED
	ResponseMetadata BatchGetWhitelistParticipantResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result BatchGetWhitelistParticipantResResult `json:"Result"`
}

type BatchGetWhitelistParticipantResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                                `json:"Version"`
	Error   *BatchGetWhitelistParticipantResResponseMetadataError `json:"Error,omitempty"`
}

type BatchGetWhitelistParticipantResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type BatchGetWhitelistParticipantResResult struct {

	// REQUIRED; 是否还有更多数据
	HasMore bool `json:"HasMore"`

	// REQUIRED; 下一批数据起始位置。-1表示无更多数据。
	NextCursor int64 `json:"NextCursor"`

	// REQUIRED; 成员
	Participants []BatchGetWhitelistParticipantResResultParticipantsItem `json:"Participants"`
}

type BatchGetWhitelistParticipantResResultParticipantsItem struct {

	// REQUIRED; 直播群群成员头像
	AvatarURL string `json:"AvatarUrl"`

	// REQUIRED; 用户加入白名单时间，单位为秒
	CreateTime int64 `json:"CreateTime"`

	// REQUIRED; 直播群群成员扩展字段
	Ext map[string]string `json:"Ext"`

	// REQUIRED; 直播群群成员昵称
	NickName string `json:"NickName"`

	// REQUIRED; 执行加入白名单操作操作人 UserId
	Operator int64 `json:"Operator"`

	// REQUIRED; 白名单成员 UserId
	ParticipantUserID int64 `json:"ParticipantUserId"`
}

type BatchModifyConversationParticipantBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 执行修改群成员信息操作人的 UserId
	Operator int64 `json:"Operator"`

	// REQUIRED; 群成员信息
	ParticipantInfos []BatchModifyConversationParticipantBodyParticipantInfosItem `json:"ParticipantInfos"`
}

type BatchModifyConversationParticipantBodyParticipantInfosItem struct {

	// REQUIRED; 进行修改的群成员的 UserId，UserId 必须大于 0。
	ParticipantUserID int64 `json:"ParticipantUserId"`

	// 禁言时间戳，表示禁言到何时，单位为秒。0表示未禁言
	BlockTime *int64 `json:"BlockTime,omitempty"`

	// 成员扩展字段
	Ext map[string]string `json:"Ext,omitempty"`

	// 成员等级
	Level *int32 `json:"Level,omitempty"`

	// 成员昵称
	NickName *string `json:"NickName,omitempty"`

	// 成员身份，可取值为：0，1，2。 0: 普通成员。 1: 群主。 2：群管理员。 默认值为0，值不合法时自动调整为默认值。
	Role *int64 `json:"Role,omitempty"`
}

type BatchModifyConversationParticipantRes struct {

	// REQUIRED
	ResponseMetadata BatchModifyConversationParticipantResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result BatchModifyConversationParticipantResResult `json:"Result"`
}

type BatchModifyConversationParticipantResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                                      `json:"Version"`
	Error   *BatchModifyConversationParticipantResResponseMetadataError `json:"Error,omitempty"`
}

type BatchModifyConversationParticipantResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type BatchModifyConversationParticipantResResult struct {

	// REQUIRED; 成功时为空，失败时返回失败用户 ID
	FailedUserIDs []int64 `json:"FailedUserIds"`
}

type BatchRemoveManagerBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 操作人 UserId
	Operator int64 `json:"Operator"`

	// REQUIRED; 要移除的管理员 UserId
	RemoveManagerUserIDs []int64 `json:"RemoveManagerUserIds"`
}

type BatchRemoveManagerRes struct {

	// REQUIRED
	ResponseMetadata BatchRemoveManagerResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result BatchRemoveManagerResResult `json:"Result"`
}

type BatchRemoveManagerResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                      `json:"Version"`
	Error   *BatchRemoveManagerResResponseMetadataError `json:"Error,omitempty"`
}

type BatchRemoveManagerResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type BatchRemoveManagerResResult struct {

	// REQUIRED; 移除失败的管理员 UserId
	FailedManagerUserIDs []int64 `json:"FailedManagerUserIds"`
}

type BatchRemoveWhitelistParticipantBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 操作人 UserId
	Operator int64 `json:"Operator"`

	// REQUIRED; 要移除白名单用户 UserId
	ParticipantUserIDs []int64 `json:"ParticipantUserIds"`
}

type BatchRemoveWhitelistParticipantRes struct {

	// REQUIRED
	ResponseMetadata BatchRemoveWhitelistParticipantResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result BatchRemoveWhitelistParticipantResResult `json:"Result"`
}

type BatchRemoveWhitelistParticipantResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                                   `json:"Version"`
	Error   *BatchRemoveWhitelistParticipantResResponseMetadataError `json:"Error,omitempty"`
}

type BatchRemoveWhitelistParticipantResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type BatchRemoveWhitelistParticipantResResult struct {

	// REQUIRED; 移除失败的 UserId
	FailedUserIDs []int64 `json:"FailedUserIds"`
}

type BatchUpdateLiveParticipantsBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 需要更新的成员资料。一次最多支持 10 个成员
	ParticipantInfos []BatchUpdateLiveParticipantsBodyParticipantInfosItem `json:"ParticipantInfos"`
}

type BatchUpdateLiveParticipantsBodyParticipantInfosItem struct {

	// REQUIRED; 群成员用户ID
	UserID int64 `json:"UserId"`

	// 群成员头像。AvatarUrl、NickName和Ext均为非必填参数，但是至少需要填一个，否则服务端会报错。
	AvatarURL *string `json:"AvatarUrl,omitempty"`

	// 群成员扩展字段。
	Ext map[string]string `json:"Ext,omitempty"`

	// 群成员昵称
	NickName *string `json:"NickName,omitempty"`
}

type BatchUpdateLiveParticipantsRes struct {

	// REQUIRED
	ResponseMetadata BatchUpdateLiveParticipantsResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result BatchUpdateLiveParticipantsResResult `json:"Result"`
}

type BatchUpdateLiveParticipantsResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                               `json:"Version"`
	Error   *BatchUpdateLiveParticipantsResResponseMetadataError `json:"Error,omitempty"`
}

type BatchUpdateLiveParticipantsResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type BatchUpdateLiveParticipantsResResult struct {
	FailedInfos []BatchUpdateLiveParticipantsResResultFailedInfosItem `json:"FailedInfos"`
}

type BatchUpdateLiveParticipantsResResultFailedInfosItem struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`

	// REQUIRED; 更新资料失败的群成员的UserId
	UserID int64 `json:"UserId"`
}

type CreateConversationBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED
	ConversationCoreInfo CreateConversationBodyConversationCoreInfo `json:"ConversationCoreInfo"`

	// REQUIRED; 群主 UserId
	OwnerUserID int64 `json:"OwnerUserId"`

	// 幂等id，如果创建时指定了此字段，并且数据库中存在此 id 对应的会话，则不会重复创建，并且接口返回的Exist字段为true。
	IdempotentID *string `json:"IdempotentId,omitempty"`

	// 信箱，用做逻辑隔离 默认值为 0
	InboxType *int32 `json:"InboxType,omitempty"`

	// 另一个成员的 UserId， 创建单聊必填
	OtherUserID *int64 `json:"OtherUserId,omitempty"`
}

type CreateConversationBodyConversationCoreInfo struct {

	// REQUIRED; 会话类型 1：单聊 2:群聊 100：直播群
	ConversationType int32 `json:"ConversationType"`

	// 会话头像 url
	AvatarURL *string `json:"AvatarUrl,omitempty"`

	// 会话描述
	Description *string `json:"Description,omitempty"`

	// map 扩展字段
	Ext map[string]string `json:"Ext,omitempty"`

	// 会话名称
	Name *string `json:"Name,omitempty"`

	// 会话公告
	Notice *string `json:"Notice,omitempty"`
}

type CreateConversationRes struct {

	// REQUIRED
	ResponseMetadata CreateConversationResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result CreateConversationResResult `json:"Result"`
}

type CreateConversationResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                      `json:"Version"`
	Error   *CreateConversationResResponseMetadataError `json:"Error,omitempty"`
}

type CreateConversationResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type CreateConversationResResult struct {

	// REQUIRED
	ConversationInfo CreateConversationResResultConversationInfo `json:"ConversationInfo"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 会话是否存在
	Exist bool `json:"Exist"`
}

type CreateConversationResResultConversationInfo struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 会话类型。 1：单聊 2:群聊
	ConversationType int32 `json:"ConversationType"`

	// REQUIRED; 信箱,用于逻辑隔离
	InboxType int32 `json:"InboxType"`

	// 群头像 url
	AvatarURL *string `json:"AvatarUrl,omitempty"`

	// 群聊创建时间戳，单位为秒
	CreateTime *int64 `json:"CreateTime,omitempty"`

	// 创群人 UserId
	CreatorUserID *int64 `json:"CreatorUserId,omitempty"`

	// 群描述
	Description *string `json:"Description,omitempty"`

	// 会话的扩展字段。
	Ext map[string]string `json:"Ext,omitempty"`

	// 会话成员数
	MemberCount *int64 `json:"MemberCount,omitempty"`

	// 修改时间戳，单位为秒
	ModifyTime *int64 `json:"ModifyTime,omitempty"`

	// 群名
	Name *string `json:"Name,omitempty"`

	// 群公告
	Notice *string `json:"Notice,omitempty"`

	// 直播群在线人数
	OnlineCount *int64 `json:"OnlineCount,omitempty"`

	// 单聊另一个成员的UserId
	OtherUserID *int64 `json:"OtherUserId,omitempty"`

	// 群主 UserId
	OwnerUserID *int64 `json:"OwnerUserId,omitempty"`

	// 会话状态。 0：正常 1：已解散
	Status *int32 `json:"Status,omitempty"`
}

type DeleteConversationMessageBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 消息 ID
	MessageID int64 `json:"MessageId"`

	// 删除方式。 0：全部用户不可见。 1：仅发送人自己可见。 默认值为0。 直播群只允许传 0。
	DeleteType *int32 `json:"DeleteType,omitempty"`
}

type DeleteConversationMessageRes struct {

	// REQUIRED
	ResponseMetadata DeleteConversationMessageResResponseMetadata `json:"ResponseMetadata"`

	// 空。此接口无需关注
	Result interface{} `json:"Result,omitempty"`
}

type DeleteConversationMessageResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                             `json:"Version"`
	Error   *DeleteConversationMessageResResponseMetadataError `json:"Error,omitempty"`
}

type DeleteConversationMessageResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type DeleteFriendBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 要删除的好友 UserId，一次最多删除 10 个好友
	DeleteUserIDs []int64 `json:"DeleteUserIds"`

	// REQUIRED; 用户 UserId
	UserID int64 `json:"UserId"`

	// 信箱，用做逻辑隔离 默认值为 0
	InboxType *int32 `json:"InboxType,omitempty"`
}

type DeleteFriendRes struct {

	// REQUIRED
	ResponseMetadata DeleteFriendResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DeleteFriendResResult `json:"Result"`
}

type DeleteFriendResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                `json:"Version"`
	Error   *DeleteFriendResResponseMetadataError `json:"Error,omitempty"`
}

type DeleteFriendResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type DeleteFriendResResult struct {

	// 删除失败的好友信息
	FailedInfos []DeleteFriendResResultFailedInfosItem `json:"FailedInfos"`
}

type DeleteFriendResResultFailedInfosItem struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`

	// REQUIRED; 好友 UserId
	UserID int64 `json:"UserId"`
}

type DeleteMessageBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 消息 ID
	MessageID int64 `json:"MessageId"`

	// REQUIRED; 群成员 ID
	ParticipantUserID int64 `json:"ParticipantUserId"`
}

type DeleteMessageRes struct {

	// REQUIRED
	ResponseMetadata DeleteMessageResResponseMetadata `json:"ResponseMetadata"`

	// 空。此接口无需关注
	Result interface{} `json:"Result,omitempty"`
}

type DeleteMessageResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                 `json:"Version"`
	Error   *DeleteMessageResResponseMetadataError `json:"Error,omitempty"`
}

type DeleteMessageResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type DestroyConversationBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 群主 UserId
	OwnerUserID int64 `json:"OwnerUserId"`
}

type DestroyConversationRes struct {

	// REQUIRED
	ResponseMetadata DestroyConversationResResponseMetadata `json:"ResponseMetadata"`

	// 空，此接口可忽略此字段。
	Result interface{} `json:"Result,omitempty"`
}

type DestroyConversationResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                       `json:"Version"`
	Error   *DestroyConversationResResponseMetadataError `json:"Error,omitempty"`
}

type DestroyConversationResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type GetBlackListBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 查询起始位置，按照添加顺序逆序查询。默认值为 0，即最后一个添加至黑名单的用户。
	Cursor int64 `json:"Cursor"`

	// REQUIRED; 查询条数，每次最多查询 20 位用户。
	Limit int64 `json:"Limit"`

	// REQUIRED; 被查询用户的 UserId
	UserID int64 `json:"UserId"`

	// 信箱，用做逻辑隔离。默认值为 0
	InboxType *int32 `json:"InboxType,omitempty"`

	// 是否需要黑名单用户总数。 false：不需要。 true：需要。 默认值为 false。
	NeedTotal *bool `json:"NeedTotal,omitempty"`
}

type GetBlackListRes struct {

	// REQUIRED
	ResponseMetadata GetBlackListResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result GetBlackListResResult `json:"Result"`
}

type GetBlackListResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                `json:"Version"`
	Error   *GetBlackListResResponseMetadataError `json:"Error,omitempty"`
}

type GetBlackListResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type GetBlackListResResult struct {

	// REQUIRED; 黑名单用户详细信息
	BlackListInfos []GetBlackListResResultBlackListInfosItem `json:"BlackListInfos"`

	// REQUIRED; 是否还有下一页
	HasMore bool `json:"HasMore"`

	// REQUIRED; 下一页起始下标。为负时表示后续没有成员数据
	NextCursor int64 `json:"NextCursor"`

	// REQUIRED; 黑名单用户总数
	TotalCount int64 `json:"TotalCount"`
}

type GetBlackListResResultBlackListInfosItem struct {

	// REQUIRED; 黑名单用户 ID
	BlackListUserID int64 `json:"BlackListUserId"`

	// REQUIRED; 加入黑名单的时间
	CreateTime int64 `json:"CreateTime"`

	// REQUIRED; 黑名单更新时间
	ModifyTime int64 `json:"ModifyTime"`

	// 黑名单用户扩展字段
	BlackListUserExt map[string]string `json:"BlackListUserExt,omitempty"`

	// 好友信息。若黑名单用户并不是查询用户好友则不返回此字段。
	FriendUserInfo *GetBlackListResResultBlackListInfosItemFriendUserInfo `json:"FriendUserInfo,omitempty"`
}

// GetBlackListResResultBlackListInfosItemFriendUserInfo - 好友信息。若黑名单用户并不是查询用户好友则不返回此字段。
type GetBlackListResResultBlackListInfosItemFriendUserInfo struct {

	// REQUIRED; 好友备注
	Alias string `json:"Alias"`

	// REQUIRED; 发起好友申请时间
	ApplyTime int64 `json:"ApplyTime"`

	// REQUIRED; 成为好友的时间
	CreateTime int64 `json:"CreateTime"`

	// REQUIRED; 好友扩展字段
	Ext interface{} `json:"Ext"`

	// REQUIRED; 好友 UserId
	FriendUserID int64 `json:"FriendUserId"`

	// REQUIRED; 好友更新时间
	ModifyTime int64 `json:"ModifyTime"`
}

type GetConversationMessagesBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 查询起始位置，即查询起始消息的 Index
	Cursor int64 `json:"Cursor"`

	// REQUIRED; 查询条数
	Limit int64 `json:"Limit"`

	// 查询方向。 0：正向查询 1：反向查询 默认值为 0。值不合法时自动调整为默认值。 直播群只能取 1。
	Reverse *int32 `json:"Reverse,omitempty"`
}

type GetConversationMessagesRes struct {

	// REQUIRED
	ResponseMetadata GetConversationMessagesResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result GetConversationMessagesResResult `json:"Result"`
}

type GetConversationMessagesResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                           `json:"Version"`
	Error   *GetConversationMessagesResResponseMetadataError `json:"Error,omitempty"`
}

type GetConversationMessagesResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type GetConversationMessagesResResult struct {

	// REQUIRED; 是否还有下一页
	HasMore bool `json:"HasMore"`

	// REQUIRED
	Messages []GetConversationMessagesResResultMessagesItem `json:"Messages"`

	// REQUIRED; 下一页起始位置
	NewCursor int64 `json:"NewCursor"`
}

type GetConversationMessagesResResultMessagesItem struct {

	// REQUIRED; 应用的唯一标志
	AppID int64 `json:"AppId"`

	// REQUIRED; 消息内容
	Content string `json:"Content"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 会话类型。 1：单聊 2:群聊
	ConversationType int32 `json:"ConversationType"`

	// REQUIRED; 消息创建时间戳，单位为毫秒
	CreateTime int64 `json:"CreateTime"`

	// REQUIRED; 消息 ID
	MessageID int64 `json:"MessageId"`

	// REQUIRED; 消息类型。 10001：文本。 10003：图片。 10004：视频 10005：文件 10006：音频 10012：自定义消息
	MsgType int32 `json:"MsgType"`

	// REQUIRED; 消息发送人 UserId
	Sender int64 `json:"Sender"`

	// REQUIRED; 消息状态，取值为0，表示消息可见。
	Status int32 `json:"Status"`

	// 消息的扩展字段
	Ext map[string]string `json:"Ext,omitempty"`

	// 消息在会话中的位置
	IndexInConversation *int64 `json:"IndexInConversation,omitempty"`

	// 引用消息
	RefMsgInfo *GetConversationMessagesResResultMessagesItemRefMsgInfo `json:"RefMsgInfo,omitempty"`
}

// GetConversationMessagesResResultMessagesItemRefMsgInfo - 引用消息
type GetConversationMessagesResResultMessagesItemRefMsgInfo struct {

	// REQUIRED; 消息引用时展示的文本内容
	Hint string `json:"Hint"`

	// REQUIRED; 被引用的消息 ID
	ReferencedMessageID int64 `json:"ReferencedMessageId"`

	// REQUIRED; 被引用的消息类型 10001：文本。 10003：图片。 10004：视频 10005：文件 10006：音频 10012：自定义消息
	ReferencedMessageType int32 `json:"ReferencedMessageType"`

	// REQUIRED; 被引用的消息状态 0:消息可见 1:消息已过期 2:消息（对用户）不可见 3:消息被撤回 4:消息本身可见，后因删除不可见
	Status int32 `json:"Status"`
}

type GetConversationSettingBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 会话成员 UserId，UserId 必须大于 0。
	ParticipantUserID int64 `json:"ParticipantUserId"`

	// 是否需要该成员在会话中的已读位置。 true：不需要。 false：需要。
	NoReadIndex *bool `json:"NoReadIndex,omitempty"`
}

type GetConversationSettingRes struct {

	// REQUIRED
	ResponseMetadata GetConversationSettingResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result GetConversationSettingResResult `json:"Result"`
}

type GetConversationSettingResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                          `json:"Version"`
	Error   *GetConversationSettingResResponseMetadataError `json:"Error,omitempty"`
}

type GetConversationSettingResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type GetConversationSettingResResult struct {

	// REQUIRED
	ConversationSettingInfo GetConversationSettingResResultConversationSettingInfo `json:"ConversationSettingInfo"`
}

type GetConversationSettingResResultConversationSettingInfo struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 会话类型。 1：单聊 2:群聊 100:直播群
	ConversationType int32 `json:"ConversationType"`

	// REQUIRED; 信箱,用于逻辑隔离
	InboxType int32 `json:"InboxType"`

	// REQUIRED; 群成员UserId
	ParticipantUserID int64 `json:"ParticipantUserId"`

	// REQUIRED; 用户已读位置
	ReadIndex int64 `json:"ReadIndex"`

	// REQUIRED; 置顶时间，单位为毫秒。0表示未置顶
	StickTopTime int64 `json:"StickTopTime"`

	// 扩展字段
	Ext map[string]string `json:"Ext,omitempty"`

	// 是否开启免打扰。 true：开启。 false：不开启
	IsMute *bool `json:"IsMute,omitempty"`

	// 是否收藏。 true：收藏。 false：不收藏
	IsSetFavorite *bool `json:"IsSetFavorite,omitempty"`

	// 是否置顶。 true：置顶。 false：不置顶
	IsStickTop *bool `json:"IsStickTop,omitempty"`

	// 收藏时间，单位为毫秒。 0表示未收藏
	SetFavoriteTime *int64 `json:"SetFavoriteTime,omitempty"`
}

type GetConversationUserCountBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`
}

type GetConversationUserCountRes struct {

	// REQUIRED
	ResponseMetadata GetConversationUserCountResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result GetConversationUserCountResResult `json:"Result"`
}

type GetConversationUserCountResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                            `json:"Version"`
	Error   *GetConversationUserCountResResponseMetadataError `json:"Error,omitempty"`
}

type GetConversationUserCountResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type GetConversationUserCountResResult struct {

	// REQUIRED; 会话成员数量
	Count int64 `json:"Count"`
}

type GetMessagesBody struct {

	// REQUIRED
	AppID int32 `json:"AppId"`

	// REQUIRED
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED
	MessageIDs []int64 `json:"MessageIds"`
}

type GetMessagesRes struct {

	// REQUIRED
	ResponseMetadata GetMessagesResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result GetMessagesResResult `json:"Result"`
}

type GetMessagesResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                               `json:"Version"`
	Error   *GetMessagesResResponseMetadataError `json:"Error,omitempty"`
}

type GetMessagesResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type GetMessagesResResult struct {

	// REQUIRED
	Messages []GetMessagesResResultMessagesItem `json:"Messages"`
}

type GetMessagesResResultMessagesItem struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 消息内容
	Content string `json:"Content"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 会话类型 1：单聊。 2：群聊。
	ConversationType int32 `json:"ConversationType"`

	// REQUIRED; 消息创建时间戳，单位为毫秒
	CreateTime int64 `json:"CreateTime"`

	// REQUIRED; 消息 ID
	MessageID int64 `json:"MessageId"`

	// REQUIRED; 消息类型。 10001：文本。 10003：图片。 10004：视频 10005：文件 10006：音频 10012：自定义消息
	MsgType int32 `json:"MsgType"`

	// REQUIRED; 消息发送人 UserId
	Sender int64 `json:"Sender"`

	// REQUIRED; 消息状态，取值为0，表示消息可见。
	Status int32 `json:"Status"`

	// 消息的扩展字段
	Ext map[string]string `json:"Ext,omitempty"`

	// 消息在会话中的位置
	IndexInConversation *int64 `json:"IndexInConversation,omitempty"`

	// 引用消息
	RefMsgInfo *GetMessagesResResultMessagesItemRefMsgInfo `json:"RefMsgInfo,omitempty"`
}

// GetMessagesResResultMessagesItemRefMsgInfo - 引用消息
type GetMessagesResResultMessagesItemRefMsgInfo struct {

	// REQUIRED; 消息引用时展示的文本内容
	Hint string `json:"Hint"`

	// REQUIRED; 被引用的消息 ID
	ReferencedMessageID int64 `json:"ReferencedMessageId"`

	// REQUIRED; 被引用的消息类型 10001：文本。 10003：图片。 10004：视频 10005：文件 10006：音频 10012：自定义消息
	ReferencedMessageType int32 `json:"ReferencedMessageType"`

	// REQUIRED; 被引用的消息状态 0:消息可见 1:消息已过期 2:消息（对用户）不可见 3:消息被撤回 4:消息本身可见，后因删除不可见
	Status int32 `json:"Status"`
}

type GetParticipantReadIndexBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`
}

type GetParticipantReadIndexRes struct {

	// REQUIRED
	ResponseMetadata GetParticipantReadIndexResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result GetParticipantReadIndexResResult `json:"Result"`
}

type GetParticipantReadIndexResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	Error *GetParticipantReadIndexResResponseMetadataError `json:"Error,omitempty"`
}

type GetParticipantReadIndexResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type GetParticipantReadIndexResResult struct {

	// REQUIRED
	ReadIndexes []GetParticipantReadIndexResResultReadIndexesItem `json:"ReadIndexes"`
}

type GetParticipantReadIndexResResultReadIndexesItem struct {

	// REQUIRED; 会话成员 UserId
	ParticipantUserID int64 `json:"ParticipantUserId"`

	// REQUIRED; 成员已读位置。若没有返回某个成员的 ReadIndex，则表示该成员的 ReadIndex 为 0。
	ReadIndex int64 `json:"ReadIndex"`
}

type GetUserConversationsBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 查询起始位置
	Cursor int64 `json:"Cursor"`

	// REQUIRED; 查询数量。最大值为 20。
	Limit int64 `json:"Limit"`

	// REQUIRED; 查询用户 UserId
	ParticipantUserID int64 `json:"ParticipantUserId"`

	// 数据来源。 0：从缓存中拉取，按会话最近活跃排序。 1:从数据库中拉取，按照创建时间正序排序。 2：拉取用户创建的直播群会话，按照创建时间逆序排序
	DataType *int32 `json:"DataType,omitempty"`

	// 信箱，用于逻辑隔离。 默认值为 0。
	InboxType *int32 `json:"InboxType,omitempty"`

	// 是否忽略会话成员数。 true：忽略。 false：不忽略。 默认值为 false。
	SkipMemberCount *bool `json:"SkipMemberCount,omitempty"`
}

type GetUserConversationsRes struct {

	// REQUIRED
	ResponseMetadata GetUserConversationsResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result GetUserConversationsResResult `json:"Result"`
}

type GetUserConversationsResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                        `json:"Version"`
	Error   *GetUserConversationsResResponseMetadataError `json:"Error,omitempty"`
}

type GetUserConversationsResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type GetUserConversationsResResult struct {

	// REQUIRED
	ConversationInfos []GetUserConversationsResResultConversationInfosItem `json:"ConversationInfos"`

	// REQUIRED; 是否还有更多的数据
	HasMore bool `json:"HasMore"`

	// REQUIRED; 下一批数据起始位置
	NextCursor int64 `json:"NextCursor"`
}

type GetUserConversationsResResultConversationInfosItem struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 会话类型。 1：单聊 2:群聊 100:直播群
	ConversationType int32 `json:"ConversationType"`

	// REQUIRED; 群聊创建时间戳，单位为秒
	CreateTime int64 `json:"CreateTime"`

	// REQUIRED; 创群人 UserId
	CreatorUserID int64 `json:"CreatorUserId"`

	// REQUIRED; 信箱,用于逻辑隔离
	InboxType int32 `json:"InboxType"`

	// REQUIRED; 修改时间戳，单位为秒
	ModifyTime int64 `json:"ModifyTime"`

	// REQUIRED; 群主 UserId
	OwnerUserID int64 `json:"OwnerUserId"`

	// 群头像 url
	AvatarURL *string `json:"AvatarUrl,omitempty"`

	// 群描述
	Description *string `json:"Description,omitempty"`

	// 会话的扩展字段。
	Ext map[string]string `json:"Ext,omitempty"`

	// 会话成员数
	MemberCount *int64 `json:"MemberCount,omitempty"`

	// 群名
	Name *string `json:"Name,omitempty"`

	// 群公告
	Notice *string `json:"Notice,omitempty"`

	// 直播群在线人数。
	OnlineCount *int64 `json:"OnlineCount,omitempty"`

	// 单聊会话另一个UserId
	OtherUserID *int64 `json:"OtherUserId,omitempty"`

	// 会话状态。 0：正常 1：已解散
	Status *int32 `json:"Status,omitempty"`
}

type IsFriendBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 需要校验用户的 UserId，一次最多检验 10 个用户
	FriendUserIDs []int64 `json:"FriendUserIds"`

	// REQUIRED; 用户 UserId
	UserID int64 `json:"UserId"`

	// 信箱，用做逻辑隔离 默认值为 0
	InboxType *int32 `json:"InboxType,omitempty"`
}

type IsFriendRes struct {

	// REQUIRED
	ResponseMetadata IsFriendResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result IsFriendResResult `json:"Result"`
}

type IsFriendResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                            `json:"Version"`
	Error   *IsFriendResResponseMetadataError `json:"Error,omitempty"`
}

type IsFriendResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type IsFriendResResult struct {

	// REQUIRED
	Infos []IsFriendResResultInfosItem `json:"Infos"`
}

type IsFriendResResultInfosItem struct {

	// REQUIRED; 校验好友的 UserId
	FriendUserID int64 `json:"FriendUserId"`

	// REQUIRED; 是否是好友。 true：是。 false：否
	IsFriend bool `json:"IsFriend"`
}

type IsInBlackListBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 需要校验的用户的 UserId，一次最多校验 10 个用户
	BlackListUserIDs []int64 `json:"BlackListUserIds"`

	// REQUIRED; 用户 ID
	UserID int64 `json:"UserId"`

	// 信箱，用做逻辑隔离。默认值为 0
	InboxType *int32 `json:"InboxType,omitempty"`
}

type IsInBlackListRes struct {

	// REQUIRED
	ResponseMetadata IsInBlackListResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result IsInBlackListResResult `json:"Result"`
}

type IsInBlackListResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                 `json:"Version"`
	Error   *IsInBlackListResResponseMetadataError `json:"Error,omitempty"`
}

type IsInBlackListResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type IsInBlackListResResult struct {

	// REQUIRED; 校验结果。 value 为 true：在黑名单中。 value 为 false：不在黑名单中
	IsInBlackListInfos map[string]bool `json:"IsInBlackListInfos"`
}

type IsUserInConversationBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 查询用户 UserId
	ParticipantUserID int64 `json:"ParticipantUserId"`
}

type IsUserInConversationRes struct {

	// REQUIRED
	ResponseMetadata IsUserInConversationResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result IsUserInConversationResResult `json:"Result"`
}

type IsUserInConversationResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                        `json:"Version"`
	Error   *IsUserInConversationResResponseMetadataError `json:"Error,omitempty"`
}

type IsUserInConversationResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type IsUserInConversationResResult struct {

	// REQUIRED; 用户是否在会话中。 true：是。 false：否。
	IsUserInConversation bool `json:"IsUserInConversation"`
}

type ListFriendBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 查询用户 UserId
	UserID int64 `json:"UserId"`

	// 查询起始位置，默认值为 0，即第一个添加的好友。
	Cursor *int64 `json:"Cursor,omitempty"`

	// 信箱，用做逻辑隔离 默认值为 0
	InboxType *int32 `json:"InboxType,omitempty"`

	// 查询条数，每次最多查询 20 位好友。默认值为 20。
	Limit *int64 `json:"Limit,omitempty"`

	// 是否需要好友总数。 false：不需要。 true：需要。 默认值为 false。
	NeedTotal *bool `json:"NeedTotal,omitempty"`
}

type ListFriendRes struct {

	// REQUIRED
	ResponseMetadata ListFriendResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result ListFriendResResult `json:"Result"`
}

type ListFriendResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                              `json:"Version"`
	Error   *ListFriendResResponseMetadataError `json:"Error,omitempty"`
}

type ListFriendResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type ListFriendResResult struct {

	// REQUIRED; 好友详细信息
	FriendInfos []ListFriendResResultFriendInfosItem `json:"FriendInfos"`

	// REQUIRED; 是否还有下一页
	HasMore bool `json:"HasMore"`

	// REQUIRED; 下一页起始位置。为负时表示后续没有成员数据
	NextCursor int64 `json:"NextCursor"`

	// REQUIRED; 好友总数
	TotalCount int64 `json:"TotalCount"`
}

type ListFriendResResultFriendInfosItem struct {

	// REQUIRED; 好友备注
	Alias string `json:"Alias"`

	// REQUIRED; 发起好友申请时间
	ApplyTime int64 `json:"ApplyTime"`

	// REQUIRED; 成为好友的时间
	CreateTime int64 `json:"CreateTime"`

	// REQUIRED; 好友 UserId
	FriendUserID int64 `json:"FriendUserId"`

	// REQUIRED; 好友更新时间
	ModifyTime int64 `json:"ModifyTime"`

	// 好友扩展字段
	Ext map[string]string `json:"Ext,omitempty"`
}

type ModifyConversationBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话信息
	ConversationCoreInfo ModifyConversationBodyConversationCoreInfo `json:"ConversationCoreInfo"`
}

// ModifyConversationBodyConversationCoreInfo - 会话信息
type ModifyConversationBodyConversationCoreInfo struct {

	// REQUIRED; 需要修改的会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// 会话头像 url
	AvatarURL *string `json:"AvatarUrl,omitempty"`

	// 会话描述
	Description *string `json:"Description,omitempty"`

	// 会话的扩展字段
	Ext map[string]string `json:"Ext,omitempty"`

	// 会话名
	Name *string `json:"Name,omitempty"`

	// 会话公告
	Notice *string `json:"Notice,omitempty"`

	// 群主的UserId
	OwnerUserID *int64 `json:"OwnerUserId,omitempty"`
}

type ModifyConversationRes struct {

	// REQUIRED
	ResponseMetadata ModifyConversationResResponseMetadata `json:"ResponseMetadata"`

	// 空，此接口可忽略此字段。
	Result interface{} `json:"Result,omitempty"`
}

type ModifyConversationResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求��公共参数。
	Version string                                      `json:"Version"`
	Error   *ModifyConversationResResponseMetadataError `json:"Error,omitempty"`
}

type ModifyConversationResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type ModifyConversationSettingBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED
	ConversationSettingInfo ModifyConversationSettingBodyConversationSettingInfo `json:"ConversationSettingInfo"`
}

type ModifyConversationSettingBodyConversationSettingInfo struct {

	// REQUIRED; 需要修改的会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 需要修改的会话成员UserId
	ParticipantUserID int64 `json:"ParticipantUserId"`

	// 扩展字段
	Ext map[string]string `json:"Ext,omitempty"`

	// 是否开启免打扰。 true：开启。 false：不开启 默认值为 false。
	IsMute *bool `json:"IsMute,omitempty"`

	// 是否需要修改收藏时间
	IsSetFavorite *bool `json:"IsSetFavorite,omitempty"`

	// 是否需要修改置顶时间
	IsStickTop *bool `json:"IsStickTop,omitempty"`

	// 收藏时间戳，单位为毫秒
	SetFavoriteTime *int64 `json:"SetFavoriteTime,omitempty"`

	// 置顶时间戳，单位为毫秒
	StickTopTime *int64 `json:"StickTopTime,omitempty"`
}

type ModifyConversationSettingRes struct {

	// REQUIRED
	ResponseMetadata ModifyConversationSettingResResponseMetadata `json:"ResponseMetadata"`

	// 空，此接口可忽略此字段。
	Result interface{} `json:"Result,omitempty"`
}

type ModifyConversationSettingResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                             `json:"Version"`
	Error   *ModifyConversationSettingResResponseMetadataError `json:"Error,omitempty"`
}

type ModifyConversationSettingResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type ModifyMessageBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 消息的扩展字段
	Ext map[string]string `json:"Ext"`

	// REQUIRED; 消息 ID
	MessageID int64 `json:"MessageId"`
}

type ModifyMessageRes struct {

	// REQUIRED
	ResponseMetadata ModifyMessageResResponseMetadata `json:"ResponseMetadata"`

	// 空，此接口可忽略此字段。
	Result interface{} `json:"Result,omitempty"`
}

type ModifyMessageResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                 `json:"Version"`
	Error   *ModifyMessageResResponseMetadataError `json:"Error,omitempty"`
}

type ModifyMessageResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type ModifyParticipantReadIndexBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 修改成员所属 UserId，UserId 必须大于 0。
	ParticipantUserID int64 `json:"ParticipantUserId"`

	// REQUIRED; 成员已读位置。传入的ReadIndex必须要大于该会话成员目前的ReadIndex。
	ReadIndex int64 `json:"ReadIndex"`
}

type ModifyParticipantReadIndexRes struct {

	// REQUIRED
	ResponseMetadata ModifyParticipantReadIndexResResponseMetadata `json:"ResponseMetadata"`

	// 空，此接口可忽略此字段。
	Result interface{} `json:"Result,omitempty"`
}

type ModifyParticipantReadIndexResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                              `json:"Version"`
	Error   *ModifyParticipantReadIndexResResponseMetadataError `json:"Error,omitempty"`
}

type ModifyParticipantReadIndexResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type QueryLiveParticipantStatusBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 查询群成员 UserId
	ParticipantUserIDs []int64 `json:"ParticipantUserIds"`
}

type QueryLiveParticipantStatusRes struct {

	// REQUIRED
	ResponseMetadata QueryLiveParticipantStatusResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result QueryLiveParticipantStatusResResult `json:"Result"`
}

type QueryLiveParticipantStatusResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                              `json:"Version"`
	Error   *QueryLiveParticipantStatusResResponseMetadataError `json:"Error,omitempty"`
}

type QueryLiveParticipantStatusResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type QueryLiveParticipantStatusResResult struct {

	// REQUIRED
	Status []QueryLiveParticipantStatusResResultStatusItem `json:"Status"`
}

type QueryLiveParticipantStatusResResultStatusItem struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 该成员所属会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 加群时间，单位为秒
	CreateTime int64 `json:"CreateTime"`

	// REQUIRED; 是否在群中
	IsInConversation bool `json:"IsInConversation"`

	// REQUIRED; 该成员 UserId
	ParticipantUserID int64 `json:"ParticipantUserId"`
}

type QueryOnlineStatusBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 查询用户的 UserId。一次最多查询 100 个成员。
	UserIDs []int64 `json:"UserIds"`
}

type QueryOnlineStatusRes struct {

	// REQUIRED
	ResponseMetadata QueryOnlineStatusResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result QueryOnlineStatusResResult `json:"Result"`
}

type QueryOnlineStatusResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                     `json:"Version"`
	Error   *QueryOnlineStatusResResponseMetadataError `json:"Error,omitempty"`
}

type QueryOnlineStatusResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type QueryOnlineStatusResResult struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 用户在线状态
	QueryOnlineStatusResults map[string][]QueryOnlineStatusResultItems `json:"QueryOnlineResults"`
}

type QueryOnlineStatusResultItems struct {
	ConnID        string `json:"ConnId"`
	UserID        int64  `json:"UserId"`
	DeviceID      int64  `json:"DeviceId"`
	ClientVersion int64  `json:"ClientVersion"`
	Platform      string `json:"Platform"`
}

type RecallMessageBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 消息 ID
	MessageID int64 `json:"MessageId"`

	// REQUIRED; 撤回消息的会话成员 ID
	ParticipantUserID int64 `json:"ParticipantUserId"`

	// 撤回消息的会话成员角色。 0：普通会话成员。 1：群主。 2：管理员。
	RecallRole *int32 `json:"RecallRole,omitempty"`
}

type RecallMessageRes struct {

	// REQUIRED
	ResponseMetadata RecallMessageResResponseMetadata `json:"ResponseMetadata"`

	// 空，此接口可忽略此字段。
	Result *RecallMessageResResult `json:"Result,omitempty"`
}

type RecallMessageResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                 `json:"Version"`
	Error   *RecallMessageResResponseMetadataError `json:"Error,omitempty"`
}

type RecallMessageResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

// RecallMessageResResult - 空，此接口可忽略此字段。
type RecallMessageResResult struct {

	// REQUIRED; 异动列表
	Items []interface{} `json:"Items"`
}

type RemoveBlackListBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 需要移出黑名单的用户的 ID，一次最多移出 10 个用户
	BlackListUserIDs []int64 `json:"BlackListUserIds"`

	// REQUIRED; 用户 ID
	UserID int64 `json:"UserId"`

	// 信箱，用做逻辑隔离。默认值为 0
	InboxType *int32 `json:"InboxType,omitempty"`
}

type RemoveBlackListRes struct {

	// REQUIRED
	ResponseMetadata RemoveBlackListResResponseMetadata `json:"ResponseMetadata"`

	// 空。此接口无需关注
	Result interface{} `json:"Result,omitempty"`
}

type RemoveBlackListResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                   `json:"Version"`
	Error   *RemoveBlackListResResponseMetadataError `json:"Error,omitempty"`
}

type RemoveBlackListResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type ScanConversationParticipantListBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 查询起始位置
	Cursor int64 `json:"Cursor"`

	// REQUIRED; 查询人数
	Limit int64 `json:"Limit"`

	// 是否只查询群成员 UserId。 true：是。 false：否。 默认值为 false。
	OnlyUserIDs *bool `json:"OnlyUserIds,omitempty"`

	// 按照进群的时间正序或逆序查询。 false: 正序。 true：逆序。 默认值为：fasle。
	Reverse *bool `json:"Reverse,omitempty"`

	// 直播群专用字段。是否需要获得直播群在线成员列表。 true：是。 false：否。 默认值为 false。 当此字段设置 true 时，Reverse 和 OnlyUserIds 字段均会失效，该接口按成员进入直播群时间逆序返回在线成员列表。
	ScanOnlineParticipant *bool `json:"ScanOnlineParticipant,omitempty"`
}

type ScanConversationParticipantListRes struct {

	// REQUIRED
	ResponseMetadata ScanConversationParticipantListResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result ScanConversationParticipantListResResult `json:"Result"`
}

type ScanConversationParticipantListResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                                   `json:"Version"`
	Error   *ScanConversationParticipantListResResponseMetadataError `json:"Error,omitempty"`
}

type ScanConversationParticipantListResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type ScanConversationParticipantListResResult struct {

	// REQUIRED; 是否还有下一页
	HasMore bool `json:"HasMore"`

	// REQUIRED; 下一页起始位置。为负时表示后续没有成员数据
	NextCursor int32 `json:"NextCursor"`

	// REQUIRED
	Participants []ScanConversationParticipantListResResultParticipantsItem `json:"Participants"`
}

type ScanConversationParticipantListResResultParticipantsItem struct {

	// REQUIRED; 直播群群成员头像
	AvatarURL string `json:"AvatarUrl"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 操作人对应的 UserId
	Operator int64 `json:"Operator"`

	// REQUIRED; 群成员 ID
	ParticipantUserID int64 `json:"ParticipantUserId"`

	// REQUIRED; 成员身份。 0: 普通成员。 1: 群主。 2：群管理员。
	Role int64 `json:"Role"`

	// 禁言时间戳，单位为秒。0表示不禁言
	BlockTime *int64 `json:"BlockTime,omitempty"`

	// 成员进群时间对应时间戳，单位为秒
	CreateTime *int64 `json:"CreateTime,omitempty"`

	// 成员的扩展字段
	Ext map[string]string `json:"Ext,omitempty"`

	// 成员等级
	Level *int32 `json:"Level,omitempty"`

	// 成员昵称
	NickName *string `json:"NickName,omitempty"`

	// 成员状态。 0：正常 1：退出
	Status *int32 `json:"Status,omitempty"`
}

type SendMessageBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 消息内容。当你给客户端发消息时，Content 内容需符合客户端格式，详细信息请参看消息格式。
	Content string `json:"Content"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 消息类型
	MsgType int32 `json:"MsgType"`

	// REQUIRED; 消息发送人 UserId
	Sender int64 `json:"Sender"`

	// 幂等id，如果创建时指定了此字段，并且数据库中存在此 id 对应的消息，不会重复发送。如果不指定，会随机生成一个。
	ClientMsgID *string `json:"ClientMsgId,omitempty"`

	// 消息对应时间戳，单位为毫秒。
	CreateTime *int64 `json:"CreateTime,omitempty"`

	// 消息的扩展字段
	Ext map[string]string `json:"Ext,omitempty"`

	// 消息不可见会话成员列表。VisibleUsers 和 InvisibleUsers均为空时，代表对所有人可见。
	InvisibleUsers []*int64 `json:"InvisibleUsers,omitempty"`

	// 会话中@的人
	MentionedUsers []*int64 `json:"MentionedUsers,omitempty"`

	// 消息优先级。 0：低优先级。 1:普通优先级。 2:高优先级。 该字段仅对直播群有效。为避免直播群中消息频率太多导致服务端压力过大，你可以设置消息的优先级。当前直播群下，普通优先级和低优先级消息共用频控阈值为 30 条/秒，超过部分会被服务端直接丢弃。高优消息频控阈值为
	// 10 条/秒，超过部分服务端无法保证不丢失。
	Priority *int32 `json:"Priority,omitempty"`

	// 引用消息
	RefMsgInfo *SendMessageBodyRefMsgInfo `json:"RefMsgInfo,omitempty"`

	// 消息可见会话成员列表
	VisibleUsers []*int64 `json:"VisibleUsers,omitempty"`
}

// SendMessageBodyRefMsgInfo - 引用消息
type SendMessageBodyRefMsgInfo struct {

	// REQUIRED; 消息引用时展示的文本内容
	Hint string `json:"Hint"`

	// REQUIRED; 被引用的消息 ID
	ReferencedMessageID int64 `json:"ReferencedMessageId"`
}

type SendMessageRes struct {

	// REQUIRED
	ResponseMetadata SendMessageResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result SendMessageResResult `json:"Result"`
}

type SendMessageResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                               `json:"Version"`
	Error   *SendMessageResResponseMetadataError `json:"Error,omitempty"`
}

type SendMessageResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type SendMessageResResult struct {

	// REQUIRED; 消息Id
	MessageID int64 `json:"MessageId"`
}

type UpdateBlackListBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 黑名单用户 ID 及需要更新的扩展字段。一次最多更新 10 个黑名单用户的扩展字段。
	BlackListInfos []UpdateBlackListBodyBlackListInfosItem `json:"BlackListInfos"`

	// REQUIRED; 用户 ID
	UserID int64 `json:"UserId"`

	// 信箱，用做逻辑隔离。默认值为 0
	InboxType *int32 `json:"InboxType,omitempty"`
}

type UpdateBlackListBodyBlackListInfosItem struct {

	// REQUIRED; 黑名单用户 ID
	BlackListUserID int64 `json:"BlackListUserId"`

	// 黑名单用户的扩展字段，支持新增 key 和更新已有 key。更新已有 key 时，新 value 会覆盖旧 value。
	BlackListUserExt map[string]string `json:"BlackListUserExt,omitempty"`
}

type UpdateBlackListRes struct {

	// REQUIRED
	ResponseMetadata UpdateBlackListResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result UpdateBlackListResResult `json:"Result"`
}

type UpdateBlackListResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                   `json:"Version"`
	Error   *UpdateBlackListResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateBlackListResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type UpdateBlackListResResult struct {

	// REQUIRED; 更新失败的黑名单用户信息
	FailedInfos []UpdateBlackListResResultFailedInfosItem `json:"FailedInfos"`
}

type UpdateBlackListResResultFailedInfosItem struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`

	// REQUIRED; 更新失败的黑名单用户的 ID
	UserID int64 `json:"UserId"`
}

type UpdateFriendBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 好友信息
	FriendInfos []UpdateFriendBodyFriendInfosItem `json:"FriendInfos"`

	// REQUIRED; 用户 UserId
	UserID int64 `json:"UserId"`
}

type UpdateFriendBodyFriendInfosItem struct {

	// REQUIRED; 好友扩展字段。Ext 与 Alias 为非必填字段，但是至少需要填一个，否则服务端会报错。
	Ext map[string]string `json:"Ext"`

	// REQUIRED; 好友 UserId
	FriendUserID int64 `json:"FriendUserId"`

	// 好友备注
	Alias *string `json:"Alias,omitempty"`
}

type UpdateFriendRes struct {

	// REQUIRED
	ResponseMetadata UpdateFriendResResponseMetadata `json:"ResponseMetadata"`
	Result           *UpdateFriendResResult          `json:"Result,omitempty"`
}

type UpdateFriendResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                `json:"Version"`
	Error   *UpdateFriendResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateFriendResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type UpdateFriendResResult struct {
	FailedInfos []UpdateFriendResResultFailedInfos `json:"FailedInfos"`
}

type UpdateFriendResResultFailedInfos struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`

	// REQUIRED
	UserID int64 `json:"UserId"`
}

type DestroyConversationQuery struct{}
type GetConversationMessages struct{}
type GetBlackListQuery struct{}
type BatchAddManagerQuery struct{}
type RecallMessageQuery struct{}
type BatchModifyConversationParticipantQuery struct{}
type ModifyMessageQuery struct{}
type BatchGetConversationsQuery struct{}
type DeleteConversationMessage struct{}
type IsInBlackList struct{}
type ModifyConversation struct{}
type BatchDeleteConversationParticipantQuery struct{}
type BatchGetBlockParticipantsQuery struct{}
type BatchDeleteConversationParticipant struct{}
type DeleteMessage struct{}
type UpdateBlackListQuery struct{}
type ModifyConversationSettingQuery struct{}
type BatchModifyConversationParticipant struct{}
type ModifyMessage struct{}
type CreateConversationQuery struct{}
type CreateConversation struct{}
type ScanConversationParticipantListQuery struct{}
type SendMessageQuery struct{}
type QueryLiveParticipantStatusQuery struct{}
type GetConversationUserCount struct{}
type SendMessage struct{}
type BatchAddConversationParticipant struct{}
type BatchRemoveWhitelistParticipant struct{}
type GetConversationSettingQuery struct{}
type BatchAddBlockParticipants struct{}
type GetUserConversations struct{}
type RecallMessage struct{}
type RemoveBlackList struct{}
type GetConversationMessagesQuery struct{}
type IsUserInConversation struct{}
type BatchGetConversationParticipantQuery struct{}
type GetUserConversationsQuery struct{}
type GetBlackList struct{}
type BatchRemoveManagerQuery struct{}
type QueryOnlineStatus struct{}
type RemoveBlackListQuery struct{}
type BatchAddWhitelistParticipant struct{}
type BatchAddManager struct{}
type DestroyConversation struct{}
type DeleteFriend struct{}
type BatchGetConversations struct{}
type ListFriendQuery struct{}
type BatchDeleteBlockParticipantsQuery struct{}
type IsFriendQuery struct{}
type BatchGetConversationParticipant struct{}
type BatchGetWhitelistParticipant struct{}
type GetConversationUserCountQuery struct{}
type ModifyParticipantReadIndexQuery struct{}
type BatchAddBlockParticipantsQuery struct{}
type QueryOnlineStatusQuery struct{}
type AddFriend struct{}
type DeleteMessageQuery struct{}
type IsUserInConversationQuery struct{}
type GetParticipantReadIndex struct{}
type BatchGetBlockParticipants struct{}
type ListFriend struct{}
type BatchRemoveWhitelistParticipantQuery struct{}
type BatchRemoveManager struct{}
type GetParticipantReadIndexQuery struct{}
type DeleteFriendQuery struct{}
type ModifyConversationQuery struct{}
type UpdateBlackList struct{}
type ScanConversationParticipantList struct{}
type IsFriend struct{}
type AddBlackList struct{}
type BatchAddConversationParticipantQuery struct{}
type IsInBlackListQuery struct{}
type BatchDeleteBlockParticipants struct{}
type DeleteConversationMessageQuery struct{}
type ModifyParticipantReadIndex struct{}
type BatchGetWhitelistParticipantQuery struct{}
type BatchAddWhitelistParticipantQuery struct{}
type QueryLiveParticipantStatus struct{}
type GetConversationSetting struct{}
type ModifyConversationSetting struct{}
type GetMessagesQuery struct{}
type AddFriendQuery struct{}
type GetMessages struct{}
type AddBlackListQuery struct{}
type IsFriendReq struct {
	*IsFriendQuery
	*IsFriendBody
}
type ScanConversationParticipantListReq struct {
	*ScanConversationParticipantListQuery
	*ScanConversationParticipantListBody
}
type ModifyConversationSettingReq struct {
	*ModifyConversationSettingQuery
	*ModifyConversationSettingBody
}
type BatchGetConversationsReq struct {
	*BatchGetConversationsQuery
	*BatchGetConversationsBody
}
type DeleteFriendReq struct {
	*DeleteFriendQuery
	*DeleteFriendBody
}
type GetConversationSettingReq struct {
	*GetConversationSettingQuery
	*GetConversationSettingBody
}
type GetConversationMessagesReq struct {
	*GetConversationMessagesQuery
	*GetConversationMessagesBody
}
type ListFriendReq struct {
	*ListFriendQuery
	*ListFriendBody
}
type IsInBlackListReq struct {
	*IsInBlackListQuery
	*IsInBlackListBody
}
type IsUserInConversationReq struct {
	*IsUserInConversationQuery
	*IsUserInConversationBody
}
type BatchModifyConversationParticipantReq struct {
	*BatchModifyConversationParticipantQuery
	*BatchModifyConversationParticipantBody
}
type BatchAddManagerReq struct {
	*BatchAddManagerQuery
	*BatchAddManagerBody
}
type GetParticipantReadIndexReq struct {
	*GetParticipantReadIndexQuery
	*GetParticipantReadIndexBody
}
type BatchGetBlockParticipantsReq struct {
	*BatchGetBlockParticipantsQuery
	*BatchGetBlockParticipantsBody
}
type CreateConversationReq struct {
	*CreateConversationQuery
	*CreateConversationBody
}
type UpdateBlackListReq struct {
	*UpdateBlackListQuery
	*UpdateBlackListBody
}
type GetBlackListReq struct {
	*GetBlackListQuery
	*GetBlackListBody
}
type BatchAddConversationParticipantReq struct {
	*BatchAddConversationParticipantQuery
	*BatchAddConversationParticipantBody
}
type BatchAddBlockParticipantsReq struct {
	*BatchAddBlockParticipantsQuery
	*BatchAddBlockParticipantsBody
}
type BatchRemoveManagerReq struct {
	*BatchRemoveManagerQuery
	*BatchRemoveManagerBody
}
type QueryOnlineStatusReq struct {
	*QueryOnlineStatusQuery
	*QueryOnlineStatusBody
}
type BatchDeleteBlockParticipantsReq struct {
	*BatchDeleteBlockParticipantsQuery
	*BatchDeleteBlockParticipantsBody
}
type BatchRemoveWhitelistParticipantReq struct {
	*BatchRemoveWhitelistParticipantQuery
	*BatchRemoveWhitelistParticipantBody
}
type ModifyConversationReq struct {
	*ModifyConversationQuery
	*ModifyConversationBody
}
type DeleteMessageReq struct {
	*DeleteMessageQuery
	*DeleteMessageBody
}
type DeleteConversationMessageReq struct {
	*DeleteConversationMessageQuery
	*DeleteConversationMessageBody
}
type RecallMessageReq struct {
	*RecallMessageQuery
	*RecallMessageBody
}
type AddBlackListReq struct {
	*AddBlackListQuery
	*AddBlackListBody
}
type RemoveBlackListReq struct {
	*RemoveBlackListQuery
	*RemoveBlackListBody
}
type BatchDeleteConversationParticipantReq struct {
	*BatchDeleteConversationParticipantQuery
	*BatchDeleteConversationParticipantBody
}
type BatchAddWhitelistParticipantReq struct {
	*BatchAddWhitelistParticipantQuery
	*BatchAddWhitelistParticipantBody
}
type GetConversationUserCountReq struct {
	*GetConversationUserCountQuery
	*GetConversationUserCountBody
}
type DestroyConversationReq struct {
	*DestroyConversationQuery
	*DestroyConversationBody
}
type BatchGetConversationParticipantReq struct {
	*BatchGetConversationParticipantQuery
	*BatchGetConversationParticipantBody
}
type ModifyMessageReq struct {
	*ModifyMessageQuery
	*ModifyMessageBody
}
type GetMessagesReq struct {
	*GetMessagesQuery
	*GetMessagesBody
}
type AddFriendReq struct {
	*AddFriendQuery
	*AddFriendBody
}
type SendMessageReq struct {
	*SendMessageQuery
	*SendMessageBody
}
type ModifyParticipantReadIndexReq struct {
	*ModifyParticipantReadIndexQuery
	*ModifyParticipantReadIndexBody
}
type BatchGetWhitelistParticipantReq struct {
	*BatchGetWhitelistParticipantQuery
	*BatchGetWhitelistParticipantBody
}
type QueryLiveParticipantStatusReq struct {
	*QueryLiveParticipantStatusQuery
	*QueryLiveParticipantStatusBody
}
type GetUserConversationsReq struct {
	*GetUserConversationsQuery
	*GetUserConversationsBody
}
