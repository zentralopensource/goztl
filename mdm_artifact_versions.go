package goztl

// MDMArtifactVersion represents a Zentral MDM artifact version
type MDMArtifactVersion struct {
	ArtifactID       string     `json:"artifact"`
	IOS              bool       `json:"ios"`
	IOSMaxVersion    string     `json:"ios_max_version"`
	IOSMinVersion    string     `json:"ios_min_version"`
	IPadOS           bool       `json:"ipados"`
	IPadOSMaxVersion string     `json:"ipados_max_version"`
	IPadOSMinVersion string     `json:"ipados_min_version"`
	MacOS            bool       `json:"macos"`
	MacOSMaxVersion  string     `json:"macos_max_version"`
	MacOSMinVersion  string     `json:"macos_min_version"`
	TVOS             bool       `json:"tvos"`
	TVOSMaxVersion   string     `json:"tvos_max_version"`
	TVOSMinVersion   string     `json:"tvos_min_version"`
	DefaultShard     int        `json:"default_shard"`
	ShardModulo      int        `json:"shard_modulo"`
	ExcludedTagIDs   []int      `json:"excluded_tags"`
	TagShards        []TagShard `json:"tag_shards"`
	Version          int        `json:"version"`
	Created          Timestamp  `json:"created_at"`
	Updated          Timestamp  `json:"updated_at"`
}

// MDMArtifactVersionRequest represents a request to create or update a MDM artifact version
type MDMArtifactVersionRequest struct {
	ArtifactID       string     `json:"artifact"`
	IOS              bool       `json:"ios"`
	IOSMaxVersion    string     `json:"ios_max_version"`
	IOSMinVersion    string     `json:"ios_min_version"`
	IPadOS           bool       `json:"ipados"`
	IPadOSMaxVersion string     `json:"ipados_max_version"`
	IPadOSMinVersion string     `json:"ipados_min_version"`
	MacOS            bool       `json:"macos"`
	MacOSMaxVersion  string     `json:"macos_max_version"`
	MacOSMinVersion  string     `json:"macos_min_version"`
	TVOS             bool       `json:"tvos"`
	TVOSMaxVersion   string     `json:"tvos_max_version"`
	TVOSMinVersion   string     `json:"tvos_min_version"`
	DefaultShard     int        `json:"default_shard"`
	ShardModulo      int        `json:"shard_modulo"`
	ExcludedTagIDs   []int      `json:"excluded_tags"`
	TagShards        []TagShard `json:"tag_shards"`
	Version          int        `json:"version"`
}
