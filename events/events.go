/*
Package events for community events handling.

See more https://vk.com/dev/groups_events
*/
package events // import "github.com/SevereCloud/vksdk/events"

import (
	"context"
	"encoding/json"

	"github.com/SevereCloud/vksdk/internal"
)

// Event type list
const (
	Confirmation         = "confirmation"
	MessageNew           = "message_new"
	MessageReply         = "message_reply"
	MessageEdit          = "message_edit"
	MessageAllow         = "message_allow"
	MessageDeny          = "message_deny"
	MessageTypingState   = "message_typing_state"
	PhotoNew             = "photo_new"
	PhotoCommentNew      = "photo_comment_new"
	PhotoCommentEdit     = "photo_comment_edit"
	PhotoCommentRestore  = "photo_comment_restore"
	PhotoCommentDelete   = "photo_comment_delete"
	AudioNew             = "audio_new"
	VideoNew             = "video_new"
	VideoCommentNew      = "video_comment_new"
	VideoCommentEdit     = "video_comment_edit"
	VideoCommentRestore  = "video_comment_restore"
	VideoCommentDelete   = "video_comment_delete"
	WallPostNew          = "wall_post_new"
	WallRepost           = "wall_repost"
	WallReplyNew         = "wall_reply_new"
	WallReplyEdit        = "wall_reply_edit"
	WallReplyRestore     = "wall_reply_restore"
	WallReplyDelete      = "wall_reply_delete"
	BoardPostNew         = "board_post_new"
	BoardPostEdit        = "board_post_edit"
	BoardPostRestore     = "board_post_restore"
	BoardPostDelete      = "board_post_delete"
	MarketCommentNew     = "market_comment_new"
	MarketCommentEdit    = "market_comment_edit"
	MarketCommentRestore = "market_comment_restore"
	MarketCommentDelete  = "market_comment_delete"
	GroupLeave           = "group_leave"
	GroupJoin            = "group_join"
	UserBlock            = "user_block"
	UserUnblock          = "user_unblock"
	PollVoteNew          = "poll_vote_new"
	GroupOfficersEdit    = "group_officers_edit"
	GroupChangeSettings  = "group_change_settings"
	GroupChangePhoto     = "group_change_photo"
	VkpayTransaction     = "vkpay_transaction"
	LeadFormsNew         = "lead_forms_new"
	AppPayload           = "app_payload"
	MessageRead          = "message_read"
)

// GroupEvent struct
type GroupEvent struct {
	Type    string          `json:"type"`
	Object  json.RawMessage `json:"object"`
	GroupID int             `json:"group_id"`
	EventID string          `json:"event_id"`
	Secret  string          `json:"secret"`
}

// FuncList struct
type FuncList struct {
	messageNew           []MessageNewFunc
	messageReply         []MessageReplyFunc
	messageEdit          []MessageEditFunc
	messageAllow         []MessageAllowFunc
	messageDeny          []MessageDenyFunc
	messageTypingState   []MessageTypingStateFunc
	photoNew             []PhotoNewFunc
	photoCommentNew      []PhotoCommentNewFunc
	photoCommentEdit     []PhotoCommentEditFunc
	photoCommentRestore  []PhotoCommentRestoreFunc
	photoCommentDelete   []PhotoCommentDeleteFunc
	audioNew             []AudioNewFunc
	videoNew             []VideoNewFunc
	videoCommentNew      []VideoCommentNewFunc
	videoCommentEdit     []VideoCommentEditFunc
	videoCommentRestore  []VideoCommentRestoreFunc
	videoCommentDelete   []VideoCommentDeleteFunc
	wallPostNew          []WallPostNewFunc
	wallRepost           []WallRepostFunc
	wallReplyNew         []WallReplyNewFunc
	wallReplyEdit        []WallReplyEditFunc
	wallReplyRestore     []WallReplyRestoreFunc
	wallReplyDelete      []WallReplyDeleteFunc
	boardPostNew         []BoardPostNewFunc
	boardPostEdit        []BoardPostEditFunc
	boardPostRestore     []BoardPostRestoreFunc
	boardPostDelete      []BoardPostDeleteFunc
	marketCommentNew     []MarketCommentNewFunc
	marketCommentEdit    []MarketCommentEditFunc
	marketCommentRestore []MarketCommentRestoreFunc
	marketCommentDelete  []MarketCommentDeleteFunc
	groupLeave           []GroupLeaveFunc
	groupJoin            []GroupJoinFunc
	userBlock            []UserBlockFunc
	userUnblock          []UserUnblockFunc
	pollVoteNew          []PollVoteNewFunc
	groupOfficersEdit    []GroupOfficersEditFunc
	groupChangeSettings  []GroupChangeSettingsFunc
	groupChangePhoto     []GroupChangePhotoFunc
	vkpayTransaction     []VkpayTransactionFunc
	leadFormsNew         []LeadFormsNewFunc
	appPayload           []AppPayloadFunc
	messageRead          []MessageReadFunc
	special              map[string][]func(context.Context, GroupEvent)
}

