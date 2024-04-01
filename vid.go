// Package vid provides a simple wrapper around
// github.com/blackjack/webcam library.
package vid

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/blackjack/webcam"
	"github.com/siuyin/dflt"
)

const (
	V4L2_PIX_FMT_PJPG = webcam.PixelFormat(0x47504A50)
	V4L2_PIX_FMT_YUYV = webcam.PixelFormat(0x56595559)
)

// Capture uses /dev/video0 to capture a still timestamped image.
func Capture(basename string) error {
	//eg. basename := "img-loc123-" + time.Now().Format("20060304-150405")
	cam, err := webcam.Open(dflt.EnvString("VIDEO_DEVICE", "/dev/video0"))
	if err != nil {
		return fmt.Errorf("webcam.Open: %v", err)
	}
	defer cam.Close()

	frames := 1
	if err := writeYUV(cam, basename, frames); err != nil {
		return err
	}

	if err := writePNG(basename); err != nil {
		return err
	}

	if err := os.Remove(basename + ".yuv"); err != nil {
		return fmt.Errorf("delete %s.yuv: %v", basename, err)
	}

	return nil
}

func writeYUV(cam *webcam.Webcam, basename string, frames int) error {
	format := V4L2_PIX_FMT_YUYV
	if _, _, _, err := cam.SetImageFormat(format, 1280, 720); err != nil {
		return fmt.Errorf("SetImageFormat: %v", err)
	}

	if err := cam.StartStreaming(); err != nil {
		return fmt.Errorf("StartStreaming: %v", err)
	}
	defer cam.StopStreaming()

	of, err := os.Create(basename + ".yuv")
	if err != nil {
		return fmt.Errorf("os.Create: %v", err)
	}
	defer of.Close()

	const numberOfWarmUpFrames = 3
	if err := warmUp(cam, numberOfWarmUpFrames); err != nil {
		return fmt.Errorf("warmUp: %v", err)
	}

	for n := 0; n < frames; n++ {
		if err := cam.WaitForFrame(5); err != nil {
			return fmt.Errorf("WaitForFrame: %v", err)
		}

		frame, err := cam.ReadFrame()
		if err != nil {
			return fmt.Errorf("ReadFrame: %d: %v", frame, err)
		}

		if len(frame) == 0 {
			return fmt.Errorf("ReadFrame returned empty frame")
		}

		of.Write(frame)
	}

	return nil
}

func warmUp(cam *webcam.Webcam, n int) error {
	for i := 0; i < n; i++ {
		if err := cam.WaitForFrame(5); err != nil {
			return fmt.Errorf("WaitForFrame: %v", err)
		}

		_, err := cam.ReadFrame()
		if err != nil {
			return fmt.Errorf("ReadFrame: %v", err)
		}
	}
	return nil
}

func writePNG(basename string) error {
	cmdStr := "ffmpeg"
	cmdArgs := fmt.Sprintf("-pixel_format yuyv422 -video_size 1280x720 -y -i %s.yuv -r 1 -update 1 %s.jpg", basename, basename)
	out, err := exec.Command(cmdStr, strings.Split(cmdArgs, " ")...).Output()
	if err != nil {
		log.Printf("exec FFMPEG: %s\n", out)
		return fmt.Errorf("exec FFMPEG: %v", err)
	}

	return nil
}

func writeMKV(basename string) error {
	cmdStr := "ffmpeg"
	cmdArgs := fmt.Sprintf("-pixel_format yuyv422 -video_size 1280x720 -y -i %s.yuv %s.mkv", basename, basename)
	out, err := exec.Command(cmdStr, strings.Split(cmdArgs, " ")...).Output()
	if err != nil {
		log.Printf("exec FFMPEG: %s\n", out)
		return fmt.Errorf("exec FFMPEG: %v", err)
	}

	return nil
}

// Frames captures n frames of video.
func Frames(basename string, n int) error {
	cam, err := webcam.Open(dflt.EnvString("VIDEO_DEVICE", "/dev/video0"))
	if err != nil {
		return fmt.Errorf("webcam.Open: %v", err)
	}
	defer cam.Close()

	if err := writeYUV(cam, basename, n); err != nil {
		return err
	}

	if err := writeMKV(basename); err != nil {
		return err
	}

	if err := os.Remove(basename + ".yuv"); err != nil {
		return fmt.Errorf("delete %s.yuv: %v", basename, err)
	}

	return nil
}
