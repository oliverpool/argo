package argo

import "encoding/base64"

// AddURI adds a new download
// - uris is an array of HTTP/FTP/SFTP/BitTorrent URIs (strings) pointing to the same resource. If you mix URIs pointing to different resources, then the download may fail or be corrupted without aria2 complaining. When adding BitTorrent Magnet URIs, uris must have only one element and it should be BitTorrent Magnet URI. options is a struct and its members are pairs of option name and value.
// This method returns the GID of the newly registered download.
func (c Client) AddURI(uris []string, options ...Option) (reply GIDwithID, err error) {
	opt := c.mergeOptions(options...)
	err = c.Caller.CallWithID("aria2.addUri", &reply.GID, &reply.ID, uris, opt)
	return
}

// AddTorrent adds a BitTorrent download by uploading the content of a ".torrent" file. If you want to add a BitTorrent Magnet URI, use the AddUri() method instead.
// This method returns the GID of the newly registered download
// If --rpc-save-upload-metadata is true, the uploaded data is saved as a file named as the hex string of SHA-1 hash of data plus ".torrent" in the directory specified by --dir option. E.g. a file name might be 0a3893293e27ac0490424c06de4d09242215f0a6.torrent. If a file with the same name already exists, it is overwritten!
// If the file cannot be saved successfully or --rpc-save-upload-metadata is false, the downloads added by this method are not saved by --save-session.
func (c Client) AddTorrent(content []byte, options ...Option) (reply GIDwithID, err error) {
	return c.AddTorrentWithWebSeed(content, []string{}, options...)
}

// AddTorrentWithWebSeed adds a BitTorrent download by uploading the content of a ".torrent" file. If you want to add a BitTorrent Magnet URI, use the AddUri() method instead.
// uris is an array of URIs (string). uris is used for Web-seeding. For single file torrents, the URI can be a complete URI pointing to the resource; if URI ends with /, name in torrent file is added. For multi-file torrents, name and path in torrent are added to form a URI for each file.
// This method returns the GID of the newly registered download
// If --rpc-save-upload-metadata is true, the uploaded data is saved as a file named as the hex string of SHA-1 hash of data plus ".torrent" in the directory specified by --dir option. E.g. a file name might be 0a3893293e27ac0490424c06de4d09242215f0a6.torrent. If a file with the same name already exists, it is overwritten!
// If the file cannot be saved successfully or --rpc-save-upload-metadata is false, the downloads added by this method are not saved by --save-session.
func (c Client) AddTorrentWithWebSeed(content []byte, uris []string, options ...Option) (reply GIDwithID, err error) {
	opt := c.mergeOptions(options...)
	b64 := base64.StdEncoding.EncodeToString(content)
	err = c.Caller.CallWithID("aria2.addTorrent", &reply.GID, &reply.ID, b64, uris, opt)
	return
}

// AddMetalink adds a Metalink download by uploading a ".metalink" file. metalink is a base64-encoded string which contains the contents of the ".metalink" file.
// This method returns an array of GIDs of newly registered downloads.
// If --rpc-save-upload-metadata is true, the uploaded data is saved as a file named hex string of SHA-1 hash of data plus ".metalink" in the directory specified by --dir option.
// E.g. a file name might be 0a3893293e27ac0490424c06de4d09242215f0a6.metalink.
// If a file with the same name already exists, it is overwritten!
// If the file cannot be saved successfully or --rpc-save-upload-metadata is false, the downloads added by this method are not saved by --save-session.
func (c Client) AddMetalink(content []byte, options ...Option) (reply GIDs, err error) {
	opt := c.mergeOptions(options...)
	b64 := base64.StdEncoding.EncodeToString(content)
	err = c.Caller.CallWithID("aria2.addMetalink", &reply.GIDs, &reply.ID, b64, opt)
	return
}

// Remove the download denoted by gid (string).
// If the specified download is in progress, it is first stopped. The status of the removed download becomes removed.
// This method returns GID of removed download.
func (c Client) Remove(gid GID) (reply GID, err error) {
	err = c.Caller.Call("aria2.remove", &reply, gid)
	return
}