// NewFuncList returns a new FuncList
func NewFuncList() *FuncList {
	return &FuncList{
		special: make(map[string][]func(context.Context, GroupEvent)),
	}
}

// Handler group event handler
func (fl *FuncList) Handler(ctx context.Context, e GroupEvent) error { // nolint:gocyclo
	ctx = context.WithValue(ctx, internal.GroupIDKey, e.GroupID)
	ctx = context.WithValue(ctx, internal.EventIDKey, e.EventID)

	if sliceFunc, ok := fl.special[e.Type]; ok {
		for _, f := range sliceFunc {
			f(ctx, e)
		}
	}

	switch e.Type {
	case MessageNew:
		var obj MessageNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.messageNew {
			f(ctx, obj)
		}
	case MessageReply:
		var obj MessageReplyObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.messageReply {
			f(ctx, obj)
		}
	case MessageEdit:
		var obj MessageEditObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.messageEdit {
			f(ctx, obj)
		}
	case MessageAllow:
		var obj MessageAllowObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.messageAllow {
			f(ctx, obj)
		}
	case MessageDeny:
		var obj MessageDenyObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.messageDeny {
			f(ctx, obj)
		}
	case MessageTypingState: // На основе ответа
		var obj MessageTypingStateObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.messageTypingState {
			f(ctx, obj)
		}
	case PhotoNew:
		var obj PhotoNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.photoNew {
			f(ctx, obj)
		}
	case PhotoCommentNew:
		var obj PhotoCommentNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.photoCommentNew {
			f(ctx, obj)
		}
	case PhotoCommentEdit:
		var obj PhotoCommentEditObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.photoCommentEdit {
			f(ctx, obj)
		}
	case PhotoCommentRestore:
		var obj PhotoCommentRestoreObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.photoCommentRestore {
			f(ctx, obj)
		}
	case PhotoCommentDelete:
		var obj PhotoCommentDeleteObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.photoCommentDelete {
			f(ctx, obj)
		}
	case AudioNew:
		var obj AudioNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.audioNew {
			f(ctx, obj)
		}
	case VideoNew:
		var obj VideoNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.videoNew {
			f(ctx, obj)
		}
	case VideoCommentNew:
		var obj VideoCommentNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.videoCommentNew {
			f(ctx, obj)
		}
	case VideoCommentEdit:
		var obj VideoCommentEditObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.videoCommentEdit {
			f(ctx, obj)
		}
	case VideoCommentRestore:
		var obj VideoCommentRestoreObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.videoCommentRestore {
			f(ctx, obj)
		}
	case VideoCommentDelete:
		var obj VideoCommentDeleteObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.videoCommentDelete {
			f(ctx, obj)
		}
	case WallPostNew:
		var obj WallPostNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.wallPostNew {
			f(ctx, obj)
		}
	case WallRepost:
		var obj WallRepostObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.wallRepost {
			f(ctx, obj)
		}
	case WallReplyNew:
		var obj WallReplyNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.wallReplyNew {
			f(ctx, obj)
		}
	case WallReplyEdit:
		var obj WallReplyEditObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.wallReplyEdit {
			f(ctx, obj)
		}
	case WallReplyRestore:
		var obj WallReplyRestoreObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.wallReplyRestore {
			f(ctx, obj)
		}
	case WallReplyDelete:
		var obj WallReplyDeleteObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.wallReplyDelete {
			f(ctx, obj)
		}
	case BoardPostNew:
		var obj BoardPostNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.boardPostNew {
			f(ctx, obj)
		}
	case BoardPostEdit:
		var obj BoardPostEditObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.boardPostEdit {
			f(ctx, obj)
		}
	case BoardPostRestore:
		var obj BoardPostRestoreObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.boardPostRestore {
			f(ctx, obj)
		}
	case BoardPostDelete:
		var obj BoardPostDeleteObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.boardPostDelete {
			f(ctx, obj)
		}
	case MarketCommentNew:
		var obj MarketCommentNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.marketCommentNew {
			f(ctx, obj)
		}
	case MarketCommentEdit:
		var obj MarketCommentEditObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.marketCommentEdit {
			f(ctx, obj)
		}
	case MarketCommentRestore:
		var obj MarketCommentRestoreObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.marketCommentRestore {
			f(ctx, obj)
		}
	case MarketCommentDelete:
		var obj MarketCommentDeleteObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.marketCommentDelete {
			f(ctx, obj)
		}
	case GroupLeave:
		var obj GroupLeaveObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.groupLeave {
			f(ctx, obj)
		}
	case GroupJoin:
		var obj GroupJoinObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.groupJoin {
			f(ctx, obj)
		}
	case UserBlock:
		var obj UserBlockObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.userBlock {
			f(ctx, obj)
		}
	case UserUnblock:
		var obj UserUnblockObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.userUnblock {
			f(ctx, obj)
		}
	case PollVoteNew:
		var obj PollVoteNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.pollVoteNew {
			f(ctx, obj)
		}
	case GroupOfficersEdit:
		var obj GroupOfficersEditObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.groupOfficersEdit {
			f(ctx, obj)
		}
	case GroupChangeSettings:
		var obj GroupChangeSettingsObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.groupChangeSettings {
			f(ctx, obj)
		}
	case GroupChangePhoto:
		var obj GroupChangePhotoObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.groupChangePhoto {
			f(ctx, obj)
		}
	case VkpayTransaction:
		var obj VkpayTransactionObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.vkpayTransaction {
			f(ctx, obj)
		}
	case LeadFormsNew:
		var obj LeadFormsNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.leadFormsNew {
			f(ctx, obj)
		}
	case AppPayload:
		var obj AppPayloadObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.appPayload {
			f(ctx, obj)
		}
	case MessageRead:
		var obj MessageReadObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.messageRead {
			f(ctx, obj)
		}
	}
	// NOTE: like_add like_remove
	return nil
}

