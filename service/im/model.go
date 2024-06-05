package im

type AddBlackListBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 需要加入黑名单的 UserId 及每个用户对应的扩展信息，一次最多添加 10 个用户至黑名单
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

	// REQUIRED; 添加黑名单失败的用户信息
	FailedInfos []AddBlackListResResultFailedInfosItem `json:"FailedInfos"`
}

type AddBlackListResResultFailedInfosItem struct {

	// REQUIRED; 错误码 [https://www.volcengine.com/docs/6348/412253]
	Code string `json:"Code"`

	// REQUIRED; 错误信息 [https://www.volcengine.com/docs/6348/412253]
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

	// 信箱，用做逻辑隔离。默认值为 0
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

	// REQUIRED; 错误码 [https://www.volcengine.com/docs/6348/412253]
	Code string `json:"Code"`

	// REQUIRED; 错误信息 [https://www.volcengine.com/docs/6348/412253]
	Message string `json:"Message"`

	// REQUIRED; 好友 UserId
	UserID int64 `json:"UserId"`
}

type BatchAddBlockParticipantsBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; key 为群成员 ID，类型：String，只能传userId，value 为禁言或者拉黑时长，类型：int，单位为秒。
	ParticipantBlockInfos map[string]int64 `json:"ParticipantBlockInfos"`

	// 操作行为
	// * 0：禁言。用户无法在直播群中发言。
	// * 1：拉黑。用户无法加入直播群。
	// 默认值为 0
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

	// 是否开启屏障。如设置屏障，新加入用户无法看到历史会话消息。
	// * false：不开启。
	// * true：开启。
	// 默认值为false。
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

	// 成员身份，可取值为：0，1，2。
	// * 0：普通成员。
	// * 1：群主。添加群主时，需确保会话中的 ownerUid 与群主的 UserId 相同。
	// * 2：群管理员。
	// 默认值为0，值不合法时自动调整为默认值。
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

	// REQUIRED; 错误信息
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

	// REQUIRED; 要添加为管理员的群成员 UserId
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

	// 操作行为。
	// * 0：取消禁言。
	// * 1：取消拉黑。
	// 默认值为 0。
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

	// 操作行为
	// * 0：获取禁言列表
	// * 1：获取拉黑列表
	// 默认为 0。
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

	// REQUIRED; 下一批数据起始位置
	NextCursor int64 `json:"NextCursor"`

	// REQUIRED; 禁言/拉黑成员列表
	Participants []BatchGetBlockParticipantsResResultParticipantsItem `json:"Participants"`
}

type BatchGetBlockParticipantsResResultParticipantsItem struct {

	// REQUIRED; 禁言/拉黑成员头像
	AvatarURL string `json:"AvatarUrl"`

	// REQUIRED; 禁言/拉黑到何时对应时间戳，单位为秒
	BlockTime int64 `json:"BlockTime"`

	// REQUIRED; 禁言/拉黑设置对应的时间戳，单位为秒
	CreateTime int64 `json:"CreateTime"`

	// REQUIRED; 禁言/拉黑成员扩展字段。key 的类型为 string，value 的类型为 string。
	Ext map[string]string `json:"Ext"`

	// REQUIRED; 直播群群成员标记
	Marks []string `json:"Marks"`

	// REQUIRED; 禁言/拉黑成员昵称
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

	// REQUIRED; 成功时返回查询会话成员信息，失败时为空。
	Participants []BatchGetConversationParticipantResResultParticipantsItem `json:"Participants"`
}

type BatchGetConversationParticipantResResultParticipantsItem struct {

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 操作人对应的 UserId
	Operator int64 `json:"Operator"`

	// REQUIRED; 群成员 ID
	ParticipantUserID int64 `json:"ParticipantUserId"`

	// REQUIRED; 成员身份。
	// * 0：普通成员。
	// * 1：群主。
	// * 2：群管理员。
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

	// 成员状态。
	// * 0：正常
	// * 1：退出
	Status *int32 `json:"Status,omitempty"`
}

type BatchGetConversationsBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID []int64 `json:"ConversationShortId"`

	// 是否忽略获取会话成员数。
	// * true： 忽略。
	// * false：不忽略。
	// 默认值为 true。
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

	// REQUIRED; 会话详细信息
	ConversationCoreInfos []BatchGetConversationsResResultConversationCoreInfosItem `json:"ConversationCoreInfos"`
}

type BatchGetConversationsResResultConversationCoreInfosItem struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 会话Id，字符串类型，含义跟ConversationShortId一样，用来定位唯一的一个会话，历史原因，目前大部分接口都在使用ConversationShortId，但是仍然有比较比较老的接口会使用到ConversationId，如果目前您接入的功能没有使用到ConversationId直接忽略即可
	ConversationID string `json:"ConversationId"`

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

	// 单聊会话另一个 UserId
	OtherUserID *int64 `json:"OtherUserId,omitempty"`

	// 会话状态。 0：正常 1：已解散
	Status *int32 `json:"Status,omitempty"`
}

type BatchGetUserBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 查询用户 UserId，一次最多查询50个用户
	UserIDs []int64 `json:"UserIds"`
}

type BatchGetUserRes struct {

	// REQUIRED
	ResponseMetadata BatchGetUserResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result BatchGetUserResResult `json:"Result"`
}

type BatchGetUserResResponseMetadata struct {

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
	Error   *BatchGetUserResResponseMetadataError `json:"Error,omitempty"`
}

type BatchGetUserResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type BatchGetUserResResult struct {

	// REQUIRED; 未查到的用户uid，原因：未注册或者已注销
	NotFoundUsers []int64 `json:"NotFoundUsers"`

	// REQUIRED; 用户详细信息
	Users []BatchGetUserResResultUsersItem `json:"Users"`
}

type BatchGetUserResResultUsersItem struct {

	// REQUIRED; 用户所属应用
	AppID int32 `json:"AppId"`

	// REQUIRED; 用户注册时间戳，单位为毫秒
	CreateTime int64 `json:"CreateTime"`

	// REQUIRED; map[string]string
	Ext map[string]string `json:"Ext"`

	// REQUIRED; 更新时间戳，单位为毫秒
	ModifyTime int64 `json:"ModifyTime"`

	// REQUIRED; 用户昵称
	NickName string `json:"NickName"`

	// REQUIRED; 用户头像
	Portrait string `json:"Portrait"`

	// REQUIRED; 用户 UserId
	UserID int64 `json:"UserId"`

	// []string
	Tags []string `json:"Tags,omitempty"`
}

type BatchGetWhitelistParticipantBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 查询起始位置
	Cursor int64 `json:"Cursor"`

	// REQUIRED; 每批查询条数。最大值为20。
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

	// REQUIRED; 白名单成员信息
	Participants []BatchGetWhitelistParticipantResResultParticipantsItem `json:"Participants"`
}

