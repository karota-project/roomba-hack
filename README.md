kcapture
============

## Feature
- 接続したUSBCAMからV4L2で映像を取得しffmpegでエンコーディングしてffserverで配信する

## Usage
```bash
func Start(cmd string, args []string) (err error) 
func Stop(proc string) (err error) 
```