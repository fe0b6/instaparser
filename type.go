package instaparser

type PostInfo struct {
	Id        string
	Published int
	Likes     int
	Comments  int
	Text      string
	Owner     PostDataOwner
}

type PostData struct {
	EntryData PostDataEntryData `json:"entry_data"`
}

type PostDataEntryData struct {
	PostPage []PostDataPostPage `json:"PostPage"`
}

type PostDataPostPage struct {
	Graphql PostDataGraphql `json:"graphql"`
}

type PostDataGraphql struct {
	ShortcodeMedia PostDataShortcodeMedia `json:"shortcode_media"`
}

type PostDataShortcodeMedia struct {
	Id                   string                     `json:"id"`
	TakenAtTimestamp     int                        `json:"taken_at_timestamp"`
	EdgeMediaToCaption   PostDataEdgeMediaToCaption `json:"edge_media_to_caption"`
	EdgeMediaToComment   DataCount                  `json:"edge_media_to_comment"`
	EdgeMediaPreviewLike DataCount                  `json:"edge_media_preview_like"`
	Owner                PostDataOwner              `json:"owner"`
}

type PostDataEdgeMediaToCaption struct {
	Edges []PostDataEdges `json:"edges"`
}

type PostDataEdges struct {
	Node PostDataEdgesNode `json:"node"`
}

type PostDataEdgesNode struct {
	Text string `json:"text"`
}

type PostDataOwner struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

type ProfileInfo struct {
	FollowedBy int
}

type ProfileData struct {
	EntryData ProfileDataEntryData `json:"entry_data"`
}

type ProfileDataEntryData struct {
	ProfilePage []ProfileDataProfilePage `json:"ProfilePage"`
}

type ProfileDataProfilePage struct {
	Graphql ProfileDataGraphql `json:"graphql"`
}

type ProfileDataGraphql struct {
	User ProfileDataUser `json:"user"`
}

type ProfileDataUser struct {
	FollowedBy DataCount `json:"edge_followed_by"`
}

type DataCount struct {
	Count int `json:"count"`
}