type BatchGetWhitelistParticipantResResultParticipantsItem struct {

	// REQUIRED; 白名单成员头像
	AvatarURL string `json:"AvatarUrl"`

	// REQUIRED; 用户加入白名单时间，单位为秒
	CreateTime int64 `json:"CreateTime"`

	// REQUIRED; 白名单成员扩展字段。key 的类型为 string，value 的类型为 string。
	Ext map[string]string `json:"Ext"`

	// REQUIRED; 白名单成员昵称
	NickName string `json:"NickName"`

	// REQUIRED; 执行加入白名单操作操作人 UserId
	Operator int64 `json:"Operator"`

	// REQUIRED; 白名单成员 UserId
	ParticipantUserID int64 `json:"ParticipantUserId"`

	// 直播群群成员标记
	Marks []string `json:"Marks,omitempty"`
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

	// 禁言时间戳，表示禁言到何时，单位为秒。0表示未禁言。
	BlockTime *int64 `json:"BlockTime,omitempty"`

	// 成员扩展字段
	Ext map[string]string `json:"Ext,omitempty"`

	// 成员等级
	Level *int32 `json:"Level,omitempty"`

	// 成员昵称
	NickName *string `json:"NickName,omitempty"`

	// 成员身份，可取值为：0，1，2。
	// * 0：普通成员。
	// * 1：群主。
	// * 2：群管理员。
	// 默认值为0，值不合法时自动调整为默认值。
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

	// REQUIRED; 需要更新的成员资料。一次最多支持 10 个成员。
	ParticipantInfos []BatchUpdateLiveParticipantsBodyParticipantInfosItem `json:"ParticipantInfos"`

	// 更新群成员标记。默认值为 1。1：添加2：删除
	MarkAction *int32 `json:"MarkAction,omitempty"`
}

type BatchUpdateLiveParticipantsBodyParticipantInfosItem struct {

	// REQUIRED; 群成员用户 ID
	UserID int64 `json:"UserId"`

	// 群成员头像。AvatarUrl、NickName和Ext均为非必填参数，但是至少需要填一个，否则服务端会报错。
	AvatarURL *string `json:"AvatarUrl,omitempty"`

	// 群成员扩展字段。key 的类型为 String，value 的类型为 String。
	Ext map[string]string `json:"Ext,omitempty"`

	// 群成员标记
	Marks []string `json:"Marks,omitempty"`

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

	// 更新资料失败的群成员信息
	FailedInfos []*BatchUpdateLiveParticipantsResResultFailedInfosItem `json:"FailedInfos,omitempty"`
}

type BatchUpdateLiveParticipantsResResultFailedInfosItem struct {

	// REQUIRED; 错误码 [https://www.volcengine.com/docs/6348/412253]
	Code string `json:"Code"`

	// REQUIRED; 错误描述
	Message string `json:"Message"`

	// REQUIRED; 更新资料失败的群成员的 UserId
	UserID int64 `json:"UserId"`
}

type BatchUpdateUserBody struct {

	// REQUIRED; 用户的唯一标识
	AppID int32 `json:"AppId"`

	// REQUIRED; 用户资料。一次最多 10 个用户
	Users []BatchUpdateUserBodyUsersItem `json:"Users"`
}

type BatchUpdateUserBodyUsersItem struct {

	// REQUIRED; 扩展字段。key 的类型为 string，value 的类型为 string。
	Ext map[string]string `json:"Ext"`

	// REQUIRED; 昵称
	NickName string `json:"NickName"`

	// REQUIRED; 头像 url
	Portrait string `json:"Portrait"`

	// REQUIRED; 标签
	Tags []string `json:"Tags"`

	// REQUIRED; 用户id
	UserID int64 `json:"UserId"`
}

type BatchUpdateUserRes struct {

	// REQUIRED
	ResponseMetadata BatchUpdateUserResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *BatchUpdateUserResResult `json:"Result,omitempty"`
}

type BatchUpdateUserResResponseMetadata struct {

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
	Error   *BatchUpdateUserResResponseMetadataError `json:"Error,omitempty"`
}

type BatchUpdateUserResResponseMetadataError struct {

	// REQUIRED; 错误码 [https://www.volcengine.com/docs/6348/412253]
	Code string `json:"Code"`

	// REQUIRED; 错误信息 [https://www.volcengine.com/docs/6348/412253]
	Message string `json:"Message"`
}

// BatchUpdateUserResResult - 视请求的接口而定
type BatchUpdateUserResResult struct {

	// REQUIRED; 更新失败的用户信息
	FailedInfos []BatchUpdateUserResResultFailedInfosItem `json:"FailedInfos"`
}

type BatchUpdateUserResResultFailedInfosItem struct {

	// REQUIRED; 错误码。参考错误码 [https://www.volcengine.com/docs/6348/412253]
	Code string `json:"Code"`

	// REQUIRED; 错误信息。
	Message string `json:"Message"`

	// REQUIRED; 更新失败的用户 UserId
	UserID int64 `json:"UserId"`
}

type BatchUpdateUserTagsBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 标签更新方式。 0：对于重复 key，覆盖 value，对于新 key，进行添加 1：删除重复 key 默认值为 0。
	Op int32 `json:"Op"`

	// REQUIRED; 用户标签
	Tags []string `json:"Tags"`

	// REQUIRED; 用户userId列表
	UserIDs []int64 `json:"UserIds"`
}

type BatchUpdateUserTagsRes struct {

	// REQUIRED
	ResponseMetadata BatchUpdateUserTagsResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result BatchUpdateUserTagsResResult `json:"Result"`
}

type BatchUpdateUserTagsResResponseMetadata struct {

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
	Error   *BatchUpdateUserTagsResResponseMetadataError `json:"Error,omitempty"`
}

type BatchUpdateUserTagsResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type BatchUpdateUserTagsResResult struct {

	// REQUIRED; 更新失败的用户信息
	FailedInfos []BatchUpdateUserTagsResResultFailedInfosItem `json:"FailedInfos"`
}

type BatchUpdateUserTagsResResultFailedInfosItem struct {

	// REQUIRED; 错误码。参考:错误码 [https://www.volcengine.com/docs/6348/412253]
	Code string `json:"Code"`

	// REQUIRED; 错误信息。
	Message string `json:"Message"`

	// REQUIRED; 用户 UserId
	UserID int64 `json:"UserId"`
}

type CreateConversationBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话详细信息
	ConversationCoreInfo CreateConversationBodyConversationCoreInfo `json:"ConversationCoreInfo"`

	// REQUIRED; 会话成员 UserId
	OwnerUserID int64 `json:"OwnerUserId"`

	// 幂等id，如果创建时指定了此字段，并且数据库中存在此 id 对应的会话，则不会重复创建，并且接口返回的Exist字段为true。
	IdempotentID *string `json:"IdempotentId,omitempty"`

	// 信箱，用做逻辑隔离 默认值为 0
	InboxType *int32 `json:"InboxType,omitempty"`

	// 另一个成员的 UserId， 创建单聊必填
	OtherUserID *int64 `json:"OtherUserId,omitempty"`
}

// CreateConversationBodyConversationCoreInfo - 会话详细信息
type CreateConversationBodyConversationCoreInfo struct {

	// REQUIRED; 会话类型
	// * 1：单聊
	// * 2：群聊
	// * 100：直播群
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

	// REQUIRED; 会话详细信息
	ConversationInfo CreateConversationResResultConversationInfo `json:"ConversationInfo"`

	// REQUIRED; 会话Id，字符串类型，含义跟ConversationShortId一样，用来定位唯一的一个会话，历史原因，目前大部分接口都在使用ConversationShortId，但是仍然有比较比较老的接口会使用到ConversationId，如果目前您接入的功能没有使用到ConversationId直接忽略即可
	ConversationID string `json:"ConversationId"`

	// REQUIRED; 会话id
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 会话是否存在
	Exist bool `json:"Exist"`
}

// CreateConversationResResultConversationInfo - 会话详细信息
type CreateConversationResResultConversationInfo struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 会话Id，字符串类型，含义跟ConversationShortId一样，用来定位唯一的一个会话，历史原因，目前大部分接口都在使用ConversationShortId，但是仍然有比较比较老的接口会使用到ConversationId，如果目前您接入的功能没有使用到ConversationId直接忽略即可
	ConversationID string `json:"ConversationId"`

	// REQUIRED; 会话类型。
	// * 1：单聊
	// * 2：群聊
	ConversationType int32 `json:"ConversationType"`

	// REQUIRED; 信箱，用于逻辑隔离
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

	// 单聊另一个成员的 UserId
	OtherUserID *int64 `json:"OtherUserId,omitempty"`

	// 群主 UserId
	OwnerUserID *int64 `json:"OwnerUserId,omitempty"`

	// 会话状态。
	// * 0：正常
	// * 1：已解散
	Status *int32 `json:"Status,omitempty"`
}

type DeleteConversationMessageBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 消息 ID
	MessageID int64 `json:"MessageId"`

	// 删除方式。
	// * 0：全部用户不可见。
	// * 1：仅发送人自己可见。
	// 默认值为0。 直播群只允许传 0。
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

	// 信箱，用做逻辑隔离。默认值为 0
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
	FailedInfos []DeleteFriendResResultFailedInfosItem `json:"FailedInfos,omitempty"`
}

type DeleteFriendResResultFailedInfosItem struct {

	// REQUIRED; 错误码 [https://www.volcengine.com/docs/6348/412253]
	Code string `json:"Code"`

	// REQUIRED; 错误信息 [https://www.volcengine.com/docs/6348/412253]
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

type GetAppTokenBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; Token过期时间 单位毫秒
	ExpireTime int64 `json:"ExpireTime"`

	// REQUIRED; 用户UserId
	UserID int64 `json:"UserId"`
}

type GetAppTokenRes struct {

	// REQUIRED
	ResponseMetadata GetAppTokenResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetAppTokenResResult `json:"Result,omitempty"`
}

type GetAppTokenResResponseMetadata struct {

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
	Error   *GetAppTokenResResponseMetadataError `json:"Error,omitempty"`
}

type GetAppTokenResResponseMetadataError struct {

	// REQUIRED; 错误码 [https://www.volcengine.com/docs/6348/412253]
	Code string `json:"Code"`

	// REQUIRED; 错误信息 [https://www.volcengine.com/docs/6348/412253]
	Message string `json:"Message"`
}

// GetAppTokenResResult - 视请求的接口而定
type GetAppTokenResResult struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 生成的Token
	Token string `json:"Token"`

	// REQUIRED; 用户UserId
	UserID int64 `json:"UserId"`
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

	// 是否需要黑名单用户总数。
	// * false：不需要。
	// * true：需要。
	// 默认值为 false。
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
	Ext map[string]string `json:"Ext"`

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

	// 查询方向。
	// * 0：正向查询
	// * 1：反向查询
	// 默认值为 0。值不合法时自动调整为默认值。 直播群只能取 1。
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

	// REQUIRED; 消息详细信息
	Messages []GetConversationMessagesResResultMessagesItem `json:"Messages"`

	// REQUIRED; 下一页起始位置
	NewCursor int64 `json:"NewCursor"`
}

type GetConversationMessagesResResultMessagesItem struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 消息内容
	Content string `json:"Content"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 会话类型
	// * 1：单聊。
	// * 2：群聊。
	ConversationType int32 `json:"ConversationType"`

	// REQUIRED; 消息创建时间戳，单位为毫秒
	CreateTime int64 `json:"CreateTime"`

	// REQUIRED; 消息 ID
	MessageID int64 `json:"MessageId"`

	// REQUIRED; 消息类型。
	// * 10001：文本。
	// * 10003：图片。
	// * 10004：视频
	// * 10005：文件
	// * 10006：音频
	// * 10012：自定义消息
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

	// REQUIRED; 被引用的消息类型
	// * 10001：文本。
	// * 10003：图片。
	// * 10004：视频
	// * 10005：文件
	// * 10006：音频
	// * 10012：自定义消息
	ReferencedMessageType int32 `json:"ReferencedMessageType"`

	// REQUIRED; 被引用的消息状态
	// * 0：消息可见
	// * 1：消息已过期
	// * 2：消息（对用户）不可见
	// * 3：消息被撤回
	// * 4：消息本身可见，后因删除不可见
	Status int32 `json:"Status"`
}

type GetConversationMarksBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`
}

type GetConversationMarksRes struct {

	// REQUIRED
	ResponseMetadata GetConversationMarksResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetConversationMarksResResult          `json:"Result,omitempty"`
}

type GetConversationMarksResResponseMetadata struct {

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
	Error   *GetConversationMarksResResponseMetadataError `json:"Error,omitempty"`
}

type GetConversationMarksResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type GetConversationMarksResResult struct {

	// REQUIRED; 标记类型
	MarkTypes []string `json:"MarkTypes"`
}

type GetConversationSettingBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 会话成员 UserId，UserId 必须大于 0。
	ParticipantUserID int64 `json:"ParticipantUserId"`

	// 是否需要该成员在会话中的已读位置。
	// * true：不需要。
	// * false：需要。
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

	// REQUIRED; 用户会话设置
	ConversationSettingInfo GetConversationSettingResResultConversationSettingInfo `json:"ConversationSettingInfo"`
}