// ForceRemove the download denoted by gid.
// This method behaves just like Remove() except that this method removes the download without performing any actions which take time, such as contacting BitTorrent trackers to unregister the download first.
func (c Client) ForceRemove(gid GID) (reply GID, err error) {
	err = c.Caller.Call("aria2.forceRemove", &reply, gid)
	return
}

// Pause the download denoted by gid (string).
// The status of paused download becomes paused. If the download was active, the download is placed in the front of waiting queue. While the status is paused, the download is not started. To change status to waiting, use the Unpause() method.
// This method returns GID of removed download.
func (c Client) Pause(gid GID) (reply GID, err error) {
	err = c.Caller.Call("aria2.pause", &reply, gid)
	return
}

// PauseAll is equal to calling Pause() for every active/waiting download.
// This methods returns OK.
func (c Client) PauseAll() (reply Ok, err error) {
	err = c.Caller.Call("aria2.pauseAll", &reply)
	return
}

// ForcePause behaves just like Pause() except that this method pauses downloads without performing any actions which take time, such as contacting BitTorrent trackers to unregister the download first.
// This method returns GID of removed download.
func (c Client) ForcePause(gid GID) (reply GID, err error) {
	err = c.Caller.Call("aria2.forcePause", &reply, gid)
	return
}

// ForcePauseAll is equal to calling ForcePause() for every active/waiting download.
// This methods returns OK.
func (c Client) ForcePauseAll() (reply Ok, err error) {
	err = c.Caller.Call("aria2.forcePauseAll", &reply)
	return
}

// Unpause changes the status of the download denoted by gid (string) from paused to waiting, making the download eligible to be restarted.
// This method returns GID of removed download.
func (c Client) Unpause(gid GID) (reply GID, err error) {
	err = c.Caller.Call("aria2.unpause", &reply, gid)
	return
}

// UnpauseAll is equal to calling Unpause() for every active/waiting download.
// This methods returns OK.
func (c Client) UnpauseAll() (reply Ok, err error) {
	err = c.Caller.Call("aria2.unpauseAll", &reply)
	return
}

// TellStatus returns the URIs used in the download denoted by gid (string).
// keys is an array of strings. If specified, the response contains only keys in the keys array. If keys is empty or omitted, the response contains all keys. This is useful when you just want specific keys and avoid unnecessary transfers.
func (c Client) TellStatus(gid GID, keys ...string) (reply StatusInfo, err error) {
	if len(keys) > 0 {
		err = c.Caller.Call("aria2.tellStatus", &reply, gid, keys)
	} else {
		err = c.Caller.Call("aria2.tellStatus", &reply, gid)
	}
	return
}

// GetURIs returns the URIs used in the download denoted by gid (string).
func (c Client) GetURIs(gid GID) (reply []URIInfo, err error) {
	err = c.Caller.Call("aria2.getUris", &reply, gid)
	return
}

// GetFiles returns the URIs used in the download denoted by gid (string).
func (c Client) GetFiles(gid GID) (reply []FileInfo, err error) {
	err = c.Caller.Call("aria2.getFiles", &reply, gid)
	return
}

// GetPeers returns a list peers of the download denoted by gid (string). This method is for BitTorrent only.
func (c Client) GetPeers(gid GID) (reply []PeerInfo, err error) {
	err = c.Caller.Call("aria2.getPeers", &reply, gid)
	return
}

// GetServers returns currently connected HTTP(S)/FTP/SFTP servers of the download denoted by gid (string).
func (c Client) GetServers(gid GID) (reply []ServerInfo, err error) {
	err = c.Caller.Call("aria2.getServers", &reply, gid)
	return
}

// TellActive returns a list of active downloads.
// For the keys parameter, please refer to the aria2.tellStatus() method.
func (c Client) TellActive(keys ...string) (reply []StatusInfo, err error) {
	if len(keys) > 0 {
		err = c.Caller.Call("aria2.tellActive", &reply, keys)
	} else {
		err = c.Caller.Call("aria2.tellActive", &reply)
	}
	return
}