// OnEvent handler
func (fl *FuncList) OnEvent(eventType string, f func(context.Context, GroupEvent)) {
	if fl.special == nil {
		fl.special = make(map[string][]func(context.Context, GroupEvent))
	}

	fl.special[eventType] = append(fl.special[eventType], f)
}

// MessageNew handler
func (fl *FuncList) MessageNew(f MessageNewFunc) {
	fl.messageNew = append(fl.messageNew, f)
}

// MessageReply handler
func (fl *FuncList) MessageReply(f MessageReplyFunc) {
	fl.messageReply = append(fl.messageReply, f)
}

// MessageEdit handler
func (fl *FuncList) MessageEdit(f MessageEditFunc) {
	fl.messageEdit = append(fl.messageEdit, f)
}

// MessageAllow handler
func (fl *FuncList) MessageAllow(f MessageAllowFunc) {
	fl.messageAllow = append(fl.messageAllow, f)
}

// MessageDeny handler
func (fl *FuncList) MessageDeny(f MessageDenyFunc) {
	fl.messageDeny = append(fl.messageDeny, f)
}

// MessageTypingState handler
func (fl *FuncList) MessageTypingState(f MessageTypingStateFunc) {
	fl.messageTypingState = append(fl.messageTypingState, f)
}

// PhotoNew handler
func (fl *FuncList) PhotoNew(f PhotoNewFunc) {
	fl.photoNew = append(fl.photoNew, f)
}

// PhotoCommentNew handler
func (fl *FuncList) PhotoCommentNew(f PhotoCommentNewFunc) {
	fl.photoCommentNew = append(fl.photoCommentNew, f)
}

// PhotoCommentEdit handler
func (fl *FuncList) PhotoCommentEdit(f PhotoCommentEditFunc) {
	fl.photoCommentEdit = append(fl.photoCommentEdit, f)
}

// PhotoCommentRestore handler
func (fl *FuncList) PhotoCommentRestore(f PhotoCommentRestoreFunc) {
	fl.photoCommentRestore = append(fl.photoCommentRestore, f)
}

// PhotoCommentDelete handler
func (fl *FuncList) PhotoCommentDelete(f PhotoCommentDeleteFunc) {
	fl.photoCommentDelete = append(fl.photoCommentDelete, f)
}

// AudioNew handler
func (fl *FuncList) AudioNew(f AudioNewFunc) {
	fl.audioNew = append(fl.audioNew, f)
}

// VideoNew handler
func (fl *FuncList) VideoNew(f VideoNewFunc) {
	fl.videoNew = append(fl.videoNew, f)
}

// VideoCommentNew handler
func (fl *FuncList) VideoCommentNew(f VideoCommentNewFunc) {
	fl.videoCommentNew = append(fl.videoCommentNew, f)
}

// VideoCommentEdit handler
func (fl *FuncList) VideoCommentEdit(f VideoCommentEditFunc) {
	fl.videoCommentEdit = append(fl.videoCommentEdit, f)
}

// VideoCommentRestore handler
func (fl *FuncList) VideoCommentRestore(f VideoCommentRestoreFunc) {
	fl.videoCommentRestore = append(fl.videoCommentRestore, f)
}

// VideoCommentDelete handler
func (fl *FuncList) VideoCommentDelete(f VideoCommentDeleteFunc) {
	fl.videoCommentDelete = append(fl.videoCommentDelete, f)
}

