kcapture
============

## Feature
- Roomba-Hack
- ffmpeg / ffserver utils
- 接続したUSBCAMからV4L2で映像を取得しffmpegでエンコーディングしてffserverで配信する

## Usage
- 定義

<pre>
func Start(cmd string, args []string) (isSuccessed bool, err error) 
func Stop(proc string) (isSuccessed bool, err error) 
</pre>