// TellWaiting returns list of waiting downloads, including paused ones.
// offset is an integer and specifies the offset from the download waiting at the front.
// num is an integer and specifies the max.
// number of downloads to be returned. For the keys parameter, please refer to the TellStatus() method.
//
// If offset is a positive integer, this method returns downloads in the range of [offset, offset + num).
//
// offset can be a negative integer. offset == -1 points last download in the waiting queue and offset == -2 points the download before the last download, and so on. Downloads in the response are in reversed order then.
//
// For example, imagine three downloads "A","B" and "C" are waiting in this order. TellWaiting(0, 1) returns ["A"]. TellWaiting(1, 2) returns ["B", "C"]. TellWaiting(-1, 2) returns ["C", "B"].
func (c Client) TellWaiting(offset int, num int, keys ...string) (reply []StatusInfo, err error) {
	if len(keys) > 0 {
		err = c.Caller.Call("aria2.tellWaiting", &reply, offset, num, keys)
	} else {
		err = c.Caller.Call("aria2.tellWaiting", &reply, offset, num)
	}
	return
}

// TellStopped returns a list of stopped downloads.
// offset is an integer and specifies the offset from the least recently stopped download.
// num is an integer and specifies the max. number of downloads to be returned.
//
// For the keys parameter, please refer to the TellStatus() method.
//
// offset and num have the same semantics as described in the TellWaiting() method.
func (c Client) TellStopped(offset int, num int, keys ...string) (reply []StatusInfo, err error) {
	if len(keys) > 0 {
		err = c.Caller.Call("aria2.tellStopped", &reply, offset, num, keys)
	} else {
		err = c.Caller.Call("aria2.tellStopped", &reply, offset, num)
	}
	return
}

// ChangePosition changes the position of the download denoted by gid in the queue.
// pos is an integer.
// strategy is a PositionStrategy.
// If the destination position is less than 0 or beyond the end of the queue, it moves the download to the beginning or the end of the queue respectively.
// The response is an integer denoting the resulting position.
func (c Client) ChangePosition(gid GID, pos int, strategy PositionStrategy) (reply NewPosition, err error) {
	err = c.Caller.Call("aria2.changePosition", &reply, gid, pos, strategy)
	return
}

// ChangeURI removes the URIs in delUris from and appends the URIs in addUris to download denoted by gid.
// delUris and addUris are lists of strings. A download can contain multiple files and URIs are attached to each file.
// fileIndex is used to select which file to remove/attach given URIs. fileIndex is 1-based.
// URIs are appended to the back of the list (to choose position, see ChangeURIWithPosition).
// When removing an URI, if the same URIs exist in download, only one of them is removed for each URI in delUris. In other words, if there are three URIs http://example.org/aria2 and you want remove them all, you have to specify (at least) 3 http://example.org/aria2 in delUris.
// This method returns a list which contains two integers. The first integer is the number of URIs deleted. The second integer is the number of URIs added.
func (c Client) ChangeURI(gid GID, fileIndex int, delURIs, addURIs []string) (reply DeletionAddition, err error) {
	err = c.Caller.Call("aria2.changeUri", &reply, gid, fileIndex, delURIs, addURIs)
	return
}

// ChangeURIWithPosition removes the URIs in delUris from and appends the URIs in addUris to download denoted by gid.
// delUris and addUris are lists of strings. A download can contain multiple files and URIs are attached to each file.
// fileIndex is used to select which file to remove/attach given URIs. fileIndex is 1-based.
// position is used to specify where URIs are inserted in the existing waiting URI list. position is 0-based.
// To append URIs to the back of the list, see ChangeURI.
// This method first executes the removal and then the addition. position is the position after URIs are removed, not the position when this method is called.
// When removing an URI, if the same URIs exist in download, only one of them is removed for each URI in delUris. In other words, if there are three URIs http://example.org/aria2 and you want remove them all, you have to specify (at least) 3 http://example.org/aria2 in delUris.
// This method returns a list which contains two integers. The first integer is the number of URIs deleted. The second integer is the number of URIs added.
func (c Client) ChangeURIWithPosition(gid GID, fileIndex int, delURIs, addURIs []string, position int) (reply DeletionAddition, err error) {
	err = c.Caller.Call("aria2.changeUri", &reply, gid, fileIndex, delURIs, addURIs, position)
	return
}