// GetConversationSettingResResultConversationSettingInfo - 用户会话设置
type GetConversationSettingResResultConversationSettingInfo struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 会话Id，字符串类型，含义跟ConversationShortId一样，用来定位唯一的一个会话，历史原因，目前大部分接口都在使用ConversationShortId，但是仍然有比较比较老的接口会使用到ConversationId，如果目前您接入的功能没有使用到ConversationId直接忽略即可
	ConversationID string `json:"ConversationId"`

	// REQUIRED; 会话类型。
	// * 1：单聊
	// * 2：群聊
	ConversationType int32 `json:"ConversationType"`

	// REQUIRED; 信箱，用于逻辑隔离
	InboxType int32 `json:"InboxType"`

	// REQUIRED; 群成员UserId
	ParticipantUserID int64 `json:"ParticipantUserId"`

	// REQUIRED; 用户已读位置
	ReadIndex int64 `json:"ReadIndex"`

	// REQUIRED; 置顶时间，单位为毫秒。0表示未置顶
	StickTopTime int64 `json:"StickTopTime"`

	// 扩展字段。 key 的数据类型为 String，value 的数据类型为 String。
	Ext map[string]string `json:"Ext,omitempty"`

	// 是否开启免打扰。
	// * true：开启。
	// * false：不开启
	IsMute *bool `json:"IsMute,omitempty"`

	// 是否收藏。
	// * true：收藏。
	// * false：不收藏
	IsSetFavorite *bool `json:"IsSetFavorite,omitempty"`

	// 是否置顶。
	// * true：置顶。
	// * false：不置顶
	IsStickTop *bool `json:"IsStickTop,omitempty"`

	// 收藏时间，单位为毫秒。0表示未收藏
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

type GetMessagesReadReceiptBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话Id
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 消息Id
	MessageIDs []int64 `json:"MessageIds"`
}

type GetMessagesReadReceiptRes struct {

	// REQUIRED
	ResponseMetadata GetMessagesReadReceiptResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetMessagesReadReceiptResResult `json:"Result,omitempty"`
}

type GetMessagesReadReceiptResResponseMetadata struct {

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
	Error   *GetMessagesReadReceiptResResponseMetadataError `json:"Error,omitempty"`
}

type GetMessagesReadReceiptResResponseMetadataError struct {

	// REQUIRED
	Code string `json:"Code"`

	// REQUIRED
	Message string `json:"Message"`
}

// GetMessagesReadReceiptResResult - 视请求的接口而定
type GetMessagesReadReceiptResResult struct {

	// REQUIRED; 已读回执详情
	ReadReceipt []GetMessagesReadReceiptResResultReadReceiptItem `json:"ReadReceipt"`
}

type GetMessagesReadReceiptResResultReadReceiptItem struct {

	// REQUIRED; 会话Id
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 消息Id
	MessageID int64 `json:"MessageId"`

	// REQUIRED; 消息已读的UserId列表
	ReadUserIDs []int64 `json:"ReadUserIds"`

	// REQUIRED; 单聊中消息的接收方是否已读（只有单聊会话这个字段才有意义）
	ReceiverIsRead bool `json:"ReceiverIsRead"`

	// REQUIRED; 消息未读的UserId列表
	UnReadUserIDs []int64 `json:"UnReadUserIds"`
}

type GetMessagesBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 消息 ID
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

	// REQUIRED; 消息详细信息
	Messages []GetMessagesResResultMessagesItem `json:"Messages"`
}

type GetMessagesResResultMessagesItem struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 消息内容
	Content string `json:"Content"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 会话类型
	// * 1：单聊。
	// * 2：群聊。
	ConversationType int32 `json:"ConversationType"`

	// REQUIRED; 消息创建时间戳，单位为毫秒
	CreateTime int64 `json:"CreateTime"`

	// REQUIRED; 消息 ID
	MessageID int64 `json:"MessageId"`

	// REQUIRED; 消息类型。
	// * 10001：文本。
	// * 10003：图片。
	// * 10004：视频
	// * 10005：文件
	// * 10006：音频
	// * 10012：自定义消息
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

	// REQUIRED; 被引用的消息类型
	// * 10001：文本。
	// * 10003：图片。
	// * 10004：视频
	// * 10005：文件
	// * 10006：音频
	// * 10012：自定义消息
	ReferencedMessageType int32 `json:"ReferencedMessageType"`

	// REQUIRED; 被引用的消息状态
	// * 0：消息可见
	// * 1：消息已过期
	// * 2：消息（对用户）不可见
	// * 3：消息被撤回
	// * 4：消息本身可见，后因删除不可见
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
	Version string                                           `json:"Version"`
	Error   *GetParticipantReadIndexResResponseMetadataError `json:"Error,omitempty"`
}

type GetParticipantReadIndexResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

type GetParticipantReadIndexResResult struct {

	// REQUIRED; 会话成员信息
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

	// 数据来源。
	// * 0：从缓存中拉取，按会话最近活跃排序。
	// * 1：从数据库中拉取，按照创建时间正序排序。
	// * 2：拉取用户创建的直播群会话，按照创建时间逆序排序
	DataType *int32 `json:"DataType,omitempty"`

	// 信箱，用于逻辑隔离。 默认值为 0。
	InboxType *int32 `json:"InboxType,omitempty"`

	// 是否忽略会话成员数。
	// * true：忽略。
	// * false：不忽略。 默认值为 false。
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

	// REQUIRED; 会话详细信息
	ConversationInfos []GetUserConversationsResResultConversationInfosItem `json:"ConversationInfos"`

	// REQUIRED; 是否还有更多的数据
	HasMore bool `json:"HasMore"`

	// REQUIRED; 下一批数据起始位置
	NextCursor int64 `json:"NextCursor"`
}

type GetUserConversationsResResultConversationInfosItem struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话Id，字符串类型，含义跟ConversationShortId一样，用来定位唯一的一个会话，历史原因，目前大部分接口都在使用ConversationShortId，但是仍然有比较比较老的接口会使用到ConversationId，如果目前您接入的功能没有使用到ConversationId直接忽略即可
	ConversationID string `json:"ConversationId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 会话类型。
	// * 1：单聊
	// * 2：群聊
	ConversationType int32 `json:"ConversationType"`

	// REQUIRED; 群聊创建时间戳，单位为秒
	CreateTime int64 `json:"CreateTime"`

	// REQUIRED; 创群人 UserId
	CreatorUserID int64 `json:"CreatorUserId"`

	// REQUIRED; 信箱，用于逻辑隔离
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

	// 单聊会话另一个 UserId
	OtherUserID *int64 `json:"OtherUserId,omitempty"`

	// 会话状态。
	// * 0：正常
	// * 1：已解散
	Status *int32 `json:"Status,omitempty"`
}

type IsFriendBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 需要校验用户的 UserId，一次最多检验 10 个用户
	FriendUserIDs []int64 `json:"FriendUserIds"`

	// REQUIRED; 用户 UserId
	UserID int64 `json:"UserId"`

	// 信箱，用做逻辑隔离。默认值为 0
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

	// REQUIRED; 好友关系详情
	Infos []IsFriendResResultInfosItem `json:"Infos"`
}

type IsFriendResResultInfosItem struct {

	// REQUIRED; 校验好友的 UserId
	FriendUserID int64 `json:"FriendUserId"`

	// REQUIRED; 是否是好友。
	// * true：是。
	// * false：否
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

	// REQUIRED; 校验结果。key 为 uid，value 为 true：在黑名单中。 value 为 false：不在黑名单中
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

	// REQUIRED; 用户是否在会话中。
	// * true：是。
	// * false：否。
	IsUserInConversation bool `json:"IsUserInConversation"`
}

type ListFriendBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 查询用户 UserId
	UserID int64 `json:"UserId"`

	// 查询起始位置，默认值为 0，即第一个添加的好友。
	Cursor *int64 `json:"Cursor,omitempty"`

	// 信箱，用做逻辑隔离。默认值为 0
	InboxType *int32 `json:"InboxType,omitempty"`

	// 查询条数，每次最多查询 20 位好友。默认值为 20。
	Limit *int64 `json:"Limit,omitempty"`

	// 是否需要好友总数。
	// * false：不需要。
	// * true：需要。
	// 默认值为 false。
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

type MarkConversationBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 标记类型
	MarkTypes []string `json:"MarkTypes"`

	// 操作类型。默认值为 1。
	// * 1:新增
	// * 2:删除。
	MarkAction *int32 `json:"MarkAction,omitempty"`

	// 操作人UserId
	Operator *int64 `json:"Operator,omitempty"`
}

type MarkConversationRes struct {

	// REQUIRED
	ResponseMetadata MarkConversationResResponseMetadata `json:"ResponseMetadata"`

	// 空。此接口无需关注
	Result interface{} `json:"Result,omitempty"`
}

type MarkConversationResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                    `json:"Version"`
	Error   *MarkConversationResResponseMetadataError `json:"Error,omitempty"`
}

type MarkConversationResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
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

	// 会话的扩展字段。key 的数据类型为 String，value 的数据类型为 String。
	Ext map[string]string `json:"Ext,omitempty"`

	// 会话名
	Name *string `json:"Name,omitempty"`

	// 会话公告
	Notice *string `json:"Notice,omitempty"`

	// 会话拥有人 UserId
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

	// REQUIRED; 请求的版本号，属于请求的公共参数。
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

	// REQUIRED; 用户会话设置
	ConversationSettingInfo ModifyConversationSettingBodyConversationSettingInfo `json:"ConversationSettingInfo"`
}

// ModifyConversationSettingBodyConversationSettingInfo - 用户会话设置
type ModifyConversationSettingBodyConversationSettingInfo struct {

	// REQUIRED; 需要修改的会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 需要修改的会话成员 UserId
	ParticipantUserID int64 `json:"ParticipantUserId"`

	// 扩展字段
	Ext map[string]string `json:"Ext,omitempty"`

	// 是否开启免打扰。
	// * true：开启。
	// * false：不开启
	// 默认值为 false。
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

	// REQUIRED; 消息 ID
	MessageID int64 `json:"MessageId"`

	// 消息内容。修改时Content内容需符合客户端格式，详细信息请参看消息格式 [https://www.volcengine.com/docs/6348/372181#server]。
	Content *string `json:"Content,omitempty"`

	// 消息的扩展字段
	Ext map[string]string `json:"Ext,omitempty"`
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

	// REQUIRED; 成员在直播群的状态
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

	// 撤回消息的会话成员角色。
	// * 0：普通会话成员。
	// * 1：群主。
	// * 2：管理员。
	RecallRole *int32 `json:"RecallRole,omitempty"`
}

type RecallMessageRes struct {

	// REQUIRED
	ResponseMetadata RecallMessageResResponseMetadata `json:"ResponseMetadata"`

	// 空，此接口可忽略此字段。
	Result interface{} `json:"Result,omitempty"`
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

type RegisterUsersBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 注册用户的信息。一次最多注册 20 个用户。
	Users []RegisterUsersBodyUsersItem `json:"Users"`
}

type RegisterUsersBodyUsersItem struct {

	// REQUIRED; 注册用户的 UserId
	UserID int64 `json:"UserId"`

	// 1. 字段个数不能超过20个；
	// 2. key不能以"s:"开头（im的保留key），长度不能超过50字节
	// 3. value长度不能超过500字节
	Ext map[string]string `json:"Ext,omitempty"`

	// 长度不能超过100字节
	NickName *string `json:"NickName,omitempty"`

	// 长度不能超过500字节
	Portrait *string `json:"Portrait,omitempty"`

	// 用于全员广播
	Tags []string `json:"Tags,omitempty"`
}

type RegisterUsersRes struct {

	// REQUIRED
	ResponseMetadata RegisterUsersResResponseMetadata `json:"ResponseMetadata"`

	// 空。此接口无需关注。
	Result *RegisterUsersResResult `json:"Result,omitempty"`
}

type RegisterUsersResResponseMetadata struct {

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
	Error   *RegisterUsersResResponseMetadataError `json:"Error,omitempty"`
}

type RegisterUsersResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

// RegisterUsersResResult - 空。此接口无需关注。
type RegisterUsersResResult struct {

	// REQUIRED; 注册失败的用户信息
	FailedInfos []RegisterUsersResResultFailedInfosItem `json:"FailedInfos"`
}

type RegisterUsersResResultFailedInfosItem struct {

	// REQUIRED; 错误码。参错误码 [412253]
	Code string `json:"Code"`

	// REQUIRED; 错误详细信息
	Message string `json:"Message"`

	// REQUIRED; 失败用户的 UserId
	UserID int64 `json:"UserId"`
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

	// 标记类型。仅支持查询具有该标记的在线成员。
	MarkType *string `json:"MarkType,omitempty"`

	// 是否只查询群成员 UserId。
	// * true：是。
	// * false：否。
	// 默认值为 false。
	OnlyUserIDs *bool `json:"OnlyUserIds,omitempty"`

	// 按照进群的时间正序或逆序查询。
	// * false：正序。
	// * true：逆序。
	// 默认值为：fasle。
	Reverse *bool `json:"Reverse,omitempty"`

	// 直播群专用字段。是否需要获得直播群在线成员列表。
	// * true：是：拉取直播群在线成员列表（包含群主和管理员，只要在线就会返回，非直播群不会返回任何数据）。
	// * false：否：拉取成员列表（非直播群），拉取群主管理员（直播群）。
	// 默认值为 false。
	// 当此字段设置 true 时，Reverse 和 OnlyUserIds 字段均会失效，该接口按成员进入直播群时间逆序返回在线成员列表。
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

	// REQUIRED; 群成员详细信息
	Participants []ScanConversationParticipantListResResultParticipantsItem `json:"Participants"`
}

