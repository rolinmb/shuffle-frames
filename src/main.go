package main

import (
  "os"
  "os/exec"
  "strconv"
  "time"
)

const (
  VIDEOIN = "video-in/teaser0_06022024.mp4"
  OUTNAME = "video-out/result.mp4"
)

func main() {
  ts := strconv.FormatInt(time.Now().Unix(), 10)
  ffmpegCmd := exec.Command(
    "ffmpeg",
    "-i", VIDEOIN,
    "-vf", "shuffleframes=seed="+ts,
    "-c:a", "copy",
    OUTNAME,
  )
  ffmpegCmd.Stdout = os.Stdout
  ffmpegCmd.Stderr = os.Stderr
  if err := ffmpegCmd.Run(); err != nil {
    panic(err)
  }
}
