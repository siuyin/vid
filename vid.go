package vid

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/blackjack/webcam"
	"github.com/siuyin/dflt"
)

const (
	V4L2_PIX_FMT_PJPG = webcam.PixelFormat(0x47504A50)
	V4L2_PIX_FMT_YUYV = webcam.PixelFormat(0x56595559)
)

func Capture() error {
	cam, err := webcam.Open(dflt.EnvString("VIDEO_DEVICE", "/dev/video0"))
	if err != nil {
		return fmt.Errorf("webcam.Open: %v", err)
	}
	defer cam.Close()

	//output := dflt.EnvString("IMAGE_FILE", "myImage")
	output := "img-loc123-" + time.Now().Format("20060304-150405")
	if err := writeYUV(cam, output); err != nil {
		return err
	}

	if err := writePNG(output); err != nil {
		return err
	}

	return nil
}

func writeYUV(cam *webcam.Webcam, basename string) error {
	format := V4L2_PIX_FMT_YUYV
	if _, _, _, err := cam.SetImageFormat(format, 1280, 720); err != nil {
		return fmt.Errorf("SetImageFormat: %v", err)
	}

	if err := cam.StartStreaming(); err != nil {
		return fmt.Errorf("StartStreaming: %v", err)
	}
	defer cam.StopStreaming()

	if err := cam.WaitForFrame(5); err != nil {
		return fmt.Errorf("WaitForFrame: %v", err)
	}

	frame, err := cam.ReadFrame()
	if err != nil {
		return fmt.Errorf("ReadFrame: %v", err)
	}

	if len(frame) == 0 {
		return fmt.Errorf("ReadFrame returned empty frame")
	}

	of, err := os.Create(basename + ".yuv")
	if err != nil {
		return fmt.Errorf("os.Create: %v", err)
	}
	defer of.Close()

	of.Write(frame)

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