type ScanConversationParticipantListResResultParticipantsItem struct {

	// REQUIRED; 群成员头像
	AvatarURL string `json:"AvatarUrl"`

	// REQUIRED; 会话 ID
	ConversationShortID int64 `json:"ConversationShortId"`

	// REQUIRED; 操作人对应的 UserId
	Operator int64 `json:"Operator"`

	// REQUIRED; 群成员 ID
	ParticipantUserID int64 `json:"ParticipantUserId"`

	// REQUIRED; 成员身份。
	// * 0：普通成员。
	// * 1：群主。
	// * 2：群管理员。
	Role int64 `json:"Role"`

	// 禁言时间戳，单位为秒。0表示不禁言
	BlockTime *int64 `json:"BlockTime,omitempty"`

	// 成员进群时间对应时间戳，单位为秒
	CreateTime *int64 `json:"CreateTime,omitempty"`

	// 成员的扩展字段
	Ext map[string]string `json:"Ext,omitempty"`

	// 成员等级
	Level *int32 `json:"Level,omitempty"`

	// 成员标记
	Marks []string `json:"Marks,omitempty"`

	// 成员昵称
	NickName *string `json:"NickName,omitempty"`

	// 成员状态。
	// * 0：正常
	// * 1：退出
	Status *int32 `json:"Status,omitempty"`
}

type SendMessageBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 消息内容。当你给客户端发消息时，Content内容需符合客户端格式，详细信息请参看消息格式 [https://www.volcengine.com/docs/6348/372181#server]。
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

	// 消息不可见会话成员列表。VisibleUsers和InvisibleUsers均为空时，代表对所有人可见。
	InvisibleUsers []*int64 `json:"InvisibleUsers,omitempty"`

	// 会话中@的人
	MentionedUsers []*int64 `json:"MentionedUsers,omitempty"`

	// 消息优先级。
	// * 0：低优先级。
	// * 1：普通优先级。
	// * 2：高优先级。
	// 该字段仅对直播群有效。为避免直播群中消息频率太多导致服务端压力过大，你可以设置消息的优先级。当前直播群下，普通优先级和低优先级消息共用频控阈值为 30 条/秒，超过部分会被服务端直接丢弃。高优消息频控阈值为 10 条/秒，超过部分服务端无法保证不丢失。
	Priority *int32 `json:"Priority,omitempty"`

	// 引用消息。该接口中，该字段只需传入ReferencedMessageId和Hint参数
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

type BatchSendMessageBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 消息内容。当你给客户端发消息时，Content 内容需符合客户端格式，详细信息请参看消息格式
	Content string `json:"Content"`

	// REQUIRED; 消息类型
	MsgType int32 `json:"MsgType"`

	// REQUIRED; 消息接收人 UserId 列表
	Receiver []int64 `json:"Receiver"`

	// REQUIRED; 消息发送人 UserId
	Sender int64 `json:"Sender"`

	// 消息的扩展字段，key 的数据类型为 String，value 的数据类型为 String
	Ext map[string]string `json:"Ext,omitempty"`

	// 信箱，用做逻辑隔离 默认值为 0
	InboxType *int32 `json:"InboxType,omitempty"`
}

type BatchSendMessageRes struct {

	// REQUIRED
	ResponseMetadata BatchSendMessageResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result BatchSendMessageResResult `json:"Result"`
}

type BatchSendMessageResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                    `json:"Version"`
	Error   *BatchSendMessageResResponseMetadataError `json:"Error,omitempty"`
}

type BatchSendMessageResResponseMetadataError struct {

	// REQUIRED
	Code string `json:"Code"`

	// REQUIRED
	Message string `json:"Message"`
}

type BatchSendMessageResResult struct {

	// REQUIRED; 消息Id列表，key为用户UserId，value为消息Id
	UserMessageID map[string]int64 `json:"UserMessageId"`
}

type UnRegisterUsersBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 注销用户 UserId。一次最多注销 10 个用户。
	UserIDs []int64 `json:"UserIds"`
}

type UnRegisterUsersRes struct {

	// REQUIRED
	ResponseMetadata UnRegisterUsersResResponseMetadata `json:"ResponseMetadata"`

	// 空。此接口无需关注
	Result *UnRegisterUsersResResult `json:"Result,omitempty"`
}

type UnRegisterUsersResResponseMetadata struct {

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
	Error   *UnRegisterUsersResResponseMetadataError `json:"Error,omitempty"`
}

type UnRegisterUsersResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}

// UnRegisterUsersResResult - 空。此接口无需关注
type UnRegisterUsersResResult struct {

	// REQUIRED; 注销失败的用户信息
	FailedInfos []UnRegisterUsersResResultFailedInfosItem `json:"FailedInfos"`
}

type UnRegisterUsersResResultFailedInfosItem struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：错误码 [https://www.volcengine.com/docs/6348/412253]
	Message string `json:"Message"`

	// REQUIRED; 失败的uid
	UserID int64 `json:"UserId"`
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
	// 假设 ID 为 10002 的黑名单用户原 ext 为 {"key1":"value1"}，
	// * 若调用本接口通过 BlackListUserExt 传入 {"key1":"value2"}，此时 10002 用户的 ext 将更新为 {"key1":"value2"}；
	// * 若调用本接口通过 BlackListUserExt 传入 {"key2":"value2"}，此时 10002 用户的 ext 将更新为 {"key1":"value1", "key2":"value2"}。
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

	// REQUIRED; 错误码 [https://www.volcengine.com/docs/6348/412253]
	Code string `json:"Code"`

	// REQUIRED; 错误信息 [https://www.volcengine.com/docs/6348/412253]
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

	// 信箱，用做逻辑隔离。默认值为0
	InboxType *int32 `json:"InboxType,omitempty"`
}

type UpdateFriendBodyFriendInfosItem struct {

	// REQUIRED; 好友扩展字段。Ext与Alias为非必填字段，但是至少需要填一个，否则服务端会报错。
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

	// 更新失败的好友信息
	FailedInfos []UpdateFriendResResultFailedInfos `json:"FailedInfos,omitempty"`
}