// GetOption returns options of the download denoted by gid.
// The response is a struct where keys are the names of options.
// Note that this method does not return options which have no default value and have not been set on the command-line, in configuration files or RPC methods.
func (c Client) GetOption(gid GID) (reply Option, err error) {
	err = c.Caller.Call("aria2.getOption", &reply, gid)
	return
}

// ChangeOption changes options of the download denoted by gid (string) dynamically.
//
// The options listed in Input File subsection are available, except for following options:
// dry-run
// metalink-base-uri
// parameterized-uri
// pause
// piece-length
// rpc-save-upload-metadata
//
// Except for the following options, changing the other options of active download makes it restart (restart itself is managed by aria2, and no user intervention is required):
// bt-max-peers
// bt-request-peer-speed-limit
// bt-remove-unselected-file
// force-save
// max-download-limit
// max-upload-limit
//
// This method returns OK for success.
func (c Client) ChangeOption(gid GID, options ...Option) (reply Ok, err error) {
	opt := c.mergeOptions(options...)
	err = c.Caller.Call("aria2.changeOption", &reply, gid, opt)
	return
}

// GetGlobalOption returns the global options.
//  The response is a struct. Its keys are the names of options. Values are strings.
// Note that this method does not return options which have no default value and have not been set on the command-line, in configuration files or RPC methods.
// Because global options are used as a template for the options of newly added downloads, the response contains keys returned by the GetOption() method.
func (c Client) GetGlobalOption() (reply Option, err error) {
	err = c.Caller.Call("aria2.getGlobalOption", &reply)
	return
}

// ChangeGlobalOption changes global options dynamically.
//
// The following options are available:
//
// bt-max-open-files
// download-result
// keep-unfinished-download-result
// log
// log-level
// max-concurrent-downloads
// max-download-result
// max-overall-download-limit
// max-overall-upload-limit
// optimize-concurrent-downloads
// save-cookies
// save-session
// server-stat-of
//
// In addition, options listed in the Input File subsection are available, except for following options:
// checksum
// index-out
// out
// pause
//select-file
//
// With the log option, you can dynamically start logging or change log file. To stop logging, specify an empty string("") as the parameter value.
// Note that log file is always opened in append mode.
func (c Client) ChangeGlobalOption(options ...Option) (reply Ok, err error) {
	opt := c.mergeOptions(options...)
	err = c.Caller.Call("aria2.changeGlobalOption", &reply, opt)
	return
}

// GetGlobalStat returns global statistics such as the overall download and upload speeds.
func (c Client) GetGlobalStat() (reply GlobalStatInfo, err error) {
	err = c.Caller.Call("aria2.getGlobalStat", &reply)
	return
}

// PurgeDownloadResult purges completed/error/removed downloads to free memory.
func (c Client) PurgeDownloadResult() (reply Ok, err error) {
	err = c.Caller.Call("aria2.purgeDownloadResult", &reply)
	return
}

// RemoveDownloadResult removes a completed/error/removed download denoted by gid from memory.
func (c Client) RemoveDownloadResult(gid GID) (reply Ok, err error) {
	err = c.Caller.Call("aria2.removeDownloadResult", &reply, gid)
	return
}

// GetVersion returns the version of aria2 and the list of enabled features.
func (c Client) GetVersion() (reply VersionInfo, err error) {
	err = c.Caller.Call("aria2.getVersion", &reply)
	return
}

// GetSessionInfo returns session information.
func (c Client) GetSessionInfo() (reply SessionInfo, err error) {
	err = c.Caller.Call("aria2.getSessionInfo", &reply)
	return
}

// Shutdown shuts down aria2.
func (c Client) Shutdown() (reply Ok, err error) {
	err = c.Caller.Call("aria2.shutdown", &reply)
	return
}

// ForceShutdown shuts down aria2. This method behaves like Shutdown without performing any actions which take time, such as contacting BitTorrent trackers to unregister downloads first.
func (c Client) ForceShutdown() (reply Ok, err error) {
	err = c.Caller.Call("aria2.forceShutdown", &reply)
	return
}

// SaveSession saves the current session to a file specified by the --save-session option.
func (c Client) SaveSession() (reply Ok, err error) {
	err = c.Caller.Call("aria2.saveSession", &reply)
	return
}