// WallPostNew handler
func (fl *FuncList) WallPostNew(f WallPostNewFunc) {
	fl.wallPostNew = append(fl.wallPostNew, f)
}

// WallRepost handler
func (fl *FuncList) WallRepost(f WallRepostFunc) {
	fl.wallRepost = append(fl.wallRepost, f)
}

// WallReplyNew handler
func (fl *FuncList) WallReplyNew(f WallReplyNewFunc) {
	fl.wallReplyNew = append(fl.wallReplyNew, f)
}

// WallReplyEdit handler
func (fl *FuncList) WallReplyEdit(f WallReplyEditFunc) {
	fl.wallReplyEdit = append(fl.wallReplyEdit, f)
}

// WallReplyRestore handler
func (fl *FuncList) WallReplyRestore(f WallReplyRestoreFunc) {
	fl.wallReplyRestore = append(fl.wallReplyRestore, f)
}

// WallReplyDelete handler
func (fl *FuncList) WallReplyDelete(f WallReplyDeleteFunc) {
	fl.wallReplyDelete = append(fl.wallReplyDelete, f)
}

// BoardPostNew handler
func (fl *FuncList) BoardPostNew(f BoardPostNewFunc) {
	fl.boardPostNew = append(fl.boardPostNew, f)
}

// BoardPostEdit handler
func (fl *FuncList) BoardPostEdit(f BoardPostEditFunc) {
	fl.boardPostEdit = append(fl.boardPostEdit, f)
}

// BoardPostRestore handler
func (fl *FuncList) BoardPostRestore(f BoardPostRestoreFunc) {
	fl.boardPostRestore = append(fl.boardPostRestore, f)
}

// BoardPostDelete handler
func (fl *FuncList) BoardPostDelete(f BoardPostDeleteFunc) {
	fl.boardPostDelete = append(fl.boardPostDelete, f)
}

// MarketCommentNew handler
func (fl *FuncList) MarketCommentNew(f MarketCommentNewFunc) {
	fl.marketCommentNew = append(fl.marketCommentNew, f)
}

// MarketCommentEdit handler
func (fl *FuncList) MarketCommentEdit(f MarketCommentEditFunc) {
	fl.marketCommentEdit = append(fl.marketCommentEdit, f)
}

// MarketCommentRestore handler
func (fl *FuncList) MarketCommentRestore(f MarketCommentRestoreFunc) {
	fl.marketCommentRestore = append(fl.marketCommentRestore, f)
}

// MarketCommentDelete handler
func (fl *FuncList) MarketCommentDelete(f MarketCommentDeleteFunc) {
	fl.marketCommentDelete = append(fl.marketCommentDelete, f)
}

// GroupLeave handler
func (fl *FuncList) GroupLeave(f GroupLeaveFunc) {
	fl.groupLeave = append(fl.groupLeave, f)
}

// GroupJoin handler
func (fl *FuncList) GroupJoin(f GroupJoinFunc) {
	fl.groupJoin = append(fl.groupJoin, f)
}

// UserBlock handler
func (fl *FuncList) UserBlock(f UserBlockFunc) {
	fl.userBlock = append(fl.userBlock, f)
}

// UserUnblock handler
func (fl *FuncList) UserUnblock(f UserUnblockFunc) {
	fl.userUnblock = append(fl.userUnblock, f)
}

// PollVoteNew handler
func (fl *FuncList) PollVoteNew(f PollVoteNewFunc) {
	fl.pollVoteNew = append(fl.pollVoteNew, f)
}

// GroupOfficersEdit handler
func (fl *FuncList) GroupOfficersEdit(f GroupOfficersEditFunc) {
	fl.groupOfficersEdit = append(fl.groupOfficersEdit, f)
}

// GroupChangeSettings handler
func (fl *FuncList) GroupChangeSettings(f GroupChangeSettingsFunc) {
	fl.groupChangeSettings = append(fl.groupChangeSettings, f)
}

// GroupChangePhoto handler
func (fl *FuncList) GroupChangePhoto(f GroupChangePhotoFunc) {
	fl.groupChangePhoto = append(fl.groupChangePhoto, f)
}

// VkpayTransaction handler
func (fl *FuncList) VkpayTransaction(f VkpayTransactionFunc) {
	fl.vkpayTransaction = append(fl.vkpayTransaction, f)
}

// LeadFormsNew handler
func (fl *FuncList) LeadFormsNew(f LeadFormsNewFunc) {
	fl.leadFormsNew = append(fl.leadFormsNew, f)
}

// AppPayload handler
func (fl *FuncList) AppPayload(f AppPayloadFunc) {
	fl.appPayload = append(fl.appPayload, f)
}

// MessageRead handler
func (fl *FuncList) MessageRead(f MessageReadFunc) {
	fl.messageRead = append(fl.messageRead, f)
}

// NOTE: like_add like_remove