// UpdateFriendResResultFailedInfos - 更新失败的好友信息
type UpdateFriendResResultFailedInfos struct {

	// REQUIRED; 错误码 [https://www.volcengine.com/docs/6348/412253]
	Code string `json:"Code"`

	// REQUIRED; 错误信息 [https://www.volcengine.com/docs/6348/412253]
	Message string `json:"Message"`

	// REQUIRED; 好友 UserId
	UserID int64 `json:"UserId"`
}

type UserBroadcastBody struct {

	// REQUIRED; 应用的唯一标志
	AppID int32 `json:"AppId"`

	// REQUIRED; 消息内容。关于消息类型和消息内容，参看 消息相关 [https://www.volcengine.com/docs/6348/372181]
	Content string `json:"Content"`

	// REQUIRED; 消息过期时间，单位为秒，范围为[1,604800]
	MsgTTL int64 `json:"MsgTTL"`

	// REQUIRED; 消息类型。
	MsgType int32 `json:"MsgType"`

	// REQUIRED; 发送方
	Sender int64 `json:"Sender"`

	// 消息的扩展字段。key 的类型为 String，value 的类型为 String。
	Ext map[string]string `json:"Ext,omitempty"`

	// 信箱，用做逻辑隔离。默认值为 0。
	InboxType *int32 `json:"InboxType,omitempty"`

	// 筛选用户时，Tag 匹配关系。 0：或 1：且 默认值为 0。
	TagOp *int32 `json:"TagOp,omitempty"`

	// 筛选标签。不填表示全员广播。
	Tags []string `json:"Tags,omitempty"`
}

type UserBroadcastRes struct {

	// REQUIRED
	ResponseMetadata UserBroadcastResResponseMetadata `json:"ResponseMetadata"`

	// 空，此接口可忽略此字段。
	Result interface{} `json:"Result,omitempty"`
}

type UserBroadcastResResponseMetadata struct {

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
	Error   *UserBroadcastResResponseMetadataError `json:"Error,omitempty"`
}

type UserBroadcastResResponseMetadataError struct {

	// REQUIRED; 错误码。参考:https://www.volcengine.com/docs/6348/412253
	Code string `json:"Code"`

	// REQUIRED; 错误信息。参考：https://www.volcengine.com/docs/6348/412253
	Message string `json:"Message"`
}
type GetConversationSetting struct{}
type GetUserConversations struct{}
type SendMessage struct{}
type BatchRemoveManagerQuery struct{}
type CreateConversation struct{}
type RemoveBlackListQuery struct{}
type ScanConversationParticipantListQuery struct{}
type GetBlackList struct{}
type AddFriendQuery struct{}
type BatchAddWhitelistParticipantQuery struct{}
type BatchSendMessage struct{}
type BatchSendMessageQuery struct{}
type RecallMessage struct{}
type UserBroadcast struct{}
type BatchRemoveWhitelistParticipant struct{}
type BatchUpdateUserQuery struct{}
type IsUserInConversationQuery struct{}
type QueryOnlineStatus struct{}
type GetAppTokenQuery struct{}
type GetParticipantReadIndexQuery struct{}
type ListFriend struct{}
type QueryOnlineStatusQuery struct{}
type BatchDeleteBlockParticipantsQuery struct{}
type BatchGetConversationParticipant struct{}
type IsInBlackList struct{}
type GetConversationMarksQuery struct{}
type BatchGetBlockParticipantsQuery struct{}
type CreateConversationQuery struct{}
type IsFriendQuery struct{}
type BatchDeleteConversationParticipantQuery struct{}
type GetUserConversationsQuery struct{}
type GetConversationMessages struct{}
type BatchAddConversationParticipantQuery struct{}
type AddBlackListQuery struct{}
type GetAppToken struct{}
type MarkConversation struct{}
type BatchGetConversationParticipantQuery struct{}
type RemoveBlackList struct{}
type BatchAddManagerQuery struct{}
type GetConversationSettingQuery struct{}
type BatchUpdateUserTagsQuery struct{}
type RegisterUsers struct{}
type ModifyParticipantReadIndexQuery struct{}
type ModifyParticipantReadIndex struct{}
type BatchAddWhitelistParticipant struct{}
type UpdateFriendQuery struct{}
type ModifyConversationSetting struct{}
type QueryLiveParticipantStatus struct{}
type IsInBlackListQuery struct{}
type IsUserInConversation struct{}
type DestroyConversation struct{}
type RecallMessageQuery struct{}
type UnRegisterUsersQuery struct{}
type BatchUpdateUser struct{}
type GetConversationUserCountQuery struct{}
type BatchGetConversationsQuery struct{}
type ModifyMessage struct{}
type ModifyConversationSettingQuery struct{}
type SendMessageQuery struct{}
type DeleteFriendQuery struct{}
type GetBlackListQuery struct{}
type BatchAddConversationParticipant struct{}
type DestroyConversationQuery struct{}
type GetConversationMessagesQuery struct{}
type DeleteConversationMessage struct{}
type ScanConversationParticipantList struct{}
type ListFriendQuery struct{}
type GetMessagesQuery struct{}
type UpdateFriend struct{}
type BatchUpdateLiveParticipants struct{}
type UnRegisterUsers struct{}
type GetMessagesReadReceiptQuery struct{}
type BatchGetBlockParticipants struct{}
type BatchGetConversations struct{}
type DeleteMessage struct{}
type GetMessages struct{}
type MarkConversationQuery struct{}
type GetParticipantReadIndex struct{}
type QueryLiveParticipantStatusQuery struct{}
type UpdateBlackListQuery struct{}
type BatchGetUser struct{}
type BatchRemoveWhitelistParticipantQuery struct{}
type GetConversationUserCount struct{}
type BatchGetWhitelistParticipant struct{}
type BatchAddManager struct{}
type ModifyConversation struct{}
type BatchAddBlockParticipants struct{}
type UpdateBlackList struct{}
type GetConversationMarks struct{}
type BatchAddBlockParticipantsQuery struct{}
type BatchRemoveManager struct{}
type IsFriend struct{}
type BatchUpdateLiveParticipantsQuery struct{}
type BatchGetUserQuery struct{}
type BatchDeleteConversationParticipant struct{}
type RegisterUsersQuery struct{}
type BatchGetWhitelistParticipantQuery struct{}
type DeleteConversationMessageQuery struct{}
type BatchUpdateUserTags struct{}
type BatchModifyConversationParticipant struct{}
type BatchDeleteBlockParticipants struct{}
type ModifyMessageQuery struct{}
type UserBroadcastQuery struct{}
type BatchModifyConversationParticipantQuery struct{}
type ModifyConversationQuery struct{}
type AddBlackList struct{}
type DeleteMessageQuery struct{}
type DeleteFriend struct{}
type AddFriend struct{}
type GetMessagesReadReceipt struct{}
type BatchDeleteConversationParticipantReq struct {
	*BatchDeleteConversationParticipantQuery
	*BatchDeleteConversationParticipantBody
}
type BatchAddManagerReq struct {
	*BatchAddManagerQuery
	*BatchAddManagerBody
}
type GetConversationSettingReq struct {
	*GetConversationSettingQuery
	*GetConversationSettingBody
}
type DestroyConversationReq struct {
	*DestroyConversationQuery
	*DestroyConversationBody
}
type ModifyMessageReq struct {
	*ModifyMessageQuery
	*ModifyMessageBody
}
type IsUserInConversationReq struct {
	*IsUserInConversationQuery
	*IsUserInConversationBody
}
type BatchGetConversationParticipantReq struct {
	*BatchGetConversationParticipantQuery
	*BatchGetConversationParticipantBody
}
type ListFriendReq struct {
	*ListFriendQuery
	*ListFriendBody
}
type ModifyConversationSettingReq struct {
	*ModifyConversationSettingQuery
	*ModifyConversationSettingBody
}
type SendMessageReq struct {
	*SendMessageQuery
	*SendMessageBody
}

type BatchSendMessageReq struct {
	*BatchSendMessageQuery
	*BatchSendMessageBody
}

type GetMessagesReq struct {
	*GetMessagesQuery
	*GetMessagesBody
}
type UnRegisterUsersReq struct {
	*UnRegisterUsersQuery
	*UnRegisterUsersBody
}
type BatchUpdateUserReq struct {
	*BatchUpdateUserQuery
	*BatchUpdateUserBody
}
type BatchModifyConversationParticipantReq struct {
	*BatchModifyConversationParticipantQuery
	*BatchModifyConversationParticipantBody
}
type BatchAddConversationParticipantReq struct {
	*BatchAddConversationParticipantQuery
	*BatchAddConversationParticipantBody
}
type IsFriendReq struct {
	*IsFriendQuery
	*IsFriendBody
}
type GetConversationMarksReq struct {
	*GetConversationMarksQuery
	*GetConversationMarksBody
}
type GetUserConversationsReq struct {
	*GetUserConversationsQuery
	*GetUserConversationsBody
}
type RecallMessageReq struct {
	*RecallMessageQuery
	*RecallMessageBody
}
type RegisterUsersReq struct {
	*RegisterUsersQuery
	*RegisterUsersBody
}
type MarkConversationReq struct {
	*MarkConversationQuery
	*MarkConversationBody
}
type GetParticipantReadIndexReq struct {
	*GetParticipantReadIndexQuery
	*GetParticipantReadIndexBody
}
type GetConversationMessagesReq struct {
	*GetConversationMessagesQuery
	*GetConversationMessagesBody
}
type QueryLiveParticipantStatusReq struct {
	*QueryLiveParticipantStatusQuery
	*QueryLiveParticipantStatusBody
}
type UserBroadcastReq struct {
	*UserBroadcastQuery
	*UserBroadcastBody
}
type BatchAddBlockParticipantsReq struct {
	*BatchAddBlockParticipantsQuery
	*BatchAddBlockParticipantsBody
}
type RemoveBlackListReq struct {
	*RemoveBlackListQuery
	*RemoveBlackListBody
}
type UpdateFriendReq struct {
	*UpdateFriendQuery
	*UpdateFriendBody
}
type BatchGetBlockParticipantsReq struct {
	*BatchGetBlockParticipantsQuery
	*BatchGetBlockParticipantsBody
}
type BatchGetWhitelistParticipantReq struct {
	*BatchGetWhitelistParticipantQuery
	*BatchGetWhitelistParticipantBody
}
type AddFriendReq struct {
	*AddFriendQuery
	*AddFriendBody
}
type ModifyParticipantReadIndexReq struct {
	*ModifyParticipantReadIndexQuery
	*ModifyParticipantReadIndexBody
}
type BatchRemoveWhitelistParticipantReq struct {
	*BatchRemoveWhitelistParticipantQuery
	*BatchRemoveWhitelistParticipantBody
}
type UpdateBlackListReq struct {
	*UpdateBlackListQuery
	*UpdateBlackListBody
}
type GetAppTokenReq struct {
	*GetAppTokenQuery
	*GetAppTokenBody
}
type ScanConversationParticipantListReq struct {
	*ScanConversationParticipantListQuery
	*ScanConversationParticipantListBody
}
type BatchDeleteBlockParticipantsReq struct {
	*BatchDeleteBlockParticipantsQuery
	*BatchDeleteBlockParticipantsBody
}
type BatchUpdateLiveParticipantsReq struct {
	*BatchUpdateLiveParticipantsQuery
	*BatchUpdateLiveParticipantsBody
}
type BatchGetConversationsReq struct {
	*BatchGetConversationsQuery
	*BatchGetConversationsBody
}
type GetConversationUserCountReq struct {
	*GetConversationUserCountQuery
	*GetConversationUserCountBody
}
type IsInBlackListReq struct {
	*IsInBlackListQuery
	*IsInBlackListBody
}
type BatchGetUserReq struct {
	*BatchGetUserQuery
	*BatchGetUserBody
}
type ModifyConversationReq struct {
	*ModifyConversationQuery
	*ModifyConversationBody
}
type CreateConversationReq struct {
	*CreateConversationQuery
	*CreateConversationBody
}
type BatchUpdateUserTagsReq struct {
	*BatchUpdateUserTagsQuery
	*BatchUpdateUserTagsBody
}
type GetBlackListReq struct {
	*GetBlackListQuery
	*GetBlackListBody
}
type BatchAddWhitelistParticipantReq struct {
	*BatchAddWhitelistParticipantQuery
	*BatchAddWhitelistParticipantBody
}
type DeleteMessageReq struct {
	*DeleteMessageQuery
	*DeleteMessageBody
}
type QueryOnlineStatusReq struct {
	*QueryOnlineStatusQuery
	*QueryOnlineStatusBody
}
type BatchRemoveManagerReq struct {
	*BatchRemoveManagerQuery
	*BatchRemoveManagerBody
}
type DeleteConversationMessageReq struct {
	*DeleteConversationMessageQuery
	*DeleteConversationMessageBody
}
type DeleteFriendReq struct {
	*DeleteFriendQuery
	*DeleteFriendBody
}
type AddBlackListReq struct {
	*AddBlackListQuery
	*AddBlackListBody
}

type GetMessagesReadReceiptReq struct {
	*GetMessagesReadReceiptQuery
	*GetMessagesReadReceiptBody
}
